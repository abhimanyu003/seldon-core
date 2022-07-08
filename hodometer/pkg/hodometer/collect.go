package hodometer

import (
	"context"
	"fmt"
	"io"
	"reflect"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	pb "github.com/seldonio/seldon-core/hodometer/apis"
)

const (
	subscriberName = "hodometer"

	maxNodesPerCall = 100
)

type Collector interface {
	Collect(ctx context.Context, level MetricsLevel) *UsageMetrics
}

type kubernetesMetrics struct {
	version   string
	isGlobal  bool
	nodeCount uint
}

type schedulerMetrics struct {
	version string
}

type experimentMetrics struct {
	count uint
}

type pipelineMetrics struct {
	count uint
}

type serverMetrics struct {
	count             uint
	replicaCount      uint
	multimodelEnabled uint
	overcommitEnabled uint
	memoryBytes       uint
}

type modelMetrics struct {
	count uint
}

var _ Collector = (*SeldonCoreCollector)(nil)

type SeldonCoreCollector struct {
	schedulerClient pb.SchedulerClient
	k8sClient       kubernetes.Interface
	logger          logrus.FieldLogger
	clusterId       string
}

func NewSeldonCoreCollector(
	logger logrus.FieldLogger,
	schedulerHost string,
	schedulerPort uint,
	clusterId string,
) (*SeldonCoreCollector, error) {
	logger = logger.WithField("source", "SeldonCoreCollector")

	connectOptions := []grpc.DialOption{
		grpc.WithInsecure(), // TODO - support TLS
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", schedulerHost, schedulerPort),
		connectOptions...,
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewSchedulerClient(conn)

	clusterId = parseOrCreateClusterId(clusterId)

	k8sClient := getKubernetesClientset(logger)

	return &SeldonCoreCollector{
		schedulerClient: client,
		k8sClient:       k8sClient,
		logger:          logger,
		clusterId:       clusterId,
	}, nil
}

func parseOrCreateClusterId(clusterId string) string {
	id, err := uuid.Parse(clusterId)
	if err != nil {
		id = uuid.New()
	}
	return id.String()
}

func getKubernetesClientset(logger logrus.FieldLogger) *kubernetes.Clientset {
	logger = logger.WithField("func", "getKubernetesClientset")

	c, err := rest.InClusterConfig()
	if err != nil {
		logger.WithError(err).Warn("cannot create Kubernetes client")
		return nil
	}

	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		logger.WithError(err).Warn("cannot create Kubernetes client")
		return nil
	}

	return clientset
}

func (scc *SeldonCoreCollector) Collect(ctx context.Context, level MetricsLevel) *UsageMetrics {
	logger := scc.logger.WithField("func", "Collect")

	kubernetesResults := make(chan *kubernetesMetrics)
	schedulerResults := make(chan *schedulerMetrics)
	experimentResults := make(chan *experimentMetrics)
	pipelineResults := make(chan *pipelineMetrics)
	serverResults := make(chan *serverMetrics)
	modelResults := make(chan *modelMetrics)

	go func() {
		logger.Info("collecting Kubernetes details")
		kubernetesResults <- scc.collectKubernetes(ctx)
		close(kubernetesResults)
	}()
	go func() {
		logger.Info("collecting scheduler details")
		schedulerResults <- scc.collectScheduler(ctx)
		close(schedulerResults)
	}()
	go func() {
		logger.Info("collecting experiments")
		experimentResults <- scc.collectExperiments(ctx)
		close(experimentResults)
	}()
	go func() {
		logger.Info("collecting pipelines")
		pipelineResults <- scc.collectPipelines(ctx)
		close(pipelineResults)
	}()
	go func() {
		logger.Info("collecting servers")
		serverResults <- scc.collectServers(ctx)
		close(serverResults)
	}()
	go func() {
		logger.Info("collecting models")
		modelResults <- scc.collectModels(ctx)
		close(modelResults)
	}()

	kubernetesDetails := <-kubernetesResults
	schedulerDetails := <-schedulerResults
	experiments := <-experimentResults
	pipelines := <-pipelineResults
	servers := <-serverResults
	models := <-modelResults

	logger.Info("collating metrics")

	um := &UsageMetrics{}
	um.CollectorVersion = BuildVersion
	um.CollectorGitCommit = GitCommit
	um.ClusterId = scc.clusterId
	um.NodeCount = 1

	if kubernetesDetails != nil {
		um.NodeCount = kubernetesDetails.nodeCount
		um.KubernetesMetrics.KubernetesVersion = kubernetesDetails.version
		um.KubernetesMetrics.IsGlobalInstallation = kubernetesDetails.isGlobal
	}
	if schedulerDetails != nil {
		um.SeldonCoreVersion = schedulerDetails.version
	}
	if level > metricsLevelCluster {
		if experiments != nil {
			um.ResourceMetrics.ExperimentCount = experiments.count
		}
		if pipelines != nil {
			um.ResourceMetrics.PipelineCount = pipelines.count
		}
		if servers != nil {
			um.ResourceMetrics.ServerCount = servers.count
			um.ResourceMetrics.ServerReplicaCount = servers.replicaCount
		}
		if models != nil {
			um.ResourceMetrics.ModelCount = models.count
		}
	}
	if level > metricsLevelResource {
		if servers != nil {
			um.FeatureMetrics.MultimodelEnabledCount = servers.multimodelEnabled
			um.FeatureMetrics.OvercommitEnabledCount = servers.overcommitEnabled
			um.FeatureMetrics.ServerMemoryGbSum = float32(servers.memoryBytes) / 1e9
		}
	}

	logger.Debugf("%#v", *um)

	return um
}

func (scc *SeldonCoreCollector) collectKubernetes(ctx context.Context) *kubernetesMetrics {
	if scc.k8sClient == nil || reflect.ValueOf(scc.k8sClient).IsNil() {
		return nil
	}

	km := &kubernetesMetrics{}
	scc.updateKubernetesVersion(km)
	scc.updateKubernetesNodes(ctx, km)
	return km
}

func (scc *SeldonCoreCollector) updateKubernetesVersion(metrics *kubernetesMetrics) {
	logger := scc.logger.WithField("func", "updateKubernetesVersion")

	version, err := scc.k8sClient.Discovery().ServerVersion()
	if err != nil {
		logger.WithError(err).Error("unable to retrieve server version")
		return
	}
	metrics.version = version.GitVersion
}

func (scc *SeldonCoreCollector) updateKubernetesNodes(ctx context.Context, metrics *kubernetesMetrics) {
	logger := scc.logger.WithField("func", "updateKubernetesNodes")

	nodeClient := scc.k8sClient.CoreV1().Nodes()
	continuation := ""
	nodeCount := 0

	for {
		listOps := metaV1.ListOptions{Limit: maxNodesPerCall, Continue: continuation}
		nodes, err := nodeClient.List(ctx, listOps)
		if err != nil {
			logger.WithError(err).Error("unable to retrieve node details")
			return
		}

		nodeCount += len(nodes.Items)

		if nodes.Continue == "" {
			break
		}
		continuation = nodes.Continue
	}
	metrics.nodeCount = uint(nodeCount)
}

func (scc *SeldonCoreCollector) collectScheduler(ctx context.Context) *schedulerMetrics {
	logger := scc.logger.WithField("func", "collectScheduler")

	request := &pb.SchedulerStatusRequest{SubscriberName: subscriberName}
	response, err := scc.schedulerClient.SchedulerStatus(ctx, request)
	if err != nil {
		logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
		return nil
	}

	sm := &schedulerMetrics{}
	sm.version = response.ApplicationVersion
	return sm
}

func (scc *SeldonCoreCollector) collectExperiments(ctx context.Context) *experimentMetrics {
	logger := scc.logger.WithField("func", "collectExperiments")

	request := &pb.ExperimentStatusRequest{
		SubscriberName: subscriberName,
		Name:           nil,
	}
	subscription, err := scc.schedulerClient.ExperimentStatus(ctx, request)
	if err != nil {
		logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
		return nil
	}

	numExperiments := uint(0)
	for {
		exp, err := subscription.Recv()
		if err == io.EOF {
			return &experimentMetrics{count: numExperiments}
		}
		if err != nil {
			logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
			return nil
		}

		if exp.Active {
			numExperiments++
		}
	}
}

func (scc *SeldonCoreCollector) collectPipelines(ctx context.Context) *pipelineMetrics {
	logger := scc.logger.WithField("func", "collectPipelines")

	request := &pb.PipelineStatusRequest{
		SubscriberName: subscriberName,
		Name:           nil, // Request all pipelines
		AllVersions:    false,
	}

	subscription, err := scc.schedulerClient.PipelineStatus(ctx, request)
	if err != nil {
		logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
		return nil
	}

	metrics := &pipelineMetrics{}
	for {
		p, err := subscription.Recv()
		if err == io.EOF {
			return metrics
		}
		if err != nil {
			logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
			return nil
		}

		updatePipelineMetrics(metrics, p)
	}
}

func updatePipelineMetrics(metrics *pipelineMetrics, status *pb.PipelineStatusResponse) {
	if isPipelineActive(status) {
		metrics.count++
	}
}

func isPipelineActive(p *pb.PipelineStatusResponse) bool {
	if p == nil || len(p.Versions) == 0 {
		return false
	}

	isActive := false
	for _, v := range p.Versions {
		if v == nil || v.State == nil {
			continue
		}

		if v.State.Status == pb.PipelineVersionState_PipelineCreate ||
			v.State.Status == pb.PipelineVersionState_PipelineCreating ||
			v.State.Status == pb.PipelineVersionState_PipelineReady {
			isActive = true
		}
	}
	return isActive
}

func (scc *SeldonCoreCollector) collectServers(ctx context.Context) *serverMetrics {
	logger := scc.logger.WithField("func", "collectServers")

	request := &pb.ServerStatusRequest{
		SubscriberName: subscriberName,
		Name:           nil,
	}
	subscription, err := scc.schedulerClient.ServerStatus(ctx, request)
	if err != nil {
		logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
		return nil
	}

	metrics := &serverMetrics{}
	for {
		s, err := subscription.Recv()
		if err == io.EOF {
			return metrics
		}
		if err != nil {
			logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
			return nil
		}

		updateServerMetrics(metrics, s)
	}
}

func updateServerMetrics(metrics *serverMetrics, status *pb.ServerStatusResponse) {
	if status.ExpectedReplicas > 0 || status.AvailableReplicas > 0 {
		metrics.count++
		if status.ExpectedReplicas > 0 {
			metrics.replicaCount += uint(status.ExpectedReplicas)
		} else {
			metrics.replicaCount += uint(status.AvailableReplicas)
		}

		for _, r := range status.Resources {
			if r.OverCommitPercentage > 0 {
				// Overcommitting is redundant/useless without multi-model serving
				metrics.overcommitEnabled++
				metrics.multimodelEnabled++
			} else if r.NumLoadedModels > 1 {
				metrics.multimodelEnabled++
			}

			metrics.memoryBytes += uint(r.TotalMemoryBytes)
		}
	}
}

func (scc *SeldonCoreCollector) collectModels(ctx context.Context) *modelMetrics {
	logger := scc.logger.WithField("func", "collectModels")

	request := &pb.ModelStatusRequest{
		SubscriberName: subscriberName,
		Model:          nil,
	}
	subscription, err := scc.schedulerClient.ModelStatus(ctx, request)
	if err != nil {
		logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
		return nil
	}

	metrics := &modelMetrics{}
	for {
		m, err := subscription.Recv()
		if err == io.EOF {
			return metrics
		}
		if err != nil {
			logger.WithError(err).Error("unable to fetch from Seldon Core scheduler")
			return nil
		}

		if !m.Deleted {
			metrics.count++
		}
	}
}
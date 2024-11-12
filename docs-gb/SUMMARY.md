# Table of contents

* [Home](README.md)
* [Getting Started](getting-started/README.md)
  * [Docker Installation](getting-started/docker-installation.md)
  * [Kubernetes Installation](getting-started/kubernetes-installation/README.md)
    * [Ansible](getting-started/kubernetes-installation/ansible.md)
    * [Helm](getting-started/kubernetes-installation/helm.md)
    * [Security](getting-started/kubernetes-installation/security/README.md)
      * [AWS MSK mTLS](getting-started/kubernetes-installation/security/aws-msk-mtls.md)
      * [AWS MSK SASL](getting-started/kubernetes-installation/security/aws-msk-sasl.md)
      * [Azure Event Hub SASL Example](getting-started/kubernetes-installation/security/azure-event-hub-sasl.md)
      * [Confluent Cloud Oauth 2.0 Example](getting-started/kubernetes-installation/security/confluent-oauth.md)
      * [Confluent Cloud SASL Example](getting-started/kubernetes-installation/security/confluent-sasl.md)
      * [Strimzi mTLS Example](getting-started/kubernetes-installation/security/strimzi-mtls.md)
      * [Strimzi SASL Example](getting-started/kubernetes-installation/security/strimzi-sasl.md)
      * [Reference](getting-started/kubernetes-installation/security/reference.md)
  * [Configuration](getting-started/configuration.md)
  * [Seldon CLI](getting-started/cli.md)
* [APIs](apis/README.md)
  * [Internal](apis/internal/README.md)
    * [Chainer](apis/internal/chainer.md)
    * [Agent](apis/internal/agent.md)
  * [Inference](apis/inference/README.md)
    * [Open Inference Protocol](apis/inference/v2.md)
  * [Scheduler](apis/scheduler.md)
* [Architecture](architecture/README.md)
  * [DataFlow](architecture/dataflow.md)
* [Examples](examples/README.md)
  * [Local examples](examples/local-examples.md)
  * [Kubernetes examples](examples/k8s-examples.md)
  * [Huggingface models](examples/huggingface.md)
  * [Model zoo](examples/model-zoo.md)
  * [Artifact versions](examples/multi-version.md)
  * [Pipeline examples](examples/pipeline-examples.md)
  * [Pipeline to pipeline examples](examples/pipeline-to-pipeline.md)
  * [Explainer examples](examples/explainer-examples.md)
  * [Custom Servers](examples/custom-servers.md)
  * [Local experiments](examples/local-experiments.md)
  * [Experiment version examples](examples/experiment-versions.md)
  * [Inference examples](examples/inference.md)
  * [Tritonclient examples](examples/tritonclient-examples.md)
  * [Batch Inference examples (kubernetes)](examples/batch-examples-k8s.md)
  * [Batch Inference examples (local)](examples/batch-examples-local.md)
  * [Checking Pipeline readiness](examples/pipeline-ready-and-metadata.md)
  * [Multi-Namespace Kubernetes](examples/k8s-clusterwide.md)
  * [Huggingface speech to sentiment with explanations pipeline](examples/speech-to-sentiment.md)
  * [Production image classifier with drift and outlier monitoring](examples/cifar10.md)
  * [Production income classifier with drift, outlier and explanations](examples/income.md)
  * [Conditional pipeline with pandas query model](examples/pandasquery.md)
  * [Kubernetes Server with PVC](examples/k8s-pvc.md)
* [Kubernetes](kubernetes/README.md)
  * [Scaling](kubernetes/scaling.md)
  * [Autoscaling](kubernetes/autoscaling.md)
  * [HPA Autoscaling in single-model serving](kubernetes/hpa-rps-autoscaling.md)
  * [Tracing](kubernetes/tracing.md)
  * [Storage Secrets](kubernetes/storage-secrets.md)
  * [Kafka](kubernetes/kafka.md)
  * [Metrics](kubernetes/metrics.md)
  * [Resources](kubernetes/resources/README.md)
    * [Model](kubernetes/resources/model.md)
    * [Experiment](kubernetes/resources/experiment.md)
    * [Pipeline](kubernetes/resources/pipeline.md)
    * [Server](kubernetes/resources/server.md)
    * [Server Config](kubernetes/resources/serverconfig.md)
    * [Server Runtime](kubernetes/resources/seldonruntime.md)
    * [Seldon Config](kubernetes/resources/seldonconfig.md)
  * [Service Meshes](kubernetes/service-meshes/README.md)
    * [Ambassador](kubernetes/service-meshes/ambassador.md)
    * [Istio](kubernetes/service-meshes/istio.md)
    * [Traefik](kubernetes/service-meshes/traefik.md)
* [Resource allocation](resource-allocation/README.md)
  * [Example: Serving models on dedicated GPU nodes](resource-allocation/example-serving-models-on-dedicated-gpu-nodes.md)
* [Models](models/README.md)
  * [Multi-Model Serving](models/mms.md)
  * [Inference Artifacts](models/inference-artifacts.md)
  * [rClone](models/rclone.md)
  * [Parameterized Models](models/parameterized-models/README.md)
  * [Pandas Query](models/parameterized-models/pandasquery.md)
* [Metrics](metrics/README.md)
  * [Usage](metrics/usage.md)
  * [Operational](metrics/operational.md)
  * [Local Metrics](metrics/local-metrics-test.md)
* [Development](development/README.md)
  * [License](development/licenses.md)
  * [Release](development/release.md)
* [CLI](cli/README.md)
  * [Seldon](cli/seldon.md)
  * [Config](cli/seldon\_config.md)
    * [Config Activate](cli/seldon\_config\_activate.md)
    * [Config Deactivate](cli/seldon\_config\_deactivate.md)
    * [Config Add](cli/seldon\_config\_add.md)
    * [Config List](cli/seldon\_config\_list.md)
    * [Config Remove](cli/seldon\_config\_remove.md)
  * [Experiment](cli/seldon\_experiment.md)
    * [Experiment Start](cli/seldon\_experiment\_start.md)
    * [Experiment Status](cli/seldon\_experiment\_status.md)
    * [Experiment List](cli/seldon\_experiment\_list.md)
    * [Experiment Stop](cli/seldon\_experiment\_stop.md)
  * [Model](cli/seldon\_model.md)
    * [Model Status](cli/seldon\_model\_status.md)
    * [Model Load](cli/seldon\_model\_load.md)
    * [Model List](cli/seldon\_model\_list.md)
    * [Model Infer](cli/seldon\_model\_infer.md)
    * [Model Metadata](cli/seldon\_model\_metadata.md)
    * [Model Unload](cli/seldon\_model\_unload.md)
  * [Pipeline](cli/seldon\_pipeline.md)
    * [Pipeline Load](cli/seldon\_pipeline\_load.md)
    * [Pipeline Status](cli/seldon\_pipeline\_status.md)
    * [Pipeline List](cli/seldon\_pipeline\_list.md)
    * [Pipeline Inspect](cli/seldon\_pipeline\_inspect.md)
    * [Pipeline Infer](cli/seldon\_pipeline\_infer.md)
    * [Pipeline Unload](cli/seldon\_pipeline\_unload.md)
  * [Server](cli/seldon\_server.md)
    * [Server List](cli/seldon\_server\_list.md)
    * [Server Status](cli/seldon\_server\_status.md)
* [Pipelines](pipelines.md)
* [Experiments](experiments.md)
* [Servers](servers.md)
* [Inference](inference.md)
* [Outlier Detection](outlier.md)
* [Drift Detection](drift.md)
* [Explainers](explainers.md)
* [Performance Tests](performance-tests.md)
* [Upgrading](upgrading.md)
* [FAQ](faqs.md)
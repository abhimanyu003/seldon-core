# Seldon Config

{% hint style="info" %}
This section is for advanced usage where you want to define how seldon is installed in each namespace.
{% endhint %}

The SeldonConfig resource defines the core installation components installed by Seldon. If you wish to
install Seldon, you can use the [SeldonRuntime](seldonruntime.md) resource which allows easy
overriding of some parts defined in this specification. In general, we advise core DevOps to use
the default SeldonConfig or customize it for their usage. Individual installation of Seldon can
then use the SeldonRuntime with a few overrides for special customisation needed in that namespace.

The specification contains core PodSpecs for each core component and a section for general configuration
including the ConfigMaps that are created for the Agent (rclone defaults), Kafka and Tracing (open telemetry).

```go
// operator/apis/mlops/v1alpha1/seldonconfig_types.go
// :language: golang
// :start-after: // SeldonConfigSpec
// :end-before: // SeldonConfigStatus
```

Some of these values can be overridden on a per namespace basis via the SeldonRuntime resource. Labels and annotations
can also be set at the component level - these will be merged with the labels and annotations from the SeldonConfig
resource in which they are defined and added to the component's corresponding Deployment, or StatefulSet.

The default configuration is shown below.

```yaml
# operator/config/seldonconfigs/default.yaml
```
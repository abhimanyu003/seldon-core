package io.seldon.dataflow.kafka

import io.seldon.dataflow.mtls.CertificateConfig
import io.seldon.dataflow.mtls.Provider
import org.apache.kafka.clients.consumer.ConsumerConfig
import org.apache.kafka.clients.producer.ProducerConfig
import org.apache.kafka.common.config.SslConfigs
import org.apache.kafka.common.config.TopicConfig
import org.apache.kafka.common.security.auth.SecurityProtocol
import org.apache.kafka.streams.StreamsConfig
import java.util.*

data class KafkaStreamsParams(
    val bootstrapServers: String,
    val numPartitions: Int,
    val replicationFactor: Int,
    val security: KafkaSecurityParams,
)

data class KafkaSecurityParams(
    val securityProtocol: SecurityProtocol,
    val certConfig: CertificateConfig,
)

data class KafkaDomainParams(
    val useCleanState: Boolean,
    val joinWindowMillis: Long,
)

const val KAFKA_MAX_MESSAGE_BYTES = 1_000_000_000

val kafkaTopicConfig = mapOf(
    TopicConfig.MAX_MESSAGE_BYTES_CONFIG to KAFKA_MAX_MESSAGE_BYTES.toString()
)

fun getKafkaAdminProperties(params: KafkaStreamsParams): KafkaAdminProperties {
    return Properties().apply {
        this[StreamsConfig.BOOTSTRAP_SERVERS_CONFIG] = params.bootstrapServers
        this[StreamsConfig.SECURITY_PROTOCOL_CONFIG] = params.security.securityProtocol.toString()
    }
}

fun getKafkaProperties(params: KafkaStreamsParams): KafkaProperties {
    // See https://docs.confluent.io/platform/current/streams/developer-guide/config-streams.html


    return Properties().apply {
        // TODO - add version to app ID?  (From env var.)
        this[StreamsConfig.APPLICATION_ID_CONFIG] = "seldon-dataflow"
        this[StreamsConfig.BOOTSTRAP_SERVERS_CONFIG] = params.bootstrapServers
        this[StreamsConfig.PROCESSING_GUARANTEE_CONFIG] = "at_least_once"
        this[StreamsConfig.NUM_STREAM_THREADS_CONFIG] = 1

        // Security
        this[StreamsConfig.SECURITY_PROTOCOL_CONFIG] = params.security.securityProtocol.toString()
        if (params.security.securityProtocol == SecurityProtocol.SSL) {
            val keyStoreConfig = Provider.keyStoresFromCertificates(params.security.certConfig)

            this[SslConfigs.SSL_KEYSTORE_LOCATION_CONFIG] = keyStoreConfig.keyStoreLocation
            this[SslConfigs.SSL_KEYSTORE_PASSWORD_CONFIG] = keyStoreConfig.keyStorePassword
            this[SslConfigs.SSL_TRUSTSTORE_LOCATION_CONFIG] = keyStoreConfig.trustStoreLocation
            this[SslConfigs.SSL_TRUSTSTORE_PASSWORD_CONFIG] = keyStoreConfig.trustStorePassword
        }

        // Testing
        this[StreamsConfig.REPLICATION_FACTOR_CONFIG] = params.replicationFactor
        this[StreamsConfig.CACHE_MAX_BYTES_BUFFERING_CONFIG] = 0
        this[StreamsConfig.COMMIT_INTERVAL_MS_CONFIG] = 1

        this[ConsumerConfig.AUTO_OFFSET_RESET_CONFIG] = "earliest"
        this[ConsumerConfig.MAX_PARTITION_FETCH_BYTES_CONFIG] = KAFKA_MAX_MESSAGE_BYTES
        this[ConsumerConfig.FETCH_MAX_BYTES_CONFIG] = KAFKA_MAX_MESSAGE_BYTES
        this[ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG] = 60000

        this[ProducerConfig.LINGER_MS_CONFIG] = 0
        this[ProducerConfig.MAX_REQUEST_SIZE_CONFIG] = KAFKA_MAX_MESSAGE_BYTES
    }
}

fun KafkaProperties.withAppId(name: String): KafkaProperties {
    val properties = KafkaProperties()

    properties.putAll(this.toMap())
    this[StreamsConfig.APPLICATION_ID_CONFIG] = "seldon-dataflow-$name"
    // TODO add k8s host name to ensure static membership is only used for consumers from the same pod restarting?
    //
    // If set, allows static membership which would allow restarts within SESSION_TIMEOUT_MS_CONFIG
    // to happen with no rebalance
    this[ConsumerConfig.GROUP_INSTANCE_ID_CONFIG] = "seldon-dataflow-$name"

    return properties
}

fun KafkaProperties.withStreamThreads(n: Int): KafkaProperties {
    val properties = KafkaProperties()

    properties.putAll(this.toMap())
    this[StreamsConfig.NUM_STREAM_THREADS_CONFIG] = n

    return properties
}

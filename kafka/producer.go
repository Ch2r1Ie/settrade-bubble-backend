package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func NewSyncProducerGuarantee(addrs []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	producer, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		log.Panic(err)
	}
	return producer
}

func NewSyncProducerFirenForget(addrs []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	producer, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		log.Panic(err)
	}
	return producer
}

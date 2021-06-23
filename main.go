package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"studygo2/logtransfer/config"
	"studygo2/logtransfer/es"
	"studygo2/logtransfer/kafka"
)

func main(){
	//加载配置文件

	config,err:=config.ReadConfigFile()
	if err != nil {
		fmt.Println("read config fail:",err)
	}
	//fmt.Println(config)

	//初始化Es
	EsClient,err:=es.Init(config.ElesticsServerAddr)
	if err != nil {
		fmt.Println("init Es fail:",err)
	}


	//初始化kafak'
/*	kafka.Init([]string{config.KafkaServerAddr},config.Topic)

	if err != nil {
		fmt.Println("init kafka fail:",err)
	}
	partitionList,err:=consumer.Partitions("web_log")*/
	consumer,err:=sarama.NewConsumer([]string{config.KafkaServerAddr},nil)
	//fmt.Printf("main consumer is :%d",consumer)
	//从kafka取日志数据发往ES
    kafka.ComsumerKafkaSentToEs(consumer,config.Topic,EsClient)
	/*var i int=3
	fmt.Printf("%d\n",&i)
	test(i)
	test2(&i)*/
	for{

	}
}

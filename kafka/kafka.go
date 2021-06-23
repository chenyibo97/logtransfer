package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/olivere/elastic/v7"
	"studygo2/logtransfer/es"
)

type Logdata struct {
	Data string `json:"value"`
}
func ComsumerKafkaSentToEs(consumer sarama.Consumer,topic string,Client *elastic.Client)(err error){
	partitionList,err:=consumer.Partitions(topic)
	//fmt.Printf("kafka consumer in fun2 is : %d\n",consumer)
	//fmt.Println(partitionList)
	for partition:=range partitionList{
		pc,err:=consumer.ConsumePartition(topic,int32(partition),sarama.OffsetNewest)
		if err != nil {
			fmt.Println("fail to start consumer to prtition:",err)
		}
		//defer pc.AsyncClose()

		go func(pc sarama.PartitionConsumer){
			fmt.Println("a groutine has start watch")
			defer fmt.Println("a groutine exit")
			logdata:=new(Logdata)
			for msg:=range pc.Messages(){
				//json.Unmarshal(msg.Value,logdata)
				logdata.Data=string(msg.Value)
				//fmt.Println(string(msg.Value))
				es.SendToEs(Client,topic,logdata)
				fmt.Printf("partition:%d offset：%d  key:%v value:%v",msg.Partition,msg.Offset,msg.Key,string(msg.Value))
			}

			fmt.Println("groutine exie")
		}(pc)
	}
	return
}
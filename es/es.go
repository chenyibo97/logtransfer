package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"strings"
)

//初始化Es
func Init(elasticAddr string)(Client *elastic.Client,err error){
	if !strings.HasPrefix(elasticAddr,"http://"){
		elasticAddr="http://"+elasticAddr
	}
	Client,err=elastic.NewClient(elastic.SetURL(elasticAddr))
	if err != nil {
		return
	}

   return
}
//接收数据发送到kafka
func SendToEs(Client *elastic.Client,indexstr string,data interface{}){
	_,err:=Client.Index().Index(indexstr).BodyJson(data).Do(context.Background())
	if err != nil {
		panic(err)
	}
}
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	KafkaServerAddr    string `json:"kafkaServerAddr"`
	Topic              string `json:"topic"`
	ElesticsServerAddr string `json:"elesticsServerAddr"`
}
func ReadConfigFile() (*Config,error){
	config:=new(Config)
    file,err:=os.Open("./config/config.json")
	if err != nil {
		return nil,err
	}

	defer file.Close()
    decoder:=json.NewDecoder(file)
    err=decoder.Decode(config)
	if err != nil {
		fmt.Println("decode err:",err)
		return nil, err
	}
	return config,nil
}
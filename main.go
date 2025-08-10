package main

import (
	"context"
	"etcdTest/config"
	"log"
	"time"
)

func main() {
	configServer := config.NewConfigServer("etcd")
	conf, err := configServer.GetConfiguration()
	if err != nil {
		log.Println("There was an error here on getting the configuration: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	configServer.WatchConfig(ctx, conf)

	ticker := time.NewTicker(2 * time.Second)

	for range ticker.C {
		log.Println("This is my ConnectionString config: ", conf.ConnectionString)
	}
}

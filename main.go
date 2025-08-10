package main

import (
	"context"
	"etcdTest/config"
	"etcdTest/core"
	"etcdTest/database"
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

	database, err := database.NewDatabase("postgres", conf.ConnectionString)
	if err != nil {
		log.Println("Error on creating the new database client: ", err)
	}
	notifyChan := database.WatchConnectionString()
	configServer.WatchConfig(ctx, conf, notifyChan)

	userFactory := core.NewUserFactory(database)
	ticker := time.NewTicker(2 * time.Second)

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		user, err := userFactory.GetUser(ctx, 1)
		if err != nil {
			log.Println("Could not get the user: ", user)
		}

		log.Println("This is my ConnectionString config: ", conf.ConnectionString)
		log.Println("This is the user name: ", user.Name)
	}
}

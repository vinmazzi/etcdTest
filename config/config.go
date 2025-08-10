package config

import (
	"etcdTest/core"
	"log"
)

func NewConfigServer(configServerType string) core.ConfigServer {
	switch configServerType {
	case "etcd":
		var cServer *EtcdConfigServer
		var err error
		params := EtcdParams{
			Hosts: []string{"127.0.0.1:2379"},
		}
		if cServer, err = NewEtcdConfigServer(params); err != nil {
			log.Println("There is an error here on the new etcd config server: ", err)
		}
		return cServer
	}

	return nil

}

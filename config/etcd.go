package config

import (
	"context"
	"crypto/sha256"
	"errors"
	"etcdTest/core"
	"log"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

type EtcdParams struct {
	Hosts []string
}

type EtcdConfigServer struct {
	*etcd.Client
}

func NewEtcdConfigServer(params EtcdParams) (*EtcdConfigServer, error) {
	client, err := etcd.New(etcd.Config{
		Endpoints: params.Hosts,
	})

	ecs := &EtcdConfigServer{
		Client: client,
	}

	if err != nil {
		CouldNotCreateEtcdServer := errors.New("Could Not create the etcdSever")
		err = errors.Join(CouldNotCreateEtcdServer, err)
		log.Println(err)
		return nil, err
	}

	return ecs, nil
}

func (ecs *EtcdConfigServer) GetConfiguration() (*core.Config, error) {
	conf := core.Config{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := ecs.Get(ctx, "/apps/etcdTest/connectionString")
	if err != nil {
		log.Println("There is an error getting this configuration")
		return nil, err
	}

	for _, v := range resp.Kvs {
		conf.ConnectionString = string(v.Value)
		// log.Println("This is what I found on this key:", string(v.Value))
	}

	return &conf, nil
}

func Getsha256Hash(value []byte) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write(value); err != nil {
		log.Println("There is an error getting the hash: ", err)
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func (ecs *EtcdConfigServer) WatchConfig(ctx context.Context, config *core.Config) {
	var hash string
	var err error

	ticker := time.NewTicker(time.Second)
	if hash, err = Getsha256Hash([]byte(config.ConnectionString)); err != nil {
		log.Println(err)
	}

	go func() {
		for range ticker.C {
			cNew, err := ecs.GetConfiguration()
			if err != nil {
				log.Println("There is an error trying to get the configuration on watch: ", err)
			}
			newHash, err := Getsha256Hash([]byte(cNew.ConnectionString))
			if err != nil {
				log.Println("There is an error here: ", err)
			}

			if newHash != hash {
				log.Println("Changing the config object")
				*config = *cNew
				hash = newHash
			}
		}
	}()
}

package database

import (
	"etcdTest/core"
	"log"
)

func NewDatabase(databaseType string, connectionString string) (core.Database, error) {
	switch databaseType {
	case "postgres":
		database, err := NewPostgresDatabase(connectionString)
		if err != nil {
			log.Println("There is an error on creating the new databse: ", err)
			return nil, err
		}

		return database, nil
	}

	return nil, nil
}

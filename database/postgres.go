package database

import (
	"context"
	"database/sql"
	"errors"
	"etcdTest/core"
	_ "github.com/lib/pq"
	"log"
)

type PostgresConnectionParams struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Sslmode      string
}

type PostgresDatabase struct {
	*sql.DB
}

func (cli *PostgresDatabase) WatchConnectionString() chan string {
	config := make(chan string)

	go func() {
		var err error
		var nCli *PostgresDatabase
		for connString := range config {
			nCli, err = NewPostgresDatabase(connString)
			if err != nil {
				log.Println("Could not Create new Cli: ", err)
			}
			err = nCli.Ping()
			if err != nil {
				log.Println("Could not connect to the database: ", err)
			}
			*cli = *nCli
		}
	}()

	return config
}

func NewPostgresDatabase(connString string) (*PostgresDatabase, error) {
	// connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", params.User, params.Password, params.DatabaseName, params.Host, params.Sslmode)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println("Could create connection with the database: ", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Println("There was an error connecting to the database: ", err)
		return nil, err
	}

	postgresDb := &PostgresDatabase{
		DB: db,
	}

	return postgresDb, nil
}

func (cli PostgresDatabase) GetUser(ctx context.Context, id int) (*core.User, error) {
	user := &core.User{}
	query := "SELECT name FROM users where id=$1"

	err := cli.QueryRowContext(ctx, query, id).Scan(&user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Could not find user with this ID")
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

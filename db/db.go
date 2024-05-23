package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"log"
)

func ConnectDb() (*sql.DB, error) {
	var dbName string = "postgres"
	var dbUser string = "postgres"
	var dbHost string = "travel-db-dev.cpjhxwdsxtwl.sa-east-1.rds.amazonaws.com"
	var dbPort int = 5432
	var dbEndpoint string = fmt.Sprintf("%s:%d", dbHost, dbPort)
	var region string = "sa-east-1"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	authenticationToken, err := auth.BuildAuthToken(
		context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)
	if err != nil {
		panic("failed to create authentication token: " + err.Error())
	}
	authenticationToken = "postgres"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, authenticationToken, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//create table
	result, err := db.Exec(`CREATE TABLE IF NOT EXISTS users ("id" SERIAL PRIMARY KEY, "first_name" VARCHAR(50), "last_name" VARCHAR(100), "email" VARCHAR(100), "password" VARCHAR(100), "city" VARCHAR(50), "country" VARCHAR(3), "currency" VARCHAR(3), "language" VARCHAR(5), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}
	print(result.RowsAffected())
	fmt.Println("Connected to RDS")
	return db, nil
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	_ "github.com/OSBC-LLC/go-rest-template/docs"
	"github.com/OSBC-LLC/go-rest-template/internal"
	"github.com/OSBC-LLC/go-rest-template/internal/config"
	"github.com/heroku/x/hmetrics"
	_ "github.com/lib/pq"

	kit_utils "github.com/sailsforce/gomicro-kit/utils"
)

var port = "8080"

func init() {
	if err := kit_utils.InitEnv(); err != nil {
		log.Println("error loading .env: ", err)
	}
}

// @title       Orch-Rest-Template
// @version     1.0.0
// @description Micro-service written in Golang that reads data from a PostgreSQL database.

// @contact.name SuperNova
// @contact.url  https://gus.lightning.force.com/lightning/r/ADM_Scrum_Team__c/a00EE00000PCtdSYAT/view

// @host     herokuapp.com
// @BasePath /
func main() {
	go hmetrics.Report(context.Background(), hmetrics.DefaultEndpoint, nil) //nolint:errcheck // provided by documentation

	// setup config
	c := config.ServiceConfig{}
	err := c.DefaultConfig()
	if err != nil {
		log.Fatalf("error creating service config: %v\n", err)
	}

	// get router
	router := internal.Routes(c)

	if os.Getenv("MIGRATE") == "true" {
		if err = internal.MigrateTables(c); err != nil {
			log.Fatalf("error migrating tables: %v\n", err)
		}
	}

	// start service
	c.Logger.Info("micro-service running on PORT: ", os.Getenv("PORT"))
	c.Logger.Debug(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

package main

import (
	"flag"
	"log"

	"go-pre-interview/api"
	"go-pre-interview/model"
	"go-pre-interview/server"
)

var (
	dbPath     = flag.String("db", "", "path to the sqlite database (required)")
	schemaPath = flag.String("schema", "", "path to the sqlite schema sql (required)")
)

func main() {
	flag.Parse()
	if *dbPath == "" || *schemaPath == "" {
		log.Fatalf("path to database, schema file must be supplied")
	}
	datastore := model.InitDatastore(*dbPath, *schemaPath)
	api := api.Init(datastore)
	server.StartServer(api)
}

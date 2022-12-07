package main

import (
	"log"

	"github.com/jsimonetti/berkeleydb"

	"github.com/ekzjuperi/request-saver/api"
	"github.com/ekzjuperi/request-saver/configs"
	"github.com/ekzjuperi/request-saver/utils"
)

func main() {
	// get config.
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("configs.GetConfig() err: %v", err)
	}

	// get db connection.
	dbConn, err := berkeleydb.NewDB()
	if err != nil {
		log.Fatalf("berkeleydb.NewDB() err: %v", err)
	}

	err = dbConn.Open(cfg.DBPath, berkeleydb.DbHash, berkeleydb.DbCreate)
	if err != nil {
		log.Fatalf("dbConn.Open() err: %v", err)
	}
	defer dbConn.Close()

	// start API.
	go api.NewAPI(dbConn, cfg.Port).Start()

	// exit if get signal from OS.
	<-utils.GetOsSignalChan()

	log.Printf("service stop work")
}

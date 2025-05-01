package main

import (
	"github.com/TeslaMode1X/advProg2Final/review/config"
	dbInstance "github.com/TeslaMode1X/advProg2Final/review/internal/infrastructure/db"
	"github.com/TeslaMode1X/advProg2Final/review/internal/server"
	"log"
	"os"
)

func main() {
	cfg := config.InitConfig()

	db := dbInstance.NewPostgresDatabase(&cfg)

	db.Migrate()

	l := log.New(os.Stdout, "recipe-rpc", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	server.NewGrpcServer(&cfg, db, l).Start()
}

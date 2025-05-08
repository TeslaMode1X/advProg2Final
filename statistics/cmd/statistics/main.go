package main

import (
	"github.com/TeslaMode1X/advProg2Final/statistics/config"
	dbInstance "github.com/TeslaMode1X/advProg2Final/statistics/internal/infrastructure/db"
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/server"
	"log"
	"os"
)

func main() {
	cfg := config.InitConfig()

	db := dbInstance.NewPostgresDatabase(&cfg)

	db.Migrate()

	l := log.New(os.Stdout, "statistics-rpc", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	server.NewGinServer(&cfg, db, l).Start()
}

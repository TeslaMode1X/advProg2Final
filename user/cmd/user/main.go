package main

import (
	"github.com/TeslaMode1X/advProg2Final/user/config"
	dbInstance "github.com/TeslaMode1X/advProg2Final/user/internal/infrastructure/db"
	"github.com/TeslaMode1X/advProg2Final/user/internal/server"
	"log"
	"os"
)

func main() {
	cfg := config.InitConfig()

	db := dbInstance.NewPostgresDatabase(&cfg)

	db.Migrate()

	l := log.New(os.Stdout, "user-rpc ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	server.NewGrpcServer(&cfg, db, l).Start()
}

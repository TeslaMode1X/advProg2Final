package main

import (
	"github.com/TeslaMode1X/advProg2Final/recipe/config"
	dbInstance "github.com/TeslaMode1X/advProg2Final/recipe/internal/infrastructure/db"
)

func main() {
	cfg := config.InitConfig()

	db := dbInstance.NewPostgresDatabase(&cfg)

	db.Migrate()

	//l := log.New(os.Stdout, "recipe-rpc", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
}

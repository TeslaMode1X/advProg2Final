package main

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/user/config"
	dbInstance "github.com/TeslaMode1X/advProg2Final/user/internal/infrastructure/db"
)

func main() {
	cfg := config.InitConfig()

	db := dbInstance.NewPostgresDatabase(&cfg)

	db.Migrate()

	fmt.Printf("%+v\n %+v", db, cfg)
}

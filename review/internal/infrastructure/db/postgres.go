package db

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/review/config"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

type postgresDatabase struct {
	DB *gorm.DB
}

func (p *postgresDatabase) GetDB() *gorm.DB {
	return p.DB
}

func (p *postgresDatabase) Migrate() {
	if err := p.DB.Migrator().AutoMigrate(&dao.ReviewEntity{}); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}
}

func NewPostgresDatabase(conf *config.Config) interfaces.Database {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			conf.DB.Host,
			conf.DB.User,
			conf.DB.Password,
			conf.DB.DatabaseName,
			conf.DB.Port,
			conf.DB.SSLMode)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %v", err))
		}

		dbInstance = &postgresDatabase{DB: db}
	})

	return dbInstance
}

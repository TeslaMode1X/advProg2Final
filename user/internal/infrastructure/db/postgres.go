package db

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/user/config"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
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
	if err := p.GetDB().Migrator().AutoMigrate(&model.UserEntity{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	var count int64
	p.GetDB().Model(&model.UserEntity{}).Count(&count)
	if count == 0 {
		users := []model.UserEntity{
			{
				Username: "admin",
				Password: "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3", // 123
				Email:    "admin@test.com",
			},
		}

		if err := p.GetDB().Create(&users).Error; err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
		log.Println("Successfully seeded users")
	} else {
		log.Println("Database already seeded, skipping initial data insertion")
	}
}

func NewPostgresDatabase(conf *config.Config) interfaces.Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			conf.DB.Host,
			conf.DB.User,
			conf.DB.Password,
			conf.DB.DatabaseName,
			conf.DB.Port,
			conf.DB.SSLMode,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %v", err))
		}

		dbInstance = &postgresDatabase{DB: db}
	})

	return dbInstance
}

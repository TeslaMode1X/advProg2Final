package db

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/user/config"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/repository/dao"
	"github.com/gofrs/uuid"
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
	if err := p.GetDB().Migrator().AutoMigrate(&dao.UserEntity{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	userUUID, _ := uuid.NewV4()

	var count int64
	p.GetDB().Model(&dao.UserEntity{}).Count(&count)
	if count == 0 {
		users := []dao.UserEntity{
			{
				ID:       userUUID,
				Username: "admin",
				Password: "$2a$10$UzIKarUF/Xctup91UXKo.OjA6hxP3bnnE.3p6AMwe/0wMpcn1GAoG", // 12345678
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

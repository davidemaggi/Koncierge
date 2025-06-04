package db

import (
	"github.com/davidemaggi/koncierge/internal/utils"
	"github.com/davidemaggi/koncierge/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

var (
	db     *gorm.DB
	once   sync.Once
	dbFile string
)

func Init() {

	dbFile = utils.GetKonciergeDBPath()

	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Dialector{
			DriverName: "sqlite",
			DSN:        dbFile,
		}, &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatalf("failed to connect to SQLite database: %v", err)
		}
	})
}

func Migrate() error {
	if db == nil {
		log.Fatal("database not initialized. Call db.Init() first.")
	}
	err := GetDB().AutoMigrate(&models.KubeConfigEntity{})

	if err != nil {
		return err
	}

	err = GetDB().AutoMigrate(&models.ForwardEntity{})
	if err != nil {
		return err
	}

	err = GetDB().AutoMigrate(&models.AdditionalConfigEntity{})
	if err != nil {
		return err
	}

	return err
}

// GetDB returns the singleton gorm DB instance.
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("database not initialized. Call db.Init() first.")
	}
	return db
}

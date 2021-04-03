package persistence

import (
	"fmt"
	"ws-baseline-golang-v2/domain/entity"
	"ws-baseline-golang-v2/infrastructure/persistence"
	"ws-baseline-golang-v2/infrastructure/repository"
	"ws-baseline-golang-v2/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DbHelper struct {
	ProductRepository repository.ProductRepository
	db                *gorm.DB
}

func InitDbHelper() (*DbHelper, error) {

	var host = utils.GetStrEnv("DB_HOST")
	var port = utils.GetIntEnv("DB_PORT")
	var user = utils.GetStrEnv("DB_USER")
	var password = utils.GetStrEnv("DB_PASSWORD")
	var dbname = utils.GetStrEnv("DB_NAME")
	var drive = utils.GetStrEnv("DB_DRIVER")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := gorm.Open(drive, psqlInfo)

	if err != nil {
		panic(err)
	}
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(600)
	sqlDB.SetMaxOpenConns(0)
	db.LogMode(true)
	db.AutoMigrate()
	return &DbHelper{
		ProductRepository: persistence.InitProductRepositoryImpl(db),
		db:                db,
	}, nil
}

func (s *DbHelper) Close() error {
	return s.db.Close()
}

func (s *DbHelper) Automigrate() error {
	return s.db.AutoMigrate(&entity.Product{}).Error
}

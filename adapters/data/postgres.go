package data

import (
	"fmt"
	"log"
	"microservice/helpers"
	"microservice/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

//PostgresAdapter sets up an in-memory hashmap acting as a data repository
type PostgresAdapter struct {
	DB *gorm.DB
}

//NewPostgresAdapter creates a new memory adapter
func NewPostgresAdapter() (*PostgresAdapter, error) {

	adapter := &PostgresAdapter{}

	DbHost := helpers.GetEnvVar("DB_HOST")
	DbPort := helpers.GetEnvVar("DB_PORT")
	DbUser := helpers.GetEnvVar("DB_USER")
	DbName := helpers.GetEnvVar("DB_NAME")
	DbPassword := helpers.GetEnvVar("DB_PASSWORD")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	DB, err := gorm.Open("postgres", DBURL)
	if err != nil {
		return adapter, fmt.Errorf("cannot connect to postgres database: %v", err)
	}
	log.Printf("Connected to the postgres database!")

	DB.Debug().AutoMigrate(&models.BandMention{})

	adapter.DB = DB

	log.Println("PostgresAdapter created...")
	return adapter, nil
}

//Find returns all of the bands that have a name that match given param
func (pa *PostgresAdapter) Find(name string) ([]models.BandMention, error) {
	results := []models.BandMention{}

	err := pa.DB.Debug().Model(models.BandMention{}).Where("name LIKE ?", fmt.Sprintf("%%%v%%", name)).Find(&results).Error
	return results, err
}

//Increase will add one to the mentions on a BandMention model
func (pa *PostgresAdapter) Increase(bm *models.BandMention) (*models.BandMention, error) {

	bm.Mentions = bm.Mentions + 1

	err := pa.DB.Model(models.BandMention{}).Where("name = ?", bm.Name).UpdateColumn("mentions", bm.Mentions).Error

	return bm, err
}

//Create adds a new BandMention on the stored data with an initial value of 1 mention
func (pa *PostgresAdapter) Create(name string) error {

	newBm := &models.BandMention{
		Mentions: 1,
		Name:     name,
	}

	err := pa.DB.Debug().Create(&newBm).Error
	if err != nil {
		return err
	}

	return nil
}

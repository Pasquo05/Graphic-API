package funcapisort

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type money struct {
	Id        uint `gorm:"primary key;autoIncrement"`
	Currfrom  string
	Currto    string
	Value     float64
	Createdat string
}

var moneyTable []money

var db *gorm.DB

func NewConnection() {

	var err error
	dsn := "host=localhost user=postgres password=trippi2005 dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection filed")
	}

	fmt.Println(db)
	fmt.Println("CONNECTED")

	err = db.Find(&moneyTable).Error

	if err != nil {
		fmt.Println("Errore", err)
	}

}

func GetConversion(_ interface{}) (interface{}, error) {
	fmt.Println(moneyTable)
	return moneyTable, nil
}

func EmptyDecoder(r *http.Request) (interface{}, error) {
	return nil, nil
}

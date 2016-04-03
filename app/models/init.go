package models

import(
	"log"
	_ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

// TODO make all these constants apart of the environment


// Global Database Refrence
//var db *gorm.DB

func GetDatabaseConnection() gorm.DB {

	// TODO enable ssl
	// TODO enable other MYSQL string variables

	dbconn, err := gorm.Open("mysql", "root:pass123@tcp(localhost:3306)/amillerdb?charset=utf8")

	if err != nil {
		log.Panic("Cannot connect to database: ")
		log.Panic(err)
	}

	// Get database handle
	dbconn.DB()
	// TODO toggle dev/prod mode
	dbconn.LogMode(true)
	return dbconn
}

func InitDatabaseConnection() {
	log.Println("InitDatabaseConnection()")

	db := GetDatabaseConnection()

	// Ping
	err := db.DB().Ping()

	if err != nil {
		log.Panic("Cannot ping database: ")
		log.Panic(err)
	} else {
		log.Println("Able to pint the datbase!")
	}

	// Max Idle Connections
	db.DB().SetMaxIdleConns(10)

	// Max Open Connections
	db.DB().SetMaxOpenConns(100)

	log.Println("Database connection successfully created =D")

	log.Println("Creating and Migrating Tables...")
	// TODO Loop through all model objects and create them automatically.
	log.Println("Posts...")
	db.CreateTable(&Post{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Post{})

	log.Println("Tags...");
	db.CreateTable(&Tag{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Tag{})
}

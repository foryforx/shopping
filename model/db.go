package model

import (
	"database/sql"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	mgo "gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

type dbconn struct {
	DB  *gorm.DB
	SDB *sql.DB
}

var dbinst *dbconn
var once sync.Once

func GetDBInstance() *dbconn {
	once.Do(func() {
		d := initDb()
		dbinst = &dbconn{DB: d, SDB: d.DB()}
	})
	return dbinst
}

func initDb() *gorm.DB {
	// Openning file
	if _, err := os.Stat("./data.db"); os.IsNotExist(err) {
		// path/to/whatever does not exist
	}
	db, err := gorm.Open("sqlite3", "./data.db")

	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}

	// Creating the table
	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
		db.CreateTable(&Cart{})
		db.CreateTable(&Promotion{})
		db.CreateTable(&Login{})

		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Product{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Cart{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Promotion{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Login{})

	}
	return db
}

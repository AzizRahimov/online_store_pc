package db

import (
	"database/sql"
	"fmt"
	"github.com/AzizRahimov/online_store_pc/utils"
	_ "github.com/lib/pq"
	"log"
)

// StartDbConnection - connect to db
func StartDbConnection() *sql.DB {

	settingParams := utils.AppSettings.PostgresParams

	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		settingParams.Server, settingParams.Port,
		settingParams.User, settingParams.DataBase,
		settingParams.Password)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal("Couldn't connect to database", err.Error())
	}

	return db
}

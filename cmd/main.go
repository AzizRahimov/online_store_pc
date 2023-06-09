package main

import (
	"github.com/AzizRahimov/online_store_pc/db"
	"github.com/AzizRahimov/online_store_pc/pkg/routes"
	"github.com/AzizRahimov/online_store_pc/utils"
)

func main() {
	utils.ReadSettings()                    // read config from file
	dbConnection := db.StartDbConnection()  // connect to db
	routes.RunServerAndRoutes(dbConnection) // run server and routes

}

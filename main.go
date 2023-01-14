package main

import (
	"github.com/lutfiharidha/sequis-test/app"
	"github.com/lutfiharidha/sequis-test/database"
)

// @title           Sequis-Test
// @version         1.0
// @description     For a social network application, friendship management is a common feature. The application will need features like friend request, approve or reject friend request, list friend requests, list friends, block friend, common friend between user.
// @host      		localhost:8081
// @in                          header
func main() {
	db := app.NewSQL().SetupDatabaseConnection() //setup database connection

	database.Migrator(db) //migrate table
	app.Router()          //start server
}

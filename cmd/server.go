/*
Copyright © 2021 The Haul Authors

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Assembly struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Final       bool   `json:"final"`
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the haul server",
	Run: func(cmd *cobra.Command, args []string) {
		connectDB()
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	declareFlags()
}

func declareFlags() {
	serverCmd.Flags().IntP("port", "p", 8080, "The network port to bind the server to (config: \"server.port\")")
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))

	serverCmd.Flags().String("db_type", "sqlite", "The database type (config: \"db.type\")")
	viper.BindPFlag("db.type", serverCmd.Flags().Lookup("db_type"))

	serverCmd.Flags().String("db_path", "haul.db", "The database path (config: \"db.path\")")
	viper.BindPFlag("db.path", serverCmd.Flags().Lookup("db_path"))
}

func runServer() {
	fmt.Println("server called")
	fmt.Println(viper.GetInt("server.port"))
}

func connectDB() {
	var dialector gorm.Dialector

	// Check db type
	switch viper.Get("db.type") {
	case "sqlite":
		dialector = sqlite.Open(viper.GetString("db.path"))
	}

	// Connection
	db, db_err := gorm.Open(dialector, &gorm.Config{})
	if db_err != nil {
		log.Print(db_err)
		log.Fatal("[ERR] Unable to connect to database")
	}

	sqlDB, sqlDB_err := db.DB()
	if sqlDB_err != nil {
		log.Print(sqlDB_err)
		log.Print("[ERR] Invalid database connection object")
		log.Fatal("[ASSISTANCE] Verify your database connection values (flag or config)")
	}
	defer sqlDB.Close()

	// Tests
	log.Print("[I] Pinging db")
	sqlDB.Ping()
	log.Print("[I] Stats-ing db")
	sqlDB.Stats()

	// Migrate the schemas
	log.Print("[I] Migrating assets schema")
	db.AutoMigrate(&Asset{})
	log.Print("[I] Migrating assemblies schema")
	db.AutoMigrate(&Assembly{})

	// Ok
	log.Print("[OK] Database connection successful!")
}

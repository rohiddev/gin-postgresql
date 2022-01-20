package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"

	// _ "github.com/jinzhu/gorm/dialects/sqlite" // Import for side effect https://golang.org/doc/effective_go.html#blank
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	//"github.com/joho/godotenv" //using database.env

	"github.com/spf13/viper"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}




func viperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}


func SetupModels() *gorm.DB {

	viper.SetConfigName("database") // config file name without extension
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	// Set default value
	viper.SetDefault("app.linetoken", "DefaultLineTokenValue")


	viper_user := viper.GetString("POSTGRES_PASSWORD")
	log.Printf("viper : %s = %s \n", "POSTGRES_USER", viper_user)



	prosgret_conname := fmt.Sprintf("host=127.0.0.1 port=5432 user=rohiddev dbname=rohiddev password=Mad9yar24 sslmode=disable")

	fmt.Println("conname is\t\t", prosgret_conname)

	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})

	// Initialize value
	m := Book{Author: "author1", Title: "title1"}

	db.Create(&m)

	return db
}

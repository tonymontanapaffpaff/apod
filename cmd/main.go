package main

import (
	"fmt"
	"os"

	"github.com/tonymontanapaffpaff/apod/pkg/api"
	"github.com/tonymontanapaffpaff/apod/pkg/data"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	logLevel       = os.Getenv("LOG_LEVEL")
	serverEndpoint = os.Getenv("SERVER_ENDPOINT")
	dbHost         = os.Getenv("DB_HOST")
	dbPort         = os.Getenv("DB_PORT")
	dbUser         = os.Getenv("DB_USER")
	dbName         = os.Getenv("DB_NAME")
	dbPassword     = os.Getenv("DB_PASSWORD")
	sslMode        = os.Getenv("SSL_MODE")
)

func init() {
	if logLevel == "" {
		logLevel = "debug"
	}
	if serverEndpoint == "" {
		serverEndpoint = "9090"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbUser == "" {
		dbUser = "postgres"
	}
	if dbName == "" {
		dbName = "apod"
	}
	if dbPassword == "" {
		dbPassword = "root"
	}
	if sslMode == "" {
		sslMode = "disable"
	}
}

func main() {
	log.Debugf("dbName=%s", dbName)
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	db, err := getConnection(dbHost, dbPort, dbUser, dbName, dbPassword, sslMode)
	if err != nil {
		log.Fatalf("Can't connect to database, error: %v", err)
	}
	r := gin.Default()
	pictureData := data.NewPictureData(db)
	albumData := data.NewAlbumData(db)
	api.ServePictureResource(r, *pictureData, *albumData)
	api.ServeAlbumResource(r, *albumData)
	err = r.Run(":" + serverEndpoint)
	if err != nil {
		log.Fatalf("Server has been crashed, err: %v", err)
	}
}

func getConnection(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return connection, nil
}

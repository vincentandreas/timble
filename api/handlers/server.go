package handlers

import (
	"fmt"
	"time"

	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/vincentandreas/dealls/models"
	"github.com/vincentandreas/dealls/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dbSetup() (*gorm.DB) {

	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DBNAME")
	host := os.Getenv("POSTGRES_HOST")

	dbParams := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", username, password, dbName, host)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbParams,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})


	if err != nil{
		panic(err)
	}
	return conn
}

func SetupServer() (*gin.Engine, *gorm.DB) {
	
	conn := dbSetup()
	conn.AutoMigrate(
		&models.User{},
		&models.UserPreference{},
		&models.History{},
	)

	r := gin.Default()

	storeSecret := os.Getenv("SESSION_SECRET")
	store := memstore.NewStore([]byte(storeSecret)) //todo change
	store.Options(
		sessions.Options{ Path: "/", 
		MaxAge: int((time.Hour * 24).Seconds()), 
		 HttpOnly: true, })

	r.Use(sessions.Sessions("timble",store))


	repo := repo.NewImplementation(conn)
	h := NewBaseHandler(repo)
	HandleRequests(h, r)
	return r, conn
}
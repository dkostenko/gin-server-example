package api

import (
	"github.com/dkostenko/gin-server-example/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	dbpkg "github.com/dkostenko/gin-server-example/db"
)

var db gorm.DB

func New() *gin.Engine {
	db = dbpkg.DB
	router := gin.Default()

	router.GET("/users", usersList)

	return router
}

// GET /users handler
func usersList(c *gin.Context) {
	users := []models.User{}
	db.Find(&users)

	c.JSON(200, gin.H{"status": "ok", "users": users})
}

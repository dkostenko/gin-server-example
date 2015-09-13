package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var DB = make([]*user, 0)
var counter = 0

type user struct {
	Id       int
	Name     string
	Password string
}

type userUpdateParams struct {
	Name string `form:"name" json:"name"`
}

type userCreateParams struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type userResopnsePartial struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	// Routes
	router.GET("/users", usersList)
	router.GET("/users/:id", userShow)
	router.POST("/users", userCreate)
	router.PUT("/users/:id", userUpdate)

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}

// PUT /users handler
func userUpdate(c *gin.Context) {
	var params userUpdateParams

	if err := c.Bind(&params); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(404, gin.H{"status": "not found"})
		return
	}

	user := DB[id-1]

	if user == nil {
		c.JSON(404, gin.H{"status": "not found"})
		return
	}

	user.Name = params.Name

	c.JSON(201, gin.H{"status": "ok", "user": &userResopnsePartial{Id: user.Id, Name: user.Name}})
}

// GET /users handler
func usersList(c *gin.Context) {
	users := make([]*userResopnsePartial, 0)

	for _, user := range DB {
		users = append(users, &userResopnsePartial{Id: user.Id, Name: user.Name})
	}

	c.JSON(200, gin.H{"status": "ok", "users": users})
}

// GET /users/:id handler
func userShow(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(404, gin.H{"status": "not found"})
		return
	}

	user := DB[id-1]

	if user == nil {
		c.JSON(404, gin.H{"status": "not found"})
		return
	}

	c.JSON(200, gin.H{"status": "ok", "user": &userResopnsePartial{Id: user.Id, Name: user.Name}})
}

// POST /users handler
func userCreate(c *gin.Context) {
	var params userCreateParams

	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	counter++
	user := &user{Id: counter, Name: params.Name, Password: params.Password}
	DB = append(DB, user)

	c.JSON(201, gin.H{"status": "ok", "user": &userResopnsePartial{Id: user.Id, Name: user.Name}})
}

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId   int
	UserName string
	Age      int
}

var user = []User{}

func Start() {
	router := gin.Default()
	router.GET("/getUser", GetUser)
	router.Run(":8080")

}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, user)
}

func GetUserById(c *gin.Context) {
	ID := c.Param("id")
	UID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	for _, userId := range user {
		if userId.UserId == UID {
			c.JSON(http.StatusAccepted, UID)
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Not Found",
	})
}

func main() {
	Start()
}

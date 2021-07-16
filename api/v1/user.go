package v1

import "github.com/gin-gonic/gin"

func CreateToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createToken",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateUser",
	})
}

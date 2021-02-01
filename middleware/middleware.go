package middleware

import (
	"YongQing-API/database"
	"YongQing-API/errors"
	"YongQing-API/model"
	"log"

	"github.com/gin-gonic/gin"
)

type HeaderAuth struct {
	OpenID    string `json:"OpenID"`
	TimeStamp string `json:"TimeStamp"`
	Signature string `json:"Signature"`
}

func UserLogin(c *gin.Context) {
	var h HeaderAuth
	err := c.ShouldBindHeader(&h)
	errors.HandleError(err)
	log.Println(h)

	var userID int
	database.DB.Model(model.User{}).Where("open_id=?", h.OpenID).Select("id").Scan(&userID)
	if userID == 0 {
		c.AbortWithStatusJSON(200, gin.H{"status": "FAIL", "message": "LoginNeeded", "data": nil})
	}
	c.Set("userID", userID)
	c.Next()
}

package user

import (
	"api-server/handler"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&r); err != nil {
		handler.Error(c, errors.Wrap(err, "binding request body to the struct failed"))
		return
	}

	if r.Username == "" {
		handler.Error(c, errors.New("username is empty"))
		return
	}
	if r.Password == "" {
		handler.Error(c, errors.New("passwd is empty"))
		return
	}

	handler.Success(c, gin.H{"id": 100})
}

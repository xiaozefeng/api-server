package handler

import (
	"api-server/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.Result{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	err = errors.Cause(err)
	//log.Errorf("handler error: %T, %v", err, err)
	log.Error(err)
	c.JSON(http.StatusOK, model.Result{
		Code:    0,
		Message: err.Error(),
		Data:    nil,
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, fn func() (obj interface{}, err error)) {
	obj, err := fn()
	genericJSONResponse(c, obj, err, http.StatusOK, http.StatusInternalServerError)
}

func CreatedJSONResponse(c *gin.Context, err error) {
	genericJSONResponse(c, "", err, http.StatusCreated, http.StatusInternalServerError)
}

func DeletedJSONResponse(c *gin.Context, err error) {
	genericJSONResponse(c, "", err, http.StatusNoContent, http.StatusInternalServerError)
}

func genericJSONResponse(c *gin.Context, obj interface{}, err error, successCode, failureCode int) {
	code := successCode
	if err != nil {
		code = failureCode
	}

	c.JSON(code, obj)
}

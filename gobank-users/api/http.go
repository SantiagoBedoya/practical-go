package api

import (
	"net/http"

	"github.com/SantiagoBedoya/gobank-users-api/users"
	"github.com/SantiagoBedoya/gobank-utils/httperrors"
	"github.com/gin-gonic/gin"
)

// Handler define interfaces for handlers
type Handler interface {
	Create(*gin.Context)
	FindByID(*gin.Context)
	Login(*gin.Context)
}

type handler struct {
	srv users.Service
}

// NewHandler return the implementation of Handler interface
func NewHandler(srv users.Service) Handler {
	return &handler{srv: srv}
}

func (h *handler) Create(c *gin.Context) {
	var data users.User
	if err := c.ShouldBindJSON(&data); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid JSON body")
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}
	u, err := h.srv.Create(&data)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, u)
}
func (h *handler) Login(c *gin.Context) {
	var data users.User
	if err := c.ShouldBindJSON(&data); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid JSON body")
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}
	u, err := h.srv.Login(&data)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, u)
}
func (h *handler) FindByID(c *gin.Context) {
	userID := c.Param("id")
	u, err := h.srv.FindByID(userID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"local/shop/backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

const (
	apiPrefix = "/api"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/health", h.health)

	user := router.Group(makeupRoute("user"))
	{
		user.POST("", h.createUser)
	}

/*	category := router.Group(makeupRoute("category"))
	{
		category.POST("")
	}

	product := router.Group(makeupRoute("product"))
	{
		product.POST("")
	}

	order := router.Group(makeupRoute("order"))
	{
		order.POST("")
	}

	report := router.Group(makeupRoute("report"))
	{
		report.GET("")
	}*/

	return router
}

func (h *Handler) health(c *gin.Context) {
}

func makeupRoute(group string) string {
	return fmt.Sprintf("%s/%s", apiPrefix, group)
}
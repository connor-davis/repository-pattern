package api

import (
	"github.com/connor-davis/repository-pattern/internal/api/routes"
	"github.com/connor-davis/repository-pattern/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	router       fiber.Router
	todosService services.TodosService
}

func NewApi(router fiber.Router, todosService services.TodosService) *Api {
	return &Api{
		router:       router,
		todosService: todosService,
	}
}

func (api *Api) Setup() {
	r := routes.NewRoutes(api.router, api.todosService)
	r.Setup()
}

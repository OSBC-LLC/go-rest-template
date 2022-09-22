package controllers

import (
	"net/http"

	"github.com/OSBC-LLC/go-rest-template/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	kit_logger "github.com/sailsforce/gomicro-kit/logger"
)

func TodoRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", GetTodo)

	return r
}

// GetTodo godoc
// @Summary     Get list of todos
// @Description Show the full list of todos.
// @Produce     json
// @Success     200 {array} models.Todo
// @Router      /api/todo [get]
func GetTodo(rw http.ResponseWriter, r *http.Request) {
	logger := kit_logger.GetLogEntry(r)
	reqId := middleware.GetReqID(r.Context())

	logger.Info("starting todo request...")

	logger.Info(services.TodoReverseString("the quick brown fox jumped over the lazy dog"))

	logger.Info("todo request finished.")
	render.JSON(rw, r, []byte(reqId+" | Well hello there"))
}

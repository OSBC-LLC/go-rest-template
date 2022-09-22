package controllers

import (
	"net/http"
	"os"

	"github.com/OSBC-LLC/go-rest-template/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	kit_logger "github.com/sailsforce/gomicro-kit/logger"
	kit_models "github.com/sailsforce/gomicro-kit/models"
)

func HeartbeatRoutes(c config.ServiceConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", GetHeartbeat(c))

	return r
}

// GetHeartbeat godoc
// @Summary     Get health status of service
// @Description Shows if database is online, dyno metadata, and health status
// @Produce     json
// @Success     200 {object} models.Heartbeat
// @Router      /api/heartbeat [get]
func GetHeartbeat(c config.ServiceConfig) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		logger := kit_logger.GetLogEntry(r)
		reqId := middleware.GetReqID(r.Context())
		dbOnline := false

		logger.Info("starting heartbeat request...")
		logger.Debug("pinging database...", c.RV.DatabaseURL)

		if c.RV.DatabaseURL != "" {
			sqlDB, _ := c.Psql.DB()
			if err := sqlDB.Ping(); err == nil {
				dbOnline = true
			}
		}

		response := kit_models.Heartbeat{
			RequestID:      reqId,
			DatabaseOnline: dbOnline,
			AppName:        os.Getenv("HEROKU_APP_NAME"),
			ReleaseDate:    os.Getenv("HEROKU_RELEASE_CREATED_AT"),
			ReleaseVersion: os.Getenv("HEROKU_RELEASE_VERSION"),
			Slug:           os.Getenv("HEROKU_SLUG_COMMIT"),
			Message:        "heartbeat",
		}

		logger.Info("heartbeat request finished.")
		render.JSON(rw, r, response)
	}
}

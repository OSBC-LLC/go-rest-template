package internal

import (
	"net/http"
	"os"

	"github.com/OSBC-LLC/go-rest-template/internal/config"
	"github.com/OSBC-LLC/go-rest-template/internal/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"

	kit_logger "github.com/sailsforce/gomicro-kit/logger"
	kit_middleware "github.com/sailsforce/gomicro-kit/middleware"
)

func Routes(c config.ServiceConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.RequestID,
		kit_logger.NewStructuredLogger(c.Logger),
		kit_middleware.Headers,
		middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/heartbeat", kit_middleware.NewRelicWrapper(
			controllers.HeartbeatRoutes(c), c.NewRelic))
		r.Mount("/todo", kit_middleware.NewRelicWrapper(
			kit_middleware.ValidateHmac(
				controllers.TodoRoutes()), c.NewRelic))
	})

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("X-Frame-Options", "DENY")
		rw.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Logger.Debug(rw.Write([]byte("online")))
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+os.Getenv("PORT")+"/swagger/doc.json"),
	))

	return r
}

func MigrateTables(c config.ServiceConfig) error {
	return c.Psql.AutoMigrate()
}

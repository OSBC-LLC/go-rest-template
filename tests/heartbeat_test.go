package tests

import (
	"log"
	"net/http/httptest"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/OSBC-LLC/go-rest-template/internal/config"
	"github.com/OSBC-LLC/go-rest-template/internal/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kit_models "github.com/sailsforce/gomicro-kit/models"
)

var _ = Describe("Heartbeat", func() {

	c := config.ServiceConfig{}
	router := &chi.Mux{}

	BeforeEach(func() {
		// load custom env variables for tests
		os.Setenv("HEROKU_APP_NAME", "test-app-name")
		os.Setenv("HEROKU_RELEASE_VERSION", "test-version")
		os.Setenv("HEROKU_SLUG_COMMIT", "test-slug")
		os.Setenv("HEROKU_RELEASE_CREATED_AT", "test-created-at")

	})

	AfterEach(func() {
		os.Clearenv()
	})

	Describe("calling heartbeat", func() {
		Context("no database", func() {

			BeforeEach(func() {
				if err := c.DefaultConfig(); err != nil {
					log.Fatalf("error setting up test config: %v", err)
				}
				router = controllers.HeartbeatRoutes(c)
			})

			AfterEach(func() {
				c = config.ServiceConfig{}
			})

			It("should return heartbeat model", func() {
				var respBody kit_models.Heartbeat
				rr := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/", nil)
				router.ServeHTTP(rr, req)
				// parse resp body
				if err := render.DecodeJSON(rr.Body, &respBody); err != nil {
					log.Fatalf("error decoding json body: %v", err)
				}
				// validate
				Expect(rr.Code).To(Equal(200))
				Expect(respBody.DatabaseOnline).To(Equal(false))
				Expect(respBody.AppName).To(Equal("test-app-name"))
				Expect(respBody.ReleaseVersion).To(Equal("test-version"))
				Expect(respBody.Slug).To(Equal("test-slug"))
				Expect(respBody.ReleaseDate).To(Equal("test-created-at"))
				Expect(respBody.Message).To(Equal("heartbeat"))
			})
		})
		Context("with mock database", func() {

			BeforeEach(func() {
				if err := c.DefaultConfig(); err != nil {
					log.Fatalf("error setting up test config: %v", err)
				}

				db, _, err := sqlmock.New()
				if err != nil {
					log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				c.SetDatabase(db, 4)
				rv := config.RuntimeVariables{DatabaseURL: "foo"}
				c.SetRuntimeVariables(rv)

				router = controllers.HeartbeatRoutes(c)
			})

			AfterEach(func() {
				os.Clearenv()
				c = config.ServiceConfig{}
			})

			It("should return db ping", func() {
				var respBody kit_models.Heartbeat
				rr := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/", nil)
				router.ServeHTTP(rr, req)
				// parse resp body
				if err := render.DecodeJSON(rr.Body, &respBody); err != nil {
					log.Fatalf("error decoding json body: %v", err)
				}
				// validate
				Expect(rr.Code).To(Equal(200))
				Expect(respBody.DatabaseOnline).To(Equal(true))
				Expect(respBody.AppName).To(Equal("test-app-name"))
				Expect(respBody.ReleaseVersion).To(Equal("test-version"))
				Expect(respBody.Slug).To(Equal("test-slug"))
				Expect(respBody.ReleaseDate).To(Equal("test-created-at"))
				Expect(respBody.Message).To(Equal("heartbeat"))
			})
		})
	})
})

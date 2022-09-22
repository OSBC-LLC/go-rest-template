package tests

import (
	"log"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/OSBC-LLC/go-rest-template/internal/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("Service Config", func() {

	dbURL := "postgres://postgres:admin@localhost:5432/postgres"
	c := config.ServiceConfig{}

	AfterEach(func() {
		c = config.ServiceConfig{}
	})

	Context("setting runtime variables", func() {
		It("set the new rv struct", func() {
			rv := config.RuntimeVariables{
				DatabaseURL:         "foo",
				LogLevel:            "bar",
				NewRelicAppName:     "nran",
				NewRelicLicense:     "nrl",
				NewRelicDisplayName: "nrdn",
				IsTest:              "true",
			}
			c.SetRuntimeVariables(rv)
			Expect(c.RV.DatabaseURL).To(Equal("foo"))
			Expect(c.RV.LogLevel).To(Equal("bar"))
			Expect(c.RV.NewRelicAppName).To(Equal("nran"))
			Expect(c.RV.NewRelicLicense).To(Equal("nrl"))
			Expect(c.RV.NewRelicDisplayName).To(Equal("nrdn"))
			Expect(c.RV.IsTest).To(Equal("true"))
		})
	})
	Context("setting the database", func() {
		It("should set the config database", func() {
			db, _, err := sqlmock.New()
			if err != nil {
				log.Fatalf("an error happened when creating mock db: %v", err)
			}
			c.SetDatabase(db, 4)
			Expect(c.RV.DatabaseURL).To(Equal(""))
			Expect(c.Psql).ToNot(BeNil())
		})
	})
	Context("new service logger", func() {
		It("should create a logger with lvl debug", func() {
			logger := config.NewServiceLogger(config.RuntimeVariables{LogLevel: "debug"})
			Expect(c.RV.DatabaseURL).To(Equal(""))
			Expect(logger.Level.String()).To(Equal("debug"))
		})
		It("should create a logger with default lvl info", func() {
			logger := config.NewServiceLogger(config.RuntimeVariables{LogLevel: "foo"})
			Expect(c.RV.DatabaseURL).To(Equal(""))
			Expect(logger.Level.String()).To(Equal("info"))
		})
	})
	Context("new db conn", func() {
		It("should create new db conn for config", func() {
			rv := config.RuntimeVariables{DatabaseURL: dbURL, IsTest: "true"}
			db, err := config.NewDBConn(rv, logrus.New())
			Expect(db).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
	})
	Context("new db conn: unhappy path", func() {
		Describe("no is_test flag", func() {
			It("should return an error", func() {
				rv := config.RuntimeVariables{DatabaseURL: dbURL}
				db, err := config.NewDBConn(rv, logrus.New())
				Expect(db).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})
		Describe("is_test flag set to false", func() {
			It("should return an error", func() {
				rv := config.RuntimeVariables{DatabaseURL: dbURL, IsTest: "false"}
				db, err := config.NewDBConn(rv, logrus.New())
				Expect(db).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})
		Describe("is_test flag set to foo", func() {
			It("should return an error", func() {
				rv := config.RuntimeVariables{DatabaseURL: dbURL, IsTest: "foo"}
				db, err := config.NewDBConn(rv, logrus.New())
				Expect(db).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Context("default config func", func() {
		BeforeEach(func() {
			os.Setenv("DATABASE_URL", dbURL)
			os.Setenv("LOG_LEVEL", "foo")
			os.Setenv("NEW_RELIC_APP_NAME", "test app name")
			os.Setenv("NEW_RELIC_LICENSE", "20202023994485574")
			os.Setenv("NEW_RELIC_DISPLAY_NAME", "ACME")
			os.Setenv("IS_TEST", "true")
		})

		AfterEach(func() {
			os.Clearenv()
		})

		It("should return service config with defaults", func() {
			c.DefaultConfig()
			Expect(c.RV.DatabaseURL).To(Equal(dbURL))
			Expect(c.RV.LogLevel).To(Equal("foo"))
			Expect(c.RV.NewRelicAppName).To(Equal("test app name"))
			Expect(c.RV.NewRelicLicense).To(Equal("20202023994485574"))
			Expect(c.RV.NewRelicDisplayName).To(Equal("ACME"))
			Expect(c.Psql).ToNot(BeNil())
			Expect(c.Logger).ToNot(BeNil())
			Expect(c.Logger.Level.String()).To(Equal("info"))
			Expect(c.NewRelic).To(BeNil())
		})
	})
})

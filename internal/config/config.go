package config

import (
	"database/sql"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	kit_utils "github.com/sailsforce/gomicro-kit/utils"
)

type ServiceConfig struct {
	RV       RuntimeVariables
	Logger   *logrus.Logger
	Psql     *gorm.DB
	NewRelic *newrelic.Application
}

type RuntimeVariables struct {
	DatabaseURL         string
	LogLevel            string
	NewRelicAppName     string
	NewRelicLicense     string
	NewRelicDisplayName string
	IsTest              string
}

func (c *ServiceConfig) SetRuntimeVariables(rv RuntimeVariables) {
	c.RV = rv
}

func (c *ServiceConfig) SetDatabase(db *sql.DB, loglvl int) error {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(loglvl)),
	})
	if err != nil {
		return err
	}
	c.Psql = gormDB
	return nil
}

func (c *ServiceConfig) DefaultConfig() error {
	// create RuntimeVariables
	rv := RuntimeVariables{
		os.Getenv("DATABASE_URL"),
		os.Getenv("LOG_LEVEL"),
		os.Getenv("NEW_RELIC_APP_NAME"),
		os.Getenv("NEW_RELIC_LICENSE"),
		os.Getenv("NEW_RELIC_DISPLAY_NAME"),
		os.Getenv("IS_TEST"),
	}
	c.RV = rv

	// create logger
	c.Logger = NewServiceLogger(rv)

	// check if DB URL exists, then create connection
	if rv.DatabaseURL != "" {
		c.Logger.Info("setting up Gorm...")
		db, err := NewDBConn(rv, c.Logger)
		if err != nil {
			return err
		}
		c.Psql = db
	}

	// check if all Newrelic variables are set, then create connection
	if rv.NewRelicAppName != "" && rv.NewRelicLicense != "" && rv.NewRelicDisplayName != "" {
		relic, err := NewServiceNewRelicConn(rv, c.Logger)
		if err != nil {
			return err
		}
		c.NewRelic = relic
	}

	return nil
}

func NewServiceLogger(rv RuntimeVariables) *logrus.Logger {
	newLog := logrus.New()
	newLog.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	newLog.SetOutput(os.Stdout)
	logLvl, err := logrus.ParseLevel(rv.LogLevel)
	if err != nil {
		newLog.Info("using default log level")
		logLvl = 4
	}
	newLog.SetLevel(logLvl)

	return newLog
}

func NewDBConn(rv RuntimeVariables, logger *logrus.Logger) (*gorm.DB, error) {
	logger.Debug("Connecting to db: ", kit_utils.GetDSN(rv.DatabaseURL))
	db, err := ConnectToDB(rv, logger.GetLevel())
	if err != nil {
		logger.Error("error connecting to database: ", err)
		return nil, err
	}
	return db, nil
}

func NewServiceNewRelicConn(rv RuntimeVariables, logger *logrus.Logger) (*newrelic.Application, error) {
	relic, err := newrelic.NewApplication(
		newrelic.ConfigAppName(rv.NewRelicAppName),
		newrelic.ConfigLicense(rv.NewRelicLicense),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = true
			cfg.HostDisplayName = rv.NewRelicDisplayName
		},
	)
	if err != nil {
		logger.Error("error setting up new relic logs: ", err)
		return nil, err
	}
	return relic, nil
}

func ConnectToDB(rv RuntimeVariables, logLvl logrus.Level) (*gorm.DB, error) {
	database, _, _ := sqlmock.New()
	var err error
	if !(rv.IsTest == "true") {
		database, err = sql.Open("postgres", kit_utils.GetDSN(rv.DatabaseURL))
		if err != nil {
			return nil, err
		}
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLvl)),
	})

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}

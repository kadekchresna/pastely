package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kadekchresna/pastely/config"
	driver_db "github.com/kadekchresna/pastely/driver/db"
	"github.com/kadekchresna/pastely/helper/logger"
	echopprof "github.com/kadekchresna/pastely/helper/pprof"
	"github.com/kadekchresna/pastely/internal/repository"
	v1 "github.com/kadekchresna/pastely/internal/web/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

const (
	STAGING     = `stg`
	PRODUCTIOON = `prd`
)

func initConfig() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("error load ENV, %s", err.Error()))
	}
}

func init() {
	if os.Getenv("APP_ENV") != PRODUCTIOON {

		// init invoke env before everything
		cobra.OnInitialize(initConfig)

	}

	// adding command invokable
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "web",
	Short: "Running Web Service",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

type WebApp struct {
	DB *gorm.DB
}

type Handlers struct {
}

func run() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := config.InitConfig()
	app := WebInit(config)
	handlers := WebV1Dependencies(app)

	v1.InitAPI(e, v1.Handlers(handlers))

	e.Any("/metrics", echo.WrapHandler(promhttp.Handler()))

	echopprof.RegisterPprofRoutes(e)
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", config.AppPort),
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	logger.Log().Info(fmt.Sprintf("%s service started...", config.AppName))
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	logger.Log().Info(fmt.Sprintf("%s service finished", config.AppName))
}

func WebInit(config config.Config) WebApp {
	db := driver_db.InitDB(config)

	return WebApp{
		DB: db,
	}
}

func WebV1Dependencies(app WebApp) v1.Handlers {
	_ = repository.NewPasteRepo(app.DB)

	return v1.Handlers{}
}

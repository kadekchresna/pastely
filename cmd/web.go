package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kadekchresna/pastely/config"
	cfg "github.com/kadekchresna/pastely/config"
	"github.com/kadekchresna/pastely/driver/cache"
	driver_db "github.com/kadekchresna/pastely/driver/db"
	"github.com/kadekchresna/pastely/helper/logger"
	echopprof "github.com/kadekchresna/pastely/helper/pprof"
	"github.com/kadekchresna/pastely/helper/transaction"

	v1_log_repo "github.com/kadekchresna/pastely/internal/v1/repository/log"
	v1_log_usecase "github.com/kadekchresna/pastely/internal/v1/usecase/log"
	v1_log_web "github.com/kadekchresna/pastely/internal/v1/web/log"

	v1_paste_repo "github.com/kadekchresna/pastely/internal/v1/repository/paste"
	v1_paste_usecase "github.com/kadekchresna/pastely/internal/v1/usecase/paste"
	v1_web "github.com/kadekchresna/pastely/internal/v1/web"
	v1_paste_web "github.com/kadekchresna/pastely/internal/v1/web/paste"

	v2_log_repo "github.com/kadekchresna/pastely/internal/v2/repository/log"

	v2_paste_repo "github.com/kadekchresna/pastely/internal/v2/repository/paste"
	v2_paste_usecase "github.com/kadekchresna/pastely/internal/v2/usecase/paste"
	v2_web "github.com/kadekchresna/pastely/internal/v2/web"
	v2_paste_web "github.com/kadekchresna/pastely/internal/v2/web/paste"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

func initConfig() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("error load ENV, %s", err.Error()))
	}
}

func init() {
	if os.Getenv("APP_ENV") != config.PRODUCTION {

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
	DB     config.DB
	Config cfg.Config
	Cache  cache.Cache
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
	v1Handlers := WebV1Dependencies(app)

	v1_web.InitAPI(e, v1Handlers)

	v2Handlers := WebV2Dependencies(app)

	v2_web.InitAPI(e, v2Handlers)

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

func WebInit(config cfg.Config) WebApp {
	masterDB := driver_db.InitDB(config.DatabaseMasterDSN)
	slaveDB := driver_db.InitDB(config.DatabaseSlaveDSN)
	analyticDB := driver_db.InitDB(config.DatabaseAnalyticDSN)

	cache := cache.InitCache(config)

	return WebApp{
		DB: cfg.DB{
			MasterDB:   masterDB,
			SlaveDB:    slaveDB,
			AnalyticDB: analyticDB,
		},
		Cache:  cache,
		Config: config,
	}
}

func WebV1Dependencies(app WebApp) v1_web.Handlers {
	pasteRepo := v1_paste_repo.NewPasteRepo(app.DB)
	transactionRepo := transaction.NewTransactionRepo(app.DB)
	logRepo := v1_log_repo.NewLogRepo(app.DB)

	pasteUsecases := v1_paste_usecase.NewPasteUsecase(app.Config, pasteRepo, transactionRepo, logRepo)
	logUsecase := v1_log_usecase.NewLogUsecase(app.Config, logRepo)

	pasteHandler := v1_paste_web.NewPasteHandler(pasteUsecases)
	logHandler := v1_log_web.NewLogHandler(logUsecase)

	return v1_web.Handlers{
		Paste: pasteHandler,
		Log:   logHandler,
	}
}

func WebV2Dependencies(app WebApp) v2_web.Handlers {
	pasteRepo := v2_paste_repo.NewPasteRepo(app.DB)
	transactionRepo := transaction.NewTransactionRepo(app.DB)
	logRepo := v2_log_repo.NewLogRepo(app.DB)
	pasteUsecases := v2_paste_usecase.NewPasteUsecase(app.Config, pasteRepo, transactionRepo, logRepo, app.Cache)
	pasteHandler := v2_paste_web.NewPasteHandler(pasteUsecases)

	return v2_web.Handlers{
		Paste: pasteHandler,
	}
}

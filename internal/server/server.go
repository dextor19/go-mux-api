package server

import (
	"fmt"
	"muxtemp/internal/db"
	"muxtemp/internal/product"
	"muxtemp/pkg/config"
	"muxtemp/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Start() {
	log.New()
	config.LoadAppConfig()
	db.Connect(config.AppConfig.ConnectionString)
	db.Migrate()
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()
	product.RegisterRoutes(api)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router)
		if err != nil {
			log.Logger.Fatal("failed to serve http server")
		}
	}()

	log.Logger.Info("Starting Server", zap.String("port", config.AppConfig.Port))
	<-done

	log.Logger.Info("Stopped Server")
}

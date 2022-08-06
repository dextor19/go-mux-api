package main

import (
	"fmt"
	"muxtemp/internal/db"
	"muxtemp/internal/product"
	"muxtemp/pkg/config"
	"muxtemp/pkg/log"
	"net/http"
	"go.uber.org/zap"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	log.New()
	config.LoadAppConfig()
	db.Connect(config.AppConfig.ConnectionString)
	db.Migrate()
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()
	product.RegisterRoutes(api)
	log.Logger.Info("Starting Server", zap.String("port", config.AppConfig.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router)
	if err != nil {
		log.Logger.Fatal("failed to serve http server")
	}

}


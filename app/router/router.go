package router

import (
	"time"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	// "github.com/spf13/viper"
	"github.com/corentindeboisset/golang-api/app/controller"
)


func GetRouter() (*chi.Mux, error) {
	r := chi.NewRouter()

	// Use some middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	r.Use(middleware.Timeout(60 * time.Second))

	// if viper.GetBool("profiler.enable") {
	// 	r.Mount(viper.GetString(), middleware.Profiler)
	// }

	userController, err := controller.InitializeUserController()
	if err != nil {
		return nil, err
	}
	settingController, err := controller.InitializeSettingController()
	if err != nil {
		return nil, err
	}

	r.Get("/", userController.UserPage)
	r.Get("/setting", settingController.SettingPage)

	return r, nil
}

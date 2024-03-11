package main

import (
	"log"
	"net/http"
	"svar-widgets/grid-backend-go/data"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jinzhu/configor"
	"github.com/unrolled/render"
)

var format = render.New()

// Config is the structure that stores the settings for this backend app
var Config AppConfig

func main() {
	configor.New(&configor.Config{ENVPrefix: "APP", Silent: true}).Load(&Config, "config.yml")

	dao := data.NewDAO(Config.DB, Config.Server.URL, Config.BinaryData)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	if len(Config.Server.Cors) > 0 {
		c := cors.New(cors.Options{
			AllowedOrigins:   Config.Server.Cors,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Remote-Token"},
			AllowCredentials: true,
			MaxAge:           300,
		})
		r.Use(c.Handler)
	}

	initRoutes(r, dao)

	log.Printf("Starting webserver at port " + Config.Server.Port)
	err := http.ListenAndServe(Config.Server.Port, r)
	if err != nil {
		log.Println(err.Error())
	}
}

package main

import (
	"log"

	"github.com/ghaidafasya24/ws-ghaida2024/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/ghaidafasya24/ws-ghaida2024/url"


	_ "github.com/ghaidafasya24/ws-ghaida2024/docs"
	"github.com/gofiber/fiber/v2"
	
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/ghaidafasya24
// @contact.email 714220031@std.ulbi.ac.id

// @host ws-ghaida2024-10da574be531.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}

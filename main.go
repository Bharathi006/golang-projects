package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
// Initialize default config

app.Use(cors.New())
// Or extend your config for customization
app.Use(cors.New(cors.Config{
    AllowOrigins: "https://gofiber.io, https://gofiber.net",
    AllowHeaders:  "Origin, Content-Type, Accept",
}))

app.Use(cors.New())

app.Use(cors.New(cors.Config{
    AllowOriginsFunc: func(origin string) bool {
        return os.Getenv("ENVIRONMENT") == "development"
    },
}))
// Config defines the config for middleware.
type Config struct {
    
    Next func(c *fiber.Ctx) bool

    AllowOriginsFunc func(origin string) bool

   
    AllowOrigins string

    
    AllowMethods string

    
    AllowHeaders string

    AllowCredentials bool

    
    ExposeHeaders string

    
    MaxAge int
}
var ConfigDefault = Config{
    Next:         nil,
    AllowOriginsFunc: nil,
    AllowOrigins: "*",
    AllowMethods: strings.Join([]string{
        fiber.MethodGet,
        fiber.MethodPost,
        fiber.MethodHead,
        fiber.MethodPut,
        fiber.MethodDelete,
        fiber.MethodPatch,
    }, ","),
    AllowHeaders:     "",
    AllowCredentials: false,
    ExposeHeaders:    "",
    MaxAge:           0,
}
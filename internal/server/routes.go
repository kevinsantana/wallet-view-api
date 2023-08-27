package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/kevinsantana/wallet-view/internal/rest"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc fiber.Handler
	Public      bool
	Scopes      []string
}

var healthCheck = Routes{
	{
		Name:        "Healthz",
		Method:      http.MethodGet,
		Pattern:     "/healthz",
		HandlerFunc: rest.Liveness,
		Public:      true,
	},
	{
		Name:        "Readiness",
		Method:      http.MethodGet,
		Pattern:     "/readiness",
		HandlerFunc: rest.Readiness,
		Public:      true,
	},
}

func Router() *fiber.App {
	r := fiber.New(fiber.Config{
		Prefork:               false,
		CaseSensitive:         false,
		StrictRouting:         false,
		ServerHeader:          "*",
		AppName:               "Wallet View Api",
		Immutable:             false,
		DisableStartupMessage: true,
	})

	api := r.Group("/")
	for _, route := range healthCheck {
		api.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	return r
}
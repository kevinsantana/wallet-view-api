package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/goccy/go-json"

	"github.com/kevinsantana/wallet-view-api/internal/rest"
	"github.com/kevinsantana/wallet-view-api/internal/rest/middlewares"
	"github.com/kevinsantana/wallet-view-api/internal/rest/handlers"
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

var walletBalance = Routes{
	{
		Name:        "WalletBalance",
		Method:      http.MethodGet,
		Pattern:     "/walletBalance/:address/:currency",
		HandlerFunc: handlers.GetWalletBalance,
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
		Immutable:             true,
		DisableStartupMessage: true,
		ErrorHandler:          middlewares.ErrorHandler(),
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})

	api := r.Group("/")
	for _, route := range healthCheck {
		api.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	v1 := api.Group("/api/v1")

	var routes []Route
	routes = append(routes, walletBalance...)

	for _, route := range routes {
		v1.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	r.Use(middlewares.RouteNotFound())

	return r
}

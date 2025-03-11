package app

import "github.com/gin-gonic/gin"

type Route func(app *App)

func RouteLogin(endpoint string) Route {
	return func(app *App) {
		app.gin.GET(endpoint, func(c *gin.Context) {

		})
	}
}

func ProfileRoute(endpoint string) Route {
	return func(app *App) {
		app.gin.GET(endpoint, func(c *gin.Context) {

		})
	}
}

package app

import (
	"github.com/ShudderStorm/go-github-tracker/internal/cache"
	"github.com/ShudderStorm/go-github-tracker/pkg/oauth"
	"github.com/gin-gonic/gin"
)

type App struct {
	host        string
	oauth       *oauth.Client
	accessCache *cache.Cache[cache.Access]
	stateCache  *cache.Cache[cache.State]
	gin         *gin.Engine
}

type CacheInitializer interface {
	InitAccessCache() *cache.Cache[cache.Access]
	InitStateCache() *cache.Cache[cache.State]
}

type OAuthInitializer interface {
	InitOAuth(opts ...oauth.Option) *oauth.Client
}

type Option func(*App)

func WithOAuth(init OAuthInitializer) Option {
	return func(app *App) {
		app.oauth = init.InitOAuth()
	}
}

func WithCache(init CacheInitializer) Option {
	return func(app *App) {
		app.accessCache = init.InitAccessCache()
		app.stateCache = init.InitStateCache()
	}
}

func WithRoutes(routes ...Route) Option {
	return func(app *App) {
		for _, route := range routes {
			route(app)
		}
	}
}

func New(opts ...Option) *App {
	app := &App{}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

func (app *App) Run() {
	app.gin.Run()
}

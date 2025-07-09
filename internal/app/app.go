package app

import (
	"context"
	"github.com/biryanim/denet_tz/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func (a *App) Run() error {

	err := a.runHTTPServer()
	if err != nil {
		log.Fatalf("failed to run http server: %v", err)
	}

	return nil
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("local.env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	router := gin.Default()

	public := router.Group("/")
	{
		public.POST("/register", a.serviceProvider.AuthImpl(ctx).Register)
		public.POST("/login", a.serviceProvider.AuthImpl(ctx).Login)
	}

	protected := router.Group("/")
	protected.Use(a.serviceProvider.AuthImpl(ctx).AuthMiddleware())
	{
		protected.GET("/users/:id/status", a.serviceProvider.UserImpl(ctx).GetStatus)
		protected.GET("/users/leaderboard", a.serviceProvider.UserImpl(ctx).GetLeaderboard)
		protected.POST("/users/:id/task/complete", a.serviceProvider.UserImpl(ctx).CompleteTask)
		protected.POST("/users/:id/referrer", a.serviceProvider.UserImpl(ctx).AddReferrer)
	}

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: router,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())
	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

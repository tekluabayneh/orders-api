package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()

	go func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Errorf("failed to close redis: %w", err)
		}
	}()

	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	fmt.Println("starting server")

	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to run the servser %w\n", err)
		}
		close(ch)
	}()
	select {
	case err := <-ch:
		return err
	case <-ch:
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

	return nil
}

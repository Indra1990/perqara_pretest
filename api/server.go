package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	_ "perqara_testing/docs"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/uptrace/bun"
)

type Server struct {
	db   *bun.DB
	http *http.Server
}

func NewServer(db *bun.DB) *Server {
	router := Routes(db)
	server := &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      20 * time.Second,
		Addr:              fmt.Sprintf("%v:%v", viper.GetString("httpserver.host"), viper.GetString("httpserver.port")),
		Handler:           router,
	}

	server.SetKeepAlivesEnabled(true)

	srv := &Server{
		db:   db,
		http: server,
	}

	return srv
}

func Routes(db *bun.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
	))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome perqara testing"))
	})

	router.Get("/test", deliveryVendingMachine(db).Test)
	router.Get("/items", deliveryVendingMachine(db).ListItems)
	router.Post("/item/create", deliveryVendingMachine(db).CreateItem)
	router.Get("/item/{itemId}", deliveryVendingMachine(db).DetailItem)
	router.Put("/item/{itemId}", deliveryVendingMachine(db).UpdateItem)
	router.Delete("/item/{itemId}", deliveryVendingMachine(db).DeleteItem)

	return router
}

func (s *Server) Start(ctx context.Context) (err error) {
	fmt.Println("App: Starting ...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err = s.http.ListenAndServe(); err != nil {
			return
		}
	}()
	log.Println("server run localhost:8081")

	<-ctx.Done()

	return
}

func (s *Server) Stop(ctx context.Context) (err error) {
	fmt.Println("App: Stopping...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err = s.http.Shutdown(ctx); err != nil {
			return
		}
	}()
	wg.Wait()

	return
}

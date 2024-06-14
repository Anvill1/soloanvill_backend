package app

import (
	"context"
	"hello/api"
	"hello/internals/app/db"
	"hello/internals/app/handlers"
	"hello/internals/app/processors"
	"hello/internals/cfg"
	"hello/middleware"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
	db     *pgx.Conn
}

func NewServer(config cfg.Cfg, ctx context.Context) *Server {
	server := new(Server)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *Server) Serve() {
	log.Println("Starting server...")
	var err error
	server.db, err = pgx.Connect(server.ctx, server.config.GetDBString())
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	go func(dbStorage *db.DBStorage) { // Вызывает фукнцию проверки и инициализации БД
		if !dbStorage.CheckDB(server.config.DefaultTable) {
			if err != nil {
				log.Fatalln(err)
			}
			log.WithFields(log.Fields{
				"migrations": "true",
			}).Info("Run init.sql script")
			dbStorage.InitDB("migrations/init.sql")
		}
	}(db.NewDBStorage(server.db))

	dbStorage := db.NewDeployStorage(server.db)
	userProcessor := processors.NewDeployProccessor(dbStorage, &server.config)
	deployHandler := handlers.NewDeployHandler(userProcessor)

	routes := api.CreateRoutes(deployHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Shutdown() {
	log.Printf("Server stopped")

	cxtShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()
	var err error
	if err = server.srv.Shutdown(cxtShutDown); err != nil {
		log.Fatalf("server Shutdown failed:#{err}")
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}

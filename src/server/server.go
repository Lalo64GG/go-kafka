package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/go-kafka/src/routes"
	"github.com/lalo64/go-kafka/src/service"
	"github.com/lalo64/go-kafka/src/config"

	"log"
)

type Server struct {
	engine *gin.Engine
	host string
	port string
	httpAddr string
}



func NewServer(host, port string) Server {
	srv := Server{
		engine: gin.Default(),
		host: host,
		port: port,
		httpAddr: host + ":" + port,
	}

	db, _ := config.InitDB()

	// 

	srv.engine.GET("/ping", func (ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	temperatureService := service.NewTemperatureService(db)

	routes.TemperatureRoutes(srv.engine.Group("/api"), temperatureService)
	

	return srv
}


func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nfsarch33/go-chatgpt-prompt-splitter/pkg/Internal/handlers"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Initialize logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	r := gin.Default()

	// Load templates relative to /app directory
	r.LoadHTMLGlob("./static/templates/*")
	// Serve static files
	r.Static("./static", "./static")

	// Routes
	// GET /
	r.GET("/", handlers.GetIndex)
	// POST /
	r.POST("/", handlers.PostPrompt)

	// listens and serves on 0.0.0.0:8080 by default, if PORT not specified in .env file
	err := r.Run()
	if err != nil {
		log.Fatal("Error running server: ", err)
		return
	}
}

package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nfsarch33/go-chatgpt-prompt-splitter/pkg/Internal/handlers"
	"github.com/nfsarch33/go-chatgpt-prompt-splitter/pkg/network"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	// Get port from environment variables
	port := os.Getenv("PORT")

	r := gin.Default()

	// Add middleware for error recovery
	r.Use(gin.Recovery())

	// Load templates relative to /app directory
	r.LoadHTMLGlob("./static/templates/*")
	// Serve static files
	r.Static("./static", "./static")

	// Routes
	// GET /
	r.GET("/", handlers.GetHome)
	// POST /
	r.POST("/", handlers.PostPrompt)

	addr := fmt.Sprintf(":%s", port)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Channel to signal when the server should shutdown
	quit := make(chan os.Signal, 2)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Info("Starting server...")
		if e := srv.ListenAndServe(); e != nil && e != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", e)
		}
	}()

	network.WaitForServer(fmt.Sprintf("http://localhost:%s", port))
	log.Info("Server is ready to handle requests")

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // disable headless mode
		// Add more options here if needed
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	log.Info("Opening browser")
	err = chromedp.Run(
		ctx,
		chromedp.Navigate(fmt.Sprintf("http://localhost:%s", port)))
	if err != nil {
		log.Panic(err)
	}

	// Wait for the context to be done, which happens when the browser is closed
	log.Info("Waiting for browser to close")
	<-ctx.Done()
	log.Info("Browser closed")
	// Send a signal to shutdown the server
	log.Info("Signaling server shutdown")
	quit <- syscall.SIGTERM

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Waiting for server to shutdown")

	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Panic("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

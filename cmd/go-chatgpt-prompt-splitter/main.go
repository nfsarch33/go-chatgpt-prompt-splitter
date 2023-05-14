package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/nfsarch33/go-chatgpt-prompt-splitter/pkg/promptsplitter"
	log "github.com/sirupsen/logrus"
)

var rdb *redis.Client

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	opt, err := redis.ParseURL(os.Getenv("UPSTASH_REDIS_URL"))
	if err != nil {
		panic(err)
	}

	rdb = redis.NewClient(opt)
}

func main() {
	r := gin.Default()

	// Load templates relative to /app directory
	r.LoadHTMLGlob("./static/*")
	// Serve static files
	r.Static("./static", "./static")

	r.GET("/", func(c *gin.Context) {
		counter, err := rdb.Incr(c, "visit_counter").Result()
		if err != nil {
			log.Error("Error incrementing visit_counter: ", err)
		}
		hash, err := promptsplitter.GenerateRandomString(8)
		if err != nil {
			log.Error("Error generating random string: ", err)
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": err,
			})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"visit_count": counter,
			"hash":        hash,
		})
	})

	r.POST("/", func(c *gin.Context) {
		prompt := c.PostForm("prompt")
		splitLength, _ := strconv.Atoi(c.PostForm("split_length"))
		fileData := promptsplitter.SplitPrompt(prompt, splitLength)
		hash, err := promptsplitter.GenerateRandomString(8)
		if err != nil {
			log.Error("Error generating random string: ", err)
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": err,
			})
			return
		}
		counter, err := rdb.Incr(c, "visit_counter").Result()
		if err != nil {
			log.Error("Error incrementing visit_counter: ", err)
		}

		ginH := gin.H{
			"Prompt":      prompt,
			"SplitLength": splitLength,
			"FileData":    fileData,
			"Hash":        hash,
			"VisitCount":  counter,
		}

		log.Info("Split prompt hash: ", hash)

		c.HTML(http.StatusOK, "index.html", ginH)
	})

	// listens and serves on 0.0.0.0:8080 by default, if PORT not specified in .env file
	err := r.Run()
	if err != nil {
		log.Fatal("Error running server: ", err)
		return
	}
}

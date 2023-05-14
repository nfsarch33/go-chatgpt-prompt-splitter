package handlers

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"github.com/nfsarch33/go-chatgpt-prompt-splitter/pkg/promptsplitter"
	log "github.com/sirupsen/logrus"
)

var rdb *redis.Client

func GetIndex(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	opt, err := redis.ParseURL(os.Getenv("UPSTASH_REDIS_URL"))
	if err != nil {
		panic(err)
	}

	rdb = redis.NewClient(opt)

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
}

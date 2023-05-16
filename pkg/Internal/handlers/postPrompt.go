package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nfsarch33/go-chatgpt-prompt-splitter/pkg/promptsplitter"
	log "github.com/sirupsen/logrus"
)

func PostPrompt(c *gin.Context) {
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

	ginH := gin.H{
		"Prompt":        prompt,
		"SplitLength":   splitLength,
		"FileDataSlice": fileData,
		"Hash":          hash,
	}

	log.Info(
		"POST / serving prompt: ",
		"Request: Split length: ",
		splitLength,
		" Prompt: ",
		prompt,
		" Number of splits: ",
		len(fileData),
		" Total length: ",
		len(prompt),
	)

	c.HTML(http.StatusOK, "index.html", ginH)
}

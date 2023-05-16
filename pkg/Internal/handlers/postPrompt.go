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

	promptLength := len(prompt)
	if promptLength == 0 {
		log.Error("Prompt length is 0")
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Prompt length is 0",
		})
		return
	}

	if promptLength < splitLength {
		log.Error("Prompt length is less than split length")
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Prompt length is less than split length",
		})
		return
	}

	numberOfSplits := len(fileData)

	log.Info(
		"POST / serving prompt: ",
		"Request: Split length: ",
		splitLength,
		" Prompt: ",
		prompt[0:25],
		" Number of splits: ",
		numberOfSplits,
		" Total length: ",
		promptLength,
	)

	ginH := gin.H{
		"Hash":           hash,
		"Prompt":         prompt,
		"FileDataSlice":  fileData,
		"SplitLength":    splitLength,
		"NumberOfSplits": numberOfSplits,
		"PromptLength":   promptLength,
	}

	c.HTML(http.StatusOK, "index.html", ginH)
}

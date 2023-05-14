package promptsplitter

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	mrand "math/rand"
	"os"

	log "github.com/sirupsen/logrus"
)

var runFullStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var myRand *mrand.Rand // Declare the Rand instance

type FileData struct {
	Name       string
	Content    string
	PartNumber int
	TotalParts int
	PartHash   string // Added PartHash
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	seed := int64(binary.LittleEndian.Uint64(b[:]))
	src := mrand.NewSource(seed)
	myRand = mrand.New(src) // Initialize the Rand instance
}

func SplitPrompt(text string, splitLength int) []FileData {
	if splitLength <= 0 {
		log.Panic("Max length must be greater than 0.")
	}

	textLen := len(text)
	numParts := (textLen + splitLength - 1) / splitLength
	fileData := make([]FileData, 0)

	for i := 0; i < numParts; i++ {
		start := i * splitLength
		end := min((i+1)*splitLength, textLen)
		content := ""
		if i == numParts-1 {
			content = fmt.Sprintf(
				"[START PART %d/%d]\n%s\n[END PART %d/%d]\nALL PARTS SENT. "+
					"Now you can continue processing the request.",
				i+1,
				numParts,
				text[start:end],
				i+1,
				numParts,
			)
		} else {
			content = fmt.Sprintf("Do not answer yet. This is just another part of the text I want to send you. "+
				"Just receive and acknowledge as \"Part %d/%d received\" and wait for the next part."+
				"\n[START PART %d/%d]\n%s\n[END PART %d/%d]\nRemember not answering yet. "+
				"Just acknowledge you received this part with the message \"Part %d/%d received\" and wait for the"+
				" next part.", i+1, numParts, i+1, numParts, text[start:end], i+1, numParts, i+1, numParts)
		}

		// Generate a random UUID for each part and assign it to PartHash
		newRandom, err := uuid.NewRandom()
		if err != nil {
			log.Fatal(err)
		}

		fileData = append(fileData, FileData{
			Name:       fmt.Sprintf("split_%03d_of_%03d.txt", i+1, numParts),
			Content:    content,
			PartNumber: i + 1,    // Added PartNumber
			TotalParts: numParts, // Added TotalParts
			PartHash:   newRandom.String(),
		})
	}

	log.Info("SplitPrompt completed successfully")

	return fileData
}

func GenerateRandomString(length int) (string, error) {
	if length < 0 {
		log.Error("invalid length: ", length)
		return "", fmt.Errorf("invalid length: %d", length)
	}

	lettersAndDigits := []rune(runFullStr)
	s := make([]rune, length)
	for i := range s {
		s[i] = lettersAndDigits[myRand.Intn(len(lettersAndDigits))] // Use myRand instead of mrand
	}

	log.Info("GenerateRandomString completed successfully")

	return string(s), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

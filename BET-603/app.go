package main

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type WordData struct {
    Mutex            sync.Mutex
    SearchCount    int
    TotalFrequency int
    DocFrequency   int
	LastTF           int
	LastDF           int
}

type Words struct {
	Content []string `json:"content"`
}
var (
	mu sync.Mutex
	wordDataMap = make(map[string]*WordData)
)


func main() {
    r := gin.Default()

    r.POST("/search", func(c *gin.Context) {
        var words Words 
        if err := c.BindJSON(&words); err != nil {
            c.JSON(400, gin.H{"error": err})
            return
        }

        var wg sync.WaitGroup

        for _, word := range words.Content {
            wg.Add(1)
            go func(w string) {
                defer wg.Done()
                searchWord(w)
            }(word)
        }

        wg.Wait()

        c.JSON(200, wordDataMap)
    })

    r.Run(":8090")
}


func searchWord(word string) (*WordData, error) {
	mu.Lock()
	defer mu.Unlock()

	
	data, exists := wordDataMap[word]
	if !exists {
		data = &WordData{}
		wordDataMap[word] = data
	}

	
	files := []string{"./test/example1.txt", "./test/example2.txt", "./test/example3.txt"}
	tf := 0
	df := 0

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Error reading file %s: %v", file, err)
			continue
		}
		// Count the occurrences of the word in the file.
		tf += strings.Count(string(content), word)
		// Increment DF if the word is found at least once in the file.
		if strings.Contains(string(content), word) {
			df++
		}
	}

	// Update WordData.
	data.Mutex.Lock()
	data.SearchCount++
	data.LastTF = tf
	data.LastDF = df
	data.TotalFrequency += tf
	data.DocFrequency += df
	data.Mutex.Unlock()

	return data, nil
}
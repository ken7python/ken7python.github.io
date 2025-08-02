package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const githubAPI = "https://api.github.com/graphql"

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func fetchContributions(login string, from, to string) (map[string]interface{}, error) {
	token := os.Getenv("GITHUB_TOKEN")

	payload := GraphQLRequest{
		Query: contributionQuery,
		Variables: map[string]interface{}{
			"login": login,
			"from":  from,
			"to":    to,
		},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", githubAPI, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &result)

	return result, nil
}

func main() {
	// .env読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ .envファイルが読み込めませんでした")
	}

	r := gin.Default()

	r.GET("/grass", func(c *gin.Context) {
		user := "ken7python"

		to := time.Now()
		from := to.AddDate(-1, 0, 0)

		data, err := fetchContributions(
			user,
			from.Format(time.RFC3339),
			to.Format(time.RFC3339),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	log.Println("🚀 サーバー起動中: http://localhost:8000")
	r.Run(":8000")
}

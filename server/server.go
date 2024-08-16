package main

import (
	"bbb/models"
	"bbb/rabbitmq"
	"bbb/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	go rabbitmq.Start()

	repository.Migrate()

	router := gin.Default()
	router.GET("/health_check", health)
	router.POST("/api/votos", postVote)

	router.Run("localhost:8080")
}

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func postVote(c *gin.Context) {
	var vote models.Vote

	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rabbitmq.Call(vote.Candidate)

	fmt.Printf("Received vote: %+v\n", vote)

	c.IndentedJSON(http.StatusAccepted, nil)
}

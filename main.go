package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/qotd/random", getRandomQuestion)
	r.Run("0.0.0.0:8080")
}

type question struct {
	ID       string `JSON : "id"`
	Question string `JSON : "question"`
	Author   string `JSON : "author"`
}

func getRandomQuestion(c gin.Context) {

}

var questions = map[string]question{
	"b2cb3984-2b99-48b9-a92f-0241d2e1b992": {"b2cb3984-2b99-48b9-a92f-0241d2e1b992", "What's your favorite color?", "Zack"},
	"06ffc204-ca87-4ae5-a43e-58c73158ae7a": {"06ffc204-ca87-4ae5-a43e-58c73158ae7a", "What's your favorite food?", "Shrimp"},
	"46388d63-386f-4e6e-842b-c1501cee4087": {"46388d63-386f-4e6e-842b-c1501cee4087", "What's your favorite movie?", "Lobster"},
}

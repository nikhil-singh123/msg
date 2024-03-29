package main

import (
	"msg/handler"

	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	handler.Setup()

	router := gin.Default()
	//router.LoadHTMLGlob("templates/**/*")
	router.GET("/", homePage)

	router.GET("/sign-up", handler.SignUp)
	router.POST("/sign-up", handler.PostSignUp)

	router.POST("/library", handler.Role)

	router.GET("/add-book", handler.AddBook)
	router.POST("/add-book", handler.PostAddBook) 

	router.GET("/remove-book", handler.RemoveBook)
	router.POST("/remove-book", handler.RemoveBook) 

	router.GET("/update-book", handler.UpdateBook)
	router.POST("/update-book", handler.PostUpdateBook) 

	router.GET("/list-issue-request", handler.ListIssueRequests)

	router.GET("/approve-reject-issue-request", handler.ApproveRejectIssueRequest)
	router.POST("/approve", handler.PostApproveRejectIssueRequest) 

	router.GET("/search-book", handler.SearchBook)
	router.POST("/search-book", handler.PostSearchBook) 

	router.GET("/raise-issue-request", handler.RaiseIssueRequest)
	router.POST("/raise-issue", handler.PostRaiseIssueRequest)

	router.Run(":4001")
}
func homePage(c *gin.Context) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(c.Writer, nil)
}

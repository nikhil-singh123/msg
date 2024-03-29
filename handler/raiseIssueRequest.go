package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RaiseIssue struct {
	BookID int `json:"bookid"`
}

func RaiseIssueRequest(c *gin.Context) {
	const tmpl = `
	<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:300px;width:400px">
	<div style="padding-top: 40px;padding-left: 50px;display:flex;flex-direction:column;width:300px;text-align:center;">
	<h1> Raise Issue Request</h1>
	<div style="display: block;text-align: center;">
	<form action="/raise-issue" onsubmit="myFunction()" method="post" style="display: flex;flex-direction: column;width: 300px;text-align: center;">
	Fill Book ID (ISBN) <input type="text" name="bookid" style="margin-top:8px;" required>
	&nbsp;
	<input type="submit" value="Request">
	</form>
	</div>  </div>
	</div>
	</div>

	<script>
    function myFunction() {
    alert("Request Submited");
    }
    </script>
	
	</body>
	`
	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)

}

func PostRaiseIssueRequest(c *gin.Context) {

	var book RaiseIssue
	book.BookID, _ = strconv.Atoi(c.PostForm("bookid"))
	postBody, _ := json.Marshal(book)
	responseBody := bytes.NewBuffer(postBody)

	url := "http://localhost:3001/reader/raise-issue-request/" + Email

	resp, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resp.Status)
	defer resp.Body.Close()

	const tmpl = `
	<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:300px;width:400px">
	<div style="padding-top: 40px;padding-left: 50px;display:flex;flex-direction:column;width:300px;text-align:center;">
	<h1> Raise Issue Request</h1>
	<div style="display: block;text-align: center;">
	<form action="/raise-issue" onsubmit="myFunction()" method="post" style="display: flex;flex-direction: column;width: 300px;text-align: center;">
	Fill Book ID (ISBN) <input type="text" name="bookid" style="margin-top:8px;" required>
	&nbsp;
	<input type="submit" value="Request">
	</form>
	</div>  </div>
	</div>
	</div>
	<script>
    function myFunction() {
    alert(" Request Submited ");
    }
    </script>
	</body>
	`
	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)

}

package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Query string `json:"query"`
}

func SearchBook(c *gin.Context) {
	const tmpl = `
	<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:300px;width:400px">
	<div style="padding-top: 40px;padding-left: 50px;display:flex;flex-direction:column;width:300px;text-align:center;">
	<h1> Search Book</h1>
	<div style="display: block;text-align: center;">
	<form action="/search-book" method="post" style="display: flex;flex-direction: column;width: 300px;text-align: center;">
	Fill the Title/Author/Publisher<input type="text" name="query" style="margin-top:8px;" required>
	&nbsp;
	<input type="submit" value="Get">
	</form>
	</div>  </div>
	</div>
	</div>
	</body>
	`
	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)
}

func PostSearchBook(c *gin.Context) {
	var query Query
	query.Query = c.PostForm("query")

	postBody, _ := json.Marshal(query)
	responseBody := bytes.NewBuffer(postBody)

	hi := io.Reader(responseBody)

	url := "http://localhost:3001/reader/search-book/" + Email

	r, err := http.NewRequestWithContext(c, "GET", url, hi)

	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resp.Status)

	respData, _ := io.ReadAll(resp.Body)

	var re []BookInventory
	if err := json.Unmarshal([]byte(respData), &re); err != nil {
		panic(err)
	}

	const tmpl = `
	<body >
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
    
    <ul style="background-color: lightblue;text-decoration: none;">
	<h1>Book Details</h1>
        {{range .}}
			
			<div class="card">
				<div class="container" style="background-color: lightblue;text-decoration: none;">
		  			<h4 style="margin-right:30px;"><b>Title : {{ .Title }}</b></h4> 
		  			<p>ISBN Number (Book ID) : {{ .ISBN }}</p>
					<p>Author  : {{ .Author }} </p>
					<p>Publisher  : {{ .Publisher }} </p>
					<p>TotalCopies : {{ .TotalCopies }} </p>
					<p>Availablecopies : {{ .AvailableCopies }} </p>
				</div>
	  		</div>
		
        {{end}}
    </ul>
	</div>
	</body>
	`

	// t, _ := template.ParseFiles("listIssueRequest.html")

	t, _ := template.New("webpage").Parse(tmpl)

	t.Execute(c.Writer, re)

	defer resp.Body.Close()

}

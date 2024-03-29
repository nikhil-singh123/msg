package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

type NoBook struct {
	ISBN       int    `json:"isbn"`
	NumberBook uint64 `json:"numberbook"`
}

func RemoveBook(c *gin.Context) {
	if c.Request.Method == "GET" {
		const tmpl = `
	<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:300px;width:400px">
	<h1>Remove Book</h1>
	<form action="/remove-book" onsubmit="myFunction()" method="post" style="padding-top: 40px;
	padding-left: 50px;display:flex;flex-direction:column;width:300px;text-align:center;">
	    ISBN Number(Book ID)<input type="number" name="isbn" required>
	    Number of Books<input type="number" name="nobook" required>
	    &nbsp;
	        <input type="submit" value="Remove Book">
	</form>
	</div>
	</div>
	<script>
    function myFunction() {
    alert("The Book is removed");
    }
    </script>
	</body>
	`
		t, _ := template.New("webpage").Parse(tmpl)
		t.Execute(c.Writer, nil)
	} else {
		var noBook NoBook
		noBook.ISBN, _ = strconv.Atoi(c.PostForm("isbn"))
		n, _ := strconv.Atoi(c.PostForm("nobook"))
		noBook.NumberBook = uint64(n)

		postBody, _ := json.Marshal(noBook)
		responseBody := bytes.NewBuffer(postBody)

		hi := io.Reader(responseBody)

		url := "http://localhost:3001/admin/remove-book/" + Email

		r, err := http.NewRequestWithContext(c, "PUT", url, hi)

		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(r)
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
		<h1>Remove Book</h1>
		<form action="/remove-book" onsubmit="myFunction()" method="post" style="padding-top: 40px;
		padding-left: 50px;display:flex;flex-direction:column;width:300px;text-align:center;">
			ISBN Number(Book ID)<input type="number" name="isbn" required>
			Number of Books<input type="number" name="nobook" required> 
			&nbsp;
				<input type="submit" value="Remove Book">
		</form>
		</div>
		</div>
		<script>
        function myFunction() {
        alert("The Book is removed");
        }  
        </script>
		</body>
	    `

		t, _ := template.New("webpage").Parse(tmpl)
		t.Execute(c.Writer, nil)

	}
}

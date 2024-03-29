package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateBook(c *gin.Context) {
	const tmpl = `<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:470px;width:400px">
	<h3 style="margin-bottom: 0px;"> Update Book In Library</h3>
	<form action="/update-book" onsubmit="myFunction()" method="post" style="padding-top: 30px;padding-left: 50px;
	display:flex;flex-direction:column;width:300px;text-align:center;">
        ISBN<input type="number" name="isbn" required>
        LibID<input type="number" name="libid">
        Title<input type="text" name="title">
        Author<input type="text" name="author">
        Publisher<input type="text" name="publisher">
        Version<input type="text" name="version">
        TotalCopies<input type="number" name="totalcopies">
        AvailableCopies<input type="number" name="availablecopies">
        &nbsp;
            <input type="submit" value="Update-Book">
    </form>
	</div>
	</div>
	<script>
    function myFunction() {
    alert("The Book is Updated");
    }
    </script>
	</body>`

	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)

}

func PostUpdateBook(c *gin.Context) {
	var book BookInventory
	book.ISBN, _ = strconv.Atoi(c.PostForm("isbn"))
	book.LibID, _ = strconv.Atoi(c.PostForm("libid"))
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Publisher = c.PostForm("publisher")
	book.Version = c.PostForm("version")
	total, _ := strconv.Atoi(c.PostForm("totalcopies"))
	book.TotalCopies = uint64(total)
	avail, _ := strconv.Atoi(c.PostForm("availablecopies"))
	book.AvailableCopies = uint64(avail)

	postBody, _ := json.Marshal(book)
	responseBody := bytes.NewBuffer(postBody)

	hi := io.Reader(responseBody)

	url := "http://localhost:3001/admin/update-book/" + Email

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

	const tmpl = `<body>
	
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:470px;width:400px">
	<h3 style="margin-bottom: 0px;"> Update Book In Library</h3>
	<form action="/Update-book" onsubmit="myFunction()" method="post" style="padding-top: 30px;padding-left: 50px;
	display:flex;flex-direction:column;width:300px;text-align:center;">
        ISBN<input type="number" name="isbn" required>
        LibID<input type="number" name="libid">
        Title<input type="text" name="title">
        Author<input type="text" name="author">
        Publisher<input type="text" name="publisher">
        Version<input type="text" name="version">
        TotalCopies<input type="number" name="totalcopies">
        AvailableCopies<input type="number" name="availablecopies">
        &nbsp;
            <input type="submit" value="Update-Book">
    </form>
	</div>
	</div>
	<script>
    function myFunction() {
    alert("The Book is Updated");
    }
    </script>
	</body>`

	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)

}

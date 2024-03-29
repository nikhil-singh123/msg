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

type BookInventory struct {
	ISBN            int    `json:"isbn"`
	LibID           int    `json:"libid"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     uint64 `json:"totalcopies"`
	AvailableCopies uint64 `json:"availablecopies"`
}

func AddBook(c *gin.Context) {
	const editTmpl = `
	<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:470px;width:400px">
	<h3 style="margin-bottom: 0px;"> Add Book In Library</h3>
	<form action="/add-book" onsubmit="myFunction()" method="post" style="padding-top: 30px;padding-left: 50px;
	display:flex;flex-direction:column;width:300px;text-align:center;">
        ISBN<input type="number" name="isbn" required>
        LibID<input type="number" name="libid" required>
        Title<input type="text" name="title" required>
        Author<input type="text" name="author" required>
        Publisher<input type="text" name="publisher" required>
        Version<input type="text" name="version" required>
        TotalCopies<input type="number" name="totalcopies" required>
        AvailableCopies<input type="number" name="availablecopies" required>
        &nbsp;
            <input type="submit" value="Add-Book">
    </form>
	</div>
	</div>
	<script>
    function myFunction() {
    alert("The book was added");
    }
    </script>
	</body>
	`

	t, _ := template.New("page").Parse(editTmpl)
	t.Execute(c.Writer, nil)

}

func PostAddBook(c *gin.Context) {

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

	url := "http://localhost:3001/admin/add-book/" + Email

	resp, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resp.Status)
	defer resp.Body.Close()

	const editTmpl = `
	<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:470px;width:400px">
	<h3 style="margin-bottom: 0px;"> Add Book In Library</h3>
	<form action="/add-book" onsubmit="myFunction()" method="post" style="padding-top: 30px;padding-left: 50px;
	display:flex;flex-direction:column;width:300px;text-align:center;">
        ISBN<input type="number" name="isbn" required>
        LibID<input type="number" name="libid" required>
        Title<input type="text" name="title" required>
        Author<input type="text" name="author" required>
        Publisher<input type="text" name="publisher" required>
        Version<input type="text" name="version" required>
        TotalCopies<input type="number" name="totalcopies" required>
        AvailableCopies<input type="number" name="availablecopies" required>
        &nbsp;
            <input type="submit" value="Add-Book">
    </form>
	</div>
	</div>
	<script>
    function myFunction() {
    alert("The Book is Added");
    }
    </script>
	</body>
	`

	t, _ := template.New("page").Parse(editTmpl)
	t.Execute(c.Writer, nil)

}

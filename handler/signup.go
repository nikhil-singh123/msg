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

func SignUp(c *gin.Context) {
	const tmpl = `
	<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:470px;width:400px">
	<h3 style="margin-bottom: 0px;"> Sign Up</h3>
	<form action="/sign-up" onsubmit="myFunction()" method="post" style="padding-top: 30px;padding-left: 50px;
	display:flex;flex-direction:column;width:300px;text-align:center;">
	Name <input type="text" name="name" style="margin-top:8px;" required>
	Email<input type="text" name="email" style="margin-top:8px;" required>
	Phone Number <input type="number" name="phonenumber" style="margin-top:8px;" required>
	Library ID <input type="text" name="libid" style="margin-top:8px;" required>
	&nbsp;
	<input type="submit" value="sign-up">
	</form>
	</div>
	</div>
	<script>
    function myFunction() {
    alert("The user was added");
    }
    </script>
	</body>
	`
	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)
}

func PostSignUp(c *gin.Context) {
	var user Users

	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.ContactNumber, _ = strconv.Atoi(c.PostForm("phonenumber"))
	user.LibID, _ = strconv.Atoi(c.PostForm("libid"))
	user.Role = "reader"

	postBody, _ := json.Marshal(user)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:3001/admin/add-user/admin@gmail.com", "application/json", responseBody)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resp.Status)
	defer resp.Body.Close()

	const tmpl = `
	<body >
	<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
	<div style="display: flex;justify-content: center;">
	
	<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:470px;width:400px">
	<h3 style="margin-bottom: 0px;"> Sign Up</h3>
	<form action="/sign-up" onsubmit="myFunction()" method="post" style="padding-top: 30px;padding-left: 50px;
	display:flex;flex-direction:column;width:300px;text-align:center;">
	Name <input type="text" name="name" style="margin-top:8px;" required>
	Email<input type="text" name="email" style="margin-top:8px;" required>
	Phone Number <input type="number" name="phonenumber" style="margin-top:8px;" required>
	Library ID <input type="text" name="libid" style="margin-top:8px;" required>
	&nbsp;
	<input type="submit" value="sign-up">
	</form>
	</div>
	</div>
	<script>
    function myFunction() {
    alert("The user was added");
    }
    </script>
	</body>
	`
	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)

}

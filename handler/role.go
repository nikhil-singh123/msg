package handler

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

var Email string

func Role(c *gin.Context) {
	var user Users

	const wrongMail = `
	<body>
	<a href="http://localhost:4001/" style="background-color: lightblue;">Home Page</a>
	<h1 style="background-color: lightblue; text-align:center;">Wrong User</h1>
	</body>
	`

	email := c.PostForm("email")

	Email = email

	if err := DB.Where("email = ?", Email).First(&user).Error; err != nil {
		t, _ := template.New("webpage").Parse(wrongMail)
		t.Execute(c.Writer, nil)
		return
	}

	if user.Role == "admin" {
		const tmpl = `
		<html>
		<head>
		</head>
		<body>
		<a href="http://localhost:4001/" style="color: antiquewhite;text-decoration: none;">Home Page</a>
		
		<div style="margin-top: 5%;">

		<h1 style="background-color: lightblue;text-align: center;"> Admin Portal </h1>

		<div style="display: flex;justify-content: space-around;margin-top: 80px;">


		<div style="border: 2px white solid;border-radius: 5px;">
		
		<div style="width: 180px;background-color: lightblue;height: 140px;text-align: center;padding-top: 10px;border-radius: 5px;">
   		<a href="http://localhost:4001/add-book" style=" color:black; text-decoration: none;">		
		<h2>Add Book</h2>
		</a>
		</div>
		</div>

		<div style="border: 2px white solid;border-radius: 5px;">
		
		<div style="width: 180px;background-color: lightblue;height: 140px;text-align: center;padding-top: 10px;border-radius: 5px;">
   		<a href="http://localhost:4001/remove-book" style=" color:black; text-decoration: none;">
		<h2>remove-book</h2>
		</a> 
		</div>
  		</div>

		<div style="border: 2px white solid;border-radius: 5px;">
		
		<div style="width: 180px;background-color: lightblue;height: 140px;text-align: center;padding-top: 10px;border-radius: 5px;">
   		<a href="http://localhost:4001/update-book" style=" color:black; text-decoration: none;">
		<h2>Update Book</h2>
		</a> 
		</div>
  		</div>

		<div style="border: 2px white solid;border-radius: 5px;">
		
		<div style="width: 180px;background-color: lightblue;height: 140px;text-align: center;padding-top: 10px;border-radius: 5px;">
   		<a href="http://localhost:4001/list-issue-request" style=" color:black; text-decoration: none;">
		<h2>List of Issue Request</h2>
		</a> 
		</div>
  		</div>
		
		<div style="border: 2px white solid;border-radius: 5px;">
		
		<div style="width: 180px;background-color: lightblue;height: 140px;text-align: center;padding-top: 10px;border-radius: 5px;">
   		<a href="http://localhost:4001/approve-reject-issue-request" style=" color:black; text-decoration: none;">		
		<h2>Approve The Issue Request</h2>
		</a> 
		</div>
		</div>
  		
		</div>

		</div>
		</body>
		</html>
		`

		t, _ := template.New("webpage").Parse(tmpl)
		t.Execute(c.Writer, nil)
	} else {
		const tmpl = `
		<html>
		<head>
		</head>
		<body>
		<a href="http://localhost:4001/" style="color: antiquewhite;text-decoration: none;">Home Page</a>
		
		<div style="margin-top: 5%;align-items: center;display: flex;flex-direction: column;">
		<h1 style="background-color: lightblue;text-align: center;width: 40%;"> Reader Portal </h1>

		<div style="display: flex;justify-content: space-around;margin-top: 80px;">


		<div style="border: 2px white solid;border-radius: 5px;height: 218px; margin-right: 320px;">
		
		<div style="width: 180px;background-color: lightblue;height: 80px;text-align: center;padding-top: 10px;border-radius: 5px;">
   		<a href="http://localhost:4001/search-book" style=" color:black; text-decoration: none;">		
		<h2>Search Book</h2>
		</a>
		</div>
		</div>

		<div style="border: 2px white solid;border-radius: 5px;height: 218px;">
		
		<div style="width: 180px;background-color: lightblue;height: 80px;text-align: center;padding-top: 10px;border-radius: 5px;">
   		<a href="http://localhost:4001/raise-issue-request" style=" color:black; text-decoration: none;">
		<h2>Raise Issue Request</h2>
		</a> 
		</div>
  		</div>
  		
		</div>

		</div>
		</body>
		</html>
		`

		t, _ := template.New("webpage").Parse(tmpl)
		t.Execute(c.Writer, nil)
	}

}

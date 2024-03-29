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

func ApproveRejectIssueRequest(c *gin.Context) {
	const tmpl = `
	<body>
		<a href="http://localhost:4001/" style="background-color: lightblue;text-decoration: none;">Home Page</a>
		<div style="display: flex;justify-content: center;">
		<div style="display: block;justify-content: center;text-align: center;background-color: lightblue;height:300px;width:400px">
		<div style="padding-top: 40px;padding-left: 50px;display:flex;flex-direction:column;width:300px;text-align:center;">
		<h3> Approve Issue Request</h3>
		<div style="display: block;text-align: center;">
		<form action="/approve" onsubmit="myFunction()" method="post" style="display: flex;flex-direction: column;width: 300px;text-align: center;">
		Request ID <input type="number" name="reqid" required>
		&nbsp;
		<input type="submit" value="Approve">
		</form>
		</div>  </div>
		</div>
		</div>
		<script>
    function myFunction() {
    alert("Submited");
    }
    </script>
		</body>
	`

	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)

}

type DetailReqID struct {
	ReqID int `json:"reqid"`
}

func PostApproveRejectIssueRequest(c *gin.Context) {
	var reqid DetailReqID

	reqid.ReqID, _ = strconv.Atoi(c.PostForm("reqid"))

	postBody, _ := json.Marshal(reqid)
	responseBody := bytes.NewBuffer(postBody)

	url := "http://localhost:3001/admin/approve-reject-issue-request/" + Email

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
	<h3> Approve Issue Request</h3>
	<div style="display: block;text-align: center;">
	<form action="/approve" onsubmit="myFunction()" method="post" style="display: flex;flex-direction: column;width: 300px;text-align: center;">
	Request ID <input type="number" name="reqid" required>
	&nbsp;
	<input type="submit" value="Approve">
	</form>
	</div>  </div>
	</div>
	</div>
	<script>
    function myFunction() {
    alert("Submited");
    }
    </script>
	</body>
	`

	t, _ := template.New("page").Parse(tmpl)
	t.Execute(c.Writer, nil)

}

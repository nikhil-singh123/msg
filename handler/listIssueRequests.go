package handler

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestEvents struct {
	ReqID        int        `json:"reqid"`
	BookID       int        `json:"bookid"`
	ReaderId     int        `json:"readerid"`
	RequestDate  time.Time  `json:"requestdate"`
	ApprovelDate *time.Time `json:"approveldate"`
	ApproverID   int        `json:"approverid"`
	RequestType  string     `json:"requesttype"`
}

func ListIssueRequests(c *gin.Context) {

	url := "http://localhost:3001/admin/list-issue-request/" + Email

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	respData, _ := io.ReadAll(resp.Body)

	var re []RequestEvents
	if err := json.Unmarshal([]byte(respData), &re); err != nil {
		panic(err)
	}

	const tmpl = `
	<body>
    
    <ul style="background-color: lightblue;text-decoration: none;">
	<h1>List of Issue Requests</h1>
        {{range .}}
			<div class="card">
				<div class="container" style="background-color: lightblue;text-decoration: none;">
		  			<h4><b>Request ID : {{ .ReqID }}</b></h4> 
		  			<p>Book ISBN Number (Book ID) : {{ .BookID }}</p>
					<p>Request Date : {{ .RequestDate }} </p>
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
}

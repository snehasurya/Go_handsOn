package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type restApiStruct struct {
	suite.Suite
}

func TestRestAPI(t *testing.T) {
	suite.Run(t, new(restApiStruct))
}
func (r *restApiStruct) TestHandlerFunc() {
	client := http.Client{}

	resp, _ := client.Get("http://localhost:8181/employee")
	var employee []string
	r.Equal(http.StatusOK, resp.StatusCode)
	//fmt.Println(resp)
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &employee)
	r.NotEmpty(employee)
	r.Assert().Len(employee, 2)
}

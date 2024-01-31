package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T){
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	response, err:=  testClient.Get(testServer.URL)
	if err !=nil{
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil{
		t.Error(err)
	}
	assert.Equal(t,"Hello world", string(body))
	assert.Equal(t, 200, response.StatusCode)
}
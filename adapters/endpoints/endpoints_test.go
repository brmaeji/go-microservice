package endpoints

import (
	"fmt"
	"io/ioutil"
	"log"
	"microservice/adapters/data"
	service "microservice/core"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var srv service.MentionService
var httpserver *httptest.Server

func TestMain(m *testing.M) {
	createdSrv, err := service.NewMentionsService(data.Memory)
	if err != nil {
		log.Fatalf("Could not create MentionService: %v\n", err)
	}

	log.Printf("created base service %v", srv)
	srv = createdSrv

	exitVal := m.Run()
	httpserver.Close()

	os.Exit(exitVal)
}

func TestCreateHTTPServer(t *testing.T) {
	fmt.Println("testCreateHTTPServer srv", srv)
	router, err := CreateHTTPServer(srv)
	if err != nil {
		t.Errorf("could not create server. Error: %v\n", err)
	}

	ts := httptest.NewServer(router)
	httpserver = ts
}

func TestCreate(t *testing.T) {

	createURL := fmt.Sprintf("%v/create/Ramones", httpserver.URL)

	res, err := http.Get(createURL)
	if err != nil {
		t.Errorf("could not get: %v\n", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("expected 200, got %v\n", res.StatusCode)
	}
}

func TestFind(t *testing.T) {

	findURL := fmt.Sprintf("%v/find/Ramones", httpserver.URL)

	res, err := http.Get(findURL)
	if err != nil {
		t.Errorf("could not find: %v\n", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("expected 200, got %v\n", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	expectedBody := `[{"name":"Ramones","mentions":1}]`
	fmt.Printf("%s", body)

	if expectedBody != string(body) {
		t.Errorf("expected %v, got %v\n", expectedBody, string(body))
	}
}

func TestIncrease(t *testing.T) {

	increaseURL := fmt.Sprintf("%v/increase/Ramones", httpserver.URL)

	res, err := http.Get(increaseURL)
	if err != nil {
		t.Errorf("could not increase: %v\n", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("expected 200, got %v\n", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	expectedBody := `{"name":"Ramones","mentions":2}`
	fmt.Printf("%s", body)

	if expectedBody != string(body) {
		t.Errorf("expected %v, got %v\n", expectedBody, string(body))
	}
}

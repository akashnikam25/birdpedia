package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	r, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder() // it will act as browser
	hf := http.HandlerFunc(handler)

	hf.ServeHTTP(recorder, r)

	if recorder.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			recorder.Code, http.StatusOK)
	}
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRouter(t *testing.T) {
	r := newRouter()
	mockserver := httptest.NewServer(r)
	resp, err := http.Get(mockserver.URL + "/hello")
	if err != nil {
		t.Fatal()
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal()
	}
	expected := `Hello World!`
	actual := string(b)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestNonExistingRouter(t *testing.T) {
	r := newRouter()
	mockserver := httptest.NewServer(r)
	resp, err := http.Post(mockserver.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal()
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal()
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouter1(t *testing.T) {
	r := newRouter()
	mockserver := httptest.NewServer(r)
	resp, err := http.Get(mockserver.URL + "/assets/")
	if err != nil {
		t.Fatal()
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal()
	}
	//	expected := `Hello World!`
	actual := string(b)
	fmt.Println("actual", actual, resp.Header["Content-Type"])
}

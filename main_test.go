package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostEstimate(t *testing.T) {

	initTypes()

	gin.SetMode(gin.TestMode)
	handler := Estimate
	router := gin.Default()
	router.POST("/estimate", handler)

	var jsonStr = []byte(`{"Denomination":"TestAutoEntreprise","TypeId":2,"Siret":"EST1865","Revenue":1000}`)
	req, err := http.NewRequest("POST", "/estimate", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	expected := "250"
	result := strings.TrimSuffix(resp.Body.String(), "\n")
	if result != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			result, expected)
	}

	jsonStr = []byte(`{"Denomination":"TestSAS","TypeId":1,"Adress":"21 rue louvier 34880","Siret":"HYT289654","Revenue":10000}`)
	req, err = http.NewRequest("POST", "/estimate", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	expected = "3300"
	result = strings.TrimSuffix(resp.Body.String(), "\n")
	if result != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			result, expected)
	}

	jsonStr = []byte(`{"Denomination":"TestSASMissingAdress","TypeId":1,"Siret":"HYT289654","Revenue":1000}`)
	req, err = http.NewRequest("POST", "/estimate", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	expectedStatus := http.StatusBadRequest
	if resp.Code != expectedStatus {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			resp.Body.String(), expectedStatus)
	}
	jsonStr = []byte(`{"Denomination":"TestAEMissingSiret","TypeId":2,"Revenue":1000}`)
	req, err = http.NewRequest("POST", "/estimate", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	expectedStatus = http.StatusBadRequest
	if resp.Code != expectedStatus {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			resp.Body.String(), expectedStatus)
	}

}

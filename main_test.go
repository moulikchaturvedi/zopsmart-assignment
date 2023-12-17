package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func PostMethod(c *gin.Context) {
  message := "PostMethod called"
  c.JSON(http.StatusOK, message)
}

func GetMethod(c *gin.Context) {
  message := "GetMethod called"
  c.JSON(http.StatusOK, message)
}


func TestMain(t *testing.T){
    // Switch to test mode so we don't get such noisy output
    gin.SetMode(gin.TestMode)

    r := gin.Default()
    r.POST("/students", PostMethod)

	// app := gofr.New()
	// app.POST("/students", PostStudentHandler)

    req, err := http.NewRequest(http.MethodPost, "/students", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    // Check to see if the response was what you expected
    if w.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    a := gin.Default()
    a.GET("/", GetMethod)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    areq, aerr := http.NewRequest(http.MethodGet, "/", nil)
    if aerr != nil {
        t.Fatalf("Couldn't create request: %v\n", aerr)
    }

    // Create a response recorder so you can inspect the response
    aw := httptest.NewRecorder()

    // Perform the request
    a.ServeHTTP(aw, areq)
    fmt.Println(aw.Body)

    // Check to see if the response was what you expected
    if aw.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, aw.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, aw.Code)
    }
}
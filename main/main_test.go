package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type fooTest struct {
	TestName, Key, Name, UrlParam, expectedBody string
	Req                                         *http.Request
	StatusCode                                  int
}

var getFooTests = []fooTest{
	{
		"Test Foo Get Success",
		"71d17f09-acfe-4c90-a90c-ca3a42d56ca4",
		"Banana",
		"71d17f09-acfe-4c90-a90c-ca3a42d56ca4",
		"{\"name\":\"Banana\",\"id\":\"71d17f09-acfe-4c90-a90c-ca3a42d56ca4\"}",
		mockRequest("GET", "/foo/71d17f09-acfe-4c90-a90c-ca3a42d56ca4", ""),
		200,
	},
	{
		"Test Foo Get Missing",
		"",
		"",
		"71d17f09-acfe-4c90-a90c-ca3a42d56ca4",
		"",
		mockRequest("GET", "/foo/71d17f09-acfe-4c90-a90c-ca3a42d56ca4", ""),
		404,
	},
	{
		"Test Foo Delete Success",
		"71d17f09-acfe-4c90-a90c-ca3a42d56ca4",
		"Banana",
		"71d17f09-acfe-4c90-a90c-ca3a42d56ca4",
		"",
		mockRequest("DELETE", "/foo/71d17f09-acfe-4c90-a90c-ca3a42d56ca4", ""),
		204,
	},
	{
		"Test Foo Delete Missing",
		"",
		"",
		"71d17f09-acfe-4c90-a90c-ca3a42d56ca4",
		"",
		mockRequest("DELETE", "/foo/71d17f09-acfe-4c90-a90c-ca3a42d56ca4", ""),
		404,
	},
	{
		"Test Foo Post Invalid",
		"71d17f09-acfe-4c90-a90c-ca3a42d56ca4",
		"Banana",
		"71d17f09-acfe-4c90-a90c-ca3a42d56ca4",
		"{\"message\":\"Invalid Request\"}",
		mockRequest("POST", "/foo/71d17f09-acfe-4c90-a90c-ca3a42d56ca4", "{\"name\": [\"Banana\"]}"),
		400,
	},
}

func mockRequest(verb, url, body string) *http.Request {
	// Mock up your http request
	if body != "" {
		r, _ := http.NewRequest(verb, url, strings.NewReader(body))
		return r
	}
	r, _ := http.NewRequest(verb, url, nil)
	return r
}

func TestFooGet(t *testing.T) {
	for _, test := range getFooTests {
		// TEST SETUP
		fmt.Println(test.TestName)
		mr := miniredis.RunT(t)
		if test.Key != "" {
			mr.Set(test.Key, test.Name)
		}

		Redis = redis.NewClient(&redis.Options{
			Addr:     mr.Addr(),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})

		r := test.Req
		w := httptest.NewRecorder()
		if test.UrlParam != "" {
			vars := map[string]string{
				"id": test.UrlParam,
			}
			r = mux.SetURLVars(r, vars)
		}
		// TEST EXECUTION
		switch r.Method {
		case "GET":
			GetFoo(w, r)
		case "DELETE":
			DeleteFoo(w, r)
		case "POST":
			PostFoo(w, r)
		}

		// TEST RESULTS
		strRes := ""
		if w.Body != nil {
			res, _ := io.ReadAll(w.Body)
			strRes = string(res)
		}
		if test.expectedBody != strRes {
			t.Errorf("Response Body was incorrect, got: %s, want: %s.", strRes, test.expectedBody)
		}

		if test.StatusCode != w.Code {
			t.Errorf("Respons Code was incorrect, got: %d, want: %d.", w.Code, test.StatusCode)
		}

	}

}

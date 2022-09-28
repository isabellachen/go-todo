package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func setupApi(t *testing.T) (string, func()) {
	t.Helper()
	testServer := httptest.NewServer(newMux(""))
	tempTodoFile, err := ioutil.TempFile("", "todotest")

	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		var body bytes.Buffer
		taskName := fmt.Sprintf("Task number %d", i)

		item := struct {
			Task string `json:"task"`
		}{
			Task: taskName,
		}

		if err := json.NewEncoder(&body).Encode(item); err != nil {
			t.Fatal(err)
		}

		response, err := http.Post(testServer.URL+"/todo", "application/json", &body)

		if err != nil {
			t.Fatal(err)
		}

		if response.StatusCode != http.StatusCreated {
			t.Fatalf("Failed to add initial items with %d", response.StatusCode)
		}
	}

	return testServer.URL, func() {
		testServer.Close()
		os.Remove(tempTodoFile.Name())
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name               string
		path               string
		expectedStatusCode int
		expectedContent    string
	}{
		{name: "GetRoot", path: "/", expectedStatusCode: http.StatusOK, expectedContent: "hello world"},
		{name: "ListTodos", path: "/todo", expectedStatusCode: http.StatusOK, expectedContent: "hello world"},
		{name: "NotFound", path: "/nugget", expectedStatusCode: http.StatusNotFound},
	}

	url, cleanup := setupApi(t)
	defer cleanup()

	for _, testCase := range testCases {
		res, err := http.Get(url + testCase.path)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		if res.StatusCode != testCase.expectedStatusCode {
			t.Errorf("Expected %s, got %s", http.StatusText(testCase.expectedStatusCode), http.StatusText(res.StatusCode))
		}

		switch {
		case strings.Contains(res.Header.Get("Content-Type"), "text/plain"):
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}
			if !strings.Contains(string(body), testCase.expectedContent) {
				t.Errorf("Expected %q, got %q", testCase.expectedContent, string(body))
			}
		case strings.Contains(res.Header.Get("Content-Type"), "application/json"):
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}
			fmt.Println("🍌", string(body))
		default:
			t.Fatalf("Unsupported Content-Type: %q", res.Header.Get("Content-Type"))
		}
	}
}

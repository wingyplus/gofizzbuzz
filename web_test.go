package gofizzbuzz

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFizzBuzzAllowHTTPMethod(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(FizzBuzzHandler))
	defer ts.Close()

	var r, _ = http.Get(ts.URL)
	var b, _ = ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if r.StatusCode != http.StatusMethodNotAllowed {
		t.Error("expect status Method not allowed")
	}

	if string(b) != "Method not allowed\n" {
		t.Errorf("expect body Method not allowed")
	}
}

func TestFizzBuzzParseNumberError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(FizzBuzzHandler))
	defer ts.Close()

	var r, _ = http.PostForm(ts.URL, url.Values{"number": []string{"hello"}})
	defer r.Body.Close()

	if r.StatusCode != http.StatusInternalServerError {
		t.Error("expect status Internal Server Error")
	}
}

func TestFizzBuzzPostForm(t *testing.T) {
	var tests = []struct {
		n        string
		expected string
	}{
		{n: "1", expected: "1"},
		{n: "3", expected: "Fizz"},
	}

	for _, test := range tests {
		var w = httptest.NewRecorder()
		var r, _ = http.NewRequest(
			"POST",
			"/say",
			bytes.NewBufferString(
				url.Values{"number": []string{test.n}}.Encode()),
		)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		FizzBuzzHandler(w, r)

		testResponseBody(t, w, test.expected)
		testResponseHeader(t, w)
	}
}

func testResponseBody(t *testing.T, w *httptest.ResponseRecorder, word string) {
	var body = `<!DOCTYPE html>
<html>
	<head>
		<title>FizzBuzz</title>
	</head>
	<body>
		<span>Say: %s</span>
	</body>
</html>
`
	var (
		actual   = w.Body.String()
		expected = fmt.Sprintf(body, word)
	)
	if actual != expected {
		t.Errorf("body mismatch: \n%s\nBut was:\n%s",
			expected,
			actual)
	}
}

func testResponseHeader(t *testing.T, w *httptest.ResponseRecorder) {
	if w.Header().Get("Content-Type") != "text/html" {
		t.Error("expect content-type is text/html")
	}
}

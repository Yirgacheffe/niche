package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

var feed = `<?xml version="1.0" encoding="utf-8"?>
<rss>
</rss>
`

func mockServer() *httptest.Server {

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))

}

func TestDownload(t *testing.T) {

	server := mockServer()
	defer server.Close()

	url := server.URL
	statusCode := http.StatusOK

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to  make the Get call.", ballotX, err)
			}

			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v, %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}

}

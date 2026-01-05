package main

import (
	"net/http"
	"testing"

	"snippetbox-webapp/internal/assert"
)

// TestPing() An example to understand how to use the net/http/httptest package
// to test http handlers. E.g httptest.ResponseRecorder (an implementation of 
// ResponseWriter), which records the status code, headers and body instead of 
// actually writing them to a HTTP connection.
// It checks that the response status code written by the ping handler is 200.
// It also checks that the response body written by the ping handler is OK.
func TestPing(t *testing.T) {
	// Create a new instance of our application struct.
	app := newTestApplication(t)

	// Use httptest.NewTLSServer() function to create a new test server,
	// passing in the value returned by the app.routes() method as the 
	// handler for the server.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// The netword address that the server is listening on is contained in
	// the ts.URL field. We cn use this along with the ts.Client().Get() method
	// to make a GET /ping request against the test server.
	// It returns a http.Response struct containing the response.
	code, _, body := ts.get(t, "/ping")

	// We can then check the value of the response status code and body using
	// the same pattern as before.
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
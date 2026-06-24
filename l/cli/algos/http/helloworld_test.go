package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns hello world", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		helloWorldServer(response, request)

		got := response.Body.String()
		want := "Hello, World"

		tablewriter.AssertStringGotWant(t, got, want)
	})
}

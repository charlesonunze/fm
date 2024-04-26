package server_test

import (
	"net/http"
	"testing"

	"fm/integrationtest"

	"github.com/gookit/goutil/testutil/assert"
)

func TestServer_Start(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("starts the server and listens for requests", func(t *testing.T) {
		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:7777/health")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

package handlers_test

import (
	"net/http"
	"testing"

	"fm/api/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		hh := handlers.NewHealthHandler()

		ctx, res := makeGetRequest("/health")
		err := hh.CheckHealth(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "I'm alive!", res.Body.String())
	})
}

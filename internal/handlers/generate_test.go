package handlers_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MalukiMuthusi/pulseid/internal/handlers"
	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	store := store.NewMockStore()
	assert.NotNil(t, store, "expected a valid data structure")

	router := handlers.SetUpRouter(store)

	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/generate", nil)
	if err != nil {
		assert.Fail(t, "failed to create a new request in test")
	}

	// Add Authorization Header

	credentials := base64.StdEncoding.EncodeToString([]byte("username:password"))

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", credentials))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var res models.Token

	b, err := ioutil.ReadAll(w.Body)

	if err != nil {
		assert.Fail(t, "failed to read response")
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		assert.Fail(t, "failed to unMarshal response")
	}

	assert.NotNil(t, res, "expected a valid data structure")

	assert.EqualValues(t, res.ID, 0)
}

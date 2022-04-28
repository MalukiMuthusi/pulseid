package handlers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MalukiMuthusi/pulseid/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	router := handlers.SetUpRouter()

	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/generate", nil)
	if err != nil {
		assert.Fail(t, "failed to create a new request in test")
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)

	expectedRes := map[string]interface{}{"message": "not implemented"}

	var res map[string]interface{}

	b, err := ioutil.ReadAll(w.Body)

	if err != nil {
		assert.Fail(t, "failed to read response")
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		assert.Fail(t, "failed to unMarshal response")
	}

	assert.EqualValues(t, expectedRes, res)
}

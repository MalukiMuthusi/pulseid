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
	"github.com/MalukiMuthusi/pulseid/internal/store/mock"
	"github.com/stretchr/testify/assert"
)

func TestRecall(t *testing.T) {

	store := mock.NewStore()

	type test struct {
		Name     string
		EndPoint string
		Status   int
		Res      models.RecallTokenResponse
	}

	tests := []test{
		{
			Name:     "happy case",
			EndPoint: "/recall/some_unique_token",
			Status:   http.StatusOK,
			Res: models.RecallTokenResponse{
				Success: true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodGet, test.EndPoint, nil)
			if err != nil {
				assert.Fail(t, "failed to create a new request in test")
			}

			// Add Authorization Header

			credentials := base64.StdEncoding.EncodeToString([]byte("username:password"))

			req.Header.Add("Authorization", fmt.Sprintf("Basic %s", credentials))

			w := httptest.NewRecorder()

			router := handlers.SetUpRouter(store, DebugPrintRoute)

			router.ServeHTTP(w, req)

			assert.Equal(t, test.Status, w.Code)

			var res models.RecallTokenResponse

			b, err := ioutil.ReadAll(w.Body)
			if err != nil {
				assert.Fail(t, "failed to read response")
			}

			err = json.Unmarshal(b, &res)
			if err != nil {
				assert.Fail(t, "failed to unMarshal response")
			}

			assert.EqualValues(t, test.Res, res)
		})
	}
}

package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {
	tapi := New()
	url := "/api/transaction"

	t.Run("Test  transaction", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := postRequest{
			Type:   "c",
			Amount: 100.00,
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Error("invalid json input")
		}

		req, err := http.NewRequest("POST", url, bytes.NewReader(body))
		if err != nil {
			t.Error("error creating request")
		}

		tapi.ServerHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Error("can not add credit")
		}
	})

	t.Run("Test  transaction", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := postRequest{
			Type:   "c",
			Amount: -100.00,
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Error("invalid json input")
		}

		req, err := http.NewRequest("POST", url, bytes.NewReader(body))
		if err != nil {
			t.Error("error creating request")
		}

		tapi.ServerHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("can not add credit")
		}
	})

	t.Run("Test  transaction", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload := postRequest{
			Type:   "h",
			Amount: 100.00,
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Error("invalid json input")
		}

		req, err := http.NewRequest("POST", url, bytes.NewReader(body))
		if err != nil {
			t.Error("error creating request")
		}

		tapi.ServerHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Error("can not add credit")
		}
	})

}

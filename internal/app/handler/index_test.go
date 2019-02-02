package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/bakhtik/webapp/internal/pkg/testing"
)

type CacheMock struct{}

func (cm *CacheMock) Increment(key string) (result int64, err error) {
	return 42, nil
}

func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler := &Handler{cache: &CacheMock{}}
	handler.Index().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is not 200: %v", resp.StatusCode)
	}
	want := "Web application template"
	if strings.Contains(string(body), want) == false {
		t.Errorf("Body does not contain %q", want)
	}
	want = "Visits: 42"
	if strings.Contains(string(body), want) == false {
		t.Errorf("Body does not contain %q", want)
	}
}

package examples

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_AcceptStatus(t *testing.T) {
	r := httptest.NewRequest("POST", "/some/path", strings.NewReader("hello"))
	w := httptest.NewRecorder()
	AcceptStatus(w, r)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Error(resp.Status)
	}
	// verify body
	data, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	exp := "ok"
	if got := string(data); got != exp {
		t.Errorf("got %s, expected %s", got, exp)
	}
}

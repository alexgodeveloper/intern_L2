package main

import (
	"net/http/httptest"
	"testing"
)

func Test_http(t *testing.T) {
	want := 200
	req := httptest.NewRequest("GET", "http://127.0.0.1:8080/events_for_week", nil)
	w := httptest.NewRecorder()
	EventsForWeek(w, req)

	resp := w.Result()
	if resp.StatusCode != want {
		t.Errorf("Status code %d not equal %d", resp.StatusCode, want)
	}
}

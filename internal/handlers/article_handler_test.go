package handlers

import (
	"net/http"
	"testing"
)

func TestHandlerSaveArticle(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.SaveArticle(tt.args.w, tt.args.r)
		})
	}
}

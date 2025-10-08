package main_test

import (
	main "goreader"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestJsonParser(t *testing.T) {
	jsonSample := "{\"Path\":\"sample.txt\"}"
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r       *http.Request
		want    main.ProcessJson
		wantErr bool
	}{
		{
			name:    "Succeeded",
			r:       httptest.NewRequest("POST", "/input", io.NopCloser(strings.NewReader(jsonSample))),
			want:    main.ProcessJson{Path: "sample.txt"},
			wantErr: false,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := main.JsonParser(tt.r)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("JsonParser() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("JsonParser() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

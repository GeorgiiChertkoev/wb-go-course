package main

import (
	"reflect"
	"testing"
)

func TestParseFields(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{
			name:    "single number",
			input:   "7",
			want:    []int{7},
			wantErr: false,
		},
		{
			name:    "comma separated",
			input:   "1,2,3",
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "range only",
			input:   "5-7",
			want:    []int{5, 6, 7},
			wantErr: false,
		},
		{
			name:    "mixed numbers and range",
			input:   "1,3-5,8",
			want:    []int{1, 3, 4, 5, 8},
			wantErr: false,
		},
		{
			name:    "multiple ranges",
			input:   "1-2,4-6",
			want:    []int{1, 2, 4, 5, 6},
			wantErr: false,
		},
		{
			name:    "spaces in input",
			input:   " 1 , 2-3 ",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid number",
			input:   "a",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid range",
			input:   "5-b",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "reverse range not allowed",
			input:   "5-3",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFields(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFields() got = %v, want %v", got, tt.want)
			}
		})
	}
}

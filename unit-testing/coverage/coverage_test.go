package main

import "testing"

func TestCoverage(t *testing.T) {

	type args struct {
		condition bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"no condition", args{true}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Coverage(tt.args.condition); (err != nil) != tt.wantErr {
				t.Errorf("Coverage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

// go test -cover
// go test -coverprofile=cover.out
// go tool cover -html=cover.out -o coverage.html

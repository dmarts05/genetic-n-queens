package config

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	defaultConfig := Config{
		NumRuns:         1,
		SelectionMethod: Roulette,
		PopulationSize:  16,
		MaxGenerations:  100,
		NumQueens:       8,
		MutationRate:    0.1,
		CrossOverRate:   0.5,
		Elitism:         true,
	}

	validConfig := Config{
		NumRuns:         10,
		SelectionMethod: Roulette,
		PopulationSize:  50,
		MaxGenerations:  300,
		NumQueens:       22,
		MutationRate:    0.01,
		CrossOverRate:   0.1,
		Elitism:         true,
	}

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{"Empty path", args{path: ""}, defaultConfig},
		{"Invalid path", args{path: "invalid"}, defaultConfig},
		{"Valid config", args{path: "valid.json"}, validConfig},
		{"Missing field", args{path: "invalid_missing_field.json"}, defaultConfig},
		{"Invalid field type", args{path: "invalid_field_type.json"}, defaultConfig},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.path = filepath.Join("testdata", tt.args.path)

			if got := LoadConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

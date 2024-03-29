package config

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	defaultConfig := Config{
		NumRuns:         5,
		SelectionMethod: Tournament,
		PopulationSize:  16,
		MaxGenerations:  500,
		NumQueens:       8,
		MutationRate:    0.2,
		CrossOverRate:   0.5,
		Elitism:         false,
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
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{"Empty path", args{path: ""}, defaultConfig, false},
		{"Invalid path", args{path: "invalid"}, defaultConfig, false},
		{"Valid config", args{path: "valid.json"}, validConfig, false},
		{"Missing field", args{path: "invalid_missing_field.json"}, Config{}, true},
		{"Invalid field type", args{path: "invalid_field_type.json"}, Config{}, true},
		{"Invalid field value", args{path: "invalid_value.json"}, Config{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.path = filepath.Join("testdata", tt.args.path)
			got, err := LoadConfigFromJSON(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

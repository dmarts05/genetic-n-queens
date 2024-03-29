package result

import "testing"

func TestGetBestFitness(t *testing.T) {
	type args struct {
		results []GenerationResult
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test multiple results",
			args: args{
				results: []GenerationResult{
					{
						BestFitness: 10,
					},
					{
						BestFitness: 20,
					},
					{
						BestFitness: 30,
					},
				},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBestFitness(tt.args.results); got != tt.want {
				t.Errorf("GetBestFitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWorstFitness(t *testing.T) {
	type args struct {
		results []GenerationResult
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test multiple results",
			args: args{
				results: []GenerationResult{
					{
						BestFitness: 10,
					},
					{
						BestFitness: 20,
					},
					{
						BestFitness: 30,
					},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWorstFitness(tt.args.results); got != tt.want {
				t.Errorf("GetWorstFitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMeanMeanFitness(t *testing.T) {
	type args struct {
		results []GenerationResult
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test multiple results",
			args: args{
				results: []GenerationResult{
					{
						MeanFitness: 10,
					},
					{
						MeanFitness: 20,
					},
					{
						MeanFitness: 30,
					},
				},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMeanMeanFitness(tt.args.results); got != tt.want {
				t.Errorf("GetMeanMeanFitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMeanBestFitness(t *testing.T) {
	type args struct {
		results []GenerationResult
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test multiple results",
			args: args{
				results: []GenerationResult{
					{
						BestFitness: 10,
					},
					{
						BestFitness: 20,
					},
					{
						BestFitness: 30,
					},
				},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMeanBestFitness(tt.args.results); got != tt.want {
				t.Errorf("GetMeanBestFitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMeanGenerations(t *testing.T) {
	type args struct {
		results []GenerationResult
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test multiple results",
			args: args{
				results: []GenerationResult{
					{
						Generation: 10,
					},
					{
						Generation: 20,
					},
					{
						Generation: 30,
					},
				},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMeanGenerations(tt.args.results); got != tt.want {
				t.Errorf("GetMeanGenerations() = %v, want %v", got, tt.want)
			}
		})
	}
}

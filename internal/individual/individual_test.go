package individual

import (
	"testing"

	"github.com/dmarts05/genetic-n-queens/internal/position"
)

func Test_areQueensAttacking(t *testing.T) {
	type args struct {
		pos1 position.Position
		pos2 position.Position
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same row",
			args: args{
				pos1: position.Position{Row: 0, Column: 0},
				pos2: position.Position{Row: 0, Column: 1},
			},
			want: true,
		},
		{
			name: "Same column",
			args: args{
				pos1: position.Position{Row: 0, Column: 0},
				pos2: position.Position{Row: 1, Column: 0},
			},
			want: true,
		},
		{
			name: "Same diagonal",
			args: args{
				pos1: position.Position{Row: 0, Column: 0},
				pos2: position.Position{Row: 1, Column: 1},
			},
			want: true,
		},
		{
			name: "Non-attacking",
			args: args{
				pos1: position.Position{Row: 0, Column: 0},
				pos2: position.Position{Row: 1, Column: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areQueensAttacking(tt.args.pos1, tt.args.pos2); got != tt.want {
				t.Errorf("areQueensAttacking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndividual_Fitness(t *testing.T) {
	attackingQueenPositions := []position.Position{
		{Row: 0, Column: 0},
		{Row: 1, Column: 1},
		{Row: 2, Column: 2},
		{Row: 3, Column: 3},
	}

	twoClashQueenPositions := []position.Position{
		{Row: 0, Column: 0},
		{Row: 1, Column: 0},
		{Row: 2, Column: 4},
		{Row: 3, Column: 7},
		{Row: 4, Column: 1},
		{Row: 5, Column: 3},
		{Row: 6, Column: 5},
		{Row: 7, Column: 2},
	}

	nonAttackingQueenPositions := []position.Position{
		{Row: 0, Column: 0},
		{Row: 1, Column: 6},
		{Row: 2, Column: 4},
		{Row: 3, Column: 7},
		{Row: 4, Column: 1},
		{Row: 5, Column: 3},
		{Row: 6, Column: 5},
		{Row: 7, Column: 2},
	}

	type fields struct {
		QueenPositions []position.Position
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "All Queens Attacking",
			fields: fields{
				QueenPositions: attackingQueenPositions,
			},
			want: 0,
		},
		{
			name: "2 Clash Board",
			fields: fields{
				QueenPositions: twoClashQueenPositions,
			},
			want: 26,
		},
		{
			name: "All Queens Non-Attacking",
			fields: fields{
				QueenPositions: nonAttackingQueenPositions,
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ind := New(tt.fields.QueenPositions)
			if got := ind.Fitness(); got != tt.want {
				t.Errorf("Individual.Fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

package individual

import (
	"reflect"
	"testing"
)

func Test_areQueensAttacking(t *testing.T) {
	type args struct {
		pos1 Position
		pos2 Position
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same row",
			args: args{
				pos1: Position{0, 0},
				pos2: Position{0, 1},
			},
			want: true,
		},
		{
			name: "Same column",
			args: args{
				pos1: Position{0, 0},
				pos2: Position{1, 0},
			},
			want: true,
		},
		{
			name: "Same diagonal",
			args: args{
				pos1: Position{0, 0},
				pos2: Position{1, 1},
			},
			want: true,
		},
		{
			name: "Non-attacking",
			args: args{
				pos1: Position{0, 0},
				pos2: Position{1, 2},
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

func TestIndividual_getQueenPositions(t *testing.T) {
	ind := &Individual{
		Board: [][]bool{
			{true, false, false, false},
			{false, true, false, false},
			{false, false, true, false},
			{false, false, false, true},
		},
	}
	queenPositions := ind.getQueenPositions()
	expected := []Position{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
	}
	if !reflect.DeepEqual(queenPositions, expected) {
		t.Errorf("Expected %v, got %v", expected, queenPositions)
	}
}

func TestIndividual_Fitness(t *testing.T) {
	attackingQueensBoard := [][]bool{
		{true, false, false, false},
		{false, true, false, false},
		{false, false, true, false},
		{false, false, false, true},
	}

	twoClashBoard := [][]bool{
		{true, false, false, false, false, false, false, false},
		{true, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, true},
		{false, true, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false},
		{false, false, false, false, false, true, false, false},
		{false, false, true, false, false, false, false, false},
	}

	nonAttackingQueensBoard := [][]bool{
		{true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, true},
		{false, true, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false},
		{false, false, false, false, false, true, false, false},
		{false, false, true, false, false, false, false, false},
	}

	type fields struct {
		Board [][]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "All Queens Attacking",
			fields: fields{
				Board: attackingQueensBoard,
			},
			want: 0,
		},
		{
			name: "2 Clash Board",
			fields: fields{
				Board: twoClashBoard,
			},
			want: 26,
		},
		{
			name: "All Queens Non-Attacking",
			fields: fields{
				Board: nonAttackingQueensBoard,
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ind := &Individual{
				Board: tt.fields.Board,
			}
			if got := ind.Fitness(); got != tt.want {
				t.Errorf("Individual.Fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

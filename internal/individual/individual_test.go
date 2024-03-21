package individual

import (
	"testing"
)

func TestIndividual_Fitness(t *testing.T) {
	attackingQueenPositions := []int{0, 1, 2, 3}
	twoClashQueenPositions := []int{5, 2, 4, 6, 0, 3, 7, 1}
	nonAttackingQueenPositions := []int{0, 6, 4, 7, 1, 3, 5, 2}

	type fields struct {
		QueenPositions []int
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
			ind := Individual{
				QueenPositions: tt.fields.QueenPositions,
			}
			if got := ind.Fitness(); got != tt.want {
				t.Errorf("Individual.Fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}

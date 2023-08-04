package rb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Insert(t *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		pairs []Pair
	}
	type order struct {
		pairs []Pair
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		depth  int
		order  order
	}{
		{
			name: "left balance",
			args: args{
				pairs: []Pair{
					{10, "ten"},
					{9, "nine"},
					{8, "eight"},
					{7, "seven"},
					{6, "six"},
				},
			},
			depth: 3,
			order: order{
				pairs: []Pair{
					{6, "six"},
					{7, "seven"},
					{8, "eight"},
					{9, "nine"},
					{10, "ten"},
				},
			},
		},
		{
			name: "right balance",
			args: args{
				pairs: []Pair{
					{6, "six"},
					{7, "seven"},
					{8, "eight"},
					{9, "nine"},
					{10, "ten"},
				},
			},
			depth: 3,
			order: order{
				pairs: []Pair{
					{6, "six"},
					{7, "seven"},
					{8, "eight"},
					{9, "nine"},
					{10, "ten"},
				},
			},
		},
		{
			name: "balance",
			args: args{
				pairs: []Pair{
					{7, "seven"},
					{6, "six"},
					{8, "eight"},
					{10, "ten"},
					{9, "nine"},
				},
			},
			depth: 3,
			order: order{
				pairs: []Pair{
					{6, "six"},
					{7, "seven"},
					{8, "eight"},
					{9, "nine"},
					{10, "ten"},
				},
			},
		},
		{
			name: "balance two",
			args: args{
				pairs: []Pair{
					{9, "nine"},
					{6, "six"},
					{10, "ten"},
					{8, "eight"},
					{7, "seven"},
				},
			},
			depth: 3,
			order: order{
				pairs: []Pair{
					{6, "six"},
					{7, "seven"},
					{8, "eight"},
					{9, "nine"},
					{10, "ten"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				root: tt.fields.root,
			}
			for _, p := range tt.args.pairs {
				tr.Insert(p)
			}
			assert.Equal(t, tt.depth, depth(tr.root, 0))
			assert.Equal(t, tt.order.pairs, inorderTraversal(tr.root, []Pair{}))
		})
	}
}

package rb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Insert(t *testing.T) {
	type fields struct {
		root *Node[int, string]
	}
	type args struct {
		pairs []Pair[int, string]
	}
	type order struct {
		pairs []Pair[int, string]
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
				pairs: []Pair[int, string]{
					{10, "ten"},
					{9, "nine"},
					{8, "eight"},
					{7, "seven"},
					{6, "six"},
				},
			},
			depth: 3,
			order: order{
				pairs: []Pair[int, string]{
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
				pairs: []Pair[int, string]{
					{6, "six"},
					{7, "seven"},
					{8, "eight"},
					{9, "nine"},
					{10, "ten"},
				},
			},
			depth: 3,
			order: order{
				pairs: []Pair[int, string]{
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
				pairs: []Pair[int, string]{
					{7, "seven"},
					{6, "six"},
					{8, "eight"},
					{10, "ten"},
					{9, "nine"},
				},
			},
			depth: 3,
			order: order{
				pairs: []Pair[int, string]{
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
				pairs: []Pair[int, string]{
					{9, "nine"},
					{6, "six"},
					{10, "ten"},
					{8, "eight"},
					{7, "seven"},
				},
			},
			depth: 3,
			order: order{
				pairs: []Pair[int, string]{
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
			tr := &Tree[int, string]{
				root: tt.fields.root,
			}
			for _, p := range tt.args.pairs {
				tr.Insert(p)
			}
			assert.Equal(t, tt.depth, depth(tr.root, 0))
			assert.Equal(t, tt.order.pairs, inorderTraversal(tr.root, []Pair[int, string]{}))
		})
	}
}

func TestTree_Search(t *testing.T) {
	type fields struct {
		root *Node[int, string]
	}
	type args struct {
		k int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		{
			name: "found",
			fields: fields{
				root: &Node[int, string]{
					Pair: Pair[int, string]{Key: 8, Value: "eight"},
					left: &Node[int, string]{
						Pair: Pair[int, string]{Key: 5, Value: "five"},
						right: &Node[int, string]{
							Pair: Pair[int, string]{Key: 7, Value: "seven"},
						},
					},
					right: &Node[int, string]{
						Pair: Pair[int, string]{Key: 10, Value: "ten"},
					},
				},
			},
			args: args{
				k: 7,
			},
			want:  "seven",
			want1: true,
		},
		{
			name: "not found",
			fields: fields{
				root: &Node[int, string]{
					Pair: Pair[int, string]{Key: 8, Value: "eight"},
					left: &Node[int, string]{
						Pair: Pair[int, string]{Key: 5, Value: "five"},
						right: &Node[int, string]{
							Pair: Pair[int, string]{Key: 7, Value: "seven"},
						},
					},
					right: &Node[int, string]{
						Pair: Pair[int, string]{Key: 10, Value: "ten"},
					},
				},
			},
			args: args{
				k: 40,
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree[int, string]{
				root: tt.fields.root,
			}
			got, got1 := tr.Search(tt.args.k)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

package rb

import (
	"testing"
)

func Test_rotateLeft(t *testing.T) {
	type args struct {
		n *Node[int, string]
	}
	tests := []struct {
		name string
		args args
		want *Node[int, string]
	}{
		{
			name: "rotate no parent",
			args: args{
				n: &Node[int, string]{
					Pair: Pair[int, string]{Key: 3, Value: "three"},
					right: &Node[int, string]{
						Pair: Pair[int, string]{Key: 21, Value: "twenty-one"},
						right: &Node[int, string]{
							Pair: Pair[int, string]{Key: 32, Value: "thirty-two"},
						},
					},
				},
			},
			want: &Node[int, string]{
				Pair: Pair[int, string]{Key: 21, Value: "twenty-one"},
				right: &Node[int, string]{
					Pair: Pair[int, string]{Key: 32, Value: "thirty-two"},
				},
				left: &Node[int, string]{
					Pair: Pair[int, string]{Key: 3, Value: "three"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateLeft(tt.args.n)
			equalNode(t, tt.want, got)
		})
	}
}

func Test_rotateRight(t *testing.T) {
	type args struct {
		n *Node[int, string]
	}
	tests := []struct {
		name string
		args args
		want *Node[int, string]
	}{
		{
			name: "rotate no parent",
			args: args{
				n: &Node[int, string]{
					Pair: Pair[int, string]{Key: 3, Value: "three"},
					left: &Node[int, string]{
						Pair: Pair[int, string]{Key: 21, Value: "twenty-one"},
						left: &Node[int, string]{
							Pair: Pair[int, string]{Key: 32, Value: "thirty-two"},
						},
					},
				},
			},
			want: &Node[int, string]{
				Pair: Pair[int, string]{Key: 21, Value: "twenty-one"},
				right: &Node[int, string]{
					Pair: Pair[int, string]{Key: 3, Value: "three"},
				},
				left: &Node[int, string]{
					Pair: Pair[int, string]{Key: 32, Value: "thirty-two"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight(tt.args.n)
			equalNode(t, tt.want, got)
		})
	}
}

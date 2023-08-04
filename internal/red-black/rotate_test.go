package rb

import (
	"testing"
)

func Test_rotateLeft(t *testing.T) {
	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "rotate no parent",
			args: args{
				n: &Node{
					Pair: Pair{Key: 3, Value: "three"},
					right: &Node{
						Pair: Pair{Key: 21, Value: "twenty-one"},
						right: &Node{
							Pair: Pair{Key: 32, Value: "thirty-two"},
						},
					},
				},
			},
			want: &Node{
				Pair: Pair{Key: 21, Value: "twenty-one"},
				right: &Node{
					Pair: Pair{Key: 32, Value: "thirty-two"},
				},
				left: &Node{
					Pair: Pair{Key: 3, Value: "three"},
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
		n *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "rotate no parent",
			args: args{
				n: &Node{
					Pair: Pair{Key: 3, Value: "three"},
					left: &Node{
						Pair: Pair{Key: 21, Value: "twenty-one"},
						left: &Node{
							Pair: Pair{Key: 32, Value: "thirty-two"},
						},
					},
				},
			},
			want: &Node{
				Pair: Pair{Key: 21, Value: "twenty-one"},
				right: &Node{
					Pair: Pair{Key: 3, Value: "three"},
				},
				left: &Node{
					Pair: Pair{Key: 32, Value: "thirty-two"},
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

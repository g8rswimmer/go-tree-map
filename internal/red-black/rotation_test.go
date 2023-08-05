package rb

import (
	"testing"
)

func Test_rotation_rotate(t *testing.T) {
	type fields struct {
		left      bool
		right     bool
		leftRight bool
		rightLeft bool
	}
	type args struct {
		n *Node[int, string]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node[int, string]
	}{
		{
			name: "rotate right red",
			fields: fields{
				right: true,
			},
			args: args{
				n: &Node[int, string]{
					Pair:  Pair[int, string]{Key: 3, Value: "three"},
					black: true,
					left: &Node[int, string]{
						Pair: Pair[int, string]{Key: 21, Value: "twenty-one"},
						left: &Node[int, string]{
							Pair: Pair[int, string]{Key: 32, Value: "thirty-two"},
						},
					},
				},
			},
			want: &Node[int, string]{
				Pair:  Pair[int, string]{Key: 21, Value: "twenty-one"},
				black: true,
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
			r := &rotation[int, string]{
				left:      tt.fields.left,
				right:     tt.fields.right,
				leftRight: tt.fields.leftRight,
				rightLeft: tt.fields.rightLeft,
			}
			got := r.rotate(tt.args.n)
			equalNode(t, tt.want, got)
		})
	}
}

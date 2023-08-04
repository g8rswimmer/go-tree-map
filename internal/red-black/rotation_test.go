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
		n *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name: "rotate right red",
			fields: fields{
				right: true,
			},
			args: args{
				n: &Node{
					Pair:  Pair{Key: 3, Value: "three"},
					black: true,
					left: &Node{
						Pair: Pair{Key: 21, Value: "twenty-one"},
						left: &Node{
							Pair: Pair{Key: 32, Value: "thirty-two"},
						},
					},
				},
			},
			want: &Node{
				Pair:  Pair{Key: 21, Value: "twenty-one"},
				black: true,
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
			r := &rotation{
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

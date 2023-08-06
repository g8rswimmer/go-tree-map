package tree

import (
	"testing"

	rb "github.com/g8rswimmer/go-tree-map/internal/red-black"
	"github.com/stretchr/testify/assert"
)

func TestCompleteTest_Int_String(t *testing.T) {
	type args struct {
		inputs []rb.Pair[int, string]
		key    int
	}
	type wants struct {
		value string
		size  int
		pairs []rb.Pair[int, string]
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "success",
			args: args{
				inputs: []rb.Pair[int, string]{
					{Key: 9, Value: "nine"},
					{Key: 6, Value: "six"},
					{Key: 10, Value: "ten"},
					{Key: 8, Value: "eight"},
					{Key: 7, Value: "seven"},
				},
				key: 8,
			},
			wants: wants{
				value: "eight",
				size:  5,
				pairs: []rb.Pair[int, string]{
					{Key: 6, Value: "six"},
					{Key: 7, Value: "seven"},
					{Key: 8, Value: "eight"},
					{Key: 9, Value: "nine"},
					{Key: 10, Value: "ten"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMap[int, string]()

			for _, p := range tt.args.inputs {
				m.Store(p.Key, p.Value)
			}
			assert.Equal(t, tt.wants.size, m.Len())

			v, ok := m.Load(tt.args.key)
			assert.Equal(t, true, ok)
			assert.Equal(t, tt.wants.value, v)

			ordered := []rb.Pair[int, string]{}
			handler := func(k int, v string) error {
				ordered = append(ordered, rb.Pair[int, string]{Key: k, Value: v})
				return nil
			}
			err := m.Range(handler)
			assert.NoError(t, err)
			assert.Equal(t, tt.wants.pairs, ordered)
		})
	}
}

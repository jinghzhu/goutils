package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testCasesInt64Slice = [][]int64{
	{1, 2, 3, 4},
}

func TestInt64Slice(t *testing.T) {
	for idx, in := range testCasesInt64Slice {
		if in == nil {
			continue
		}
		out := Int64Slice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := Int64ValueSlice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

var testCasesInt64ValSlice = [][]*int64{}

func TestInt64ValSlice(t *testing.T) {
	for idx, in := range testCasesInt64ValSlice {
		if in == nil {
			continue
		}
		out := Int64ValSlice(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			if in[i] == nil {
				assert.Empty(t, out[i], "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, *(in[i]), out[i], "Unexpected value at idx %d", idx)
			}
		}

		out2 := Int64Slice(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		for i := range out2 {
			if in[i] == nil {
				assert.Empty(t, *(out2[i]), "Unexpected value at idx %d", idx)
			} else {
				assert.Equal(t, in[i], out2[i], "Unexpected value at idx %d", idx)
			}
		}
	}
}

var testCasesInt64Map = []map[string]int64{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt64Map(t *testing.T) {
	for idx, in := range testCasesInt64Map {
		if in == nil {
			continue
		}
		out := Int64Map(in)
		assert.Len(t, out, len(in), "Unexpected len at idx %d", idx)
		for i := range out {
			assert.Equal(t, in[i], *(out[i]), "Unexpected value at idx %d", idx)
		}

		out2 := Int64ValueMap(out)
		assert.Len(t, out2, len(in), "Unexpected len at idx %d", idx)
		assert.Equal(t, in, out2, "Unexpected value at idx %d", idx)
	}
}

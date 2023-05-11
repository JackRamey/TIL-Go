package zeros

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

var _ = []struct {
	name string
	in   []int
	want []int
}{
	{
		name: "",
		in:   []int{0, 0, 0, 3, 1, 4, 1, 5, 9, 0, 0, 0, 0},
		want: []int{3, 1, 4, 1, 5, 9},
	},
	{
		name: "",
		in:   []int{0, 0, 0},
		want: []int{},
	},
	{
		name: "",
		in:   []int{8},
		want: []int{8},
	},
}

func Test_removeZeros(t *testing.T) {
	d := readTestData()
	for _, tt := range d.Puzzles {
		t.Run("", func(t *testing.T) {
			if gotOut := removeZeros(tt.Problem); !reflect.DeepEqual(gotOut, tt.Solution) {
				t.Errorf("removeZeros() = %v, want %v", gotOut, tt.Solution)
			}
		})
	}
}

func Test_removeZerosCursors(t *testing.T) {
	d := readTestData()
	for _, tt := range d.Puzzles {
		t.Run("", func(t *testing.T) {
			if got := removeZerosCursors(tt.Problem); !reflect.DeepEqual(got, tt.Solution) {
				t.Errorf("removeZerosCursors() = %v, want %v", got, tt.Solution)
			}
		})
	}
}

func Test_both(t *testing.T) {
	d := readTestData()
	for _, tt := range d.Puzzles {
		t.Run("", func(t *testing.T) {
			if got := removeZerosCursors(tt.Problem); !reflect.DeepEqual(got, tt.Solution) {
				t.Errorf("removeZerosCursors() = %v, want %v", got, tt.Solution)
			}
		})
	}
	for _, tt := range d.Puzzles {
		t.Run("", func(t *testing.T) {
			if got := removeZeros(tt.Problem); !reflect.DeepEqual(got, tt.Solution) {
				t.Errorf("removeZeros() = %v, want %v", got, tt.Solution)
			}
		})
	}
}

func readTestData() data {
	var d data
	err := json.Unmarshal(must(os.ReadFile("data.json")), &d)
	if err != nil {
		panic(err)
	}

	return d
}

func Test_genTestData(t *testing.T) {
	dataBytes, _ := json.Marshal(genTestData(10, 10000000))
	err := os.WriteFile("data.json", dataBytes, 0777)
	if err != nil {
		panic(err)
	}
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

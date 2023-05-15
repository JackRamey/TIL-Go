package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_valRef(t *testing.T) {
	t.Run("valRef works for primitives", func(t *testing.T) {
		expected := 5
		actual := valRef(5)
		assert.Equal(t, &expected, actual)
	})
	t.Run("valRef works for non-primitive types", func(t *testing.T) {
		type a struct {
			s string
		}
		v := a{
			s: "foo",
		}
		actual := valRef(v)
		assert.Equal(t, &v, actual)
	})
}

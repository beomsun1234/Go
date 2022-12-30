package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Hello_GetHello(t *testing.T) {
	t.Run("hello를 리턴한다", func(t *testing.T) {
		h := NewHello()
		ret := h.GetHello()

		assert.Equal(t, ret, "hello")
	})
}

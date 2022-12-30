package hello

import (
	"github.com/stretchr/testify/assert"
	"go/Go/go_test/mocks"
	"testing"
)

/*
mock 주입
*/
func Test_Hello_GetHello(t *testing.T) {
	t.Run("hello를 리턴한다", func(t *testing.T) {
		//given
		sub := new(mocks.MockHelloDependency)
		h := NewHello(sub)
		sub.On("GetHelloDependency").Return("hello")
		//when
		ret := h.GetHello()
		//then
		assert.Equal(t, ret, "hello")
	})
}

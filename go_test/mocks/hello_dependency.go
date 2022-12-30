package mocks

import "github.com/stretchr/testify/mock"

type MockHelloDependency struct {
	mock.Mock
}

func (m *MockHelloDependency) GetHelloDependency() string {
	ret := m.Mock.Called()

	res := ret.String(0)

	return res
}

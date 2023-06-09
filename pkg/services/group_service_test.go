package services

import (
	"context"
	"docker-go-project/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	repository *mocks.IRepository
}

type airMocks struct {
	service func(f *mockService)
}

func Test_GetAirActual(t *testing.T) {
	tests := []struct {
		name   string
		mocks  airMocks
		airExp string
		expErr error
	}{
		{
			name: "error flow",
			mocks: airMocks{
				service: func(f *mockService) {
					f.repository.Mock.On("Get", mock.Anything).Return("", errors.New("Error PAPI"))
				},
			},
			expErr: errors.New("Error PAPI"),
		},
		{
			name: "full flow",
			mocks: airMocks{
				service: func(f *mockService) {
					f.repository.Mock.On("Get", mock.Anything).Return("Hi PAPI", nil)
				},
			},
			airExp: "Hi PAPI",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockService{
				repository: &mocks.IRepository{},
			}
			tc.mocks.service(m)
			service := NewAirService(m.repository)
			air, err := service.GetAirActual(context.Background())
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.airExp, air)
		})
	}
}

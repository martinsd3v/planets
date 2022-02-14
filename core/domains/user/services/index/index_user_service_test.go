package index

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/martinsd3v/planets/core/domains/user/entities"
	"github.com/martinsd3v/planets/core/tools/communication"
	"github.com/martinsd3v/planets/mocks"
)

func TestService(t *testing.T) {
	comm := communication.New()

	expectedData := entities.Users{{
		UUID:  "userID",
		Name:  "userName",
		Email: "userEmail",
	}}

	useCases := map[string]struct {
		expectedResponse communication.Response
		expectedData     entities.Users
		prepare          func(repostitoryMock *mocks.MockIUserRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"success": {
			expectedData: expectedData,
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success"].Code,
				Message: comm.Mapping["success"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIUserRepository, loggerMock *mocks.MockILoggerProvider) {
				repostitoryMock.EXPECT().All(gomock.Any()).Return(expectedData, nil)
			},
		},
		"error: on repository": {
			expectedResponse: communication.Response{
				Status:  404,
				Code:    comm.Mapping["error_list"].Code,
				Message: comm.Mapping["error_list"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIUserRepository, loggerMock *mocks.MockILoggerProvider) {
				repostitoryMock.EXPECT().All(gomock.Any()).Return(entities.Users{}, errors.New("error"))
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any())
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ctx := context.Background()
			defer ctrl.Finish()

			repository := mocks.NewMockIUserRepository(ctrl)
			logger := mocks.NewMockILoggerProvider(ctrl)
			useCase.prepare(repository, logger)

			service := Service{
				Repository: repository,
				Logger:     logger,
			}
			data, response := service.Execute(ctx)

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("Expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("Expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}

			if len(data) != len(useCase.expectedData) {
				t.Errorf("Expected %s, but got %s", useCase.expectedData, data)
			}
		})
	}
}

package films

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/martinsd3v/planets/mocks"
)

func TestService(t *testing.T) {
	useCases := map[string]struct {
		expectedData int
		inputData    string
		prepare      func(loggerMock *mocks.MockILoggerProvider, clientMock *mocks.MockIHTTPClientProvider)
	}{
		"success": {
			inputData:    "planetName",
			expectedData: 2,
			prepare: func(loggerMock *mocks.MockILoggerProvider, clientMock *mocks.MockIHTTPClientProvider) {
				bodyResponse := ioutil.NopCloser(bytes.NewReader([]byte(`{"results": [{"films": ["film1", "film2"]}]}`)))
				clientMock.EXPECT().Get(gomock.Any()).Times(1).Return(&http.Response{Body: bodyResponse}, nil)
			},
		},
		"error: on Client Request": {
			inputData: "planetName",
			prepare: func(loggerMock *mocks.MockILoggerProvider, clientMock *mocks.MockIHTTPClientProvider) {
				clientMock.EXPECT().Get(gomock.Any()).Times(1).Return(nil, errors.New("error"))
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).Times(1)
			},
		},
		"error: on Response Body": {
			inputData: "planetName",
			prepare: func(loggerMock *mocks.MockILoggerProvider, clientMock *mocks.MockIHTTPClientProvider) {
				bodyResponse := ioutil.NopCloser(bytes.NewReader([]byte(`{invalidJson}`)))
				clientMock.EXPECT().Get(gomock.Any()).Times(1).Return(&http.Response{Body: bodyResponse}, nil)
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).Times(1)
			},
		},
		"success: zero results": {
			inputData: "planetName",
			prepare: func(loggerMock *mocks.MockILoggerProvider, clientMock *mocks.MockIHTTPClientProvider) {
				bodyResponse := ioutil.NopCloser(bytes.NewReader([]byte(`{"results": []}`)))
				clientMock.EXPECT().Get(gomock.Any()).Times(1).Return(&http.Response{Body: bodyResponse}, nil)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			logger := mocks.NewMockILoggerProvider(ctrl)
			client := mocks.NewMockIHTTPClientProvider(ctrl)
			useCase.prepare(logger, client)

			service := Service{
				Logger:     logger,
				HTTPClient: client,
			}
			data := service.Execute(useCase.inputData)

			if data != useCase.expectedData {
				t.Errorf("Expected %q, but got %q", useCase.expectedData, data)
			}
		})
	}
}

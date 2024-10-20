package main

import (
	"github.com/bmsandoval/counter/src/services"
	"github.com/bmsandoval/counter/src/services/mocks/Services_mocks"
	"github.com/bmsandoval/counter/src/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventStruct_ModifyCounter(t *testing.T) {
	unitTestData := []struct {
		Description string
		FilePath    string
		ResultPath  string
		Initial     int
		Increment   int
		Expect      *error
	}{
		{
			Description: "test happy path",
			FilePath:    "~/file.path",
			ResultPath:  "~/file.path",
			Initial:     0,
			Increment:   1,
			Expect:      nil,
		},
		{
			Description: "test default path",
			FilePath:    "",
			ResultPath:  "~/counter.txt",
			Initial:     0,
			Increment:   1,
			Expect:      nil,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			counterSvcMock := Services_mocks.NewMockCounterService(mockCtrl)
			counterSvcExpect := counterSvcMock.EXPECT()
			configSvcMock := Services_mocks.NewMockConfigService(mockCtrl)
			configSvcExpect := configSvcMock.EXPECT()

			configSvcExpect.LoadConfig(gomock.Any()).Return(types.Config{}, nil)
			counterSvcExpect.GetCounterPath(gomock.Any()).Return(testData.FilePath, nil)
			counterSvcExpect.GetCounter(gomock.Any()).Return(testData.Initial, nil)
			counterSvcExpect.SetCounter(gomock.Any(), gomock.Any())

			svcBundle := services.Bundle{CounterService: counterSvcMock, ConfigService: configSvcMock}

			eventHandler := EventStruct{Services: svcBundle}
			err := eventHandler.ModifyCounter(testData.Increment)
			if testData.Expect == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err, *testData.Expect)
			}
		})
	}
}

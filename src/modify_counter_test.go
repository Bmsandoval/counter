package main

import (
	"github.com/bmsandoval/counter/src/services"
	"github.com/bmsandoval/counter/src/services/mocks/Services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventStruct_ModifyCounter(t *testing.T) {
	unitTestData := []struct {
		Description string
		FilePath    string
		Initial     int
		Increment   int
		Expect      *error
	}{
		{
			Description: "can detect leading zeros in octet values",
			FilePath:    "~/file.path",
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

			counterSvcExpect.GetCounterPath(gomock.Any(), gomock.Any()).Return(testData.FilePath, nil)
			counterSvcExpect.GetCounter(testData.FilePath).Return(testData.Initial, nil)
			counterSvcExpect.SetCounter(testData.FilePath, testData.Initial+testData.Increment)

			svcBundle := services.Bundle{CounterService: counterSvcMock}

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

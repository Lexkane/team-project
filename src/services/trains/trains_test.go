package trains_test

import (
	"bytes"
	"encoding/json"
	"../../entity"
	"../trains"
	"github.com/Lexkane/TripAgent/src/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = services.NewRouter()

type TrainsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetTrains  []entity.Train
	testDataId       string
	testDataTr       string
	mockedTrainError error
}

func TestGetTrains(t *testing.T) {
	tests := []TrainsTestCase{
		{
			name: "Get_Trains_200",
			url:  "/v1/trains?departure_city=lviv&arrival_city=kyiv&price=300uah",
			want: 200,
		},
		{
			name:             "Get_Trains_404",
			url:              "/v1/trains?mcok=890",
			want:             404,
			mockedGetTrains:  []entity.Train{},
			mockedTrainError: http.ErrLineTooLong,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			entity.MockedGetData(tc.mockedGetTrains, tc.mockedTrainError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestSaveTrain(t *testing.T) {
	tests := []TrainsTestCase{
		{
			name:             "Add_To_Trip_201",
			url:              "/v1/trains",
			want:             201,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataTr:       "00000000-0000-0000-0000-000000000001",
			mockedGetTrains:  []entity.Train{},
			mockedTrainError: nil,
		},
		{
			name:             "Add_Trains_400",
			url:              "/v1/trains",
			want:             400,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataTr:       "asdsad",
			mockedGetTrains:  []entity.Train{},
			mockedTrainError: nil,
		},
		{
			name:             "Add_Trains_400_2",
			url:              "/v1/trains",
			want:             400,
			testDataId:       "asdasd",
			testDataTr:       "00000000-0000-0000-0000-000000000001",
			mockedGetTrains:  []entity.Train{},
			mockedTrainError: nil,
		},
		{
			name:             "Add_Events_400_3",
			url:              "/v1/trains",
			want:             400,
			testDataId:       "00000000-0000-0000-0000-000000000001",
			testDataTr:       "00000000-0000-0000-0000-000000000001",
			mockedTrainError: error(new(http.ProtocolError)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			var data trains.TripTrain
			data.TrainID = tc.testDataId
			data.TripID = tc.testDataTr
			body, _ := json.Marshal(data)

			entity.MockedAddToTrip(tc.mockedTrainError)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})

	}
}

func TestGetTrainFromTrip(t *testing.T) {
	tests := []TrainsTestCase{
		{
			name:            "Get_train_200",
			url:             "/v1/trains/trip/00000000-0000-0000-0000-000000000001",
			want:            200,
			mockedGetTrains: []entity.Train{},
		},
		{
			name:            "Get_Trains_400",
			url:             "/v1/trains/trip/sadsad",
			want:            400,
			mockedGetTrains: []entity.Train{},
		},
		{
			name:             "Get_Events_404",
			url:              "/v1/trains/trip/00000000-0000-0000-0000-000000000009",
			want:             404,
			mockedGetTrains:  []entity.Train{},
			mockedTrainError: http.ErrNoLocation,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			entity.MockedGetByTrip(tc.mockedGetTrains, tc.mockedTrainError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

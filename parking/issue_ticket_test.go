package parking

import (
	"reflect"
	"testing"
)

func Test_car_createTicket(t *testing.T) {
	type fields struct {
		registrationNumber string
		colour             string
	}
	tests := []struct {
		before  func()
		name    string
		fields  fields
		want    ticket
		wantErr bool
		after   func()
	}{
		{
			before: func() {
				CreateParkingLot(6)
				Park("KA-01-HH-1234", "White")

			},
			name: "success case",
			fields: fields{
				registrationNumber: "KA-01-HH-1234",
				colour:             "White",
			},
			want: ticket{
				car: car{
					registrationNumber: "KA-01-HH-1234",
					colour:             "White",
				},
				parkingSlotNumber: 2,
			},
			after: func() {
				parking = []parkingSlot{}
			},
		},
		{
			before: func() {},
			name:   "failed case",
			fields: fields{
				registrationNumber: "KA-01-HH-1234",
				colour:             "White",
			},
			wantErr: true,
			after: func() {
			},
		},
	}
	for _, tt := range tests {
		tt.before()
		t.Run(tt.name, func(t *testing.T) {
			car := car{
				registrationNumber: tt.fields.registrationNumber,
				colour:             tt.fields.colour,
			}
			got, err := car.createTicket()
			if (err != nil) != tt.wantErr {
				t.Errorf("car.createTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("car.createTicket() = %v, want %v", got, tt.want)
			}
		})
		tt.after()
	}
}

func Test_searchNearestSlot(t *testing.T) {
	tests := []struct {
		before func()
		name   string
		want   int
		after  func()
	}{
		{
			before: func() {
				CreateParkingLot(6)

			},
			name: "success case",
			want: 1,
			after: func() {
				parking = []parkingSlot{}
			},
		},
		{
			before: func() {},
			name:   "failed case",
			want:   -1,
			after: func() {
			},
		},
	}
	for _, tt := range tests {
		tt.before()
		t.Run(tt.name, func(t *testing.T) {
			if got := searchNearestSlot(); got != tt.want {
				t.Errorf("searchNearestSlot() = %v, want %v", got, tt.want)
			}
		})
		tt.after()
	}
}

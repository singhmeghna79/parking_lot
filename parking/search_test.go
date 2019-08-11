package parking

import "testing"

func TestRegistrationNumbersForCarsWithColour(t *testing.T) {
	type args struct {
		colour string
	}
	tests := []struct {
		before func()
		name   string
		args   args
		want   string
		after  func()
	}{
		{
			before: func() {
				CreateParkingLot(6)
				Park("KA-01-HH-1234", "White")
				Park("KA-01-HH-9999", "Black")
				Park("KA-01-HH-3333", "White")
			},
			name: "success case",
			args: args{
				colour: "White",
			},
			want: "KA-01-HH-1234, KA-01-HH-3333",
			after: func() {
				parking = []parkingSlot{}
			},
		},
		{
			before: func() {},
			name:   "failed case",
			args: args{
				colour: "White",
			},
			want: "Not found",
			after: func() {
			},
		},
	}
	for _, tt := range tests {
		tt.before()
		t.Run(tt.name, func(t *testing.T) {
			if got := RegistrationNumbersForCarsWithColour(tt.args.colour); got != tt.want {
				t.Errorf("RegistrationNumbersForCarsWithColour() = %v, want %v", got, tt.want)
			}
		})
		tt.after()
	}
}

func TestSlotNumbersForCarsWithColour(t *testing.T) {
	type args struct {
		colour string
	}
	tests := []struct {
		before func()
		name   string
		args   args
		want   string
		after  func()
	}{
		{
			before: func() {
				CreateParkingLot(6)
				Park("KA-01-HH-1234", "White")
				Park("KA-01-HH-9999", "Black")
				Park("KA-01-HH-3333", "White")
			},
			name: "success case",
			args: args{
				colour: "White",
			},
			want: "1, 3",
			after: func() {
				parking = []parkingSlot{}
			},
		},
		{
			before: func() {},
			name:   "failed case",
			args: args{
				colour: "White",
			},
			want: "Not found",
			after: func() {
			},
		},
	}
	for _, tt := range tests {
		tt.before()
		t.Run(tt.name, func(t *testing.T) {
			if got := SlotNumbersForCarsWithColour(tt.args.colour); got != tt.want {
				t.Errorf("SlotNumbersForCarsWithColour() = %v, want %v", got, tt.want)
			}
		})
		tt.after()
	}
}

func TestSlotNumberForRegistrationNumber(t *testing.T) {
	type args struct {
		registrationNumber string
	}
	tests := []struct {
		before func()
		name   string
		args   args
		want   string
		after  func()
	}{
		{
			before: func() {
				CreateParkingLot(6)
				Park("KA-01-HH-1234", "White")
				Park("KA-01-HH-9999", "Black")
				Park("KA-01-HH-3333", "White")
			},
			name: "success case",
			args: args{
				registrationNumber: "KA-01-HH-1234",
			},
			want: "1",
			after: func() {
				parking = []parkingSlot{}
			},
		},
		{
			before: func() {},
			name:   "failed case",
			args: args{
				registrationNumber: "KA-01-HH-1234",
			},
			want: "Not found",
			after: func() {
			},
		},
	}
	for _, tt := range tests {
		tt.before()
		t.Run(tt.name, func(t *testing.T) {
			if got := SlotNumberForRegistrationNumber(tt.args.registrationNumber); got != tt.want {
				t.Errorf("SlotNumberForRegistrationNumber() = %v, want %v", got, tt.want)
			}
		})
		tt.after()
	}
}

package parking

import (
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	type args struct {
		totalslots int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success case",
			args: args{
				totalslots: 6,
			},
			want: "Created a parking lot with 6 slots",
		},
		{
			name: "failed case",
			args: args{
				totalslots: -1,
			},
			want: "invalid parking slots",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateParkingLot(tt.args.totalslots); got != tt.want {
				t.Errorf("CreateParkingLot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPark(t *testing.T) {
	type args struct {
		regno  string
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
			name: "succes case",
			before: func() {
				CreateParkingLot(6)
			},
			args: args{
				regno:  "KA-01-HH-1234",
				colour: "White",
			},
			want: "Allocated slot number: 1",
			after: func() {
				parking = []parkingSlot{}
			},
		},
		{
			name:   "failed case",
			before: func() {},
			args: args{
				regno:  "KA-01-HH-1234",
				colour: "White",
			},
			want:  "Sorry, parking lot is full",
			after: func() {},
		},
	}
	for _, tt := range tests {
		tt.before()
		t.Run(tt.name, func(t *testing.T) {
			if got := Park(tt.args.regno, tt.args.colour); got != tt.want {
				t.Errorf("Park() = %v, want %v", got, tt.want)
			}
		})
		tt.after()
	}
}

func TestLeave(t *testing.T) {
	type args struct {
		slotno int
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
			},
			name: "success case",
			args: args{
				slotno: 1,
			},
			want: "Slot number 1 is free",
			after: func() {
				parking = []parkingSlot{}
			},
		},
	}
	for _, tt := range tests {
		tt.before()
		t.Run(tt.name, func(t *testing.T) {
			if got := Leave(tt.args.slotno); got != tt.want {
				t.Errorf("Leave() = %v, want %v", got, tt.want)
			}
		})
		tt.after()
	}
}

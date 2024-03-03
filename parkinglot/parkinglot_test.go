package parkinglot

import (
    "testing"
)

func TestParkingLotCreation(t *testing.T) {
    capacity := 6
    pl := NewParkingLot(capacity)
    if pl.Capacity != capacity {
        t.Errorf("expected parking lot capacity of %d; got %d", capacity, pl.Capacity)
    }
}

func TestParkCar(t *testing.T) {
    pl := NewParkingLot(1)
    car := &Car{RegistrationNumber: "KA-01-HH-1234", Color: "White"}
    slot := pl.Park(car)
    if slot != 1 {
        t.Errorf("expected to park car in slot 1; got slot %d", slot)
    }

    // Test parking in a full parking lot
    car2 := &Car{RegistrationNumber: "KA-01-HH-9999", Color: "Black"}
    slot = pl.Park(car2)
    if slot != -1 {
        t.Errorf("expected parking lot to be full; got slot %d", slot)
    }
}

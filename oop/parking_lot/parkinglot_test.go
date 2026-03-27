package parkinglot

import (
	"fmt"
	"testing"
)

func TestParkingLotExample(t *testing.T) {
	lot := NewParkingLot(1, 1, 11)
	// spots_per_row=11: motorcycle[0,2), compact[2,8), large[8,11)

	assert := func(name string, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("%s: got %v, want %v", name, got, want)
		}
	}

	assert("Motorcycle_1", lot.ParkVehicle("Motorcycle_1"), true)
	assert("Car_1", lot.ParkVehicle("Car_1"), true)
	assert("Car_2", lot.ParkVehicle("Car_2"), true)
	assert("Car_3", lot.ParkVehicle("Car_3"), true)
	assert("Car_4", lot.ParkVehicle("Car_4"), true)
	assert("Car_5", lot.ParkVehicle("Car_5"), true)
	assert("Bus_1 (should fail)", lot.ParkVehicle("Bus_1"), false)

	lot.UnParkVehicle("Car_5")
	assert("Bus_1 (after unpark)", lot.ParkVehicle("Bus_1"), true)
}

func TestMultiLevel(t *testing.T) {
	lot := NewParkingLot(2, 1, 4)
	// spots_per_row=4: motorcycle[0,1), compact[1,3), large[3,4)

	assert := func(name string, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("%s: got %v, want %v", name, got, want)
		}
	}

	// Type-first ordering: ALL compact spots across both levels first,
	// then large spots.
	assert("Car_1", lot.ParkVehicle("Car_1"), true)  // compact on level 0
	assert("Car_2", lot.ParkVehicle("Car_2"), true)  // compact on level 0
	assert("Car_3", lot.ParkVehicle("Car_3"), true)  // compact on level 1 (type-first!)
	assert("Car_4", lot.ParkVehicle("Car_4"), true)  // compact on level 1
	assert("Car_5", lot.ParkVehicle("Car_5"), true)  // all compact full → large spot
	assert("Car_6", lot.ParkVehicle("Car_6"), true)  // large spot
	assert("Car_7", lot.ParkVehicle("Car_7"), false) // no more spots
}

func TestMotorcycleCanParkAnywhere(t *testing.T) {
	lot := NewParkingLot(1, 1, 4)
	// spots_per_row=4: motorcycle[0,1), compact[1,3), large[3,4)

	// Fill all 4 spots with motorcycles
	for i := 1; i <= 4; i++ {
		name := "Motorcycle_" + string(rune('0'+i))
		if !lot.ParkVehicle(name) {
			t.Errorf("Motorcycle_%d should park successfully", i)
		}
	}
}

func TestBusNeedsFiveConsecutiveLargeSpots(t *testing.T) {
	lot := NewParkingLot(1, 1, 20)
	// spots_per_row=20: motorcycle[0,5), compact[5,15), large[15,20)
	// 5 large spots: exactly enough for 1 bus

	if !lot.ParkVehicle("Bus_1") {
		t.Error("Bus_1 should park successfully in 5 large spots")
	}
	if lot.ParkVehicle("Bus_2") {
		t.Error("Bus_2 should fail, no more consecutive large spots")
	}
}

func TestAntiFragmentation(t *testing.T) {
	// 1 level, 3 rows, 20 spots per row.
	// Each row: motorcycle[0,5), compact[5,15), large[15,20) → 5 large spots per row.
	// Each row can fit exactly 1 bus (5 consecutive large spots).
	lot := NewParkingLot(1, 3, 20)

	assert := func(name string, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("%s: got %v, want %v", name, got, want)
		}
	}

	// Fill ALL compact spots (10 per row × 3 rows = 30 compact spots).
	for i := 1; i <= 30; i++ {
		assert(fmt.Sprintf("Car_%d", i), lot.ParkVehicle(fmt.Sprintf("Car_%d", i)), true)
	}

	// Now more cars overflow to large spots.
	// Anti-fragmentation: they should concentrate in one row, not spread.
	// Park 5 cars in large spots → should break exactly 1 row.
	for i := 31; i <= 35; i++ {
		assert(fmt.Sprintf("Car_%d", i), lot.ParkVehicle(fmt.Sprintf("Car_%d", i)), true)
	}

	// With anti-fragmentation, the 5 cars should concentrate in 1 row,
	// leaving the other 2 rows with full bus capacity.
	// So we should be able to park 2 buses.
	assert("Bus_1", lot.ParkVehicle("Bus_1"), true)
	assert("Bus_2", lot.ParkVehicle("Bus_2"), true)
}

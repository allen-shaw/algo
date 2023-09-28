package array

import (
	"fmt"
	"testing"
)

func TestCorpFlightBookings(t *testing.T) {
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	n := 5

	res := corpFlightBookings(bookings, n)
	fmt.Println(res)
}

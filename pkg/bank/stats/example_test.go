package stats

import (
	"fmt"

	"github.com/SAE17/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100_00,
			Category: "mobile",
		},
		{
			ID:       2,
			Amount:   5_00,
			Category: "mobile",
		},
		{
			ID:       3,
			Amount:   10_00,
			Category: "mobile",
		},
		{
			ID:       4,
			Amount:   10_00,
			Category: "mobile",
		},
	}
	avg := Avg(payments)
	fmt.Println(avg)
	// Output:
	// 3125

}

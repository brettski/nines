// It's all about the nines

package main

import (
	"fmt"

    termtables	"github.com/brettski/go-termtables"
)

const hoursInYear float32 = 8760.0

var nines = [7]float32{99.0, 99.5, 99.9, 99.95, 99.99, 99.995, 99.999}

func main() {
	fmt.Printf("Hours in year set: %f\n", hoursInYear)
	table := termtables.CreateTable()
	table.AddHeaders("nine", "Year-Hours", "Year-Minutes", "Month-Hours", "Month-Minutes")
	for _, nine := range nines {
		var hours = hoursInYear - (hoursInYear * 0.01 * nine)
		table.AddRow(
			nine,
			hours,
			hours*60,
			hours/12,
			((hours / 12) * 60),
		)
	}

	fmt.Println(table.Render())
	//doTable2()
}

func doTable2() {
	table := termtables.CreateTable()

	table.AddHeaders("Name", "Age")
	table.AddRow("John", "30")
	table.AddRow("Sam", 18)
	table.AddRow("Julie", 20.14)

	fmt.Println(table.Render())

}

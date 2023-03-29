package main

import "fmt"

func TicketPlayground(height, age int) int {
	var hargaBayar int

	if (age > 12) {
		hargaBayar = 100000
	} else if (age >= 12) || height > 160 {
		hargaBayar = 60000
	} else if (age >= 10 && age <= 11) || height > 150 {
		hargaBayar = 40000
	} else if (age >= 8 && age <= 9) || height > 135 {
		hargaBayar = 25000
	} else if (age >= 5 && age <= 7) || height > 120 {
		hargaBayar = 15000
	} else if age < 5 {
		hargaBayar = -1
	}

	return hargaBayar // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(120, 10))
}

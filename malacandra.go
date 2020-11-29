//Malacandra is much nearer than
//that: we shall make it in about twenty-eight days
//C.S. Lewis
package main

import (
	"fmt"
)

func main() {
	const timeInDays = 23
	const hourInDay = 24
	const speed = 10800
	const distance = speed * (timeInDays * hourInDay)
	fmt.Println("Distance: ", distance)
}

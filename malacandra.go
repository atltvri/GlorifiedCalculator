//Malacandra is much nearer than
//that: we shall make it in about twenty-eight days
//C.S. Lewis
package main

import (
	"fmt"
)

func main() {
	const timeInDays = 28
	const hourInDay = 24
	const hourIn28Days=24*28
	const distance=56000000
	const speed = distance/hourIn28Days
	
	fmt.Println(speed,"km/h")
}

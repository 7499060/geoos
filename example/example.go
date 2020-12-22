package main

import (
	"fmt"
	"github.com/spatial-go/geoos"
)

func main() {
	// First, choose the default algorithm.
	strategy := geoos.NormalStrategy()
	// Secondly, manufacturing test data and convert it to geometry
	const wkt = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	geometry, _ := geoos.UnmarshalString(wkt)
	// Last， call the Area () method and get result.
	area, e := strategy.Area(geometry)
	if e != nil {
		fmt.Printf(e.Error())
	}
	fmt.Printf("%f", area)
	// get result 4.0
}

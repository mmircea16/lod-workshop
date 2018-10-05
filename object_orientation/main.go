package main

import "fmt"

type RectangularPlateau struct {
	topRightX int
	topRightY int
}

func (plateau RectangularPlateau) isOnThePlateau(x int, y int) bool {
	return x>=0 && x <= plateau.topRightX && y>= 0 && y <= plateau.topRightY
}

type DonutPlateau struct {
	topRightX int
	topRightY int

	holeX int
	holeY int
}

func (plateau DonutPlateau) isOnThePlateau(x int, y int) bool {
	withinBounds := x >= 0 && x <= plateau.topRightX && y >= 0 && y <= plateau.topRightY
	return withinBounds && (x != plateau.holeX || y != plateau.holeY)
}

type Plateau interface {
	isOnThePlateau(x int, y int) bool
}

func main() {
	plateau := RectangularPlateau{topRightX: 5, topRightY: 5}
	donutPlateau := &DonutPlateau{topRightX: 5, topRightY: 5, holeX: 2, holeY: 3}

	doSomethingWithAPlateau(plateau)
	doSomethingWithAPlateau(donutPlateau)
	fmt.Printf("Is within: %v\n", plateau.isOnThePlateau(3, 3))
	fmt.Printf("Is within donut: %v\n", donutPlateau.isOnThePlateau(3, 3))
	fmt.Printf("Is within donut: %v\n", donutPlateau.isOnThePlateau(2, 3))
}

func doSomethingWithAPlateau(plateau Plateau) {
	fmt.Printf("Hey I am a plateau: %v\n", plateau)
}
package funcs

import (
	"log"
	"math/rand"
	"time"
)

// global stuff
var Cells []Cell
var FoodItems []Food

func SpawnFood() []Food {
	valuesSrc := []int{5, 10, 20}
	rarities := []int{5, 4, 1}

	FoodItems = []Food{}

	var values []int
	for i := 0; i < 10; i++ {
		if len(values) < rarities[0] {
			values = append(values, valuesSrc[0])
		} else if len(values) < rarities[1]+rarities[0] {
			values = append(values, valuesSrc[1])
		} else {
			values = append(values, valuesSrc[2])
		}
	}
	for i := 0; i < len(values); i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		var pos []int
		for i := 0; i < 2; i++ {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			pos = append(pos, r.Intn(2000)-1000)
		}

		var newItem Food
		value := values[r1.Intn(len(values))]
		newItem = Food{len(FoodItems), pos, value}

		FoodItems = append(FoodItems, newItem)
	}

	log.Println(FoodItems)

	return FoodItems
}

func GetFood() []Food {
	return FoodItems
}

func InitCell(name string) Cell {
	var NewCell Cell
	var pos []int

	for i := 0; i < 2; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		pos = append(pos, r.Intn(2000)-1000)
	}
	NewCell = Cell{len(Cells), name, true, 10, 0, []Cell{}, pos}
	Cells = append(Cells, NewCell)
	log.Printf("New Cell: %v", NewCell)

	return NewCell
}

func ChangeSize(id int, size int) Cell {
	Cells[id].Size = size
	return Cells[id]
}

func Delall() {
	Cells = []Cell{}
	FoodItems = []Food{}
}

func Eat(id int, mealId int) []Cell {
	//log.Println(id, mealId)
	Cells[mealId].Alive = false
	Cells[id].Meals = append(Cells[id].Meals, Cells[mealId])
	Cells[id].Kills += 1
	return Cells
}

func GetCells() []Cell {
	return Cells
}

func ChangePos(id int, pos []int) Cell {
	Cells[id].Pos = pos
	return Cells[id]
}
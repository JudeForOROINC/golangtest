package main

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Hello")

// 	hello := widget.NewLabel("Hello Fyne!")
// 	w.SetContent(container.NewVBox(
// 		hello,
// 		widget.NewButton("Hi!", func() {
// 			hello.SetText("Welcome :)")
// 		}),
// 	))

// 	w.ShowAndRun()
// }

import (
	"fmt"

	"example.com/test/conrtoller"
	"example.com/test/engine"
	"example.com/test/view"
)

type World struct {
	locations []Location
	engine    engine.Items
}

// type LandType {

// }

type Location struct {
	id        int
	name      string
	locType   int //LandType // TODO Enum
	direction []*Location
	mobiles   []*Mob
	items     []*Item
	building  Building
}

type Item struct {
	id         int
	name       string
	value      int
	weight     int
	owner      *Mob
	isOnFloor  *Building
	isOnGroung *Location
}

type Mob struct {
	id       int
	name     string
	weight   int
	canRise  int
	location *Location
	money    int
	items    []*Item
}

type Building struct {
	id             int
	name           string
	location       *Location
	buildingtype   int
	hasBuildings   []*Building
	enterBuilding  *Building
	parentBuilding *Building
	items          []*Item
}

func (m *Mob) addItem(i *Item) {
	i.owner = m
	m.items = append(m.items, i)
	fmt.Println("was added item")
}

func main() {
	fmt.Println("hell, World!")

	// i := Item{id: 1, name: "food", weight: 1, value: 100}
	// m := Mob{id: 1, name: "Monster", canRise: 200, money: 0}
	// l := Location{id: 1, name: "port"}
	// l2 := Location{id: 2, name: "field"}

	// fmt.Println("item", i)
	// fmt.Println("monster", m)

	// m.addItem(&i)

	// fmt.Println("item", i)
	// fmt.Println("monster", m)

	// fmt.Println("locations", l, l2)

	// i1 := engine.NewItem(1, "food", 1, 2)

	// l1 := engine.Items{}

	// l1.AddItem(i1)

	// fmt.Println("locations======", l1, i1)

	// l1.DeleteItem(i1)

	// fmt.Println("locations======", l1, i1)

	// c := engine.NewLand(1, "port", "small port land")
	// c2 := engine.NewLand(2, "farm", "small farm land")

	// fmt.Println(c, c2)

	// c.LinkNowrth(c2)

	// fmt.Println(c, c2)

	// mb := engine.NewMob(4, "monster", 20)
	// c.Mobs().AddMob(mb)
	// fmt.Println(c, c2)

	w := engine.WorlExample()

	pl := w.Player()

	fmt.Println(view.Mob{*pl}.View())

	fmt.Print("I want to move out : ===================\n")

	pl.Move(engine.Out)

	fmt.Println(view.Mob{*pl}.View())

	fmt.Print("I want to move north : ===================\n")

	pl.Move(engine.North)

	fmt.Println(view.Mob{*pl}.View())

	// fmt.Print("I want to move north : ===================\n")

	// pl.Move(engine.North)

	// fmt.Println(view.Mob{*pl}.View())

	// fmt.Print("I want to move north : ===================\n")

	// pl.Move(engine.North)

	// fmt.Println(view.Mob{*pl}.View())

	// fmt.Print("I want to move north : ===================\n")

	// pl.Move(engine.North)

	// fmt.Println(view.Mob{*pl}.View())

	// fmt.Print("I want to move west : ===================\n")

	// pl.Move(engine.West)

	// fmt.Println(view.Mob{*pl}.View())

	// fmt.Print("I want to move west : ===================\n")

	// pl.Move(engine.West)

	// fmt.Println(view.Mob{*pl}.View())

	// fmt.Print("I want to move west : ===================\n")

	// pl.Move(engine.West)

	// fmt.Println(view.Mob{*pl}.View())

	// fmt.Print("I want to move west : ===================\n")

	// pl.Move(engine.West)

	// fmt.Println(view.Mob{*pl}.View())

	// fmt.Print("I want to move out : ===================\n")

	// pl.Move(engine.Out)

	// fmt.Println(view.Mob{*pl}.View())

	conrtoller.CommandListener(w)

	// var ll view.Land
	// ll = *c
	// ll := view.Land{*c}
	// fmt.Println(ll.View())

	// mb.Move(engine.North)
	// fmt.Println(c, c2)
	// fmt.Println(ll.View())

	// ticker :=

	// for {
	// 	select {
	// 	case <-ticker.C:
	// 	}
}

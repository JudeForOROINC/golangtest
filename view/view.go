package view

import (
	"strconv"

	"example.com/test/engine"
)

type Viewer interface {
	View() string
}

type Land struct {
	engine.Land
}
type Cell struct {
	engine.Cell
}
type Building struct {
	engine.Building
}
type Mobs struct {
	engine.Mobs
}
type Mob struct {
	engine.Mob
}

func (l Land) View() string {
	ret := " This is the land " + l.Name() + "\n"
	if !engine.IsNil(l.Enter()) {
		ret += " here can enter to cell " + l.Enter().(*engine.Cell).Name()
	}

	if !engine.IsNil(l.North()) {
		landThere := l.North().(*engine.Land)

		ret += "on the north there are " + landThere.Name() + "\n"
	}

	if !engine.IsNil(l.South()) {
		landThere := l.South().(*engine.Land)

		ret += "on the south there are " + landThere.Name() + "\n"
	}

	if !engine.IsNil(l.West()) {
		landThere := l.West().(*engine.Land)

		ret += "on the west there are " + landThere.Name() + "\n"
	}

	if !engine.IsNil(l.East()) {
		landThere := l.East().(*engine.Land)

		ret += "on the east there are " + landThere.Name() + "\n"
	}

	if l.Mobs() != nil {
		ret += Mobs{*l.Mobs()}.View()
	}
	return ret
}

// var isNil = isNil()

func (c Cell) View() string {
	ret := "this is cell " + c.Name() + " \n "
	if !engine.IsNil(c.Enter()) {
		ret += "has building " + c.Enter().(*engine.Building).Name() + " \n"
	}

	if !engine.IsNil(c.North()) {
		ret += "at North is " + c.North().(*engine.Cell).Name() + " \n"
	}
	if !engine.IsNil(c.South()) {
		ret += "at South is " + c.South().(*engine.Cell).Name() + " \n"
	}
	if !engine.IsNil(c.East()) {
		ret += "at East is " + c.East().(*engine.Cell).Name() + " \n"
	}
	if !engine.IsNil(c.West()) {
		ret += "at West is " + c.West().(*engine.Cell).Name() + " \n"
	}
	if !engine.IsNil(c.Exit()) {
		ret += "exit to Land " + c.Exit().(*engine.Land).Name() + " \n"
	}
	// typ := ""
	// 	val := c.North().(engine.Cell)
	// 	// val = Cell{*val}
	// 	// typ += Cell{*c.North()}.Name()
	// 	// }
	// 	ret += "At the north is " + val.Name()

	// }

	if c.Mobs() != nil {
		ret += Mobs{*c.Mobs()}.View()
	}

	// if c.Mobs() != nil && len(c.Mobs()) > 0 {
	// 	ret += "There are some monsters " + string(len(c.Mobs())) + " number /n"
	// }

	return ret
}

func (b Building) View() string {
	ret := "building " + b.Name() + "\n"
	if !engine.IsNil(b.Exit()) {
		ret += " there an exit from a building: "
		c, ok := b.Exit().(*engine.Cell)
		if ok {
			ret += c.Name()
		}
		ret += "\n"
	} else {
		ret += " no exit from building"
	}

	// ret += b.View()
	return ret
}

func (ms Mobs) View() string {
	// name := ms.Mobs
	// fmt.Println(name)
	if len(ms.Mobs.Mobs) > 0 {
		ret := "there are some monsters here " + strconv.Itoa(len(ms.Mobs.Mobs)) + " mobs number \n"
		return ret
	}
	return " no monsters here. \n"
}

func (m Mob) View() string {
	//info
	ret := " My name is "
	ret += m.Name() + " \n"
	if m.Mobs() != nil {
		ret += "Overview where am i: " + Mobs{*m.Mobs()}.View()

		ms := m.Mobs()
		loc := ms.Location()

		b, ok := loc.(*engine.Building)

		if ok {
			ret += "I am at the building : " + Building{*b}.View()
		} else {

			c, ok := loc.(*engine.Cell)
			if ok {
				ret += "I am at the Cell : " + Cell{*c}.View()
			} else {
				l, ok := loc.(*engine.Land)
				if ok {
					ret += "I am at the Land : " + Land{*l}.View()
				}
			}
		}

	}

	ret += "My current health is " + strconv.Itoa(m.Hits()) + " \n"

	return ret
}

func Main() {
	// engine.Cell

}

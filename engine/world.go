package engine

import "strconv"

type World struct {
	locations []*Land
	cells     []*Cell
	Mobs      []*Mob
	buildings []*Building
	items     []*Item
	age       int
	player    *Mob
}

// TODO made a system of links to store all worlds objects

//item must be eatable or weareble or weepon or valuable

// try to made a lyers with movement , economycs, battle, ect.

// Goal of projec = lern GoLang
// Goal of game = build an empire.

type LandType int

const (
	farmLand LandType = iota
	Forest
	hill
	Desert
)

func (w *World) CreateLocation(lt LandType, name *string, description *string) *Land {
	var dn, ds string
	switch lt {
	case farmLand:
		dn = "farmland"
		ds = "Small farmlands placed at nearest area."
	case Forest:
		dn = "Forest"
		ds = "Deep forest around you"
	case hill:
		dn = "Hills"
		ds = "High hills lies here"
	case Desert:
		dn = "Desert"
		ds = "Sad Desert of the reality"

	}
	if name != nil {
		dn = *name
	}
	if description != nil {
		ds = *description
	}
	land := Land{id: len(w.locations), name: dn, description: ds, mobs: &Mobs{}}
	land.mobs.location = &land
	w.locations = append(w.locations, &land)

	return &land
}

func (w *World) CreateCell(land *Land, name *string) *Cell {
	dn := " cell of land"
	if name != nil {
		dn = *name
	}
	cell := Cell{id: len(w.cells), name: dn, land: land, mobs: &Mobs{}}
	cell.mobs.location = &cell
	w.cells = append(w.cells, &cell)

	return &cell
}

type BuildingType int

const (
	farm BuildingType = iota
	lamber
	mine
	goldmine
	house
	towerhall
	forge
	smith
)

func (w *World) CreateBuilding(bt BuildingType, c *Cell) *Building {
	var dn string
	switch bt {
	case farm:
		dn = "farm"
	case towerhall:
		dn = "Tower"
	}
	b := Building{id: len(w.buildings), name: dn, cell: c, mobs: &Mobs{}}
	b.mobs.location = &b
	w.buildings = append(w.buildings, &b)

	return &b
}

func (w *World) CreateMob(name string, hp int, loadedWeight int) *Mob {
	m := Mob{id: len(w.Mobs), name: name, hp: hp, loadedWeight: loadedWeight}
	w.Mobs = append(w.Mobs, &m)
	return &m
}

type ItemType int

const (
	food ItemType = iota
	weapon
	armor
	wood
	ore
	gold
	iron
	instrument
)

func (w World) CreateItem(it ItemType, name *string, weight *int, value *int) *Item {
	i := Item{id: len(w.items), name: *name, value: *value, weight: *weight}
	w.items = append(w.items, &i)
	return &i
}

func WorlExample() *World {
	w := World{}
	//TODO add slice to create array of lands

	L1 := w.CreateLocation(Forest, strPointer("corner forest"), strPointer("far corner forest"))
	L2 := w.CreateLocation(farmLand, nil, nil)
	L1.LinkEast(L2)
	L3 := w.CreateLocation(farmLand, nil, nil)
	L2.LinkEast(L3)
	L4 := w.CreateLocation(hill, nil, nil)
	L3.LinkEast(L4)
	L5 := w.CreateLocation(Desert, nil, nil)
	L4.LinkEast(L5)

	L6 := w.CreateLocation(Forest, nil, nil)
	L1.LinkSouth(L6)

	L7 := w.CreateLocation(farmLand, nil, nil)
	L2.LinkSouth(L7)
	L6.LinkEast(L7)

	L8 := w.CreateLocation(farmLand, nil, nil)
	L3.LinkSouth(L8)
	L7.LinkEast(L8)

	L9 := w.CreateLocation(hill, nil, nil)
	L4.LinkSouth(L9)
	L8.LinkEast(L9)

	L10 := w.CreateLocation(hill, nil, nil)
	L5.LinkSouth(L10)
	L9.LinkEast(L10)

	L11 := w.CreateLocation(Forest, nil, nil)
	L6.LinkSouth(L11)
	//L8.LinkEast(L9)
	L12 := w.CreateLocation(hill, nil, nil)
	L7.LinkSouth(L12)
	L11.LinkEast(L12)

	L13 := w.CreateLocation(farmLand, strPointer("Town"), strPointer("Town at the middle of the lands"))
	L8.LinkSouth(L13)
	L12.LinkEast(L13)

	L14 := w.CreateLocation(farmLand, nil, nil)
	L9.LinkSouth(L14)
	L13.LinkEast(L14)

	L15 := w.CreateLocation(farmLand, nil, nil)
	L10.LinkSouth(L15)
	L14.LinkEast(L15)

	L16 := w.CreateLocation(Forest, nil, nil)
	L11.LinkSouth(L16)
	// L14.LinkEast(L15)

	L17 := w.CreateLocation(farmLand, nil, nil)
	L12.LinkSouth(L17)
	L16.LinkEast(L17)

	L18 := w.CreateLocation(farmLand, nil, nil)
	L13.LinkSouth(L18)
	L17.LinkEast(L18)

	L19 := w.CreateLocation(farmLand, nil, nil)
	L14.LinkSouth(L19)
	L18.LinkEast(L19)

	L20 := w.CreateLocation(farmLand, nil, nil)
	L15.LinkSouth(L20)
	L19.LinkEast(L20)

	L21 := w.CreateLocation(Forest, nil, nil)
	L16.LinkSouth(L21)
	// L14.LinkEast(L15)

	L22 := w.CreateLocation(Forest, nil, nil)
	L17.LinkSouth(L22)
	L21.LinkEast(L22)

	L23 := w.CreateLocation(Forest, nil, nil)
	L18.LinkSouth(L23)
	L22.LinkEast(L23)

	L24 := w.CreateLocation(Forest, nil, nil)
	L19.LinkSouth(L24)
	L23.LinkEast(L24)

	L25 := w.CreateLocation(Forest, nil, nil)
	L20.LinkSouth(L25)
	L24.LinkEast(L25)

	for _, L := range w.locations {
		Cl := w.CreateCell(L, strPointer("Cell "+strconv.Itoa(len(w.cells))))
		L.LinkEnter(Cl)
	}

	var Cells [10]*Cell

	var towerhallCell, current *Cell

	for ii := 1; ii < 100; ii++ {
		if ii == 44 {
			current = w.CreateCell(nil, strPointer("Towerhall"))
			towerhallCell = current
		} else {
			current = w.CreateCell(nil, nil)
		}
		if ii == 1 {
			L13.enterCell.LinkEast(current)
			Cells[ii%10] = current
			continue
		}
		if ii%10 != 0 {
			Cells[ii%10-1].LinkEast(current)
		}
		if ii > 9 {
			if ii == 10 {
				L13.enterCell.LinkSouth(current)
				Cells[ii%10] = current
				continue
			}
			Cells[ii%10].LinkSouth(current)
		}
		Cells[ii%10] = current

	}

	b := w.CreateBuilding(towerhall, towerhallCell)

	towerhallCell.LinkEnter(b)

	//========================now need to create player============//

	w.player = w.CreateMob("Jude", 100, 250)

	b.Mobs().AddMob(w.player)

	return &w
}

func strPointer(s string) *string {
	return &s
}

func (w World) Player() *Mob {
	return w.player
}
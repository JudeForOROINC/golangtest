package engine

//All Players or monsters is at navigation .
// main monsters func : die;

type MonsterGeter interface {
	MonsterGet()
	Die()
}

type Monster struct {
	Mob    Mob
	attack int
	loot   Items
}

type PlantGrowFase int

const (
	seed PlantGrowFase = iota
	binstok
	bush
	tree
	apples
	dead
)

type Plant struct {
	Monster  Monster
	growFace PlantGrowFase
}

type Cow struct {
	Monster Monster
	hunger  int
	age     int
	milk    int
}

type Wolf struct {
	Monster    Monster
	hunger     int
	age        int
	leadership int
}

type MonsterType int

const (
	MT_flora MonsterType = iota
	MT_grassEater
	MT_predator
)

func (M *Monster) Die() {
	room := M.Mob.mobs.location
	itm, ok := room.(Itemer)
	if ok {
		target := itm.Items()
		var it *Item
		for it = M.loot.First(); M.loot.Count() > 0; {
			it = M.loot.First()
			it.MoveTo(target)
			// it = M.loot.First()
		}
	}
}

// func MonsterFactory(w *World, typ MonsterType) *Monster {
// 	// var name string
// 	// var hp, weight int
// 	switch typ {
// 	case MT_flora:
// 		name = "Tomato"
// 		hp = 3
// 		weight = 10
// 	case MT_grassEater:
// 		name = "Cow"
// 		hp = 50
// 		weight = 70
// 	case MT_predator:
// 		name = "wolf"
// 		hp = 30
// 		weight = 30
// 	}
// 	// M := w.CreateMob(name, hp, weight)
// 	// var Mr Monster
// 	// // Mr = *M
// 	// Mr.attack = 1
// 	// Mr.loot
// }

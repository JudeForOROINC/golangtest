package engine

// type technology struct { // TODO create some common system for items, produce, ect.
// 	id                 int
// 	name               string
// 	resources_need     map[*ResourcesType]int
// 	required_buildings []Building
// }

// type ResourcesType []= [food,wood]

type Prototype struct {
	Name             string
	discription      string
	requiredLandType LandType         //Land type where can be placed
	resourcesNeeded  map[ItemType]int // Need resources to start build
	labor            int              // need labor to build
	Build            BuildingType
	BuildFase        BuildFase
}

type BuildFase int

const (
	justPlaced = iota
	materialsCollecting
	materialsLoaded
	moneyCollecting
	moneyCollected
	workStarted
	Builded
	Close
)

// func FactoryPrototypeBuilding(bt BuildingType) *Prototype {
// 	//TODO
// }

// func ChechPrototype(loc Navigation) []BuildingType {
// 	//TODO
// }

// func (p *Prototype) ValidatePlaceHere(pl Mob) bool { // Check can be started or not
// 	//TODO
// }

// func (p *Prototype) Place(master Mob) (*Prototype, string) {
// 	//TODO
// }

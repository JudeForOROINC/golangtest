package engine

type Land struct {
	id          int
	name        string
	description string
	north       *Land
	south       *Land
	east        *Land
	west        *Land
	cells       []*Cell
	enterCell   *Cell
	mobs        *Mobs
}

type Building struct {
	id    int
	name  string
	cell  *Cell
	items []*Items
	mobs  *Mobs
}

type Cell struct {
	id       int
	name     string
	building *Building
	items    Items
	north    *Cell
	south    *Cell
	east     *Cell
	west     *Cell
	mobs     *Mobs
	land     *Land
}

func (land *Land) North() Navigation {
	return land.north
}

func (land *Land) South() Navigation {
	return land.south
}
func (land *Land) East() Navigation {
	return land.east
}

func (land *Land) West() Navigation {
	return land.west
}

func (land *Land) Enter() Navigation {
	return land.enterCell
}

func (land *Land) Exit() Navigation {
	return nil
}

func (land *Land) Mobs() *Mobs {
	return land.mobs
}

func (c *Cell) Mobs() *Mobs {
	return c.mobs
}
func (b *Building) Mobs() *Mobs {
	return b.mobs
}

func NewLand(id int, name string, description string) *Land {
	land := Land{id: id, name: name, description: description, mobs: &Mobs{}}
	land.mobs.location = &land

	return &land
}

func (cell *Cell) North() Navigation {
	return cell.north
}

func (cell *Cell) South() Navigation {
	return cell.south
}
func (cell *Cell) East() Navigation {
	return cell.east
}
func (cell *Cell) West() Navigation {
	return cell.west
}

func (b *Building) Exit() Navigation {
	return b.cell
}

func (c *Cell) Enter() Navigation {
	return c.building
}
func (c *Cell) Exit() Navigation {
	return c.land
}
func (b *Building) North() Navigation {
	return nil
}
func (b *Building) South() Navigation {
	return nil
}
func (b *Building) East() Navigation {
	return nil
}
func (b *Building) West() Navigation {
	return nil
}
func (b *Building) Enter() Navigation {
	return nil
}

type Navigation interface {
	North() Navigation
	South() Navigation
	East() Navigation
	West() Navigation
	Enter() Navigation
	Exit() Navigation
	Mobs() *Mobs
}

func (c *Cell) LinkNowrth(ce *Cell) {
	c.north = ce
	ce.south = c
}

func (c *Cell) LinkSouth(ce *Cell) {
	c.south = ce
	ce.north = c
}

func (c *Cell) LinkEast(ce *Cell) {
	c.east = ce
	ce.west = c
}

func (c *Cell) LinkWest(ce *Cell) {
	c.west = ce
	ce.east = c
}

func (c *Cell) LinkEnter(b *Building) {
	c.building = b
	b.cell = c
}

func (c *Land) LinkNowrth(ce *Land) {
	c.north = ce
	ce.south = c
}

func (c *Land) LinkSouth(ce *Land) {
	c.south = ce
	ce.north = c
}

func (c *Land) LinkEast(ce *Land) {
	c.east = ce
	ce.west = c
}

func (c *Land) LinkWest(ce *Land) {
	c.west = ce
	ce.east = c
}

func (c *Land) LinkEnter(ce *Cell) {
	c.enterCell = ce
	ce.land = c
}

func (l *Land) Name() string {
	return l.name
}

func (c *Cell) Name() string {
	return c.name
}
func (b *Building) Name() string {
	return b.name
}

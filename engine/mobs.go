package engine

import "reflect"

type Mob struct {
	id           int
	name         string
	hp           int
	loadedWeight int
	items        Items
	mobs         *Mobs
	idMobs       *int
}

type Mobs struct {
	Mobs     []*Mob
	location Navigation
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
	Enter
	Out
)

func IsNil(i interface{}) bool {
	return reflect.ValueOf(i).IsNil()
}

func (m *Mob) Move(d Direction) {
	var target *Mobs
	switch d {
	case North:

		if !IsNil(m.mobs.location.North()) {
			nav := m.mobs.location.North()
			target = nav.Mobs()

		}
	case South:

		if !IsNil(m.mobs.location.South()) {
			nav := m.mobs.location.South()
			target = nav.Mobs()

		}
	case West:

		if !IsNil(m.mobs.location.West()) {
			nav := m.mobs.location.West()
			target = nav.Mobs()

		}
	case East:

		if !IsNil(m.mobs.location.East()) {
			nav := m.mobs.location.East()
			target = nav.Mobs()

		}
	case Enter:

		if !IsNil(m.mobs.location.Enter()) {
			nav := m.mobs.location.Enter()
			target = nav.Mobs()

		}

	case Out:

		if !IsNil(m.mobs.location.Exit()) {
			nav := m.mobs.location.Exit()
			target = nav.Mobs()

		}

	}

	if target != nil {
		m.mobs.RemoveMob(m)
		target.AddMob(m)
	}

}

func (m *Mobs) AddMob(mb *Mob) {
	m.Mobs = append(m.Mobs, mb)
	mb.mobs = m
	lens := len(m.Mobs)
	mb.idMobs = &lens
}

func (m *Mobs) RemoveMob(mb *Mob) {
	if mb.idMobs != nil {
		if len(m.Mobs) <= *mb.idMobs {
			slice := m.Mobs[*mb.idMobs:]
			m.Mobs = append(m.Mobs[:*mb.idMobs-1], slice...)
			mb.idMobs = nil
			mb.mobs = nil
		}
	}

}

func NewMob(id int, name string, hp int) *Mob {
	return &Mob{id: id, name: name, hp: hp}
}

func (m *Mob) Name() string {
	return m.name
}

func (m *Mob) Mobs() *Mobs {
	return m.mobs
}

func (ms *Mobs) Location() Navigation {
	return ms.location
}

func (m *Mob) Hits() int {
	return m.hp
}

func (m *Mob) Take(itm *Item) *Items {
	itm.MoveTo(&m.items)
	return &m.items
}

func (m *Mob) Items() *Items {
	return &m.items
}

func (m *Mob) Drop(itm *Item) *Items {
	loc := m.mobs.location
	itms, ok := loc.(Itemer)
	if ok {
		itm.MoveTo(itms.Items())
		//TODO report sucess
	}
	return &m.items
}

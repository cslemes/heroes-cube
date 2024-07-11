package game

import (
	"heroes-cube/internals/db"
	"heroes-cube/internals/utils"
	"sort"
)

type Person struct {
	Id        string
	Name      string
	Skills    map[string]int
	Damage    int
	HitPoints int
	Defense   int
	Race
	Class
	Slots
	Inventory
	Exp   int
	Level int
}

func (p *Person) initSkills() {

	numbers := utils.DrawDM(12, 5)
	sort.IntSlice.Sort(numbers)
	numbers = numbers[1 : len(numbers)-1]

	skills := map[string]int{
		p.PrimaryStatus:   numbers[2],
		p.SecondaryStatus: numbers[1],
		p.ThirdyStatus:    numbers[0],
	}

	p.Skills["agility"] = skills["agility"]
	p.Skills["inteligence"] = skills["inteligence"]
	p.Skills["strength"] = skills["strength"]
}

func (p *Person) initHitPoints() {
	p.HitPoints = p.Skills["strength"] + 10
}

func (p *Person) initDamage() {
	p.Damage = p.Skills[p.PrimaryStatus] + p.Level
	for _, v := range p.Slots {
		p.Damage += v.Damage
	}
}

func (p *Person) initDefense() {
	p.Defense = p.Skills["agility"]

	if p.Slots["arms"].Type == "armor" {
		p.Damage += p.Slots["arms"].Defense
	}

	if p.Slots["head"].Type == "armor" {
		p.Damage += p.Slots["head"].Defense
	}

	if p.Slots["chest"].Type == "armor" {
		p.Damage += p.Slots["chest"].Defense
	}

	if p.Slots["legs"].Type == "armor" {
		p.Damage += p.Slots["legs"].Defense
	}
}

func (p *Person) ToPersonDB() *db.Person {

	personDB := &db.Person{

		Id:          p.Id,
		Name:        p.Name,
		Strength:    p.Skills["strength"],
		Agility:     p.Skills["agility"],
		Inteligence: p.Skills["inteligence"],
		Damage:      p.Damage,
		HitPoints:   p.HitPoints,
		Defense:     p.Defense,
		Class:       p.Class.Class,
		Race:        p.Race.Race,
		Exp:         p.Exp,
		Level:       p.Level,
	}
	return personDB
}

func (p *Person) UpdateOrCreate() error {

	if err := p.Inventory.UpdateOrCreate(p.Id); err != nil {
		return err
	}

	if err := p.Slots.UpdateOrCreate(p.Id); err != nil {
		return err
	}

	personDB := p.ToPersonDB()
	personDBActual, err := db.GetPerson(p.Id, con)
	if err != nil {
		return err
	}

	if personDBActual.Id != p.Id {
		return db.CreatePerson(personDB, con)
	}
	return db.UpdatePerson(personDB, con)
}

func NewPerson(id, name, class, race string) (*Person, error) {

	checkClass := false
	for k := range Classes {
		if class == k {
			checkClass = true
			break
		}
	}

	if !checkClass {
		return nil, utils.ClassNotFound
	}

	checkRace := false
	for k := range Races {
		if race == k {
			checkRace = true
			break
		}
	}

	if !checkRace {
		return nil, utils.RaceNotFound
	}

	p := &Person{
		Id:        id,
		Name:      name,
		Skills:    map[string]int{},
		Damage:    0,
		HitPoints: 0,
		Defense:   0,
		Race:      Races[race],
		Class:     Classes[class],
		Slots:     Slots{},
		Inventory: Inventory{},
		Exp:       1,
		Level:     1,
	}

	p.initSkills()
	p.initHitPoints()
	p.initDamage()
	p.initDefense()

	return p, nil
}

func ImportPerson(idPerson string) (*Person, error) {

	personDB, err := db.GetPerson(idPerson, con)
	if err != nil {
		return nil, err
	}

	p := &Person{
		Id:        personDB.Id,
		Name:      personDB.Name,
		Damage:    personDB.Damage,
		HitPoints: personDB.HitPoints,
		Defense:   personDB.Defense,
		Exp:       personDB.Exp,
		Level:     personDB.Level,
	}

	p.Skills = map[string]int{
		"strength":    personDB.Strength,
		"agility":     personDB.Agility,
		"inteligence": personDB.Inteligence,
	}

	p.Race = Races[personDB.Race]
	p.Class = Classes[personDB.Class]

	slots, err := ImportSlots(idPerson)
	if err != nil {
		return nil, err
	}
	p.Slots = slots

	inventory, err := ImportInventory(idPerson)
	if err != nil {
		return nil, err
	}
	p.Inventory = inventory

	return p, nil

}

func GetOrCreatePerson(id, name, class, race string) (*Person, error) {

	p, err := ImportPerson(id)
	if err != nil {
		return nil, err
	}

	if p.Id != "" {
		return p, nil
	}

	p, err = NewPerson(id, name, class, race)
	if err != nil {
		return nil, err
	}

	return p, nil
}

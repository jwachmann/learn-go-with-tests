package structsmethodsinterfaces

type Animal interface {
	Name() string
}

type PartyAnimal interface {
	Animal
	Party() string
}

type PartyMonkey struct{}

func (p *PartyMonkey) Name() string {
	return "Party Monkey!"
}

func (p *PartyMonkey) Party() string {
	return "Party party party!"
}

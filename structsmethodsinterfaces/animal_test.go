package structsmethodsinterfaces

import "testing"

func TestName(t *testing.T) {
	var partyAnimal PartyAnimal
	partyAnimal = &PartyMonkey{}

	var animal Animal
	animal = partyAnimal

	got := partyAnimal.Name()
	want := "Party Monkey!"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

	got = animal.Name()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

package person

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func FindMostCriminalPerson(people map[string]Man, suspects []string) (Man, bool) {
	var mostCriminal Man
	var found bool
	for _, name := range suspects {
		if person, ok := people[name]; ok && person.Crimes > mostCriminal.Crimes {
			mostCriminal = person
			found = true
		}
	}
	return mostCriminal, found
}
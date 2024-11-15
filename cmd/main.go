package main

import (
	"log"
	"github.com/iqmag/crimeTracker/pkg/person"
)

func main() {
	people := map[string]person.Man{
		"Vin":    {Name: "Vin", LastName: "Diesel", Age: 56, Gender: "male", Crimes: 90},
		"John":   {Name: "John", LastName: "Wick", Age: 59, Gender: "male", Crimes: 0},
		"Gerard": {Name: "Gerard", LastName: "Butler", Age: 54, Gender: "male", Crimes: 1},
		"Tom":    {Name: "Tom", LastName: "Hardy", Age: 46, Gender: "female", Crimes: 0},
		"Jason":  {Name: "Jason", LastName: "Statham", Age: 56, Gender: "male", Crimes: 95},
	}

	suspects := []string{"Vin", "Gerard", "Tom", "Jason"}

	mostCriminal, found := person.FindMostCriminalPerson(people, suspects)
	if found {
		log.Printf("Наибольшее количество совершенных преступлений: (c) %s %s\n", mostCriminal.Name, mostCriminal.LastName)
	} else {
		log.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
}

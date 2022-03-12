package examples

import "fmt"

//Convert Ints, Strings, Dates, etc

func Convert(input uint8) string {
	men := input
	const women int16 = 42

	var people int
	//error
	//people = men+women
	people = int(men) + int(women)
	return string(rune(people))
}

//DaysOfTheWeek ...
func DaysOfTheWeek() {
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	//slice
	weekdays := days[0:5]
	//map
	isWeekend := map[string]bool{
		days[5]: true,
		days[6]: true,
	}
	fmt.Println(days)
	fmt.Println(weekdays)
	fmt.Println(isWeekend["tuesday"])
	fmt.Println(isWeekend["saturday"])
}

//StrutExample ...
func StrutExample() {
	type Person struct {
		name string
		age  int
	}

	// our Team struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty 'Team'
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{{name: "Forrest"}, {name: "Gordon"}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)
}

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type byIdade []user

func (u byIdade) Len() int { return len(u) }
func (u byIdade) Less(i, j int) bool { return u[i].Age < u[j].Age }
func (u byIdade) Swap(i, j int) { u[i], u[j] = u[j], u[i]}


type byLast []user

func (u byLast) Len() int	{ return len(u) }
func (u byLast) Less(i, j int) bool { return u[i].Last < u[j].Last }
func (u byLast) Swap(i, j int) { u[i], u[j] = u[j], u[i]}



func main() {
	u2 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u3 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u1 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println("original:")
	legal(users)
	
	sort.Sort(byIdade(users))
	fmt.Println("idade:")
	legal(users)
	
	sort.Sort(byLast(users))
	fmt.Println("Last name:")
	legal(users)
}
func legal(u []user) {
	for _, us := range u {
		fmt.Println("    ",us.First,us.Last)
		fmt.Println("    ",us.Age,"anos")
		for _, v := range us.Sayings {
			fmt.Println("\t->",v)
		}
	}
}
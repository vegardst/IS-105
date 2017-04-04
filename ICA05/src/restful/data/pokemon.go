package restful

/* not in use


import (
	"math/rand"
	"time"
	"strconv"
)
type Pokemon struct {
	Name string `json:"name"`
	Weight int `json:"weight"`
	ID int `json:"id"`
}

func GetPokemon() *Pokemon {
	// Create random number
	rand.Seed(time.Now().Unix())
	t := rand.Intn(150)
	// Get int to use in url
	number := strconv.Itoa(t)
	url := "http://pokeapi.co/api/v2/pokemon/" + number
	poke := new(Pokemon)
	GetJSON(url, poke)
	return poke
}
*/
package pokeapi

import (
	"time"

	"github.com/JonahLargen/Pokedex/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

var pokeCache *pokecache.Cache = pokecache.NewCache(5 * time.Minute)

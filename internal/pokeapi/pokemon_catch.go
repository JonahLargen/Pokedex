package pokeapi

import (
	"math/rand"
	"time"
)

const (
	StandardPokeballBonus = 1.0
	MaxCatchRate          = 255
)

func AttemptCatch(pokemon *PokemonResponse) bool {
	catchRate := 200 - (pokemon.BaseExperience / 2)
	if catchRate < 1 {
		catchRate = 1
	}
	if catchRate > MaxCatchRate {
		catchRate = MaxCatchRate
	}

	ballBonus := StandardPokeballBonus

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	catchValue := rng.Intn(256)

	return float64(catchValue) < float64(catchRate)*ballBonus
}

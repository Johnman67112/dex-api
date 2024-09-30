package core

import "gopkg.in/validator.v2"

type Pokemon struct {
	Number uint
	Name   string
	Type   string
	Weight float64
	Heigh  float64
}

func PokemonValidator(pokemon *Pokemon) error {
	if err := validator.Validate(pokemon); err != nil {
		return err
	}

	return nil
}

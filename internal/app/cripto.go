package app

import (
	"fmt"

	"github.com/greyfox12/GoDilpom1/pkg/cripto"
)

func InitCripto() *cripto.Auth {
	Secretkey, err := cripto.GenerateRandom(32)
	if err != nil {
		//		Logger.Fatal(fmt.Sprint("error generateRandom:  %w", err))
		panic(fmt.Sprint("error generateRandom:  %w", err))
	}
	return &cripto.Auth{Secretkey: Secretkey}

}

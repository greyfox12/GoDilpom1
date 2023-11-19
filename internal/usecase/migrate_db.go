package usecase

import "fmt"

func (u *UseCase) MigrataShemaDB() error {
	fmt.Printf("Migrate UseCase\n")
	return u.migrateShema.Execute()
}

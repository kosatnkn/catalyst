package container

import (
	"fmt"
)

// Destruct releases all necessary resources that needs to be released.
func (ctr *Container) Destruct() {

	fmt.Println("")

	fmt.Println("Closing database connections...")
	ctr.Adapters.DBAdapter.Destruct()

	fmt.Println("Closing logger...")
	ctr.Adapters.LogAdapter.Destruct()
}

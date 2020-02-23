package container

import (
	"fmt"
)

// Destruct the container releaseing resources.
func (ctr *Container) Destruct() {

	fmt.Println("")

	fmt.Println("Closing database connections...")
	ctr.Adapters.DBAdapter.Destruct()

	fmt.Println("Closing logger...")
	ctr.Adapters.LogAdapter.Destruct()
}

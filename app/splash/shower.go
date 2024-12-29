package splash

import (
	"fmt"

	"github.com/kosatnkn/catalyst/v3/metadata"
)

// Show a splash screen in one of several types.
func Show() {
	fmt.Print(metadata.BaseInfo())
	fmt.Print(metadata.BuildInfo())
	fmt.Println("---")
}

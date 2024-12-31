package splash

import (
	"fmt"

	"github.com/kosatnkn/catalyst/v3/metadata"
)

// Show a splash message.
func Show() {
	fmt.Print(metadata.BaseInfo())
	fmt.Print(metadata.BuildInfo())
	fmt.Println("---")
}

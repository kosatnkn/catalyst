package splash

import "fmt"

// A breif description of the service.
const serviceDetails string = `
Go Clean Architecture Base Project for RESTful Services

`

// Show a splash screen in one of several types.
func Show(style string) {
	fmt.Print(style)
	fmt.Print(serviceDetails)
}

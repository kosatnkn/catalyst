package splash

import "fmt"

// A breif description of the service.
const serviceDetails string = `
Go Clean Architecture Base Project for RESTful Services

`

// Show a splash screen in one of several types.
func Show(stryle string) {
	fmt.Print(stryle)
	fmt.Print(serviceDetails)
}

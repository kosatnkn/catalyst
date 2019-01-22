package unpackers

// UserUnpacker contains the unpacking structure for the user sent in request payload.
type UserUnpacker struct {
	FirstName      string            `json:"first_name" validate:"required"`
	LastName       string            `json:"last_name" validate:"required"`
	Age            int               `json:"age" validate:"gte=0,lte=130"`
	Email          string            `json:"email" validate:"required,email"`
	FavouriteColor string            `json:"favourite_color" validate:"iscolor"`          // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []AddressUnpacker `json:"addresses" validate:"required,dive,required"` // a person can have a home and cottage...
}

// RequiredFormat returns the applicable JSON format for the user data structure.
func (uu *UserUnpacker) RequiredFormat() string {
	return `{
		"first_name": <string>,
		"last_name": <string>,
		"age": <int>,
		"email": <string>,
		"favourite_color": <string>,
		"addresses": [
			{
				"street": <string>,
				"city": <string>,
				"planet": <string>,
				"phone": <string>
			}
		]
	}`
}

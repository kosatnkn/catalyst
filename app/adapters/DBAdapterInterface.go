package adapters

import "github.com/kosatnkn/db"

// DBAdapterInterface is implemented by all database adapters.
//
// TODO: Need to do a concrete definition here. Otherwise when the contract of this embedded interface changes the
// app will break. So we need to do the definition we want here. Whether that matches exactly with the interface
// presented by the third party package is an entirely different matter.
type DBAdapterInterface interface {
	db.AdapterInterface
}

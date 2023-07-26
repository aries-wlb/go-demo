package domain

import "github.com/google/wire"

// var DomainSet = wire.NewSet(
//
//	userSet,
//
// )
var DomainSet = wire.NewSet(wire.Struct(new(Domain), "*"), userSet)

type Domain struct {
	UserDomain *UserDomain
}

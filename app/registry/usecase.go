package registry

import (
	"pearshop_backend/app/usecase"

	"github.com/google/wire"
)

var (
	ProductFindUsecaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		usecase.NewProductFind,
	)
)

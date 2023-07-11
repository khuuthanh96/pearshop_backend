package registry

import (
	"github.com/google/wire"

	"pearshop_backend/app/usecase"
)

var (
	ProductFindUsecaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		usecase.NewProductFind,
	)

	ProductCreateUsecaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		usecase.NewProductCreate,
	)

	ProductUpdateUsecaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		usecase.NewProductUpdate,
	)
)

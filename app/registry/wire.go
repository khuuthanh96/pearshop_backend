//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"pearshop_backend/app/usecase"

	"github.com/google/wire"
)

func InjectedProductFindUsecase() usecase.ProductFind {
	wire.Build(ProductFindUsecaseSet)

	return nil
}

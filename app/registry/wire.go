//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"github.com/google/wire"

	"pearshop_backend/app/usecase"
)

func InjectedProductFindUsecase() usecase.ProductFind {
	wire.Build(ProductFindUsecaseSet)

	return nil
}

func InjectedProductCreateUsecase() usecase.ProductCreate {
	wire.Build(ProductCreateUsecaseSet)

	return nil
}

func InjectedProductUpdateUsecase() usecase.ProductUpdate {
	wire.Build(ProductUpdateUsecaseSet)

	return nil
}

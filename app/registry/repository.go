package registry

import (
	"pearshop_backend/app/external/storage/mysql"

	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	mysql.NewProductRepository,
)

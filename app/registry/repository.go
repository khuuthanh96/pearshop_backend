package registry

import (
	"github.com/google/wire"

	"pearshop_backend/app/external/storage/mysql"
)

var repositorySet = wire.NewSet(
	mysql.NewTxDataSQL,
	mysql.NewProductRepository,
)

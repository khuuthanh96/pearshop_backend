package registry

import (
	"pearshop_backend/pkg/gormutil"

	"github.com/google/wire"
)

// Dependency Injection: All singleton set for wire generate
var singletonSet = wire.NewSet(
	gormutil.GetDB,
)

package registry

import (
	"github.com/google/wire"

	"pearshop_backend/pkg/gormutil"
	"pearshop_backend/pkg/hashid"
)

// Dependency Injection: All singleton set for wire generate
var singletonSet = wire.NewSet(
	gormutil.GetDB,
	hashid.GetIDHasher,
)

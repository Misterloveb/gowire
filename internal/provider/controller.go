package provider

import (
	"gindemo/internal/controller"
	"github.com/google/wire"
)

var (
	indexCtl  = wire.Struct(new(controller.IndexController), "*")
	logerCtl  = wire.Struct(new(controller.LogerController), "*")
	loginCtl  = wire.Struct(new(controller.LoginController), "Server")
	systemCtl = wire.Struct(new(controller.SystemController), "*")
)

var IndexProvider = wire.NewSet(
	indexCtl,
)
var LogerProvider = wire.NewSet(
	logerCtl,
)
var LoginProvider = wire.NewSet(
	loginCtl,
)
var SystemProvider = wire.NewSet(
	systemCtl,
)

var CommonController = wire.NewSet(
	IndexProvider,
	LogerProvider,
	LoginProvider,
	SystemProvider,
)

//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Misterloveb/gowire/internal/provider"
	"github.com/Misterloveb/gowire/pkg/http"
	"github.com/google/wire"
)

func App() *http.HTTP {
	wire.Build(
		provider.Config,
		provider.CommonDao,
		provider.GinServer,
		provider.CommonServer,
		provider.CommonController,
		provider.HTTPServer,
	)
	return nil
}

package service

import (
	"leegoo/internal/data"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(data.NewKvData, NewKvService, NewGreeterService)

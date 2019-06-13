package dev

import (
	"github.com/AWildBeard/ble"
	"github.com/AWildBeard/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}

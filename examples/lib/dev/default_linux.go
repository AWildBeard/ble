package dev

import (
	"github.com/AWildBeard/ble"
	"github.com/AWildBeard/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}

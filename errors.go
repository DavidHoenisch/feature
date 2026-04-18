package main

import (
	"errors"
)

var (
	KeyNotFoundErr        error = errors.New("KeyNotFoundErr")
	KeyPresentButEmptyErr error = errors.New("KeyPresentButEmptyErr")
)

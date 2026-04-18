package main

import (
	"os"
	"strconv"
)

func getFlagFromEnv(key string) (string, error) {
	v, set := os.LookupEnv(key)

	switch set {
	case false:
		return "", KeyNotFoundErr
	default:
		if v == "" {
			return "", KeyPresentButEmptyErr
		}

		return v, nil
	}
}

func castVarToInt(v string) (int, error) {
	return strconv.Atoi(v)
}

func castVarToFloat64(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func castVarToBool(v string) (bool, error) {
	return strconv.ParseBool(v)
}

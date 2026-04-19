package main

import (
	"os"
	"strconv"
)

type ParserImpl struct {
	v string
}

func GetFlagFromEnv(key string) ParserImpl {
	v, set := os.LookupEnv(key)

	switch set {
	case false:
		return ParserImpl{}
	default:
		if v == "" {
			return ParserImpl{}
		}

		return ParserImpl{
			v: v,
		}
	}
}

func (p ParserImpl) ToInt() (int, error) {
	return strconv.Atoi(p.v)
}

func (p ParserImpl) ToFloat64() (float64, error) {
	return strconv.ParseFloat(p.v, 64)
}

func (p ParserImpl) ToBool() (bool, error) {
	return strconv.ParseBool(p.v)
}

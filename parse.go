package main

import (
	"strconv"
	"strings"
)

func parseFlagToModel(v string) (*Flag, error) {
	splitV := strings.Split(v, ":")

	category := splitV[0]
	feature := splitV[1]
	state, err := strconv.ParseBool(splitV[2])
	value := splitV[3]
	if err != nil {
		return nil, err
	}

	return &Flag{
		Category: category,
		Feature:  feature,
		State:    state,
		Value:    value,
	}, nil
}

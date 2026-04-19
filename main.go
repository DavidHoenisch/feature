package main

type Feature interface {
	GetFeatureModel() (*Flag, error)
	GetFeatureState() bool
	GetFeatureValue() (any, error)
}

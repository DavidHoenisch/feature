package main

// Model for marshalling feature flags
//
//	.[ops]: category
//	.[something]: the feature
//	.[state]: on, off (true, false)
type Flag struct {
	Category string
	Feature  string
	State    bool
}

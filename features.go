package main

import "time"

type feature struct {
	collection string
	featureId string
	validity time.Time
	action string
	attributes string
}

func (f feature) equals(other feature) bool {
	return f.collection == other.collection &&
		f.featureId == other.featureId
}


type featureInfo struct {
	collection string
	featureId string
	valid bool
	before, after, invalidBefore, invalidAfter int

}

package main

import "gonum.org/v1/gonum/graph/simple"

type UnitNode struct {
	simple.Node
	Name string `json:"name"`
	Deep uint8  `json:"deep"`
}

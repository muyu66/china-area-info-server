package main

import (
	"gonum.org/v1/gonum/graph/simple"
)

func getRootNodes(dg *simple.DirectedGraph) []UnitNode {
	return getNodesById(dg, 0)
}

func getNodesById(dg *simple.DirectedGraph, id int64) []UnitNode {
	var res []UnitNode
	var nodes = dg.To(id)
	for node := nodes; nodes.Next(); {
		var currNode = node.Node().(UnitNode)
		res = append(res, currNode)
	}
	return res
}

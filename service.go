package main

import (
	"cmp"
	"errors"
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"slices"
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

func sort(nodes *[]UnitNode) {
	slices.SortFunc(*nodes, func(i, j UnitNode) int {
		return cmp.Compare(i.ID(), j.ID())
	})
}

func checkDeep(dg *simple.DirectedGraph, nodeId int64, expectedDeep uint8) bool {
	node := dg.Node(nodeId).(UnitNode)
	return node.Deep == expectedDeep
}

func checkExist(dg *simple.DirectedGraph, nodeId int64) bool {
	node := dg.Node(nodeId)
	return node != nil
}

func getAllParentNodesById(dg *simple.DirectedGraph, id int64) (map[uint8]UnitNode, error) {
	var nodes = dg.From(id)
	var m = make(map[uint8]UnitNode)

	// 添加最底层
	var currNode = dg.Node(id).(UnitNode)
	m[currNode.Deep] = currNode

	// 往上找父节点
	err := recursionAllParentNodesById(dg, nodes, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func recursionAllParentNodesById(dg *simple.DirectedGraph, nodes graph.Nodes, m map[uint8]UnitNode) error {
	if nodes.Len() > 1 {
		return errors.New(fmt.Sprintf("递归异常, 父节点不明"))
	}
	for node := nodes; nodes.Next(); {
		if node.Node().ID() == 0 {
			break
		}
		var currNode = node.Node().(UnitNode)
		m[currNode.Deep] = currNode
		err := recursionAllParentNodesById(dg, dg.From(currNode.ID()), m)
		if err != nil {
			return err
		}
	}
	return nil
}

package datastructure

import (
	"fmt"
)

type TreeNode struct {
	Id           string
	ParentId     string
	Level        int
	Name         string
	Data         string
	Childs       Tree
	ChildNodeTag int
}

// 节点下是否包含子节点，返回子节点数量
func (node *TreeNode) HasChilds() int {
	return len((*node).Childs)
}

/*
树的层级，从1开始，使用该方法，会自动根据当前节点的Level+1
每个节点ChildNodeTag最初都是0，当向其内增加节点时，会自动+1，然后自动应该用到子节点ID上，
这样做主要是避免删除、添加节点时，签名重复。
*/
func (node *TreeNode) AddChildNode(name, data string) {
	(*node).ChildNodeTag++
	n := TreeNode{
		Id:           fmt.Sprintf("%s_%d", (*node).Id, (*node).ChildNodeTag),
		ParentId:     (*node).Id,
		Level:        (*node).Level + 1,
		Name:         name,
		Data:         data,
		Childs:       make([]*TreeNode, 0),
		ChildNodeTag: 0,
	}
	(*node).Childs = append((*node).Childs, &n)
}

// 调试用，可以递归从当前节点向下输出子节点信息
func (node *TreeNode) DebugPrintAllChilds(indent int) {
	for _, n := range node.Childs {
		fmt.Printf("[L%d]%s  %s\n", (*n).Level, (*n).Id, (*n).Data)
		n.DebugPrintAllChilds(indent + 2)
	}
}

// 新建节点返回结构体
func NewNode(id, name, data string) TreeNode {
	n := TreeNode{
		Id:           id,
		ParentId:     "",
		ChildNodeTag: 0,
		Level:        1,
		Name:         name,
		Data:         data,
		Childs:       make([]*TreeNode, 0),
	}
	return n
}

// ======Tree=======================================================================================
type Tree []*TreeNode

func MakeTree() Tree {
	tree := make([]*TreeNode, 0)
	return tree
}

// 按节点id查找节点
func (tree *Tree) FindNodeByName(nodeid string) (TreeNode, bool) {
	for _, node := range *tree {
		if nodeid == node.Id {
			return *node, true
		}
		if node.HasChilds() > 0 {
			n, ok := node.Childs.FindNodeByName(nodeid)
			if ok {
				return n, ok
			}
		}
	}
	var n TreeNode
	return n, false
}

// 根据节点查找父节点
func (tree *Tree) FindParentNode(node TreeNode) (TreeNode, bool) {
	find_level := node.Level - 1
	if find_level < 1 {
		// 如果尝试在根节点上查找其父节点，返回未找到，同时将传入节点返回
		return node, false
	}
	find_id := node.ParentId
	return tree.FindNodeByName(find_id)
}

func (tree *Tree) DeleteNode(node TreeNode) bool {
	parent_level := node.Level - 1
	if parent_level == 0 {
		remove_id := -1
		// 当前要处理的节点在根
		for i := 0; i < len(*tree); i++ {
			if node.Id == ((*tree)[i]).Id {
				remove_id = i
				break
			}
		}
		if remove_id == -1 {
			return false
		} else {
			*tree = append((*tree)[:remove_id], (*tree)[remove_id+1:]...)
			return true
		}
	}
	// 先找到其父节点
	remove_id := -1
	parentNode, ok := tree.FindParentNode(node)
	if !ok {
		return false
	}

	// 当前要处理的节点在根
	for i := 0; i < len(parentNode.Childs); i++ {
		if node.Id == (parentNode.Childs[i]).Id {
			remove_id = i
			break
		}
	}
	if remove_id == -1 {
		return false
	}
	parentNode.Childs = append((parentNode.Childs)[:remove_id], (parentNode.Childs)[remove_id+1:]...)
	return true
}

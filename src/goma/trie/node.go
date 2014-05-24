package trie

type NodeNode struct {
	Base *BaseNode
	Chck *ChckNode
}

type BaseNode struct {
	INIT_VALUE int32
}

func (this *BaseNode) ID(id int32) int32 {
	return id*-1 - 1
}

type ChckNode struct {
	TERMINATE_CODE uint16
	VACANT_CODE    uint16
	CODE_LIMMIT    uint16
}

var Node *NodeNode

func init() {
	Node = &NodeNode{}
	Node.Base = &BaseNode{}
	Node.Base.INIT_VALUE = -2147483648

	Node.Chck = &ChckNode{}
	Node.Chck.TERMINATE_CODE = 0
	Node.Chck.VACANT_CODE = 0
	Node.Chck.CODE_LIMMIT = 0xFFFF
}

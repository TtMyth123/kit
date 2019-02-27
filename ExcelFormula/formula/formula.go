package formula

import (
	"fmt"
	"strconv"

	"github.com/xuri/efp"
)

// NodeType Descriptor of node in Formula1 AST
type NodeType int8

const (
	// NodeTypeRoot Root of AST
	NodeTypeRoot NodeType = iota + 1
	// NodeTypeLiteral Literal node, value should be understood as is
	NodeTypeLiteral
	// NodeTypeInteger Literal node, value should be understood as is
	NodeTypeInteger
	// NodeTypeFloat Literal node, value should be understood as is
	NodeTypeFloat
	// NodeTypeRef Reference node, value should be dereffed
	NodeTypeRef
	// NodeTypeFunc Function call node, value must be executed
	NodeTypeFunc
	// NodeTypeOperator Infix operator
	NodeTypeOperator
)

var PRECEDENCE = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

// Node AST node
type Node struct {
	value      interface{}
	nodeType   NodeType
	children   []*Node
	parent     *Node
	infixChild *Node
}

// Formula Formula1 executable formula
type Formula struct {
	root Node
}

// NewFormula Create a new formula instance
func NewFormula(text string) *Formula {
	parser := efp.ExcelParser()
	parser.Parse(text)
	fmt.Printf("%s\n=====\n", parser.PrettyPrint())

	tokens := parser.Tokens.Items
	root := Node{
		value:    "root",
		nodeType: NodeTypeRoot,
		children: nil,
	}

	current := &root
	index := 0
	count := len(tokens)

	var token *efp.Token

	for index <= count {
		if index >= count {
			if current.infixChild != nil {
				if current.children == nil {
					current.children = []*Node{current.infixChild}
				} else {
					current.children = append(current.children, current.infixChild)
				}
				current.resetInfixChild()
			}

			if root.infixChild != nil {
				if root.children == nil {
					root.children = []*Node{root.infixChild}
				} else {
					root.children = append(root.children, root.infixChild)
					root.resetInfixChild()
				}
			}
			break
		}

		token = &tokens[index]
		tvalue := token.TValue
		ttype := token.TType
		tsubtype := token.TSubType
		var value interface{}
		var nodeType NodeType

		if tsubtype == efp.TokenSubTypeStart {
			value, nodeType = resolveNodeType(ttype, tsubtype, tvalue)
			if current.infixChild != nil {
				parent := current
				current = current.infixChild.makeNode(nodeType, value)
				current.parent = parent
			} else {
				current = current.makeNode(nodeType, value)
			}

			index++
			continue
		} else if tsubtype == efp.TokenSubTypeStop {
			if current.infixChild != nil {
				current.children = append(current.children, current.infixChild)
				current.resetInfixChild()
			}

			current = current.parent

			index++
			continue
		} else if ttype == efp.TokenTypeArgument {
			if current.infixChild != nil {
				current.children = append(current.children, current.infixChild)
				current.resetInfixChild()
			}
			index++
			continue
		} else if ttype == efp.TokenTypeOperatorInfix {
			current.makeInfixChild(tvalue)

			index++
			continue
		} else if ttype == efp.TokenTypeOperand {
			value, nodeType = resolveNodeType(ttype, tsubtype, tvalue)
			if current.infixChild != nil {
				current.infixChild.makeNode(nodeType, value)
			} else {
				current.makeNode(nodeType, value)
			}

			index++
			continue
		} else {
			index++

			continue
		}
	}

	formula := Formula{
		root: root,
	}
	return &formula
}

func (formula *Formula) GetEntryNode() *Node {
	return formula.root.children[0]
}

func (this *Node) makeNode(nodeType NodeType, value interface{}) *Node {
	node := Node{
		value:    value,
		nodeType: nodeType,
		parent:   this,
		children: nil,
	}
	if this.children == nil {
		this.children = []*Node{&node}
	} else {
		this.children = append(this.children, &node)
	}
	//100+300*200+500
	if (this.children[0].nodeType == NodeTypeOperator) && (this.nodeType == NodeTypeOperator) {
		if PRECEDENCE[this.value.(string)] > PRECEDENCE[this.children[0].value.(string)] {
			//
			node0_0 := this.children[0].children[0]
			node0_1 := this.children[0].children[1]
			node0 := this.children[0]

			curNodeOperator := Node{
				value:    this.value,
				nodeType: this.nodeType,
				parent:   this,
				children: []*Node{
					&node, node0_1,
				},
			}
			this.children = []*Node{&curNodeOperator, node0_0}
			this.value = node0.value
		}
	}

	return &node
}

func (this *Node) makeInfixChild(value string) *Node {
	if this.infixChild == nil {
		this.infixChild = &Node{
			value:    value,
			nodeType: NodeTypeOperator,
			parent:   this,
			children: []*Node{},
		}
		if len(this.children) > 0 {
			lastChild := this.LastChild()
			this.children = this.children[0 : this.ChildCount()-1]
			this.infixChild.children = []*Node{lastChild}
		}
		//} else if this.infixChild.value != value {
	} else {
		//// Infix precedence resolution
		//if PRECEDENCE[value] > PRECEDENCE[this.infixChild.value.(string)] {
		//	//begin------------------------------------
		//	//// Detach the last child and append it to the new node's children
		//	//temp := this.infixChild
		//	//node := &Node{
		//	//	value:    value,
		//	//	nodeType: NodeTypeOperator,
		//	//	parent:   this,
		//	//	children: []*Node{
		//	//		temp,
		//	//		temp.children[temp.ChildCount()-1],
		//	//	},
		//	//}
		//	//temp.children[temp.ChildCount()-1].parent = node
		//	//temp.children = temp.children[:temp.ChildCount()-2]
		//	//
		//	//this.infixChild = node
		//	//end------------------------------------
		//	//this.infixChild =
		//
		//	temp := this.infixChild
		//	node := &Node{
		//		value:    value,
		//		nodeType: NodeTypeOperator,
		//		parent:   this,
		//		children: []*Node{
		//			temp.children[temp.ChildCount()-1],
		//		},
		//	}
		//	this.infixChild = node
		//
		//} else {
		//	// Simply wrap the existing infixChild inside the new one
		//	temp := this.infixChild
		//	node := &Node{
		//		value:    value,
		//		nodeType: NodeTypeOperator,
		//		parent:   this,
		//		children: []*Node{
		//			temp,
		//		},
		//	}
		//	this.infixChild = node
		//}

		temp := this.infixChild
		node := &Node{
			value:    value,
			nodeType: NodeTypeOperator,
			parent:   this,
			children: []*Node{
				temp,
			},
		}
		this.infixChild = node
	}

	return this.infixChild
}

func (this *Node) resetInfixChild() {
	this.infixChild = nil
}

func resolveNodeType(ttype string, tsubtype string, tvalue string) (value interface{}, nodeType NodeType) {
	var _err error
	if ttype == efp.TokenTypeFunction && tsubtype == efp.TokenSubTypeStart {
		nodeType = NodeTypeFunc
		value = tvalue
		return
	} else if ttype == efp.TokenTypeSubexpression && tsubtype == efp.TokenSubTypeStart {
		nodeType = NodeTypeFunc
		value = "IDENTITY"
		return
	} else if ttype == efp.TokenTypeOperand && tsubtype == efp.TokenSubTypeRange {
		nodeType = NodeTypeRef
		value = tvalue
		return
	} else if ttype == efp.TokenTypeOperand && tsubtype == efp.TokenSubTypeText {
		nodeType = NodeTypeLiteral
		value = tvalue
		return
	} else if ttype == efp.TokenTypeOperand && tsubtype == efp.TokenSubTypeNumber {
		nodeType = NodeTypeFloat
		value, _err = strconv.ParseFloat(tvalue, 64)
		if _err != nil {
			return
		}
		return
	}

	value = tvalue
	nodeType = NodeTypeLiteral
	return
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) NodeType() NodeType {
	return node.nodeType
}

func (this *Node) ChildCount() int {
	if this.children == nil {
		return 0
	}
	return len(this.children)
}

func (this *Node) FirstChild() *Node {
	if this.children == nil || len(this.children) == 0 {
		return nil
	}

	return this.children[0]
}

func (this *Node) LastChild() *Node {
	if this.children == nil || len(this.children) == 0 {
		return nil
	}

	return this.children[len(this.children)-1]
}

func (this *Node) ChildAt(index int) *Node {
	if index < 0 {
		return nil
	} else if index >= len(this.children) {
		return nil
	}

	return this.children[index]
}

func (this *Node) HasChildren() bool {
	if this.children == nil {
		return false
	} else if len(this.children) == 0 {
		return false
	}

	return true
}

func (this *Node) Children() []*Node {
	if this.children == nil {
		return make([]*Node, 0)
	}
	return this.children[:]
}

package goqueryKit

import (
	"github.com/opesun/goquery"
	"github.com/opesun/goquery/exp/html"
	"strings"
)

type TdNode struct {
	Text string
	Node *html.Node
}

func GetNodeText(node *html.Node) string {
	s  := ""
	for _,n := range node.Child{
		//if n.Type ==  html.TextNode {
		//	s = s+n.Data
		//}
		s = s+GetNodeText(n)
	}
	if node.Type==html.TextNode {
		s = s+node.Data
	}
	return s
}

func GetTableNodes2Arr(tableNode goquery.Nodes) [][]TdNode{
	arrData := make([][]TdNode,0)
	//trNodes := tableNode.Find("tr")
	tableNode.Find("tr").Each(func (i int,trNode *goquery.Node) {

		iLen := len(trNode.Child)
		arrTd := make([]TdNode,0)
		for i:=0;i<iLen;i++{
			if trNode.Child[i].Data == "td" {
				aTdNode := TdNode{Text: GetNodeText(trNode.Child[i]),Node: trNode.Child[i]}
				arrTd = append(arrTd,aTdNode)
			}
		}
		arrData = append(arrData,arrTd)
	})

	return arrData
}

func GetAttrMap(arrAttr []html.Attribute)map[string]string {
	mp := make(map[string]string)
	for _,attr := range arrAttr{
		key:= strings.ToLower(attr.Key)
		Value:= attr.Val
		mp[key] = Value
	}

	return mp
}
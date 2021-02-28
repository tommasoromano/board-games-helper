package main

// repl.it
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type syNode struct{
    ticket []int
    next []int
}

type syMap struct{
    nodes []syNode
}
// 0=taxi, 1=bus, 2=metro, 3=taxi/bus, 4=bus/metro
var m = syMap{ 
	[]syNode{
		{[]int{0,0,4,1}, 		[]int{8,9,46,58}},
		{[]int{0,0}, 			[]int{20,10}},
		{[]int{0,0,0,1,1}, 		[]int{11,12,4,22,23}}, 
		{[]int{0,0}, 			[]int{3,13}},
		{[]int{0,0}, 			[]int{15,16}},
		{[]int{0,0}, 			[]int{29,7}},
		{[]int{0,0,1}, 			[]int{6,17,42}},
		{[]int{0,0,0}, 			[]int{1,18,19}},
		{[]int{0,0,0}, 			[]int{1,19,20}},
		{[]int{0,0,0,0}, 		[]int{2,21,34,11}},					//10
		{[]int{0,0,0}, 			[]int{10,3,22}},
		{[]int{0,0,0}, 			[]int{3,23}},
		{[]int{0,3,1,0,1,2,2,2},[]int{4,23,14,24,52,46,89,67}},
		{[]int{1,3,0}, 			[]int{13,15,25}},
		{[]int{0,0,3,3,0,1}, 	[]int{5,26,14,28,16,29}},
		{[]int{0,0,0,0}, 		[]int{5,28,15,29}},
		{[]int{0,0,0}, 			[]int{7,29,30}},
		{[]int{0,0,0}, 			[]int{8,31,42}},
		{[]int{0,0,0}, 			[]int{8,9,32}},
		{[]int{0,0,0}, 			[]int{9,2,33}},						//20
		{[]int{0,0}, 			[]int{33,10}},
		{[]int{0,0,3,3,1,1}, 	[]int{11,35,23,34,3,65}},
		{[]int{0,0,1,1,3,3}, 	[]int{12,37,3,67,22,13}},
		{[]int{0,0,0}, 			[]int{13,37,38}},
		{[]int{0,0,0}, 			[]int{14,38,39}},
		{[]int{0,0,0}, 			[]int{15,39,27}},
		{[]int{0,0,0}, 			[]int{26,40,28}},
		{[]int{0,0,0,3,3,1,1}, 	[]int{6,17,16,41,42,55,15}},
		{[]int{0,0}, 			[]int{17,42}},						//30
	},
}
func calcNextNodes(ticket string, nodes []int)(newNodes []int) {
	newNodes = make([]int, 0)
	fmt.Println(nodes)
	// loop for input nodes to check
	for i:=0;i<len(nodes);i++{
		// loop for next on map for node[i]
		//fmt.Print(len(m.nodes[nodes[i]-1].ticket))
		for j:=0;j<len(m.nodes[nodes[i]-1].ticket);j++{
			// divide for ticket
			hasNext:=false
			if ticket=="?" {
				hasNext=true
			}
			if m.nodes[nodes[i]-1].ticket[j]==0 {
				if ticket=="t" {
					hasNext=true
				}
			} else if m.nodes[nodes[i]-1].ticket[j]==1 {
				if ticket=="b" {
					hasNext=true
				}
			} else if m.nodes[nodes[i]-1].ticket[j]==2 {
				if ticket=="m" {
					hasNext=true
				}
			} else if m.nodes[nodes[i]-1].ticket[j]==3 {
				if ticket=="t"||ticket=="b" {
					hasNext=true
				}
			} else if m.nodes[nodes[i]-1].ticket[j]==4 {
				if ticket=="b"||ticket=="m" {
					hasNext=true
				}
			} else {
				
			}
			if hasNext {
				if !isInNodes(m.nodes[nodes[i]-1].next[j],newNodes) {
					newNodes = append(newNodes, m.nodes[nodes[i]-1].next[j])
				}
			}
			
		}
	}
	fmt.Print(newNodes)
	return newNodes
}
func isInNodes(i int, n []int)(bool) {
	for j:=0;j<len(n);j++{
		if n[j]==i {
			return true
		}
	}
	return false;
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	text,_:= reader.ReadString('\n')
	wrds:=strings.Split(text," ")
	nodeStart,_:=strconv.Atoi(wrds[0])
	finalNodes:=make([]int,0)
	finalNodes=append(finalNodes, nodeStart)
	for i:=1;i<len(wrds);i++ {
		finalNodes=calcNextNodes(wrds[i],finalNodes)
	}
	fmt.Print(finalNodes)
}
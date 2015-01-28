package algorithm

//SkipList
//author:Xiong Chuan Liang
//date:2014-1-28
//"github.com/xcltapestry/xclpkg/algorithm"

import (
	//"fmt"
	"math/rand"
)


const SKIPLIST_MAXLEVEL = 32 //8 
const SKIPLIST_P = 4


type Node struct {
	Forward  []Node
	Value interface{}
}


func NewNode(v interface{},level int) *Node {
	return &Node{Value: v, Forward: make([]Node, level)} 
}


type SkipList struct {
	Header *Node
	Level  int
}

func NewSkipList() *SkipList {
	return &SkipList{Level: 1, Header: NewNode(0,SKIPLIST_MAXLEVEL)}   
}

func (skipList *SkipList) Insert(key int) {

	update := make(map[int]*Node)
	node := skipList.Header

	for i := skipList.Level - 1; i >= 0; i-- {
		 for {
		 	 if node.Forward[i].Value != nil && node.Forward[i].Value.(int) < key {
		 	 	node = &node.Forward[i]
		 	 }else{
		 	 	break;
		 	 }
		 }
		 update[i] = node
	}

	level := skipList.Random_level()
	if level > skipList.Level {
		for i := skipList.Level; i < level; i++ {
			update[i] = skipList.Header
		}
		skipList.Level = level
	}

	newNode := NewNode(key,level)

	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]		
		update[i].Forward[i] = *newNode			
	}

}

func (skipList *SkipList) Random_level() int {

	level := 1
	for (rand.Int31()&0xFFFF)%SKIPLIST_P == 0 {
		level += 1
	}
	if level < SKIPLIST_MAXLEVEL {
		return level
	} else {
		return SKIPLIST_MAXLEVEL
	}
}

func (skipList *SkipList) PrintSkipList() {

	//fmt.Println("\nSkipList-------------------------------------------")
	for i := SKIPLIST_MAXLEVEL - 1; i >= 0; i-- {
		
		 //fmt.Println("level:", i)
		 node := skipList.Header.Forward[i]
		 for {
		 	 if node.Value != nil {
		 	 	//fmt.Printf("%d ", node.Value.(int))
		 	 	node = node.Forward[i]
		 	 }else{
		 	 	break;
		 	 }
		 }
		 //fmt.Println("\n--------------------------------------------------------")
	} //end for

	//fmt.Println("Current MaxLevel:", skipList.Level)
}


func (skipList *SkipList) Search(key int) *Node {

	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {

		//fmt.Println("\n Search() Level=", i)
		for {
			if node.Forward[i].Value == nil {
				break
			}

			//fmt.Printf("  %d ", node.Forward[i].Value)
			if node.Forward[i].Value.(int) == key {
				//fmt.Println("\nFound level=", i, " key=", key)			
				return &node.Forward[i]
			}

			if  node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			}else{ // > key 
				break
			}
		 } //end for find

	} //end level
	return nil
}


func (skipList *SkipList)Remove(key int) {

	update := make(map[int]*Node)
	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {
		
		for {

			if node.Forward[i].Value == nil {
				break
			}

			if node.Forward[i].Value.(int) == key {
				//fmt.Println("Remove() level=", i, " key=", key)			
				update[i] = node 
				break
			}

			if  node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			}else{ // > key 
				break
			}

		 } //end for find

	} //end level

	for i,v := range update {
		if v == skipList.Header {
			skipList.Level --
		}
 		v.Forward[i] = v.Forward[i].Forward[i]		
	}
}

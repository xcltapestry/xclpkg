package algorithm

//SkipList
//author:Xiong Chuan Liang
//date:2014-1-27

import (
	"fmt"
	"math/rand"
)

const SKIPLIST_MAXLEVEL = 8 //32
const SKIPLIST_P = 4

type SkipList struct {
	Header []List
	Level  int
}

func NewSkipList() *SkipList {
	return &SkipList{Level: 1, Header: make([]List, SKIPLIST_MAXLEVEL)}

}

func (skipList *SkipList) Insert(key int) {
	update := make(map[int]*Node)

	for i := len(skipList.Header) - 1; i >= 0; i-- {
		if skipList.Header[i].Len() > 0 {
			for e := skipList.Header[i].Front(); e != nil; e = e.Next() {
				if e.Value.(int) >= key {
					update[i] = e
					break
				}
			}
		} //Heaer[lv].List
	} //Header Level

	level := skipList.Random_level()
	if level > skipList.Level {
		skipList.Level = level
	}

	for i := 0; i < level; i++ {
		if v, ok := update[i]; ok {
			skipList.Header[i].InsertBefore(key, v)
		} else {
			skipList.Header[i].PushBack(key)
		}
	}

}

func (skipList *SkipList) Search(key int) *Node {

	for i := len(skipList.Header) - 1; i >= 0; i-- {
		if skipList.Header[i].Len() > 0 {
			for e := skipList.Header[i].Front(); e != nil; e = e.Next() {
				switch {
				case e.Value.(int) == key:
					fmt.Println("Found level=", i, " key=", key)
					return e
				case e.Value.(int) > key:
					break
				}
			} //end for

		} //end if

	} //end for
	return nil
}

func (skipList *SkipList) Delete(key int) {
	for i := len(skipList.Header) - 1; i >= 0; i-- {
		if skipList.Header[i].Len() > 0 {
			for e := skipList.Header[i].Front(); e != nil; e = e.Next() {
				switch {
				case e.Value.(int) == key:
					fmt.Println("Delete level=", i, " key=", key)
					skipList.Header[i].remove(e)

				case e.Value.(int) > key:
					break
				}
			} //end for

		} //end if

	} //end for

}

func (skipList *SkipList) PrintSkipList() {
	fmt.Println("\nSkipList-------------------------------------------")
	for i := SKIPLIST_MAXLEVEL - 1; i >= 0; i-- {
		fmt.Println("level:", i)
		if skipList.Header[i].Len() > 0 {
			for e := skipList.Header[i].Front(); e != nil; e = e.Next() {
				fmt.Printf("%d ", e.Value)
			} //end for
		} //end if
		fmt.Println("\n--------------------------------------------------------")
	} //end for
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

//////////////////////////////////////////////////////

type Node struct {
	next, prev *Node
	list       *List

	Value interface{}
}

type List struct {
	root Node
	len  int
}

func (node *Node) Next() *Node {
	if p := node.next; node.list != nil && p != &node.list.root {
		return p
	}
	return nil
}

func (node *Node) Prev() *Node {
	if p := node.prev; node.list != nil && p != &node.list.root {
		return p
	}
	return nil
}

func (list *List) Init() *List {
	list.root.next = &list.root
	list.root.prev = &list.root
	list.len = 0
	return list
}

func New() *List {
	return new(List).Init()
}

func (list *List) lazyInit() {
	if list.root.next == nil {
		list.Init()
	}
}

func (list *List) Len() int {
	return list.len
}

func (list *List) remove(e *Node) *Node {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	list.len--
	return e
}

func (list *List) Remove(e *Node) interface{} {
	if e.list == list {
		list.remove(e)
	}
	return e.Value
}

func (list *List) Front() *Node {
	if list.len == 0 {
		return nil
	}
	return list.root.next
}

func (list *List) Back() *Node {
	if list.len == 0 {
		return nil
	}
	return list.root.prev
}

func (list *List) insert(e, at *Node) *Node {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = list

	list.len++
	return e
}

func (list *List) insertValue(v interface{}, at *Node) *Node {

	return list.insert(&Node{Value: v}, at)
}

func (list *List) InsertBefore(v interface{}, mark *Node) *Node {
	if mark.list != list {
		return nil
	}
	return list.insertValue(v, mark.prev)
}

func (list *List) PushBack(v interface{}) *Node {
	list.lazyInit()
	return list.insertValue(v, list.root.prev)
}

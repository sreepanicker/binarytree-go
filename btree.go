package main

import "fmt"

import "strconv"
//structure for btree
type btree struct {
	data int
	lptr *btree
	rptr *btree
}

func (l *btree) addlLink(left *btree) {
	l.lptr = left
}
func (l *btree) addrLink(right *btree) {
	l.rptr = right
}
func createTree(first *btree, node *btree) {
	if first.data > node.data {
		//greater
		if first.lptr != nil {
			first = first.lptr
			createTree(first, node)
		} else {
			firstData := strconv.Itoa(first.data)
			nodeData := strconv.Itoa(node.data)
			fmt.Println("left " + firstData + " " + nodeData)
			first.lptr = node
		}
	} else {
		// lesser
		if first.rptr != nil {
			first = first.rptr
			createTree(first, node)
		} else {
			firstData := strconv.Itoa(first.data)
			nodeData := strconv.Itoa(node.data)
			fmt.Println("Right " + firstData + " " + nodeData)
			first.rptr = node
		}
	}
}
func displayData(start *btree) {

	if start.lptr != nil {
		displayData(start.lptr)
	}
	fmt.Println(start.data)
	if start.rptr != nil {
		displayData(start.rptr)
	}

}

func main() {

	fmt.Println("Strating the program ")
	//str
	first := btree{100, nil, nil}
	next := btree{50, nil, nil}
	createTree(&first, &next)
	third := btree{60, nil, nil}
	createTree(&first, &third)
	fourth := btree{200, nil, nil}
	createTree(&first, &fourth)
	fith := btree{30, nil, nil}
	createTree(&first, &fith)
	sixth := btree{70, nil, nil}
	createTree(&first, &sixth)
	seven := btree{150, nil, nil}
	createTree(&first, &seven)
	// displaydata
	displayData(&first)
}

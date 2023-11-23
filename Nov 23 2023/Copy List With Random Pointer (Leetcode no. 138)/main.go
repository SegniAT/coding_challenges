package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// too complicated
func copyRandomListDumb(head *Node) *Node {
	res := &Node{}
	resIter, mainIter := res, head
	oldToNewMap := make(map[*Node]*Node)

	for mainIter != nil {
		var curr *Node
		if new, ok := oldToNewMap[mainIter]; ok {
			curr = new
		} else {
			curr = &Node{
				Val: mainIter.Val,
			}
		}

		if _, ok := oldToNewMap[mainIter]; !ok {
			oldToNewMap[mainIter] = curr
		}

		var rand *Node
		if new, ok := oldToNewMap[mainIter.Random]; ok {
			rand = new
		} else {
			if mainIter.Random != nil {
				rand = &Node{
					Val: mainIter.Random.Val,
				}

				oldToNewMap[mainIter.Random] = rand
			}
		}

		curr.Random = rand

		resIter.Next = curr

		mainIter = mainIter.Next
		resIter = resIter.Next
	}

	return res.Next
}

func copyRandomList(head *Node) *Node {
	oldToNewMap := make(map[*Node]*Node)
	iter := head

	for iter != nil {
		cloned := &Node{
			Val: iter.Val,
		}
		oldToNewMap[iter] = cloned
		iter = iter.Next
	}

	iter = head
	for iter != nil {
		oldToNewMap[iter].Next = oldToNewMap[iter.Next]
		oldToNewMap[iter].Random = oldToNewMap[iter.Random]
		iter = iter.Next
	}

	return oldToNewMap[head]
}

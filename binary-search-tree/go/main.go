package main

import "fmt"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Step 1: I will start by writing a single-node search function,
// considering I like to think one step at a time.
func nodeSearch(node *BST, target int, closest int) int {

	// Step 2: At first I thought what if the node is null? So I wrote a nil check here.
	// But then I realized this is unreachable code.
	// In Step 0 we already check if tree is nil before calling this function.
	// And in Step 5 we already check if the child is nil before we recurse.
	// So a nil node can never happen in this check. I will keep it because I'm also recording my brain thinking workflow.
	// if node == nil {
	// 	return closest
	// }

	// Step 3: First of all, because I'm lazy,
	// I will assume that I'm lucky and the target value is actually equal to the node's value.
	// If so, I will just return it and end the program.
	if target == node.Value {
		// I don't think you can be closer than the same value.
		return target
	}

	// Step 4: Compare distances and update closest if the current node is nearer to the target.
	if abs(target-node.Value) < abs(target-closest) {
		closest = node.Value
	}

	// Step 5: Now we get two options.
	//
	// Option A: target > node.Value
	//   This means the target is also bigger than all the left nodes.
	//   We don't need to search on the left anymore
	//   because we would only be getting farther from our target.
	//   So we go right and call the function recursively.
	//
	// Option B: target < node.Value
	//   The only other logical option, we go left.
	if target > node.Value {
		// Step 5a: Take the right child, but there is an edge case.
		// The right child might not even exist.
		if node.Right == nil {
			return closest
		}
		return nodeSearch(node.Right, target, closest)
	} else {
		// Step 5b: Same logic, opposite direction, go left.
		// Don't expect me to write target < node.Value explicitly
		// because what other logical option could there be?
		if node.Left == nil {
			return closest
		}
		return nodeSearch(node.Left, target, closest)
	}
}

// Step 0: Entry point. Start the search from the root.
func (tree *BST) FindClosestValue(target int) int {
	// edge case.
	if tree == nil {
		return -1
	}
	return nodeSearch(tree, target, tree.Value)
}

func main() {
	// Case 1: The sample tree from the problem (target=12, expected=13)
	//
	//        10
	//       /  \
	//      5    15
	//     / \   / \
	//    2   5 13  22
	//   /       \
	//  1        14
	//
	tree1 := &BST{Value: 10,
		Left: &BST{Value: 5,
			Left:  &BST{Value: 2, Left: &BST{Value: 1}, Right: nil},
			Right: &BST{Value: 5},
		},
		Right: &BST{Value: 15,
			Left:  &BST{Value: 13, Left: nil, Right: &BST{Value: 14}},
			Right: &BST{Value: 22},
		},
	}
	fmt.Printf("Case 1: target=12, got=%d, expected=13\n", tree1.FindClosestValue(12))

	// Case 2: The breaking tree where old algo fails (target=12, expected=10)
	//
	//      10
	//     /  \
	//    5    15
	//   / \
	//  2   8
	//
	tree2 := &BST{Value: 10,
		Left: &BST{Value: 5,
			Left:  &BST{Value: 2},
			Right: &BST{Value: 8},
		},
		Right: &BST{Value: 15},
	}
	fmt.Printf("Case 2: target=12, got=%d, expected=10\n", tree2.FindClosestValue(12))

	// Case 3: Target is exact match (target=5, expected=5)
	fmt.Printf("Case 3: target=5,  got=%d, expected=5\n", tree1.FindClosestValue(5))

	// Case 4: Target is smaller than everything (target=-100, expected=1)
	fmt.Printf("Case 4: target=-100, got=%d, expected=1\n", tree1.FindClosestValue(-100))

	// Case 5: Target is bigger than everything (target=9999, expected=22)
	fmt.Printf("Case 5: target=9999, got=%d, expected=22\n", tree1.FindClosestValue(9999))

	// Case 6: Single node tree (target=42, expected=7)
	tree3 := &BST{Value: 7}
	fmt.Printf("Case 6: target=42, got=%d, expected=7\n", tree3.FindClosestValue(42))
}

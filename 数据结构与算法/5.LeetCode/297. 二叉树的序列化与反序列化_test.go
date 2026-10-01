package leetcode

import (
	. "backend/utils/algo"
	"container/list"
	"encoding/json"
	"strconv"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ans := make([]string, 0)
	if root != nil {
		queue := list.New()
		ans = append(ans, strconv.Itoa(root.Val))
		queue.PushBack(root)
		for queue.Len() > 0 {
			cur := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())

			if cur.Left != nil {
				ans = append(ans, strconv.Itoa(cur.Left.Val))
				queue.PushBack(cur.Left)
			} else {
				ans = append(ans, "null")
			}

			if cur.Right != nil {
				ans = append(ans, strconv.Itoa(cur.Right.Val))
				queue.PushBack(cur.Right)
			} else {
				ans = append(ans, "null")
			}
		}
	}

	// 过滤掉末尾的null
	strs, _ := json.Marshal(ans)
	return string(strs)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var ans []string
	json.Unmarshal([]byte(data), &ans)
	if len(ans) == 0 {
		return nil
	}
	root := generateTreeNode(ans[0])
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}

	index := 1
	for queue.Len() > 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())

		if index == len(ans) {
			continue
		}

		node.Left = generateTreeNode(ans[index])
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		index++
		node.Right = generateTreeNode(ans[index])
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		index++
	}

	return root
}

func generateTreeNode(str string) *TreeNode {
	if str == "null" {
		return nil
	}
	val, _ := strconv.Atoi(str)
	return &TreeNode{Val: val}
}

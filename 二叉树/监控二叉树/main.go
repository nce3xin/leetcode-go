package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	UNCOVER = 0
	COVER   = 1
	CAMERA  = 2
)

func minCameraCover(root *TreeNode) int {
	res := 0
	rootStatus := setCamera(root, &res)
	// 这里千万不要漏了
	// 当根节点未被覆盖时，根节点处需要安装一个摄像头，因为它没有父节点了
	if rootStatus == UNCOVER {
		res++
	}
	return res
}

// 返回root节点的状态, 一共有 3 种状态: 1. 被覆盖 2. 未覆盖 3. 放置了摄像头
func setCamera(root *TreeNode, res *int) int {
	if root == nil {
		// 空节点，只有状态是被覆盖，才能保证叶子节点无需放置摄像头，这样才能最小化摄像头的个数
		return COVER
	}
	left := setCamera(root.Left, res)
	right := setCamera(root.Right, res)
	// 后序遍历
	if left == UNCOVER || right == UNCOVER {
		*res++
		return CAMERA
	}
	if left == COVER && right == COVER {
		return UNCOVER
	}
	if left == CAMERA || right == CAMERA {
		return COVER
	}
	return -1
}

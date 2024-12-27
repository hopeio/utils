package model

type Tree struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type uint32 `json:"type"`
}

type TreeNode struct {
	ID     uint   `json:"id"`
	PID    uint   `json:"ptId"`
	Value  string `json:"value"`
	TreeID uint   `json:"treeId"`
}

package response

type List[T any] struct {
	List  []T `json:"list"`
	Total int `json:"total"`
}

package hpg

// Error はエラーデータを表す。
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

package domain

type User struct {
	ID       uint   `json:"id" binding:"-" ksql:"_id"`
	Username string `json:"username" binding:"required" ksql:"username"`
	Password string `json:"password" binding:"required" ksql:"password"`
}

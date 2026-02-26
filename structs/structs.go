package structs

// Full user stored in system
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Payload for creating user
type CreateUserPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Payload for updating user
type UpdateUserPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
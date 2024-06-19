package models

type User struct {
    ID           uint   `json:"id"`
    Username     string `json:"username"`
    Password     string `json:"password"`
    PasswordHash string `json:"-"`
    Email        string `json:"email"`
}

type Login struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

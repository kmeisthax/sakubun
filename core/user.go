package core

import (
    "github.com/twinj/uuid"
    "code.google.com/p/go.crypto/bcrypt"
)

type User interface {
    AuthenticateUser(password []byte) bool
}

type InnerUser struct {
    Id uuid.UUID
    
    Username string
    PassHash []byte
}

func (u InnerUser) AuthenticateUser(password []byte) bool {
    if (bcrypt.CompareHashAndPassword(u.PassHash, password) == nil) {
        return true
    } else {
        return false
    }
}
package generator

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

func Generate(length int)(string, error){
	password:=make([]byte, length)

	for i := 0; i < length; i++ {
		n,err:=rand.Int(rand.Reader,big.NewInt(int64(len(charset))))
		if err!=nil{
			return "",err
		}
		password[i]=charset[n.Int64()]
	}

	return string(password),nil
}
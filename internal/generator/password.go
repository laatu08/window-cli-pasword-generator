package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

const (
	uppercase  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase  = "abcdefghijklmnopqrstuvwxyz"
	digitsSet  = "0123456789"
	symbolsSet = "!@#$%^&*()-_=+[]{}<>?/|"
)

func randomChar(set string) (byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
	if err != nil {
		return 0, err
	}
	return set[n.Int64()], nil
}

func Generate(length int, upper, lower, digits, symbols bool) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	password := make([]byte, 0, length)
	charset := ""

	if upper {
		c, _ := randomChar(uppercase)
		password = append(password, c)
		charset += uppercase
	}
	if lower {
		c, _ := randomChar(lowercase)
		password = append(password, c)
		charset += lowercase
	}
	if digits {
		c, _ := randomChar(digitsSet)
		password = append(password, c)
		charset += digitsSet
	}
	if symbols {
		c, _ := randomChar(symbolsSet)
		password = append(password, c)
		charset += symbolsSet
	}

	// if charset == "" {
	// 	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	// 	// return "", errors.New("select at least one character type")
	// }

	if charset == "" {
		return "", errors.New("select at least one character type")
	}

	if length < len(password) {
		return "", errors.New("length too short for selected rules")
	}

	for len(password) < length {
		c, err := randomChar(charset)
		if err != nil {
			return "", err
		}
		password = append(password, c)
	}

	// Shuffle password
	for i := range password {
		jRand, _ := rand.Int(rand.Reader, big.NewInt(int64(len(password))))
		j := int(jRand.Int64())
		password[i], password[j] = password[j], password[i]
	}

	return string(password), nil
}

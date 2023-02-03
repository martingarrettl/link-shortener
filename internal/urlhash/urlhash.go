package urlhash

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateShortCode(url string) string {
	// This is not a secure method of salting and hashing do not replicate
	// it is okay in this application because the data only needs
	// to be unique enough so as to be unlikely to cause collisions

	// generate random number in range to use as a salt
	rng := rand.Intn(2048)
	pseudoSalt := strconv.Itoa(rng)

	// hash salted url
	salted := pseudoSalt + url
	hasher := md5.New()
	hasher.Write([]byte(salted))
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	// return first 8 characters of generated hash
	return hash[:8]
}

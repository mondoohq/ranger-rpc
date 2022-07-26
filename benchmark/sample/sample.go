package sample

import (
	"encoding/base64"
	"math/rand"
	"time"
)

var Id int64
var Message string
var Name []string

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Id = rand.Int63() * 5

	token := make([]byte, 2000000)
	rand.Read(token)

	Message = base64.StdEncoding.EncodeToString(token)
	Name = []string{"abc", "dev", "blub"}
}

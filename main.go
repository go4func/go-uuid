package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	a := 110.01

	pad := fmt.Sprintf("%012d", int64(a*100))
	fmt.Println(pad)
}

func uuID() {
	s := uuid.New().String()
	fmt.Println(s, len([]byte(s)), strings.Replace(s, "-", "", -1))
}

// 20181113013313033525
func GenerateRefID() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return fmt.Sprintf("%s%s", time.Now().Format("20060102030405"), fmt.Sprintf("%06d", r.Intn(999999)))
}

package joeusername

import (
	"math/rand"
	"strings"
	"time"
)

// Generate int
func Generate(bound int) string {
	var str strings.Builder
	words := [...]string{"Cool", "Interesting", "Right"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if bound == 0 {
		bound = r.Intn(12)
	}
	for i := 0; i < bound; i++ {
		str.WriteString(words[r.Intn(len(words))])
	}
	return str.String()
}

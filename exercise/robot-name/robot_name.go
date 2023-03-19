package robotname

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const (
	nums           = "1234567890"
	letters        = "ABCDEFGHIJKLMNOPQRSTUWXYZ"
	maxRegenerates = 26 * 26 * 10 * 10 * 10 * 10 * 10
)

type (
	Robot struct {
		name string
	}
)

var namesSet = map[string]struct{}{}

func Init() {
	rand.Seed(time.Hour.Milliseconds())
}

func generateRandomStrFrom(str string) string {
	i := rand.Intn(len(str))
	return string(str[i])
}

func generateName() string {
	var str strings.Builder
	for i := 0; i < 5; i++ {
		if i < 2 {
			str.WriteString(generateRandomStrFrom(letters))
		} else {
			str.WriteString(generateRandomStrFrom(nums))
		}
	}
	return str.String()
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	for i := 0; i < maxRegenerates; i++ {
		name := generateName()
		if _, ok := namesSet[name]; ok {
			continue
		}
		namesSet[name] = struct{}{}
		r.name = name
		return r.name, nil
	}
	return "", errors.New("max regenerates has been achieved")
}

func (r *Robot) Reset() {
	r.name = ""
}

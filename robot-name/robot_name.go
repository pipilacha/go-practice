package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const SOURCE = "ABCDEFGHIJKLMNOPQRXTUVWXYZ"
const MAX_NAMES = 26 * 26 * 1000

var robotNames = make(map[string]string, MAX_NAMES)

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := createName()
		if err != nil {
			return "", err
		}
		if _, ok := robotNames[name]; ok {
			name, _ = r.Name()
		}
		r.name = name
		robotNames[name] = name
	}
	return r.name, nil
}

func createName() (string, error) {
	if len(robotNames) == MAX_NAMES {
		return "", errors.New("maximum number of names reached")
	}
	rand.NewSource(time.Now().UnixNano())
	no := rand.Intn(1000)
	robotName := strings.ToUpper(string(SOURCE[rand.Intn(25)]) + string(SOURCE[rand.Intn(25)]))
	newName := fmt.Sprintf("%s%03d", robotName, no)
	return newName, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

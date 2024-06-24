package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var pool_names []string
var generatedNames = false

func generateCombinationPool2() {
	pool_letters := make([]string, 26*26)
	index := 0
	for i := 'A'; i <= 'Z'; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			pool_letters[index] = string(i) + string(j)
			index++
		}
	}

	pool_numbers := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		pool_numbers[i] = fmt.Sprintf("%03d", i)
	}

	pool_names = make([]string, len(pool_letters)*len(pool_numbers))
	index = 0
	for _, letters := range pool_letters {
		for _, numbers := range pool_numbers {
			pool_names[index] = letters + numbers
			index++
		}
	}

	generatedNames = true
}

func (r *Robot) Name2() (string, error) {
	if !generatedNames {
		generateCombinationPool2()
	}

	if r.name != "" {
		return r.name, nil
	}

	if len(pool_names) == 0 {
		return "", errors.New("max limit of names reached")
	}

	name_number := 0
	if len(pool_names) > 0 {
		name_number = rand.Intn(len(pool_names))
	}

	r.name = pool_names[name_number]
	pool_names = append(pool_names[:name_number], pool_names[name_number+1:]...)
	return r.name, nil
}

func (r *Robot) Reset2() {
	r.name = ""
}

/* ATTEMPT 1
PASS
ok      robotname       138.505s
*/

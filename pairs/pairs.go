package pairs

import (
	"fmt"
)

func PairsString(kv ...string) map[string]string {
	checkOdd(len(kv))

	v := map[string]string{}
	var key string
	for i, s := range kv {
		if i%2 == 0 {
			key = s
			continue
		}

		v[key] = s
	}
	return v
}

func Pairs(kv ...interface{}) map[string]interface{} {
	checkOdd(len(kv))

	v := map[string]interface{}{}
	var key string
	for i, s := range kv {
		if i%2 == 0 {
			key = fmt.Sprint(s)
			continue
		}

		v[key] = s
	}
	return v
}

func checkOdd(length int) {
	if length%2 == 1 {
		panic(fmt.Sprintf("odd number of input values: %d", length))
	}
}

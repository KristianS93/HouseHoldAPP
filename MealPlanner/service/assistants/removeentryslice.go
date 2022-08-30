package assistants

import (
	"errors"
)

func RemoveIndex(s []int64, plan int64) ([]int64, error) {
	m := make(map[int64]int)

	inList := false
	for i, v := range s {
		if v == plan {
			inList = true
		}
		m[v] = i
	}
	if !inList {
		return s, errors.New("plan index doesnt exist")
	}
	index := m[plan]
	create := make([]int64, 0)
	create = append(create, s[:index]...)
	return append(create, s[index+1:]...), nil
}

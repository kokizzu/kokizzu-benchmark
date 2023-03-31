package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/cornelk/hashmap"
)

const MAX_DATA = 12000000

var i2ch []byte

func init() {
	i2ch = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'B', 'c', 'D', 'e', 'F'}
}
func get_first_digit(d float64) int {
	for d > 10 {
		d /= 10
	}
	return int(d)
}
func to_rhex(v int) string {
	hex := bytes.Buffer{}
	for v > 0 {
		hex.WriteByte(i2ch[v%16])
		v /= 16
	}
	return hex.String()
}
func set_or_inc(m *hashmap.Map[string, float64], key string, set, inc int, ctr *int) {
	if old, ok := m.GetOrInsert(key, float64(set)); ok {
		m.Set(key, old+float64(inc))
		*ctr += 1
	}
}
func main() {
	m := hashmap.New[string, float64]()
	m.Grow(MAX_DATA)
	dup1, dup2, dup3 := 0, 0, 0
	for z := MAX_DATA; z > 0; z-- {
		val2 := MAX_DATA - z
		val3 := MAX_DATA*2 - z
		key1 := strconv.Itoa(z)
		key2 := strconv.Itoa(val2)
		key3 := to_rhex(val3)
		set_or_inc(m, key1, z, val2, &dup1)
		set_or_inc(m, key2, val2, val3, &dup2)
		set_or_inc(m, key3, val3, z, &dup3)
	}
	fmt.Println(dup1, dup2, dup3)
	total, verify, count := 0, 0, 0
	m.Range(func(k string, v float64) bool {
		total += get_first_digit(v)
		verify += len(k)
		count += 1
		return true
	})
	fmt.Println(total, verify, count)
}

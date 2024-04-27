package main

import "fmt"

type Val struct {
	x int
	y string
	z bool
}

func ChangeMapValue(dict map[string]map[string]Val, key, value string, x Val) {
	dict[key][value] = x
}

func AddValue(dict map[string]map[string]Val, key, value string, x Val) {
	if _, ok := dict[key]; !ok {
		dict[key] = make(map[string]Val)
	}
	dict[key][value] = x
}

func ListKeysInsideUpperKeys(dict map[string]map[string]Val, keys ...string) []map[string]Val {
	result := make([]map[string]Val, 0)
	for _, key := range keys {
		for k, v := range dict {
			if k == key {
				result = append(result, v)
			}
		}
	}
	return result
}

func main() {
	dict := make(map[string]map[string]Val)
	xStruct := Val{5, "Hello", true}
	AddValue(dict, "OuterKey1", "InnerKey1", xStruct)
	AddValue(dict, "OuterKey1", "InnerKey2", xStruct)
	AddValue(dict, "OuterKey1", "InnerKey3", xStruct)
	AddValue(dict, "OuterKey2", "InnerKey1", xStruct)
	AddValue(dict, "OuterKey2", "InnerKey2", xStruct)
	AddValue(dict, "OuterKey2", "InnerKey3", xStruct)
	fmt.Println(ListKeysInsideUpperKeys(dict, "OuterKey1", "OuterKey2"))
}

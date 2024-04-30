package main

import (
	"fmt"
	"sort"
)

type Str struct {
}

type Sorter struct {
	Key   string
	Value int
}

func CheckIfContains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func addRegionWithOkrugs(dict map[string]map[int]Str, region string, okrugs ...int) {
	fmt.Println("addRegionWithOkrugs function starts! \n")
	if _, ok := dict[region]; !ok {
		dict[region] = make(map[int]Str)
	}
	for _, okrug := range okrugs {
		dict[region][okrug] = Str{}
	}
	fmt.Println(dict, "\n")
	fmt.Println("addRegionWithOkrugs function ends! \n")
}

func showRegionsWithOkrug(dict map[string]map[int]Str, okrug int) []string {
	fmt.Println("showRegionsWithOkrugs function starts! \n")
	result := make([]string, 0)
	for OuterK, OuterV := range dict {
		for innnerk, _ := range OuterV {
			if innnerk == okrug {
				result = append(result, OuterK)
			}
		}
	}
	sort.Strings(result)
	fmt.Println(result, "\n")
	fmt.Println("showRegionsWithOkrugs function ends!  \n ")
	return result
}

func okrugsInRegions(dict map[string]map[int]Str, x []string) {
	fmt.Println("okrugsInRegions function starts! \n")
	result := make([]int, 0)
	for _, el := range x {
		innerDict, ok := dict[el]
		if !ok {
			continue
		}
		for k := range innerDict {
			switch CheckIfContains(result, k) {
			case true:
				continue

			case false:
				result = append(result, k)
			}

		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	fmt.Println(result, "\n")
	fmt.Println("okrugsInRegions function ends! \n")

}

func printRegions(dict map[string]map[int]Str) {
	var result []Sorter
	for k, v := range dict {
		result = append(result, Sorter{k, len(v)})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})
	fmt.Println("printRegions function starts!")
	for _, v := range result {
		fmt.Printf("%s\t%d\n", v.Key, v.Value)
	}
	fmt.Println("printRegions function ends!")

}

func main() {
	dict := make(map[string]map[int]Str)
	addRegionWithOkrugs(dict, "Piter", 1, 2, 3, 4)
	addRegionWithOkrugs(dict, "Piter", 4, 5, 6)
	addRegionWithOkrugs(dict, "Moscow", 4, 43, 23, 12)

	okrugsInRegions(dict, []string{"Piter", "Moscow"})

	showRegionsWithOkrug(dict, 4)

	printRegions(dict)

}

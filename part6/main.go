package main

import (
	"fmt"
	"sort"
)

func main() {
	//x := [5]float64{98, 93, 77, 82, 83}
	//
	//var total float64 = 0
	//for _, value := range x {
	//	total += value
	//}
	//fmt.Println(total / float64(len(x)))

	//slice1 := []int{1, 2, 3}
	//slice2 := append(slice1, 4, 5)
	//fmt.Println(slice1, slice2)

	//slice1 := []int{1, 2, 3}
	//slice2 := make([]int, 2)
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)

	//x := make(map[string]int)
	//x["key"] = 10
	//fmt.Println(x["key"])

	//elements := map[string]map[string]string{
	//	"H": map[string]string{
	//		"name":  "Hydrogen",
	//		"state": "gas",
	//	},
	//	"He": map[string]string{
	//		"name":  "Helium",
	//		"state": "gas",
	//	},
	//	"Li": map[string]string{
	//		"name":  "Lithium",
	//		"state": "solid",
	//	},
	//	"Be": map[string]string{
	//		"name":  "Beryllium",
	//		"state": "solid",
	//	},
	//	"B": map[string]string{
	//		"name":  "Boron",
	//		"state": "solid",
	//	},
	//	"C": map[string]string{
	//		"name":  "Carbon",
	//		"state": "solid",
	//	},
	//	"N": map[string]string{
	//		"name":  "Nitrogen",
	//		"state": "gas",
	//	},
	//	"O": map[string]string{
	//		"name":  "Oxygen",
	//		"state": "gas",
	//	},
	//	"F": map[string]string{
	//		"name":  "Fluorine",
	//		"state": "gas",
	//	},
	//	"Ne": map[string]string{
	//		"name":  "Neon",
	//		"state": "gas",
	//	},
	//}
	//
	//if el, ok := elements["Li"]; ok {
	//	fmt.Println(el["name"], el["state"])
	//}

	//Task1

	array1 := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(make([]int, 3, 9)) // output: [0 0 0]
	fmt.Println(array1[2:5])       // output: [c d e]

	//Task2: sort

	array2 := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	sort.Ints(array2)
	fmt.Println(array2[0]) // output: 9
}

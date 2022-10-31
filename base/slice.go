package base

import "strings"

// CheckElementExistInArray func
func CheckElementExistInArray(array []string, element string) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}

	return false
}

// CheckElementExistInArrayInt func;
func CheckElementExistInArrayInt(array []int, element int) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}

	return false
}

// CheckElementExistInArrayInt32 func;
func CheckElementExistInArrayInt32(array []int32, element int32) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}

	return false
}

// CheckElementExistInArrayInt64 func;
func CheckElementExistInArrayInt64(array []int64, element int64) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}

	return false
}

// UniqueSliceInt func;
func UniqueSliceInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// UniqueSliceInt32 func;
func UniqueSliceInt32(intSlice []int32) []int32 {
	keys := make(map[int32]bool)
	list := []int32{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// UniqueSliceInt64 func;
func UniqueSliceInt64(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// UniqueSliceString func;
func UniqueSliceString(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// DiffSliceInt32 func;
// Trả ra một mảng chứa các phần tử nằm trong mảng a nhưng không xuất hiện trong mảng b
func DiffSliceInt32(a, b []int32) []int32 {
	// Turn b into a map
	m := make(map[int32]bool)
	for _, s := range b {
		m[s] = false
	}
	// Append values from the longest slice that don't exist in the map
	var diff []int32
	for _, s := range a {
		if _, ok := m[s]; !ok {
			diff = append(diff, s)
			continue
		}
		m[s] = true
	}
	return diff
}

// DiffSliceInt32 func;
// Trả ra một mảng chứa các phần tử nằm trong mảng a nhưng không xuất hiện trong mảng b
func DiffSliceInt64(a, b []int64) []int64 {
	// Turn b into a map
	m := make(map[int64]bool)
	for _, s := range b {
		m[s] = false
	}
	// Append values from the longest slice that don't exist in the map
	var diff []int64
	for _, s := range a {
		if _, ok := m[s]; !ok {
			diff = append(diff, s)
			continue
		}
		m[s] = true
	}
	return diff
}

// MergeUniqueSliceInt32 func;
// Return a slice unique from listSource without listExclude
func MergeUniqueSliceInt32(listSource, listExclude []int32) []int32 {
	myMap := make(map[int32]bool)
	for _, id := range listSource {
		myMap[id] = true
	}

	for _, id := range listExclude {
		myMap[id] = false
	}

	listRes := make([]int32, 0)
	for id, isExcept := range myMap {
		if isExcept {
			listRes = append(listRes, id)
		}
	}

	return listRes
}

// MergeUniqueSliceInt64 func;
// Return a slice unique from listSource without listExclude
func MergeUniqueSliceInt64(listSource, listExclude []int64) []int64 {
	myMap := make(map[int64]bool)
	for _, id := range listSource {
		myMap[id] = true
	}

	for _, id := range listExclude {
		myMap[id] = false
	}

	listRes := make([]int64, 0)
	for id, isExcept := range myMap {
		if isExcept {
			listRes = append(listRes, id)
		}
	}

	return listRes
}

// CheckElementContainInArray func;
func CheckElementContainInArray(array []string, element string) bool {
	for _, e := range array {
		if strings.Contains(element, e) {
			return true
		}
	}

	return false
}

// CommonUniqueSliceInt32 func;
// Return a slice unique from listSource also belong to listExclude
func CommonUniqueSliceInt32(listSource, listExclude []int32) []int32 {
	myMap := make(map[int32]bool)
	for _, id := range listSource {
		myMap[id] = true
	}

	listRes := make([]int32, 0)
	for _, id := range listExclude {
		if _, ok := myMap[id]; ok {
			listRes = append(listRes, id)
		}
	}

	return listRes
}

// SliceInt32ToInt64 func;
func SliceInt32ToInt64(listReq []int32) []int64 {
	listRes := make([]int64, 0)
	for _, id := range listReq {
		listRes = append(listRes, int64(id))
	}
	return listRes
}

// SliceInt32ToInt64 func;
func SliceInt64ToInt32(listReq []int64) []int32 {
	listRes := make([]int32, 0)
	for _, id := range listReq {
		listRes = append(listRes, int32(id))
	}
	return listRes
}

// RemoveElementInt32 func;
func RemoveElementInt32(slice []int32, index int32) []int32 {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveElementInt64 func;
func RemoveElementInt64(slice []int64, index int) []int64 {
	return append(slice[:index], slice[index+1:]...)
}

// SliceStringRemoveElement func;
func SliceStringRemoveElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// SliceInt32RemoveElement func;
func SliceInt32RemoveElement(slice []int32, index int32) []int32 {
	return append(slice[:index], slice[index+1:]...)
}

// SliceInt64RemoveElement func;
func SliceInt64RemoveElement(slice []int64, index int) []int64 {
	return append(slice[:index], slice[index+1:]...)
}

// SliceStringRemoveElementByValue func;
func SliceStringRemoveElementByValue(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// SliceInt32RemoveElementByValue func;
func SliceInt32RemoveElementByValue(slice []int32, value int32) []int32 {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// SliceInt64RemoveElementByValue func;
func SliceInt64RemoveElementByValue(slice []int64, value int64) []int64 {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

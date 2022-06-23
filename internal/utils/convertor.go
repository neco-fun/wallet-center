package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func CovertUIntArrayToUInt64Array(arr []uint) []uint64 {
	var result []uint64
	for _, item := range arr {
		result = append(result, uint64(item))
	}
	return result
}

func ConvertUintArrayToString(arr []uint, symbol string) string {
	var temp = make([]string, len(arr))
	for k, v := range arr {
		temp[k] = fmt.Sprintf("%d", v)
	}
	return strings.Join(temp, symbol)
}

func ConvertStringToUIntArray(str string) []uint {
	var result []uint
	strArray := strings.Split(str, ",")
	for _, item := range strArray {
		intItem, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			continue
		}
		result = append(result, uint(intItem))
	}
	return result
}

func ConvertIntArrayToBigIntArray(arr []int32) []*big.Int {
	var result []*big.Int
	for _, item := range arr {
		result = append(result, big.NewInt(int64(item)))
	}
	return result
}

func ConvertUInt64ArrayToUIntArray(arr []uint64) []uint {
	var result []uint
	for _, item := range arr {
		result = append(result, uint(item))
	}
	return result
}

func GetIndexFromUIntArray(array []uint, value uint) int {
	for index, num := range array {
		if num == value {
			return index
		}
	}
	return -1
}

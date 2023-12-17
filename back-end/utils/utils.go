package utils

import "fmt"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsInArray(search string, array []string) bool {
	fmt.Print("🌻", search)
	for _, value := range array {
		fmt.Print("📖", value)
		if value == search {
			return true
		}
	}
	return false
}

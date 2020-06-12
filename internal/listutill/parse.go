package listutill

import "strings"

//ParseList parses a string like this "a, b, c" or "a,b,c" to a
//array for further processing
func ParseList(strList string) []string {
	//Trim newlines and split it through "," into a slice
	retList := strings.Split(strings.TrimSpace(strList), ",")

	//Loop through array and trim spaces
	for id, value := range retList {
		retList[id] = strings.Trim(value, " ")
	}

	return retList
}

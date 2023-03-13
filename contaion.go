package main

//import "fmt"

//func aa() {
//	a := []string{"apple", "orange", "lemon", "apple", "vine"}
//
//	str, a, err := deleteStations(a, "apple")
//	fmt.Println(str) // => "apple"
//	fmt.Println(a)   // => "[orange lemon vine]"
//	fmt.Println(err) // => "<nil>"
//
//	str, a, err = deleteStations(a, "apple")
//	fmt.Println(str) // => ""
//	fmt.Println(a)   // => "[orange lemon vine]"
//	fmt.Println(err) // => "Couldn't find"
//}
//
//func deleteStations(slice []stationInfo, s string) (string, []stationInfo, error) {
//	ret := make([]stationInfo, len(slice))
//	i := 0
//	for _, x := range slice {
//		if s != x.station {
//			ret[i] = x
//			i++
//		}
//	}
//	if len(ret[:i]) == len(slice) {
//		return "", slice, fmt.Errorf("couldn't find")
//	}
//	return s, ret[:i], nil
//}

package basics

import (
	"fmt"
	"maps"
)

func main() {
	//Maps are a built in data structure in Go that store key-value pairs.Its provides an efficient way to store  and retrieve key-value pairs.
	//Maps are unordered collections of key-value pairs, where each key is unique and maps to a specific value.

	//var mapVariable map[keyType]valueType
	//myVariable = make(map[keyType]valueType)
	//Using map literal
	// myMap = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	myMap := make(map[string]int)
	fmt.Println("Initial map:", myMap)
	myMap["key1"] = 1
	myMap["key2"] = 2
	fmt.Println("Map after adding elements:", myMap)
	//Non-existent key default value type
	fmt.Println(myMap["key7"])
	//Deleteing a key-value pair
	delete(myMap, "key1")
	fmt.Println("Map after deleting key1:", myMap)
	//Completly remove all the key-value pairs from the map
	clear(myMap)
	fmt.Println("Map after clearing all elements:", myMap)
	myMap["key1"] = 1
	myMap["key2"] = 2
	myMap["key3"] = 3
	myMap["key4"] = 4
	//Checking existence of a key
	value, unknown := myMap["key2"]
	fmt.Println("Value for key2:", value)
	fmt.Println("Value for unknown key:", unknown) // false since key doesn't exist
	//Check if maps are equal
	mapA := map[string]int{"a": 1, "b": 2}
	mapB := map[string]int{"a": 1, "b": 2}
	if maps.Equal(mapA, mapB) {
		fmt.Println("mapA and mapB are equal")
	}
	//Iterating over a map given range
	for k, v := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
}

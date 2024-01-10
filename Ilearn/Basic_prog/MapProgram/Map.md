1) What is Map in go ?
-> 
In Go, a map is a built-in data structure that provides an unordered collection of key-value pairs. It is a versatile data structure that allows you to store and retrieve values based on unique keys. The key and value can be of any comparable type.

key must be unique and always in the type which is comparable using == operator or the type which support != operator.So,we can use key like an int, float64, rune, string, comparable array and structure, pointer, etc. The data types like slice and noncomparable arrays and structs or the custom data types which are not comparable don’t use as a map key.

Synatx :

myMap := make(map[keyType]valueType)

// Shorthand syntax
myMap := map[keyType]valueType{}

2) what is Empty map and nil map ?
->
In maps, the zero value of the map is nil and a nil map doesn’t contain any key. 
If you try to add a key-value pair in the nil map, then the compiler will throw runtime error. 
Example: 
-------
var mymap map[int]string

mymap is a nil map

refer : MapProgram/1_create.go

3) how can we check existance of key in map 
    if _, ok := myMap[keytosearch]; ok {
		fmt.Printf("Key '%s' exists in the map\n", keyToCheck)
	} else {
		fmt.Printf("Key '%s' does not exist in the map\n", keyToCheck)
	}

4) how to merge two map ?
for k, v := range b {
    a[k] = v
}


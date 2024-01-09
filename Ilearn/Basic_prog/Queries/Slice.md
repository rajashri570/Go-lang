1) difference between slice cap and length
->Length: The length is the total number of elements present in the array.
Capacity: The capacity represents the maximum size upto which it can expand.

2) why slice called dynamic?
-> Slices are dynamic, which means that their size can change as you add or remove elements. Go provides several built-in functions that allow you to modify slices, such as append, copy, and delete.

3) What is zero value slice?
 ->
 In Go language, you are allowed to create a nil slice that does not contain any element in it.
 So the capacity and the length of this slice is 0. Nil slice does not contain an array reference 
 ex: var myslice []string

 4) What is the Slice literals ?
 ->
 When you know the values that goes into your slice before declaring, you can directly pass them in using slice literal. With slice literal, you do not need to use the make function, your slice gets declared and assigned its values in one step 
 Example :
 stringSlice := []string{"hello","world"}

5)whta is reflect.DeepEqual() function in go?
-> 
This function tests for deep equality between two values, which means that it recursively compares the values of structs, arrays, slices, maps, and pointers.

it is important to note that reflect.DeepEqual() has some limitations. For example, it does not correctly compare functions or channels, and it can have unexpected behavior when comparing floats or interfaces. Therefore, it is generally recommended to use the loop method
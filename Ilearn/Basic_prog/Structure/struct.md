1) what is structure in go and how it differes from class
->
structure is a user defined datatype that allows us store different types of data in single
Any real-world entity which has some set of properties/fields can be represented as a struct.
This concept is generally compared with the classes in object-oriented programming. it can be termed as a lightweight class that does not support inheritance but supports composition. 

Declaring a structure:

 type Address struct {
      name string 
      street string
      city string
      state string
      Pincode int
}
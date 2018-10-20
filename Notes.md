What is Struct?
Data Structure. Colletion of properties that are related together
Kind of like Objects in Javascript

Go is pass by value language
So anytime we pass a value to a function either
as an argument that data is copied in memory.
So function will always be working on a copy of our data structure 

Notes on Pointers:

Turn address into value with *address
Turn value into address with &value

Difference between Array and Slice in Go:

Arrays: Primitive data structure
Can't be resized
Rarely used directly

Slices:
Can grow and shrink
Used 99% of the time for lists of elements


Whenever a SLice is created, Go creates two separate data structures
An array and a structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array
When a slice is made, go makes array internally

When the slice is created, we get both an array and a data structure describing that array. The data structure describing the array *is* copied before being passed off to the function, but the underlying array is not.

Slice is a reference type because a slice contains a reference to the actualunderlying list of records

Value Type : Use pointers to change these things in a function
int, float, string, bool, structs

Reference Types:Need not worry about pointers with these
slices,maps,channels, pointers, functions 


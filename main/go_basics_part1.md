
# Go Basics - Part 1

This is my learning notes for Udemy Golang intro course.

## Hello World from Go

Every project and every file has a package name.

We have two kinds of packages, one that produces executable and one that is for reusable code.

The execuatable package should be named 'main', otherwise it will not produce an executable via 'go bulid'.


variables should be defined for the first time with type.

variables can be defined outside the scope of functions but cannot be initilized

```
var myvar string
``` 

function syntax:

```
func newCard() string {
	return "Five of Diamonds"
}

```

you can call a function in another file within the same package. but when running (and probably building) should call them all.

```
go run main.go helper.go
```


## Card Deck Example



The equivalnet of list (in python) are "array" and "slice" in go. array has fixed length.

We define a slice like this:

```
	dec := []string{"Ace of Heart" , "Two of Heart"}
```

The indexing is pretty much like python.


you don't have 'self' or 'this' to refer to objects of a class in go.

A class is defined just by the keyword 'type'.

```
type deck []string
```

Note that Every variable always have to be declared first!

By referencing to pieaces of a slice,  We created  new references that point at subsections of the slice. We never directly modified the slice that  is pointing at.



## Card Deck Example Continued

it seems that the way we separate arguments in new lines for prettyfying! is not Okay with go.

So the improt statement 'from io import ioutils' becomes 'import io/ioutils'

I don't like using a short name to represent a "receiver"  ( or "self" in python). Also don't like the way they name the variables and errors.

I found that Go linter does not like "_" in the variable names.

Also class names ( for example "deck" here) start with lowercase.

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


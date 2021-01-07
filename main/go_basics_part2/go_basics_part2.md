# Go Basics - Part 1



## Pointers

okay so the first example is good to notify how one used to python may assume that you can just modify a class instance.
By the way, even in python these things can get confusing, i.e. how the concept of pointers work.

**ToDO**  It would be interesting to refresh my memory on this pointer (mutable/immutable) variables in python sometime later.

**NOte:** One general rule in Go that will be usefull a lot, is that "Go always pass by value".

Now the way that I want to make myself remember is that "&" operator tells Go to give the address of a variable.
The '*' operator tells to give the value of a pointer.
for types, however, the '*' tells go to create a pointer type for that type.

so the second example is the "elaborate way" that helps understand what is going on.

The short way which is equivalnet to the example and common is:

```
myGuy.updateName()

func (p *person) updateName{
    *p.firstName = "joe"
}
```

so this basically means that Go allows to operate on pointer type instead of the original type.

**Note:** The slice (and some other types) which are defined by reference do not need pointing when being modified.

so that means 
```
func main(){
a := []int{ 1,2,3}

updateSlice(a)
fmt.Prinln(a)

}


func updateSlice(s []int){
    s[0] = 0
}
```
will result to '{0,2,3}'

This rule also includes 'maps', 'channels', 'pointer' and 'functions' which are known as reference types.

## Maps


so there are a few ways around creating and editing maps.
what I was looking for is that it looks like the only way to create a map all at onnce is by this syntax:

```
	colors := map[string]string{
		"red":   "ff00000",
		"green": "4bf745",
	}
```

The 'colors:= make(map[string]string)' and 'var colors map [string]string' will create empty maps which then can be filled item by item.




## Interfaces


interface syntax is:

```

type <type_name> interface {
	<func_name>(<function_argument_types>) <function_output_types>

}

```

okay so far I have not found much signs of object oriented concepts in go, like inheritance, etc.
Also as I understood, go is not much of a object oriented language.
So what came to me is how objec oriented concepts (specially by matching with python) are implemented.
This interface is a good example. Seems to me that this is somehow related to inheritance.
we link a bunch of types to a  new (probably more generic) type th
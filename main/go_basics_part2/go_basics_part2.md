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
we link a bunch of types to a  new (probably more generic) interface type through a function that they share.
So basically Go inferes that if a type(class) has a func(method) then it is considered to be part of an interface type.
But obviously this mapping of python <--> go is very flawed.

**Note** You cannnot create an Interface type variable directly. (compared to a typical concerete type)


So Interface Type **can include** other Interfaces

so in the example the response had an atribute 'Body' of type ReadCloser (which is an Interface)
**Q:** So if the object cannot be created form an interface type, how come body is of a interafce type?
Is it becasue it is indirectly creating an object with the read method and we pass an object and get our data?
Is it that an object of type Interface can be created but only have methods? (well I am reminding myself that all this is python object oriented look to the situation)
Also another point of question was that the function had some outputs but we didnot use them.
we did not do something like '_,_=func()'


```
type ReadCloser interface{
	Reader
	Closer
}

```
where Reader and Closer were Interface types as well.

```
type Reader interface{
	Read([]byte) (n int, err error)
}
```

**Q** also this syntax for the interface from the go guide seems to be different from what we saw before?!


okay so for the third examples of interfaces, I did a couole of mistakes, first I implemented two different implemnetations
for the PrintArea() function, that both had shape interface. 

The second one was that I tried to implement the function PrintAre with '(s shape)' as a receiver of the function which I got the complaint that a interface type cannot be placed as receiver. my goal was to have something like:

```
triangle.printArea()
```
I guess the extension does not apply here.

## Cocurrency vs Parralelism 

we link a bunch of types to a  new (probably more generic) type th

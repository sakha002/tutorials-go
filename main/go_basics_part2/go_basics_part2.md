# Go Basics - Part 1



## Pointers

okay so the first example is good to notify how one used to python may assume that you can just modify a class instance.
By the way, even in python these things can get confusing, i.e. how the concept of pointers work.
So It would be interesting to refresh my memory on this some time.

One general rule in Go that I feel will be guiding a lot, is that "Go always pass by value".

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
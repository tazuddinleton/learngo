## Getters & Setters
Getters and setters are not provided by default in go, but we can use it. 
- If we have a field called `name` we can have a getter like `Name()` not `GetName()`
- a setter like `SetName()`
## Interface names
Following are some of the well-known interfaces
- Do not name your interface `Reader`, `Writer`, `Formatter`, `CloseNotifier` unless you have the same signature and meaning
- One method interfaces are named the method name+er like above
## Mixed caps
Do not use underscore `like_this` to write multi-word name in go. Write `likeThis` or `LikeThis`.

## Data
### Allocation with `new(T)`
- It allocates memory for type T and zeros them out
- `new(T)` returns a `*T` 
- `new(T)` does not initializes type T
### T{prop: value} composite literal
- We can initialize a type with literals
- `Person{name: "Mr. Tom"}` will create a new instance of `Person` each time it is evaluated
- `&Person{}` is same as `new(Person)`
### Allocation with make
`make(T, args)` is different from previous two methods. It returns an initialized (not zeroed) value of type `T` not `*T`. It is used for slices, maps and channels only, because these types are representations of something else under the hood that must be initialized before use. 
- A slice uses a three item descriptor 
  - a pointer to the data contained in an array `*[1,2,3]`
  - the length of the slice
  - the capacity of the slice
  
    A psudo code representation of a slice
    ```    
    array 0x001111: [1,2,3,4,5] // somewhere in the memory

    type slice {
        *array[int]: 0x001111
        length: int
        capacity: int
    }
    ```
    
# Stack Implementation in Go

This repository showcases my Go implementation of a generic stack. I've included various operations such as pushing, popping, cloning, and more to demonstrate the stack's functionality and flexibility.

## Usage

Here is an example of how to use the stack:

```go
package main

import "fmt"

func main() {
    stack := &Stack{}

    stack.fromArray([]interface{}{1, 2, 3, 4, "5"})
    stack.Push(6)
    fmt.Println("Original stack:")
    stack.Print("h")

    clonedStack := stack.clone()
    clonedStack.Reverse()
    fmt.Println("Cloned and reversed stack:")
    clonedStack.Print("h")

    stack.Pop()
    fmt.Println("Stack after popping an element:")
    stack.Print("h")
}
```

## Functions

### fromArray

`func (s *Stack) fromArray(array []interface{}) error`

Initializes the stack from an array. Takes an array as an argument. Returns an error if the array is empty.

### clone

`func (s *Stack) clone() *Stack`

Creates and returns a deep copy of the stack. Takes no arguments.

### Push

`func (s *Stack) Push(item interface{})`

Adds an item to the top of the stack. Takes an item as an argument.

### Pop

`func (s *Stack) Pop() (interface{}, error)`

Removes and returns the top item of the stack. Takes no arguments. Returns an error if the stack is empty.

### DeleteItem

`func (s *Stack) DeleteItem(item interface{}) error`

Removes a specific item from the stack. Takes an item as an argument. Returns an error if the stack is empty or the item is not found.

### IndexOf

`func (s *Stack) IndexOf(item interface{}) (int, error)`

Returns the index of the item in the stack. Takes an item as an argument. Returns an error if the stack is empty or the item is not found.

### Peek

`func (s *Stack) Peek() (interface{}, error)`

Returns the top item without removing it. Takes no arguments. Returns an error if the stack is empty.

### Bottom

`func (s *Stack) Bottom() (interface{}, error)`

Returns the bottom item of the stack. Takes no arguments. Returns an error if the stack is empty.

### Clear

`func (s *Stack) Clear() error`

Removes all items from the stack. Takes no arguments. Returns an error if the stack is empty.

### Reverse

`func (s *Stack) Reverse() error`

Reverses the order of items in the stack. Takes no arguments. Returns an error if the stack is empty.

### Contains

`func (s *Stack) Contains(item interface{}) bool`

Checks if the stack contains a specific item. Takes an item as an argument. Returns false if the stack is empty or the item is not found.

### Size

`func (s *Stack) Size() int`

Returns the number of items in the stack. Takes no arguments.

### isEmpty

`func (s *Stack) isEmpty() bool`

Checks if the stack is empty. Takes no arguments.

### Print

`func (s *Stack) Print(direction string)`

Prints the stack items in either horizontal or vertical direction. Takes a direction string as an argument. Prints an error message if the stack is empty or an invalid direction is provided.

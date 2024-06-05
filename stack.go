package main

import (
	"errors"
	"fmt"
)

type Stack struct{
	items[] interface{}
}


//initialize from array
func (s *Stack) fromArray(array []interface{}) error{
	if len(array) == 0 {
		return errors.New("Cannot initialize a stack from an empty array")
	}

	s.items = array

	return nil

}

func (s *Stack) clone() *Stack{
	newItems := make([]interface{}, len(s.items))
	copy(newItems, s.items)

	return &Stack{
		items: newItems,
	}
}

func (s *Stack) Push(item interface{}){
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error){
	if s.isEmpty(){
		return nil, errors.New("Stack is empty")
	}else{
		poppedItem := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return poppedItem, nil
	}
}

func (s *Stack) DeleteItem(item interface{}) (error){
	if s.isEmpty(){
		return errors.New("Stack is empty")
	}

	index, err := s.IndexOf(item)

	if index == -1 || err != nil{
		return err
	}else{

		s.items = append(s.items[:index], s.items[index+1:]...)
		return nil
	}
}


func (s *Stack) IndexOf(item interface{}) (int, error){
	if s.isEmpty(){
		return -1, errors.New("Stack is empty")
	}else{
		for i, v := range s.items{
			if(v == item){
				return i, nil
			}
		}
	
		return -1, nil
	}

}

func (s *Stack) Peek() (interface{}, error){
	if s.isEmpty(){
		return nil, errors.New("Stack is empty")
	}else{
		return s.items[len(s.items)-1], nil
	}
} 

func (s *Stack) Bottom() (interface{}, error){
	if s.isEmpty(){
		return nil, errors.New("Stack is empty")
	}else{
		return s.items[0], nil
	}
} 

func (s *Stack) Clear() (error){
	if s.isEmpty(){
		return errors.New("Stack is empty")
	}else{
		s.items = []interface{}{}
		return nil
	}
} 

func (s *Stack) Reverse() (error){
	if s.isEmpty(){
		return errors.New("Stack is Empty")
	}
	for i, j := 0, len(s.items)-1; i<j; i, j = i+1, j-1 {
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}

	return nil
	
}

func (s *Stack) Contains(item interface{}) bool{
	if s.isEmpty(){
		return false
	}else{
		for _, v := range s.items{
			if item == v{
				return true
			}
		}
		return false
	}
}
 
func (s *Stack) Size(item interface{}) (int){
	return len(s.items)
}

func (s *Stack) isEmpty() bool{
	return len(s.items) == 0
}

func (s *Stack) Print(direction string){
	if(s.isEmpty()){
		fmt.Println("Stack is empty")
	}else if direction != "v" && direction != "h"{ 
		fmt.Println("Please input a correct direction")
		return
	}else{
		for i := len(s.items)-1; i >= 0; i--{
			if direction != "v"{
				fmt.Print(s.items[i])
				if i != 0{
					fmt.Print("->")
				}else{
					fmt.Print("\n")
				}
			}else{
				fmt.Println(s.items[i])
			}
		}
	}
		

}


func main(){

	stack := &Stack{}

	stack.fromArray([]interface{}{1, 2, 3, 4,"5"})
	// stack.Push(1)
	// stack.Push(2)
	// stack.Push("5")
	// stack.Push(4)

	// index, _ := stack.IndexOf("5")
	// stack.Peek()
	// stack.DeleteItem(4)
	// stack.Clear()

	// stack.Print("h")
	// stack.Reverse()


	// fmt.Println(stack.Contains("1"))

	// clonedStack := stack.clone()
	// clonedStack.Reverse()
	// stack.Print("h")
	// clonedStack.Print("h")
	stack.Print("v")
	
}
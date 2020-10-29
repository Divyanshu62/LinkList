/*

Create a LinkedList
- Push element at the end of LinkedList
- Get elementby Position in the LinkedList
- Remove last item from the LinkedList
- edit element by position in LinkedList

*/

package main

import "fmt"

type Node struct {
    next *Node
    value  int
}

type List struct {
    head *Node
    tail *Node
}

func (ls *List) PushKeyToList(key int) {

  node := &Node{
          value:  key,
      }

    if ls.head == nil {
       ls.head = node

    } else {
       ls.tail.next = node
    }
    ls.tail = node
}

func (ls *List) PrintList() {
    list := ls.head

    for list != nil {
        fmt.Printf("%+v -> ", list.value)
        list = list.next
    }

}

func (ls *List) RemoveLast() {


    list := ls.head
    if list == nil {
      ls.head = nil
      ls.tail = nil
     } else {
      if list.next == nil {
        ls.head = nil
        ls.tail = nil
      } else {
        for list != nil {

             fmt.Printf("%+v -> ", list.value)
             if list.next.next == nil {
               list.next = nil
               ls.tail = list
               break
             }

            list = list.next
         }
      }
    }
}

func (ls *List) GetElementByPos(pos int){

  list := ls.head
  count := 0
  for list != nil {
      count++
      if count == pos {
         fmt.Printf("Key is %+v\n ", list.value)
         break
      }
      list = list.next
  }
}

func (ls *List) EditElementByPos(pos int, val int){

  list := ls.head
  count := 0
  for list != nil {
      count++
      if count == pos {
         list.value = val
         break
      }
      list = list.next
  }
}


func main() {
  list := &List{}

  here:  fmt.Println("\n\n\n\nProject LinkedList\n\nSelect\n\n1 - Push element at the end of LinkedList\n\n2 - Get elementby Position in the LinkedList\n\n3 - Remove last item from the LinkedList\n\n4 - Edit element by position in LinkedList\n\n5 - Display LinkedList\n\n\n")
  fmt.Println("Enter choice : ")

  var selectOption int
  fmt.Scanln(&selectOption)

  switch selectOption {
    case 1:
      fmt.Println("Enter key to insert : ")
      var value int
      fmt.Scanln(&value)

      list.PushKeyToList(value)
      fmt.Println("\n=====================\n")
      list.PrintList()
      fmt.Println("\n=====================\n")
      goto here
    case 2:
      var pos int
      fmt.Println("Enter position to get : ")
      fmt.Scanln(&pos)
      fmt.Println("\n=====================\n")
      list.GetElementByPos(pos)
      fmt.Println("\n=====================\n")
      goto here
    case 3:
      fmt.Println("\n=====================\n")
      list.RemoveLast()
      fmt.Println("\n=====================\n")
      goto here
    case 4:
      var pos int
      var val int
      fmt.Println("Enter position to Update : ")
      fmt.Scanln(&pos)
      fmt.Println("Enter value to Update : ")
      fmt.Scanln(&val)
      fmt.Println("\n=====================\n")
      list.EditElementByPos(pos, val)
      fmt.Println("\n=====================\n")
      goto here
    case 5:
       fmt.Println("\n=====================\n")
       list.PrintList()
       fmt.Println("\n=====================\n")
      goto here
    default:
      goto here
  }
}

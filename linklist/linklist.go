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
    key  int
}

type List struct {
    head *Node
    tail *Node
}

func (L *List) PushKeyToList(key int) {

  node := &Node{
          key:  key,
      }

    if L.head == nil {
       L.head = node

    } else {
       L.tail.next = node
    }
  L.tail = node
}

func (l *List) PrintList() {
    list := l.head

    for list != nil {
        fmt.Printf("%+v -> ", list.key)
        list = list.next
    }

}

func (l *List) RemoveLast() {


    list := l.head
    if list == nil {
      l.head = nil
      l.tail = nil
     } else {
      if list.next == nil {
        l.head = nil
        l.tail = nil
      } else {
        for list != nil {

             fmt.Printf("%+v -> ", list.key)
             if list.next.next == nil {
               list.next = nil
               l.tail = list
               break
             }

            list = list.next
         }
      }
    }
}

func (l *List) GetElementByPos(pos int){

  list := l.head
  count := 0
  for list != nil {
      count++
      if count == pos {
         fmt.Printf("Key is %+v\n ", list.key)
         break
      }
      list = list.next
  }
}

func (l *List) EditElementByPos(pos int){

  list := l.head
  count := 0
  for list != nil {
      count++
      if count == pos {

        list.key = pos
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
      list.PrintList()

      goto here
    case 2:
      var pos int
      fmt.Println("Enter position to get : ")
      fmt.Scanln(&pos)
      list.GetElementByPos(pos)

      goto here
    case 3:
      list.RemoveLast()
      goto here
    case 4:
      var pos int
      fmt.Println("Enter position to Update : ")
      fmt.Scanln(&pos)
      list.EditElementByPos(pos)

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

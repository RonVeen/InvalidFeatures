package main

import (
   "fmt"
)


func main() {
   fmt.Println("It is working")

   var f1 = feature{"Iets", "id", nil, "", ""}
   var f2 = feature{"Iets", "id", nil, "", ""}
   f1.equals(f2)

}



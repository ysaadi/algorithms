package main

import "fmt"

func insertionSort(numArray []int){
  for i, v := range(numArray){
   j:=i-1
   for (j>=0 && v>numArray[j]){
    numArray[j+1] = numArray[j]
    j--
    }
  numArray[j+1] = v
  return numArray
}
}

func main(){
  numArray := [6]int{1,4,2,7,2,9}
  sortedArray := insertionSort(numArray)

  fmt.Println(sortedArray)

}

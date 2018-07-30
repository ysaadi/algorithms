package main

import (
  "fmt"
)

type binaryTree struct{
  Value int
  Parent *binaryTree
  Left *binaryTree
  Right *binaryTree
}


func insert(tree *binaryTree, parent *binaryTree, val int) *binaryTree{
  if tree == nil{
    tree = &binaryTree{val, parent, nil, nil}
  }else if tree.Value == val{
    print("duplicate value, will not insert")
    return nil
  }else if tree.Value < val{
    tree.Right = insert(tree.Right, tree, val)
  }else{tree.Left = insert(tree.Left, tree, val)}
  return tree
}

func search(tree *binaryTree, val int) *binaryTree{
  if tree == nil{
    return nil
  }else if tree.Value == val{
    return tree
  }else if val < tree.Value{
    return search(tree.Left, val)
  }else{
    return search(tree.Right, val)
  }
}

func successor(tree *binaryTree) *binaryTree{
  var successor *binaryTree
  if successor = rightSuccessor(tree); successor == nil{
    successor = leftSuccessor(tree)
  }
  if successor != nil{
    return successor
  }else{
    return tree.Parent
  }
}
//this is meant to be called first.  If there is no right tree, it will try to find a successor within the left tree.
func rightSuccessor(tree *binaryTree) *binaryTree{
  if(tree.Right == nil){
    return nil
  }else{
    successor :=tree.Right
    for successor.Left != nil {
      successor = successor.Left
    }
    return successor
  }
}

func leftSuccessor(tree *binaryTree) *binaryTree{
  if(tree.Left == nil){
    return nil
  }else{
    successor :=tree.Left
    for successor.Right != nil {
      successor = successor.Right
    }
    return successor
  }
}

func deleteVal(tree *binaryTree, val int){
  deleteNode := search(tree, val)
  deleteRight := deleteNode.Right
  deleteLeft := deleteNode.Left
  if deleteNode == nil{
    fmt.Printf("there is no node with value %d", val)

  }else {
    successor := successor(deleteNode)
    if(successor == deleteNode.Parent){
      *deleteNode = binaryTree{}
    }else{
      *deleteNode = *successor
      deleteNode.Right = deleteRight
      deleteNode.Left = deleteLeft
    }
  }
}

func deleteNode(deleteNode *binaryTree){
  deleteRight := deleteNode.Right
  deleteLeft := deleteNode.Left
  if (deleteNode == nil  || *deleteNode == binaryTree{}){
    return

  }else {
    successor := successor(deleteNode)
    if(successor == deleteNode.Parent){
      *deleteNode = binaryTree{}
    }else{
      *deleteNode = *successor
      deleteNode.Right = deleteRight
      deleteNode.Left = deleteLeft
    }
  }
}

package main

import (
	"fmt"
	ht "lab6/hash_table"
)

func main() {

	var table ht.HashTable
	table.NewHashTable()

	table.Add("hedgehog", "desert")
	table.Add("hedgehog", "indian")
	table.Add("hedgehog", "amur")
	table.Add("hedgehog", "somali")

	table.Add("squirrel", "red")
	table.Add("squirrel", "flying")
	table.Add("squirrel", "fox")
	table.Add("squirrel", "american red")

	table.Add("primate", "gibbons")

	table.Add("bear", "panda")

	listBear := table.Get("bear")
	fmt.Printf("\"bear\" hash: %d	index: %d \n", ht.Hash("bear"), ht.Hash("bear")%20)
	fmt.Printf("Hash \"primate\" %d	index: %d \n", ht.Hash("primate"), ht.Hash("bear")%20)
	for i := listBear.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

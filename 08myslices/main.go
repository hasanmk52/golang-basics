package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices!")

	var fruitList = []string{"Apple", "Mango", "Peach"}
	fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Banana", "Kiwi")
	fmt.Println("Fruit List is: ", fruitList)

	// fruitList = append(fruitList[1:]) //slice the list from index 1 to end
	// fmt.Println(fruitList) //[Mango Peach Banana Kiwi]

	// fruitList = append(fruitList[:3]) //slice the list from beginning to index 3
	// fmt.Println(fruitList) //[Apple, Mango Peach]

	fruitList = append(fruitList[1:3]) //slice the list from index 1 to 3 like substring in java
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321)

	fmt.Println("High Scores:", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted Scores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index
	var courses = []string{"react", "angular", "vue", "javascript", "python", "ruby"}
	fmt.Println("Courses:", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removing:", courses)
}
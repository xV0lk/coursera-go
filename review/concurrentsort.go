package main

import "fmt"
import "log"
import "sort"
import "sync"

// sortList : this command sorts the slices, 
func sortList (unsortedSlice []int, wg *sync.WaitGroup) []int {
	sort.Ints(unsortedSlice)
	wg.Done()  
	return unsortedSlice

}
// mergeSlices : merges the slices
func mergeSlices (list1 []int, list2 []int, list3 []int, list4 []int) []int {
	newSlice := []int{}
	newSlice = append(list1, list2...)
	newSlice = append(newSlice, list3...)
	newSlice = append(newSlice, list4...)
	sort.Ints (newSlice) //this line is practically cheating, 
	return newSlice

}

func main(){

	var newsizeofSlice int
	var itemNumber int
	var inputNumber int
	var wg sync.WaitGroup			//gives me a bug in sort function
	firstSlice := make([]int,0,4) //at least 4 in capacity
	

	fmt.Println("Enter a series of integers for sorting")
	fmt.Println("if the total number of integers you input are n mod(4)!=0 you should see extra zeros")
	fmt.Println("How many integers you'll put into the array?: ")
	fmt.Scan(&itemNumber)

	//cycle wich collects the integer
	for i:=0;i<itemNumber;i++{
		fmt.Println("Enter the integer",i+1,"of ",itemNumber)
		_, err := fmt.Scan(&inputNumber)
		if err != nil{
			log.Panic(err)
			fmt.Println("failed input")
		}

		firstSlice = append(firstSlice,inputNumber)
		fmt.Println(firstSlice)
	} //end integer collection

	//partition the slice into 4 parts, I hope they doesn't have to be even
	// UPDATE : there's a problem if the operation itemNumber/4 have a remainder (lost data)
	// so I need to make a new Slice and add zeros to complete the array

	sizeofSlice := itemNumber/4	//theorical number of slices
	remainder := itemNumber%4 	//remainder if there is one
	remaincount := 0			//just a counter, for the increments

	//remainder < 4 if remainder = 0 size of slice is correct
	// example: sizeofslice is 2, remainder is 3, there's 11 elements,
	// there are 3 elements out of the partition
	// so add 1 to size of slice to fit all 11 elements + one zero

	if remainder!= 0{
		for (itemNumber%4) != 0{
			itemNumber = itemNumber + 1 		//this should fix the bug of missing numbers
			remaincount = remaincount +1    	//at the cost of adding zeros
			firstSlice = append(firstSlice,0) 	//adding the zeros 
	}									
		newsizeofSlice = sizeofSlice+1		//should be itemNumber/4, but this should work
	}else
	{
		newsizeofSlice = sizeofSlice 		//if everything is fine
	}

	//now here STARTS the partition
	slice1 := firstSlice[:newsizeofSlice]
	slice2 := firstSlice[newsizeofSlice:2*(newsizeofSlice)]
	slice3 := firstSlice[2*(newsizeofSlice):3*(newsizeofSlice)]
	slice4 := firstSlice[3*(newsizeofSlice):]

	fmt.Println("here's the partitions",slice1,slice2,slice3,slice4)

	//here starts the goroutines!
	wg.Add(4)   //add 4 because 4 goroutines, .Done() is built in the function
	go sortList(slice1,&wg)
	go sortList(slice2,&wg)
	go sortList(slice3,&wg)
	go sortList(slice4,&wg)
	wg.Wait() // program waits here for all goroutines to end

	//Merges the slices
	mixedSlice := mergeSlices(slice1,slice2,slice3,slice4)

	//print the slice
	fmt.Println("the merged slice is: ",mixedSlice)

	
}// end func main
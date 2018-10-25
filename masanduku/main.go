package main

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/* the excel names and sheet name */
const filename string = "samplefile.xlsx"
const sheetname string = "sample"
const newfilename string = "output.xlsx"

// const filename string = "100000RecordsFull.xlsx"
// const sheetname string = "100000 Records"

func main() {
	start := time.Now()

	//read the excel file
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	//get the rows
	rows := xlsx.GetRows(sheetname)

	//remove the headers
	noheaderarray := append(rows[:0], rows[1:]...)

	//sort the fathers name column values
	fathersnamearray := make([]string, 0)

	for _, row := range noheaderarray {
		//get the fathers name column
		fathersnamearray = append(fathersnamearray, row[7])
	}
	fmt.Println("fathers names BEFORE sort", fathersnamearray)
	//sort.Strings(fathersnamearray)
	fathersnamechan := make(chan []string)
	//fathersnamechan <- fathersnamearray
	mergerSortAsync(fathersnamearray, fathersnamechan)

	//sortedStuff := <-fathersnamechan
	fmt.Println("fathers names after sort")

	fmt.Printf("took %v\n", time.Since(start))

}

func mergerSortAsync(l []string, c chan []string) {
	//commented below because am not sure what it is doing exactly
	// if len(l) < 2 {
	// 	c <- l
	// 	return
	// }
	// if len(l) < 500 { //TUNE THIS NUMER AND DONT CREATE EXTRA WORK UNLESS IT'S BIGGER
	// 	c <- mergeSort(l)
	// 	return
	// }

	mid := len(l) / 2
	c1 := make(chan []string, 1)
	c2 := make(chan []string, 1)

	go mergerSortAsync(l[:mid], c1)
	go mergerSortAsync(l[mid:], c2)

	go func() { c <- merge(<-c1, <-c2) }()

}

func mergeSort(l []string) []string {
	if len(l) < 2 {
		return l
	}
	mid := len(l) / 2
	a := mergeSort(l[:mid])
	b := mergeSort(l[mid:])
	return merge(a, b)
}

func merge(left, right []string) []string {
	var i, j int
	result := make([]string, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		result[i+j] = left[i]
		i++
	}
	for j < len(right) {
		result[i+j] = right[j]
		j++
	}
	return result
}

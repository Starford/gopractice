package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/* the excel names and sheet name */
const filename string = "samplefile.xlsx"
const sheetname string = "sample"

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
	//header := rows[0]
	//fmt.Println(header)
	noheaderarray := append(rows[:0], rows[1:]...)

	//sort the fathers name column values
	fathersnamearray := make([]string, 0)

	for _, row := range noheaderarray {
		//get the fathers name column
		fathersnamearray = append(fathersnamearray, row[7])
	}
	sort.Strings(fathersnamearray)
	//fmt.Println(fathersnamearray)

	//create new rows of data

	//fmt.Println(fathersnamearray)
	fmt.Printf("Lenghth of data in fathersnames array: %d, Cap: %d \n", len(fathersnamearray), cap(fathersnamearray))

	//fmt.Println(noheaderarray)
	fmt.Printf("Lenghth of data in all the data array: %d, Cap: %d \n", len(noheaderarray), cap(noheaderarray))

	for index := range noheaderarray {
		fmt.Println(fathersnamearray[index], noheaderarray[index])
	}

	//rows = rows[len(rows)-1]

	//fmt.Println(rows[0][7])

	fmt.Printf("took %v\n", time.Since(start))
}

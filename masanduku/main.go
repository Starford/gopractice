package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/* the excel names and sheet name */
// const filename string = "samplefile.xlsx"
// const sheetname string = "sample"
const newfilename string = "output.xlsx"

const filename string = "100000RecordsFull.xlsx"
const sheetname string = "100000 Records"

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

	//this is the header
	//header := rows[0]
	//remove the headers
	noheaderarray := append(rows[:0], rows[1:]...)

	//sort the fathers name column values
	fathersnamearray := make([]string, 0)

	for _, row := range noheaderarray {
		//get the fathers name column
		fathersnamearray = append(fathersnamearray, row[7])
	}
	sort.Strings(fathersnamearray)

	//fmt.Printf("Lenghth of data in fathersnames array: %d, Cap: %d \n", len(fathersnamearray), cap(fathersnamearray))

	//fmt.Println(noheaderarray)
	//fmt.Printf("Lenghth of data in all the data array: %d, Cap: %d \n", len(noheaderarray), cap(noheaderarray))

	//loop and create the new csv
	//fmt.Println(header)
	// for index := range noheaderarray {
	// 	fmt.Println(fathersnamearray[index])
	// 	fmt.Println(noheaderarray[index])
	// }

	//fmt.Printf("Lenghth of data in all the data array: %d, Cap: %d \n", len(rows), cap(rows))

	// donerunning := make(chan bool)
	// go func() {
	//index := xlsx.NewSheet(sheetname)
	//xlsx.SetCellValue(sheetname, "A1", "Fathers name")
	//xlsx.SetCellValue("Sheet1", "B2", 100)
	/* excel title */

	ignoreheadercounter := 1
	for i := 0; i < len(rows)-1; i++ {
		// fmt.Println(fathersnamearray[i])
		// fmt.Println(noheaderarray[i])

		indexstring := strconv.Itoa(ignoreheadercounter + 1)
		//to concat the letter
		var stringbuffer bytes.Buffer
		stringbuffer.WriteString("H")
		stringbuffer.WriteString(indexstring)

		xlsx.SetCellValue(sheetname, stringbuffer.String(), fathersnamearray[i])

		ignoreheadercounter++
	}
	xlsx.SaveAs(newfilename)
	//donerunning <- true

	//}()
	//finito := <-donerunning
	// fmt.Println("imefika")
	//fmt.Println("finito", finito)

	fmt.Printf("took %v\n", time.Since(start))

}

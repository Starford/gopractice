package main

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// const max = 1 << 11
const filename string = "samplefile.xlsx"
const sheetname string = "sample"
const newfilename string = "output.xlsx"

// const filename string = "100000RecordsFull.xlsx"
// const sheetname string = "100000 Records"

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()
	//string array of names
	//rows := []string{"the headers", "father 8", "father 2", "father 5", "father 3", "father n.."}
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

	fmt.Println("*************************************")
	//fmt.Println("no header array", fathersnamearray)
	parallelMergesort3(fathersnamearray)
	//fmt.Println(fathersnamearray)

	donerunning := make(chan bool)
	go func() {
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
		donerunning <- true

	}()
	finito := <-donerunning
	// fmt.Println("imefika")
	fmt.Println("finito", finito)

	fmt.Printf("took %v\n", time.Since(start))

}

func merge(s []string, middle int) {
	helper := make([]string, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

func mergesort(s []string) {
	if len(s) > 1 {
		middle := len(s) / 2
		mergesort(s[:middle])
		mergesort(s[middle:])
		merge(s, middle)
	}
}

func parallelMergesort3(s []string) {
	len := len(s)

	if len > 1 {
		middle := len / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			parallelMergesort3(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergesort3(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

//Without goroutines+channels, this took over 19seconds to run. Now it takes 46.8ms
func main() {
	start := time.Now()

	extractChannel := make(chan *Order)
	transformChannel := make(chan *Order)
	doneChannel := make(chan bool)

	go extract(extractChannel)
	go transform(extractChannel, transformChannel)
	go load(transformChannel, doneChannel)

	<-doneChannel
	fmt.Println(time.Since(start))
}

type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

type Order struct {
	CustomerNumber int
	PartNumber     string
	Quantity       int
	UnitCost       float64
	UnitPrice      float64
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func extract(ch chan *Order) {
	f, err := os.Open("./orders.txt")
	handleError(err)
	defer f.Close()
	r := csv.NewReader(f)
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, err = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, err = strconv.Atoi(record[2])
		handleError(err)
		ch <- order
	}
	close(ch)

}

func transform(extractChannel, transformChannel chan *Order) {
	f, err := os.Open("./productList.txt")
	handleError(err)
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	handleError(err)
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, err = strconv.ParseFloat(record[1], 64)
		handleError(err)
		product.UnitPrice, err = strconv.ParseFloat(record[2], 64)
		handleError(err)
		productList[product.PartNumber] = product
	}

	numMessages := int64(0)

	for o := range extractChannel {
		atomic.AddInt64(&numMessages, 1)
		//The following only works on Windows as ++ isn't an atomic operator:
		//numMessages++
		go func(o *Order) {
			time.Sleep(3 * time.Millisecond)
			o.UnitCost = productList[o.PartNumber].UnitCost
			o.UnitPrice = productList[o.PartNumber].UnitPrice
			transformChannel <- o
			atomic.AddInt64(&numMessages, -1)
			//The following only works on Windows as -- isn't an atomic operator:
			//numMessages--
		}(o)
	}
	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	close(transformChannel)
}

func load(transformChannel chan *Order, doneChannel chan bool) {
	f, err := os.Create("./dest.txt")
	handleError(err)
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n",
		//Formatting above is done to print to specific column widths
		"PartNumber", "Quantity", "Unit Cost",
		"Unit Price", "Total Cost", "Total Price",
	)

	numMessages := int64(0)

	for o := range transformChannel {
		atomic.AddInt64(&numMessages, 1)
		//The following only works on Windows as ++ isn't an atomic operator:
		//numMessages++
		go func(o *Order) {
			time.Sleep(1 * time.Millisecond)
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
				o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice,
				o.UnitCost*float64(o.Quantity),
				o.UnitPrice*float64(o.Quantity))
			atomic.AddInt64(&numMessages, -1)
			//The following only works on Windows as -- isn't an atomic operator:
			//numMessages--
		}(o)
	}
	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	doneChannel <- true
}

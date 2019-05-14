package main

import (
	"fmt"
	"strings"
	"io"
	"os"
	"encoding/csv"
	"bufio"
	"log"
	"strconv"
)

type RatePlan struct{
	rateAt500, rateAt1000 float32
	earlyTerminationFee, length int
}

func (rp RatePlan) Calculate(fileName string) map[string]float32{
	costMap := map[string]float32{}
	// open the usage report csv file
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// ignore the first line of head
	firstLine, _ := reader.Read()
	fmt.Println(firstLine[1], firstLine[3])
	for {
		// parse the csv file
		line, error := reader.Read()
		if error == io.EOF {
			break
		}else if error != nil{
			log.Fatal(error)
		}
		date := strings.Fields(line[1])[0]
		usage := line[3]
		month := strings.Split(date, "-")[1]
		val, err := strconv.ParseFloat(usage, 32)
		if err == nil{
			costMap[month] += float32(val)
		}
	}
	total := float32(0)
	duration := 0
	for k, v := range costMap{
		if v < 500 {
			costMap[k] = v*rp.rateAt500
		}else{
			costMap[k] = (v-500)*rp.rateAt1000+500*rp.rateAt500
		}
		duration++
		total += costMap[k]
	}
	if duration < rp.length {
		costMap["Total"] = total + float32(rp.earlyTerminationFee)
		costMap["TermintationFee"] = float32(rp.earlyTerminationFee)
	}else {
		costMap["Total"] = total
	}
	costMap["Average"] = total / float32(duration)
	return costMap
}

func main(){
	rp := RatePlan{0.13, 0.12, 175, 6}
	costMap := rp.Calculate("usage_2019-02-20.csv")
	fmt.Println(costMap)
}
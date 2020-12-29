package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	r, g, y, err := openLog()
	if err != nil {
		log.Fatal(err)
	}

	workChan := make(chan string)
	re := regexp.MustCompile(`(?m)(\w*)-`)

	go genData("red", workChan)
	go genData("green", workChan)
	go genData("yellow", workChan)

	for {
		log.Println("Pause")
		time.Sleep(time.Millisecond * 500)
		select {
		case item := <-workChan:
			match := re.FindAllStringSubmatch(item, -1)
			switch match[0][1] {
			case "red":
				_, err := r.WriteString(fmt.Sprintf("%v\n", item))
				if err != nil {
					log.Fatal(err)
				}
			case "green":
				_, err := g.WriteString(fmt.Sprintf("%v\n", item))
				if err != nil {
					log.Fatal(err)
				}
			case "yellow":
				_, err := y.WriteString(fmt.Sprintf("%v\n", item))
				if err != nil {
					log.Fatal(err)
				}
			default:
				log.Println("Chan is empty")
			}
		}
	}

}

func genData(typeData string, workChan chan string) {
	for i := 0; i >= 0; i++ {
		workChan <- fmt.Sprintf("%v-%v", typeData, i)
	}
}

func openLog() (red, green, yellow *os.File, err error) {
	fRed, err := os.Create("./red.txt")
	if err != nil {
		return nil, nil, nil, err
	}

	fGreen, err := os.Create("./green.txt")
	if err != nil {
		return nil, nil, nil, err
	}

	fYellow, err := os.Create("./yellow.txt")
	if err != nil {
		return nil, nil, nil, err
	}

	return fRed, fGreen, fYellow, err
}

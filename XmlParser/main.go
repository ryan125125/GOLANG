package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Tests struct {
	XMLName xml.Name `xml:"Tests"` //root節點
	Xmlns   string   `xml:"xmlns"` //節點上的屬性
	Test    []Test   `xml:"Test"`  //讀取Test節點，也必須要用結構化的物件才存
}

type Test struct {
	TestId      string `xml:"TestId,attr"` //節點上的屬性
	TestType    string `xml:"TestType,attr"`
	Name        string `xml:"Name"`
	CommandLine string `xml:"CommandLine"`
	Input       string `xml:"Input"`
	Output      string `xml:"Output"`
}

func main() {
	file, err := os.Open("config.xml")
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)

	test := Tests{}
	xml.Unmarshal(data, &test)
	fmt.Println(test)

	for i := 0; i < len(test.Test); i++ {
		fmt.Println("TestId: " + test.Test[i].TestId)
		fmt.Println("TestType: " + test.Test[i].TestType)
		fmt.Println("CommandLine: " + test.Test[i].CommandLine)
		fmt.Println("Input: " + test.Test[i].Input)
		fmt.Println("Output: " + test.Test[i].Output)
	}
}

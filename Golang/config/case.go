package config

import (
	"fmt"
	"os"
	"encoding/xml"
	"io/ioutil"
)

type Config struct {
	Cases	[]Case  	`xml:"case"`
}

type Case struct {
	Package		string `xml:"package"`
	Casename    string `xml:"casename"`
}

// type Configuration struct {
// 	Enabled		bool 	`xml:"enabled"`
// 	Path 		string 	`xml:"path"`
// }

func ReadCaseXML(){
	dir1,_ := os.Getwd()
	dir := dir1 + `\config\case.xml`
	fmt.Println("dir1:",dir)
	File, err := os.Open(dir)
	if err != nil {
		fmt.Println("error opening file:",err)
		return
	}
	defer File.Close()

	// var conf Configuration
	// if err := xml.NewDecoder(xmlFile).Decode(&conf); err != nil {
	// 	fmt.Println("error decode file:",err)
	// 	return
	// }
	// fmt.Println(conf)
	// fmt.Println(conf.Enabled)
	// fmt.Println(conf.Path)

	xmlFile, err := ioutil.ReadAll(File)
	if err != nil {
		fmt.Println(err)
		return 
	}

	test := Config{}
	err = xml.Unmarshal(xmlFile,&test)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println("test:",test)
	for _,c := range test.Cases {
		fmt.Println("Package:",c.Package)
		fmt.Println("Casename:",c.Casename)
	}
}
package examples

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// our struct which contains the complete
// array of all Users in the file
type Animals struct {
	XMLName xml.Name `xml:"animals"`
	Animal  []Animal `xml:"animal"`
}

// the user struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type Animal struct {
	XMLName xml.Name `xml:"animal"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Desc    Desc     `xml:"desc"`
}

// a simple struct which contains all our
// social links
type Desc struct {
	XMLName  xml.Name `xml:"desc"`
	Color    string   `xml:"color"`
	Size     string   `xml:"size"`
	Cuteness int      `xml:"cuteness"`
}

func RunXML() {
	// Open our xmlFile
	xmlFile, err := os.Open("animals.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened animals.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var animals Animals
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &animals)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(animals.Animal); i++ {
		fmt.Println("User Type: " + animals.Animal[i].Type)
		fmt.Println("User Name: " + animals.Animal[i].Name)
		fmt.Println("Cuteness: " + strconv.Itoa(animals.Animal[i].Desc.Cuteness))
	}
}

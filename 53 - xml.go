package main

import (
	"encoding/xml"
	"fmt"
)
//field tags contains directives for the encoder and decoder
type Plant struct {
	//XMLName indicates the name of the XML element representing this structure
	XMLName xml.Name `xml:"plant"` 
	//"id,attr" means that id is an XML attibute rather than a nested element
	Id	int	`xml:"id,attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	//Emit XML representing out plant structure
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	//add a generic XML header to the output, append it explicitly
	fmt.Println(xml.Header + string(out))

	var p Plant
	//parse a stream of bytes with XML into a data structure
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		//parent>child>plant fiel tag tell the encoder to nestt all plants under <parent><child>
		Plants []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))
}
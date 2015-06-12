package main

//this is a nice to have but per instructions I don't know how else I would know about converting errors intelligently
import (
	"log"

	"https://github.com/RBdevelop/GoSon"
)

//log errors about reading  data  aka dirty data
func main() {
	source, err := apiconv.ReadSourceFile("apiconv/serviceA.json")
	if err != nil {
		log.Fatalf("error reading source data: %v", err)
	}
// log errors about converting data aka no comprede
	target, err := apiconv.Convert(source)
	if err != nil {
		log.Fatalf("error converting source to target: %v", err)
	}

//log errors about writing new data aka didn't effin work
	err = apiconv.WriteTargetData(target)
	if err != nil {
		log.Fatalf("error writing data: %v", err)
	}
}

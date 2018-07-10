package dictionary

import (
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

var lookup = make(map[string]string)

//TODO process dictionary

func FindWord(w string) string {
	if val, ok := lookup[strings.ToLower(w)]; ok {
		return val
	} else {
		return w
	}
}

//Grab dictionary -- FromFile
func InitMapping() {
	decodeFile, err := os.Open("scrambled.gob")
	if err != nil {
		fmt.Println("Error-- transformation not preformed.")
		fmt.Println(err)
		os.Exit(1)

	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)
	decoder.Decode(&lookup)
}

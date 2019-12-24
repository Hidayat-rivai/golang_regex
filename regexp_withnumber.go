package main

import "fmt"
import "regexp"

func main(){
	var text = "jeruk pisang0 dan anggur"
	var regex, err = regexp.Compile("[a-z]+\\d")
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	var hasil = regex.FindAllString(text,2)
	fmt.Println(hasil)
}
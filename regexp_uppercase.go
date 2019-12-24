package main

import "fmt"
import "regexp"

func main(){
	var text = "jeruk pisang dan ANGGUR"
	var regex, err = regexp.Compile("[A-Z]+")
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	var hasil = regex.FindAllString(text,2)
	fmt.Println(hasil)
}
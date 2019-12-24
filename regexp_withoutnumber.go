package main

import "fmt"
import "regexp"

func main(){
	var text = "jeruk pisang dan 4nggur"
	var regex, err = regexp.Compile("[a-z]+\\D")
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	var hasil = regex.FindAllString(text,2)
	fmt.Println(hasil)
}
package main

import "fmt"

func main() {

	var weight float32
	var height float32
	var bmi float32

	fmt.Println("Enter your weight in kilograms")
	fmt.Scan(&weight)
	fmt.Println("Enter your height in meters (m)")
	fmt.Scan(&height)

	bmi = weight / (height * height)

	if bmi < 18.5 {
		fmt.Println(" 'You are underweight, add more potato to the broth'")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println(" 'You have a normal weight, I have healthy envy of you'")
	} else {
		fmt.Println("'You are overweight, I know, the pandemic has affected us all'")
	}

}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 1. Literal types
	// Note: Literal type copy by value

	// var theString2 string
	// theString2 = "This is string 2"
	// theString1 := "This is String"

	theString := "This variable type string"
	theInt := 305
	theBool := true
	printString(theString)
	printInt(theInt)
	printBool(theBool)
	printUnderline()
	// 2. Type in Golang are static type
	// printString(theInt) // Compile error

	// 3. interface{} type can accept all types
	var interf interface{}
	interf = theString
	theString2 := interf.(string)

	interf = theInt
	interf = theBool

	printString(theString2)

	printInterface(theString)
	printInterface(theInt)
	printInterface(theBool)
	printUnderline()

	// 4. map type
	// Note: Map variable is the pointer of map, pass by pointer of map

	// map[Key type]value type{}{} (last branket for initial instance)
	// theMap := map[string]interface{}{} 
	theMap := make(map[string]interface{})
	theMap["firstname"] = "Chaiyapong"
	theMap["lastname"] = "Lapliengtrakul"
	theMap["citizen_id"] = "1234"

	theMapInt := map[string]int{}
	theMapInt["number1"] = 1
	theMapInt["number2"] = 2
	printMapInt(theMapInt)

	printMap(theMap)
	updateMap(theMap) // This will update map that pass in to function (pass by reference)
	printMapAsJSON(theMap)

	// citizenId := theMap["citizen_id"] // citizenId is interface type; if theMap not have citizen_id(Key) -> see line 64 
	// citizenIdStr = citizenId.(string) // convert interface to String

	// Note: We can use 2 variables to get map value
	gender, ok := theMap["gender"] // ok is parameter for check key is valid
	if ok {
		fmt.Println("gender is ", gender.(string)) // Casting
	} else {
		fmt.Println("No Gender specify")
	}

	printUnderline()

	// 5. slice type (dynamic array)
	// Note: We should name slice phurally
	// Note: Slices pass by pointer of slice (just like map)
	// theSlices := []string{}
	theSlices := make([]string, 0) // another to initial slices (second parameter is slice's size)
	theSlices = append(theSlices, "item 1")
	theSlices = append(theSlices, "item 2")
	theSlices = append(theSlices, "item 3")
	printSlice(theSlices)
	updateSliceIndex0(theSlices) // This will update value of index 0 of slice ; like map pass by reference
	printSliceAsJSON(theSlices)
	printUnderline()

	// 6. Struct type for object
	// Note: It is my practice to always create struct variable as pointer (*Type)
	//   to avoid confusion in team member, so will can always assume that strut is a pointer
	theCitizen := &Citizen{
		Firstname: "Chaiyapong",
		Lastname:  "Lapliengtrakul",
		CitizenID: "1234",
	}
	// theCitizen1 := &Citizen{ // has pointer -> pass by reference
	// 	CitizenID: "567",
	// } 
	// theCitizen2 := Citizen{} // no pointer

	printCitizen(theCitizen)
	// Note: JSON marshal will use struct tag
	printCitizenAsJSON(theCitizen)
	printUnderline()

	// 7. nil value
	// Note: Literal cannot be nil (string, int, bool, ..)
	// theString = nil // Error
	// theInt = nil    // Error
	// theBool = nil   // Error

	// Note: map, slice and struct and be nil
	theMap = nil
	theSlices = nil
	theCitizen = nil

	// var theMap1 map[string]interface{}{} // default nil 
	// var theSlices1 []string // default nil 
	// var theCitizen *Citizen // default nil 

	// Note: assign value to nil map and struct will cause runtime error
	// theMap["key"] = "value" // Runtime Error
	// theCitizen.Firstname = "Chaiyapong" // Runtime Error

	// Note: append value to nil slice is OK
	theSlices = append(theSlices, "value") // OK

	// 8. Check nil or zero using len()
	if len(theMap) == 0 { // can check nill or 0 size
		fmt.Println("theMap is nil")
	}
	if theMap == nil { // check nill only
		fmt.Println("theMap is nil")
	}

	if len(theSlices) == 0 {
		fmt.Println("theSlices is nil")
	}
	if theCitizen == nil { // checl nill in struc (only 1 way)
		fmt.Println("theCitizen is nil")
	}
	printUnderline()

	// 9. Enum is just constant in golang
	var theGender GenderType 
	// theGender := Male
	theGender = Male
	theGender = Female
	theGender = Unspecify
	switch theGender {
	case Male:
		fmt.Println("Gender is Male")
	case Female:
		fmt.Println("Gender is Female")
	case Unspecify:
		fmt.Println("Gender is Unspecify")
	}
	printGender(theGender)
	printUnderline()
}

func printUnderline() {
	fmt.Println("---")
}

func printString(input string) {
	fmt.Println("string = ", input)
}

func printInt(input int) {
	fmt.Println("int = ", input)
}

func printBool(input bool) {
	fmt.Println("bool = ", input)
}

func printInterface(input interface{}) {
	fmt.Println("interface = ", input)
}

func printMapInt(input map[string]int) {
	fmt.Println("map = ", input)
}

func printMap(input map[string]interface{}) {
	fmt.Println("map = ", input)
}

func updateMap(input map[string]interface{}) {
	input["updated"] = true
}

func printMapAsJSON(input map[string]interface{}) {
	js, _ := json.Marshal(input) // any to JSON
	fmt.Println("map JSON = ", string(js)) // print string
}

func printSlice(input []string) {
	fmt.Println("slice = ", input)
}

func updateSliceIndex0(input []string) {
	input[0] = "Item 0"
}

func printSliceAsJSON(input []string) {
	js, _ := json.Marshal(input)
	fmt.Println("slice JSON = ", string(js))
}

// Citizen is type represent person in country
type Citizen struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	CitizenID string `json:"citizen_id"`
}

func printCitizen(input *Citizen) {
	fmt.Println("Citizen = ", input)
}

func printCitizenAsJSON(input *Citizen) {
	js, _ := json.Marshal(input)
	fmt.Println("Citizen JSON = ", string(js))
}

// GenderType is enum for Gender
type GenderType string

const (
	// Unspecify is gender type for Unspecify
	Unspecify GenderType = "UNSPECIFY"
	// Male is gender type for Male
	Male GenderType = "MALE"
	// Female is gender type for Female
	Female GenderType = "FEMALE"
)

func printGender(input GenderType) {
	fmt.Println("Gender = ", input)
}

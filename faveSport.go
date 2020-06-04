package main

import "fmt"

//Rating for different sports
func SportRating(sport string) string {

	if sport == "Football" || sport == "football" {
		return "Football: 100/10"
	} else if sport == "Soccer" || sport == "soccer" {
		return "Soccer: 9/10"
	} else if sport == "Pingpong" || sport == "pingpong" {
		return "Pingpong: 10/10"
	} else {
		return "Invalid Sport or Argument"
	}

}

func main() {
	fmt.Println("Aren't sports fun??")

	myfavesport := SportRating("Football")

	fmt.Println("My favorite sport is " + myfavesport)
}

package main

import "fmt"

func main() {
	fmt.Println("Aren't sports fun??")
}

//Rating for different sports
func FavoriteSport(sport string) string {

	if sport == "Football" || sport == "football" {
		return "Football: 100/10"
		/*} else if sport == "Soccer" || sport == "soccer" {
			return "Soccer: 9/10"
		} else if sport == "Pingpong" || sport == "pingpong" {
			return "Pingpong: 10/10"*/
	} else {
		return "Invalid Sport or Argument"
	}

}

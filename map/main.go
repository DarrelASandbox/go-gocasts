package main

import "fmt"

/* Using map
func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	fmt.Println(colors)
}
*/

/* Using map with var
func main() {
	// https://go.dev/blog/maps
	var colors map[string]string
	colors = make(map[string]string)
	colors["red"] = "#ff0000"

	fmt.Println(colors)
}

*/

/* Using map with delete
func main() {
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["white"] = "#fff"

	delete(colors, "red")

	fmt.Println(colors)
}

*/

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "#FFF",
		"black": "#000",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

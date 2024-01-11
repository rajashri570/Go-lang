/* Create a movie database program. Each movie has a title, director, release year, and genre.
Implement functions to add new movies, search for movies by director,
and display the list of movies released in a specific year.
*/

package main

import (
	"fmt"
	"os"
)

var movies []Movie

type Movie struct {
	title     string
	director  string
	releaseDt string
}

func (m *Movie) addMovie(nm string, dir_nm string, rel string) {

	movie := Movie{ //Query is why to use  *Movie poinrtt
		title:     nm,
		director:  dir_nm,
		releaseDt: rel,
	}
	movies = append(movies, movie)
	fmt.Println("record added successfully!")
}
func (m *Movie) searchMovie(dname string) {
	found := false
	for i := range movies {
		if movies[i].director == dname {
			fmt.Println("Movie name:", movies[i].title)
			found = true
		}
	}
	if !found {
		fmt.Println("No movies found for the director:", dname)
	}

}
func (m *Movie) listMovie(year string) {
	for i, _ := range movies {
		if movies[i].releaseDt == year {

			fmt.Println("Movies released on date: ", movies[i].title)
		}
	}
}

func main() {
	// 	var movie []Movie
	var ch int
	var mv Movie
	var name string
	var dir_name string
	var release string
	//var movie Movimovie

	for {
		fmt.Println("-----Menu-------")
		fmt.Println("1.Add movie")
		fmt.Println("2.search movie")
		fmt.Println("3.ListMovies")

		fmt.Println("\nEnter ur choice: ")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Println("Enter the movie name :")
			fmt.Scan(&name)
			fmt.Println("Enter the Director Name :")
			fmt.Scan(&dir_name)
			fmt.Println("Enter the date")
			fmt.Scan(&release)

			mv.addMovie(name, dir_name, release)

		case 2:
			fmt.Println("Enter the director name to search movies:")
			fmt.Scan(&dir_name)
			mv.searchMovie(dir_name)

		case 3:
			fmt.Println("Enter the year to list :")
			fmt.Scan(&release)
			mv.listMovie(release)

		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid input")

		}

	}
}

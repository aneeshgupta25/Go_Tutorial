package maps

import "fmt"

type Product struct {
	id string
	name string
	date string
}

type floatMap map[Product]string

func (f floatMap) output() {
	fmt.Println(f)
}

func main() {
	hobbies := [3]string{"skating", "painting", "studying"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	new_hobbies := hobbies[1:]
	fmt.Println(new_hobbies)

	// adding content of a list to another
	b := []int{1,3,4,5}
	c := []int{2,4,6,1}
	b = append(b, c...)
	fmt.Println(b, c)
		

	// map
	// make map for efficiency
	a := make(map[Product]string, 4)
	d := make(floatMap, 5)
	d.output()

	a[Product{
		id: "Hello",
		name: "Aneesh",
		date: "Gupta",
	}] = "gupta"
	a[Product{
		id: "Hello",
		name: "Aneesh",
		date: "Gupta",
	}] = "Archit"
	// a["archit"] = "gupta"
	// delete(a, "aneesh")
	fmt.Println(a)

	// iterating over arrays
	array := []string{}
	array = append(array, "aneesh")
	array = append(array, "archit")

	for index, value := range array {
		fmt.Println(index, value)	
	}

	for _, value := range array {
		// if index not required, replace with _
		fmt.Println(value)
	}

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
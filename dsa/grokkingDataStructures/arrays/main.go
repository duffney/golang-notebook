package main

func main() {
	// ints := []int{1, 2, 3}
	// // unsorted
	// ua := NewUnsortedArray[int](3)

	// // fill array
	// for _, i := range ints {
	// 	ua.Insert(i)
	// }
	// // rm 1 from array
	// ua.DeleteByIndex(0)

	// // add 6 to array, making it full again
	// if !ua.Insert(4) {
	// 	fmt.Println("insert failed, array full")
	// }

	// if !ua.Insert(5) {
	// 	fmt.Println("insert failed, array full")
	// }

	// ua.Print()

	// ua.Find(3)

	// max := ua.Max()
	// fmt.Println(max)
	// sa := NewSortedArray(4)

	// sa.Insert(1)
	// sa.Insert(3)
	// sa.Insert(2)
	// sa.Insert(4)

	// sa.Delete(2)
	// sa.Insert(5)
	// fmt.Println(sa)
	// sa.LinearSearch(3)

	bSearch := NewSortedArray(100)
	ints := make([]int, 100)
	for i := range 100 {
		ints[i] = i + 1
		bSearch.Insert(ints[i])
	}
	// fmt.Println(bSearch)
	bSearch.Find(30)
	bSearch.BinarySearchNonRecursive(30)

}

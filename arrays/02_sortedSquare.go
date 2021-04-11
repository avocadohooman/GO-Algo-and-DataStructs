package main

import "fmt"

func main() {
	array := []int{-3, -2, -1};

	fmt.Println("Result: ", SortedSquaredArray(array));
}

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	if (len(array) == 0) {
		return nil
	}
	i := 0;
	for (i < len(array)) {
		array[i] *= array[i]
		i++;
	}
	tmp := array[0];
	i = 0;
	for (i < len(array)) {
		k := 0;
		for (k < len(array)) {
			if (array[i] < array[k]) {
				tmp = array[k];
				array[k] = array[i];
				array[i] = tmp;
			}
			k++;
		}
		i++;
	}
	return array;
}

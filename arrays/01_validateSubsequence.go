package main

import "fmt"

func main() {
	array := []int{1, 1, 6, 1};
	sequence := []int{1, 1, 1, 6};
	
	// array := []int{5, 1, 22, 25, 6, -1, 8, 10};
	// sequence := []int{1, 6, -1, 10};

	fmt.Println("Result: ", IsValidSubsequence(array, sequence));
}

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.

	if (len(sequence) > len(array)) {
		return false;
	}
	if (len(sequence) == len(array)) {
		if (!compare(sequence, array)) {
			return false;
		} else {
			return true;
		}
	}
	k := 0;
	copyArray := array;
	match := 0;
	for k < len(sequence) {
		i := 0;
		for i < len(copyArray) {
			if (sequence[k] == copyArray[i]) {
				copyArray = cleanArray(copyArray, i);
				match += 1;
				break ;
			}
			i++;
		}
		k++;
	}
	if (match != len(sequence)) {
		return false;
	}
	return true;
}

func cleanArray(array []int, indexToRemove int) []int {
	
	size := len(array) - 1;
	var cleanedArray = make([]int, size)

	i := 0;
	k := 0;
	for i < len(cleanedArray) {
		if (k == indexToRemove) {
			k++;
		}
		cleanedArray[i] = array[k];
		i++;
		k++;
	}
	return cleanedArray;
}

func compare(sequence []int, compareSequence[]int) bool {
	i := 0;
	for i < len(sequence) {
		if (sequence[i] != compareSequence[i]) {
			return false;
		}
		i++;
	}
	return true;
}
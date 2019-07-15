package bliss

import (
	"fmt"
)

func ExampleDistanceFile() {
	filename1 := "file1.mp3"
	filename2 := "file2.mp3"
	_, _, distance, err := DistanceFile(filename1, filename2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Distance between %s and %s is: %f\n", filename1, filename2, distance)
}

func ExampleCosineSimilarityFile() {
	filename1 := "file1.mp3"
	filename2 := "file2.mp3"
	_, _, similarity, err := CosineSimilarityFile(filename1, filename2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Similarity between %s and %s is: %f\n", filename1, filename2, similarity)
}

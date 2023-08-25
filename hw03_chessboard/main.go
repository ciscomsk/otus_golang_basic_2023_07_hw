package main

import "fmt"

func main() {
	// var size int
	// getBoardSizeV1(&size)

	size, err := getBoardSizeV2()
	if err != nil {
		return
	}
	drawBoard(size)
}

// func getBoardSizeV1(size *int) {
//	fmt.Println("Enter a chessboard size: ")
//	_, err := fmt.Scanf("%d", size)
//	if err != nil {
//		return
//	}
//}

func getBoardSizeV2() (int, error) {
	var size int

	fmt.Println("Enter a chessboard size: ")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println("Error while reading user input")
		return 0, err
	}

	return size, nil
}

func drawBoard(size int) {
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

package main

import "fmt"

func main() {
	//var size byte
	//getBoardSizeV1(&size)

	size := getBoardSizeV2()
	drawBoard(size)
}

func getBoardSizeV1(size *int) *int {
	fmt.Println("Enter a chessboard size: ")
	_, err := fmt.Scanf("%d", size)
	if err != nil {
		return nil
	}

	return size
}

func getBoardSizeV2() int {
	var size int

	fmt.Println("Enter a chessboard size: ")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		return 0
	}

	return size
}

func drawBoard(size int) {
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if i%2 == 0 {
				if j%2 == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("#")
				}
			} else {
				if j%2 == 0 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")
	}
}

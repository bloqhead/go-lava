package main

import (
	"fmt"
	"image"
	"encoding/hex"
	"crypto/sha256"
	_ "image/jpeg"
	"os"
)

func main() {
	var imageFile string
	var desiredLength int
	var applyHash string

	// image path
	fmt.Println("Enter the path of the image you want to read: ")
	fmt.Scan(&imageFile)
	fmt.Printf("The path of your image is %s. Checking to be sure I can access it...\n", imageFile)

	// open the image
	img, err := os.Open(imageFile)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	} else {
		fmt.Printf("Successfully opened the image at %s. Proceeding...\n", imageFile)
	}

	defer img.Close()

	// hash option
	fmt.Printf("Do you want to apply a hash function to the generated string? (y/n): ")
	fmt.Scan(&applyHash)

	if applyHash == "y" || applyHash == "Y" {
		fmt.Println("Note: When applying a hash function, the generated string will have a fixed length of 64 characters.")
		fmt.Println("Enter the desired length of the generated string (cannot exceed 64): ")
	} else {
		fmt.Println("Enter the desired length of the generated string: ")
	}

	fmt.Scan(&desiredLength)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// decode the image
	imgData, _, err := image.Decode(img)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	var pixelData []byte
	imageBounds := imgData.Bounds()

	// store our pixel data
	for y := imageBounds.Min.Y; y < imageBounds.Max.Y; y++ {
		for x := imageBounds.Min.X; x < imageBounds.Max.X; x++ {
			r, g, b, a := imgData.At(x, y).RGBA()
			pixelData = append(pixelData, byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8))
		}
	}

	hexString := hex.EncodeToString(pixelData)

	// apply a hash function to the generated string if the user said yes to the option
	if applyHash == "y" || applyHash == "Y" {
		hash := sha256.New()
		hash.Write([]byte(hexString))
		hexString = hex.EncodeToString(hash.Sum(nil))
	}

	if desiredLength > len(hexString) {
		fmt.Printf("Desired length is greater than the length of the generated string\n")
		return
	} else if desiredLength < len(hexString) && desiredLength > 0 {
		hexString = hexString[:desiredLength]
	}

	// return the image data
	fmt.Printf("Generated string: %v\n", hexString)
}

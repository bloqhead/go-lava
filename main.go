package main

import (
	"fmt"
	"image"
	"encoding/hex"
	"crypto/sha512"
	_ "image/jpeg"
	"os"
)

func main() {
	var imageFile string
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
	fmt.Printf("Do you want to apply a hash function to the generated string? I use SHA512 (y/n): ")
	fmt.Scan(&applyHash)

	if applyHash == "y" || applyHash == "Y" {
		fmt.Printf("Applying a hash function to the generated string...\n")
	} else if applyHash == "n" || applyHash == "N" {
		fmt.Printf("Not applying a hash function to the generated string...\n")
	} else {
		fmt.Printf("Invalid option. Please enter 'y' or 'n'.\n")
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
		hash := sha512.New()
		hash.Write([]byte(hexString))
		hexString = hex.EncodeToString(hash.Sum(nil))
	}

	// return the image data
	fmt.Printf("Generated string: %v\n", hexString)
}

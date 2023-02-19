/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unihtml"
	"github.com/unidoc/unipdf/v3/creator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Err: provided invalid arguments. No UniHTML server path provided")
		os.Exit(1)
	}

	// Establish connection with the UniHTML Server.
	if err := unihtml.Connect(os.Args[1]); err != nil {
		fmt.Printf("Err:  Connect failed: %v\n", err)
		os.Exit(1)
	}

	// Get new PDF Creator.
	c := creator.New()

	// Read the content of the simple.html file and load it to the conversion.
	htmlDocument, err := unihtml.NewDocument("simple.html")
	if err != nil {
		fmt.Printf("Err: NewDocument failed: %v\n", err)
		os.Exit(1)
	}

	// Draw the html document file in the context of the creator.
	if err = c.Draw(htmlDocument); err != nil {
		fmt.Printf("Err: Draw failed: %v\n", err)
		os.Exit(1)
	}

	// Write the result file to PDF.
	if err = c.WriteToFile("simple.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}

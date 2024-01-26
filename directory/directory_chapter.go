/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unihtml"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

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

	// Get new PDF creator.
	c := creator.New()

	// Create new chapter in the creator.
	ch := c.NewChapter("Directory")

	// Read the content of the directory with HTML, CSS and Images files.
	htmlDocument, err := unihtml.NewDocument("data")
	if err != nil {
		fmt.Printf("Err: NewDocument failed: %v\n", err)
		os.Exit(1)
	}

	// Add this document to the context of the chapter.
	if err = ch.Add(htmlDocument); err != nil {
		fmt.Printf("Err: Adding HTML Document failed: %v\n", err)
		os.Exit(1)
	}

	// Draw the chapter in the context of the creator.
	if err = c.Draw(ch); err != nil {
		fmt.Printf("Err: Draw failed: %v\n", err)
		os.Exit(1)
	}

	// Write the results to the file.
	if err = c.WriteToFile("directory_chapter.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}

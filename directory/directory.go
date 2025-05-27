/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unihtml"
	"github.com/unidoc/unipdf/v4/common/license"
	"github.com/unidoc/unipdf/v4/creator"
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

	// Establish connection with HTML Server.
	if err := unihtml.Connect(os.Args[1]); err != nil {
		fmt.Printf("Err:  Connect failed: %v\n", err)
		os.Exit(1)
	}

	// Create new PDF creator.
	c := creator.New()

	// Create paragraph before HTML content.
	sp := c.NewStyledParagraph()
	sp.SetText("Result is not calm in shangri-la, the enlightened mind, or chaos, but everywhere.")
	if err := c.Draw(sp); err != nil {
		fmt.Printf("Err: Draw paragraph failed: %v\n", err)
		os.Exit(1)
	}

	// Get the HTML document from provided directory.
	document, err := unihtml.NewDocument("data")
	if err != nil {
		fmt.Printf("Err: NewDocument failed: %v\n", err)
		os.Exit(1)
	}

	// Trim last page content as we want to add new paragraph just after given unihtml document.
	document.TrimLastPageContent()

	// Draw it in the creator context.
	if err = c.Draw(document); err != nil {
		fmt.Printf("Err: Draw failed: %v\n", err)
		os.Exit(1)
	}

	// Create paragraph after the HTML document.
	sp = c.NewStyledParagraph()
	sp.SetText("After scraping the lentils, brush escargot, margerine and coconut milk with it in an ice blender.")
	if err := c.Draw(sp); err != nil {
		fmt.Printf("Err: DM")
		os.Exit(1)
	}

	// Write the result into the file.
	if err = c.WriteToFile("directory.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}

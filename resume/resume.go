/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unihtml"
	"github.com/unidoc/unioffice/common/license"

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

	// Connect with the UniHTML Server.
	if err := unihtml.Connect(os.Args[1]); err != nil {
		fmt.Printf("Err:  Connect failed: %v\n", err)
		os.Exit(1)
	}

	// Get new PDF creator.
	c := creator.New()

	// Create new document based on the HTML file called resume.html.
	htmlDocument, err := unihtml.NewDocument("resume.html")
	if err != nil {
		fmt.Printf("Err: NewDocument failed: %v\n", err)
		os.Exit(1)
	}

	htmlDocument.TrimLastPageContent()

	// Draw HTML Document in the Creator.
	if err = c.Draw(htmlDocument); err != nil {
		fmt.Printf("Err: Draw failed: %v\n", err)
		os.Exit(1)
	}

	// Some Paragraph used for checking where would be the place after HTML Document.
	p := c.NewParagraph("Some paragraph text used for checking the position of the paragraph")

	// Draw the Paragraph in the creator context.
	if err = c.Draw(p); err != nil {
		fmt.Printf("Err: Draw paragraph failed: %v\n", err)
		os.Exit(1)
	}

	// Write the output of the PDF creator in the resume.pdf file.
	if err = c.WriteToFile("resume.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}

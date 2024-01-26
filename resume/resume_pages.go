/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unidoc/unihtml"
	"github.com/unidoc/unihtml/sizes"
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

	// Set Page size for the Document.
	if err := htmlDocument.SetPageSize(sizes.A6); err != nil {
		fmt.Printf("Err: invalid page size: %v\n", err)
		os.Exit(1)
	}

	// Set margins for the HTML Document.
	htmlDocument.SetMargins(10, 10, 10, 10)

	// Extract Pdf pages directly from the HTML document.
	pages, err := htmlDocument.GetPdfPages(context.Background())
	if err != nil {
		fmt.Printf("Err: Getting Pages failed: %v\n", err)
		os.Exit(1)
	}

	// Add page by page to the creator context.
	for _, page := range pages {
		if err = c.AddPage(page); err != nil {
			fmt.Printf("Err: Adding page failed: %v\n", err)
			os.Exit(1)
		}
	}

	// Write the result to file.
	if err = c.WriteToFile("resume_pages.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}

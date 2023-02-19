/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/unidoc/unihtml"
	"github.com/unidoc/unihtml/sizes"
	"github.com/unidoc/unipdf/v3/model"

	"github.com/unidoc/unipdf/v3/creator"
)

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

	// Create new document based on the html file that contain external links.
	htmlDocument, err := unihtml.NewDocument("link.html")
	if err != nil {
		fmt.Printf("Err: NewDocument failed: %v\n", err)
		os.Exit(1)
	}

	// Set up the page size to the ISO A5.
	if err = htmlDocument.SetPageSize(sizes.A5); err != nil {
		fmt.Printf("Err: Setting page size failed: %v\n", err)
		os.Exit(1)
	}

	// Set up margins with the unit that fits you best Millimeter, Point, Inch.
	htmlDocument.SetMarginLeft(sizes.Millimeter(10))
	htmlDocument.SetMarginRight(sizes.Point(10))
	htmlDocument.SetMarginTop(sizes.Millimeter(10))
	htmlDocument.SetMarginBottom(sizes.Inch(0.5))

	// The unihtml module converts the data by connecting to the unihtml-server.
	// It is wise to set up the context timeout in case the client is waiting on the connection.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Convert and get all pdf pages.
	pages, err := htmlDocument.GetPdfPages(ctx)
	if err != nil {
		fmt.Printf("Err: Getting pages failed: %v\n", err)
		os.Exit(1)
	}

	// Add page one by one to the creator.
	for _, p := range pages {
		if err := c.AddPage(p); err != nil {
			fmt.Printf("Err: adding page failed: %v\n", err)
			os.Exit(1)
		}
	}

	// Define PDF document properties.
	model.SetPdfAuthor("Jacek Kucharczyk")
	model.SetPdfTitle("Margins and Properties on PDF Document")
	model.SetPdfKeywords("margins properties unihtml")
	model.SetPdfSubject("Subject")
	model.SetPdfCreationDate(time.Date(2020, 10, 10, 15, 30, 20, 0, time.Local))
	model.SetPdfModifiedDate(time.Now())

	// Write the output of the PDF creator in the result.pdf file.
	if err = c.WriteToFile("result.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}

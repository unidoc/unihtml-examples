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

	// Create new document based on the HTML file called resume.html.
	webDocument, err := unihtml.NewDocument("https://www.google.com")
	if err != nil {
		fmt.Printf("Err: NewDocument failed: %v\n", err)
		os.Exit(1)
	}

	if err = webDocument.SetPageSize(sizes.A3); err != nil {
		fmt.Printf("Err: Setting page size failed: %v\n", err)
		os.Exit(1)
	}
	webDocument.SetMargins(30, 30, 30, 30)
	webDocument.SetLandscapeOrientation()

	// The unihtml module converts the data by connecting to the unihtml-server.
	// What's more getting document from external URL requires server to connect to external website, where
	// the connection might be slow or unavailable.
	// It is wise to set up the context timeout in case the client is waiting on the connection.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Convert and get all pdf pages.
	pages, err := webDocument.GetPdfPages(ctx)
	for _, p := range pages {
		if err := c.AddPage(p); err != nil {
			fmt.Printf("Err: adding page failed: %v\n", err)
			os.Exit(1)
		}
	}

	// Write the output of the PDF creator in the weburl.pdf file.
	if err = c.WriteToFile("weburl.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}

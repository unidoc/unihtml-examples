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

	htmlContent := `
		 <!DOCTYPE html>
		 <html>
		 <head>
			 <style>
				 body {
					 background-color: #6E85F7;
					 font-size-adjust: initial;
				 }
			 </style>
		 </head>
		 <body>
		 
		 <h1>Oh Hi...</h1>
		 <p>It works!</p>
		 
		 </body>
		 </html>	
	 `

	// Get new PDF Creator.
	c := creator.New()

	// Convert the HTML content to UniHTML document.
	htmlDocument, err := unihtml.NewDocumentFromString(htmlContent)
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
	if err = c.WriteToFile("simple-from-text.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}

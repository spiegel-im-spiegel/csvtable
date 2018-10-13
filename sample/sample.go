package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/csvtable"
)

func main() {
	flag.Parse()
	if flag.NArg() > 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	var r io.Reader
	if flag.NArg() == 0 {
		r = os.Stdin
	} else {
		f, err := os.Open(flag.Arg(0)) //maybe file path
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer f.Close()
		r = f
	}

	cr := csv.NewReader(r)
	cr.Comma = '\t'
	cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	cr.TrimLeadingSpace = true // leading white space in a field is ignored.

	ct, err := csvtable.New(cr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//fmt.Println("cols :", ct.Cols())
	//fmt.Println("rows :", ct.Rows())
	cols := "city/temperature"
	w := csv.NewWriter(os.Stdout)
	w.Comma = cr.Comma
	//w.UseCRLF = true
	//w.WriteAll(ct.OutputAll())
	header, body := ct.Output(strings.Split(cols, "/"))
	w.Write(header)
	w.WriteAll(body)
}

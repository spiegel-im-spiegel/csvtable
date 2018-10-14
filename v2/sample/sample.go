package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	csvtable "github.com/spiegel-im-spiegel/csvtable/v2"
)

func readWriteTSV(r io.Reader, w io.Writer) error {
	cr := csv.NewReader(r)
	cr.Comma = '\t'
	cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	cr.TrimLeadingSpace = true // leading white space in a field is ignored.

	cw := csv.NewWriter(w)
	cw.Comma = cr.Comma
	//cw.UseCRLF = true
	defer cw.Flush()

	ct, err := csvtable.New(cr, true)
	if err != nil {
		return err
	}

	//header := []string{}
	//cw.Write(ct.Header())
	header := strings.Split("city/temperature", "/")
	cw.Write(header)
	for {
		if err := ct.Next(); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		dt := ct.GetByName(header)
		cw.Write(dt)
	}
	return nil
}

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

	if err := readWriteTSV(r, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

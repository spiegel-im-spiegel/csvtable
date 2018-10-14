package csvtable

import (
	"encoding/csv"
	"errors"
)

//Errors in csvtable package
var (
	ErrNoData = errors.New("no data")
)

//CsvTable is CSV/TSV table class
type CsvTable struct {
	reader *csv.Reader
	header []string
	col    map[string]int
	row    []string
}

//New returns new CsvTable instance
func New(r *csv.Reader, isHeader bool) (*CsvTable, error) {
	ct := &CsvTable{reader: r, col: map[string]int{}}
	if isHeader {
		h, err := ct.reader.Read()
		if err != nil {
			return ct, err
		}
		ct.header = h
		for i, e := range ct.header {
			ct.col[e] = i
		}
	}
	return ct, nil
}

//IsHeader returns true if csv header exist
func (ct *CsvTable) IsHeader() bool {
	if ct == nil {
		return false
	}
	if len(ct.header) == 0 {
		return false
	}
	return true
}

//Header returns header record
func (ct *CsvTable) Header() []string {
	if !ct.IsHeader() {
		return []string{}
	}
	return ct.header
}

//Cols returns count of columns
func (ct *CsvTable) Cols() int {
	if ct == nil {
		return 0
	}
	if ct.IsHeader() {
		return len(ct.header)
	}
	return len(ct.row)
}

//Next reads next record
func (ct *CsvTable) Next() error {
	if ct == nil {
		return ErrNoData
	}
	d, err := ct.reader.Read()
	if err != nil {
		ct.row = nil
		return err
	}
	ct.row = d
	return nil
}

//GetByName returns element by column name
func (ct *CsvTable) GetByName(cols []string) []string {
	if ct == nil {
		return []string{}
	}
	if len(cols) == 0 {
		return ct.Get(nil)
	}
	dt := make([]string, 0, len(cols))
	for _, s := range cols {
		dt = append(dt, ct.getElement(ct.colnum(s)))
	}
	return dt
}

//Get returns elements
func (ct *CsvTable) Get(cols []int) []string {
	if ct == nil {
		return []string{}
	}
	if len(cols) == 0 {
		return ct.row
	}
	dt := make([]string, 0, len(cols))
	for _, c := range cols {
		dt = append(dt, ct.getElement(c))
	}
	return dt
}

func (ct *CsvTable) colnum(s string) int {
	if ct == nil {
		return -1
	}
	if !ct.IsHeader() {
		return -1
	}
	if c, ok := ct.col[s]; ok {
		return c
	}
	return -1
}

func (ct *CsvTable) getElement(col int) string {
	if ct == nil {
		return ""
	}
	if col < 0 {
		return ""
	}
	rowdata := ct.row
	if col >= len(rowdata) {
		return ""
	}
	return rowdata[col]
}

package csvtable

import (
	"encoding/csv"
	"errors"
)

var (
	NoData = errors.New("no data")
)

//CsvTable is CSV/TSV table class
type CsvTable struct {
	header []string
	col    map[string]int
	body   [][]string
}

//New returns new CsvTable instance
func New(r *csv.Reader) (*CsvTable, error) {
	ct := &CsvTable{}
	err := ct.readAll(r)
	return ct, err
}

//Cols returns count of columns
func (ct *CsvTable) Cols() int {
	if ct == nil {
		return 0
	}
	return len(ct.header)
}

//Rows returns count of rows
func (ct *CsvTable) Rows() int {
	if ct == nil {
		return 0
	}
	return len(ct.body)
}

//Get returns element in table
func (ct *CsvTable) Get(row int, cols []string) []string {
	if len(cols) == 0 {
		return []string{}
	}
	dt := make([]string, 0, len(cols))
	for _, s := range cols {
		dt = append(dt, ct.getElement(row, ct.colnum(s)))
	}
	return dt
}

//Output returns data in columns
func (ct *CsvTable) Output(cols []string) (header []string, body [][]string) {
	if ct == nil {
		header = cols
		body = [][]string{}
		return
	}
	if len(cols) == 0 {
		header, body = ct.OutputAll()
		return
	}
	header = cols
	body = make([][]string, 0, ct.Rows()+1)
	for i := 0; i < ct.Rows(); i++ {
		body = append(body, ct.Get(i, header))
	}
	return
}

//OutputAll returns all data
func (ct *CsvTable) OutputAll() (header []string, body [][]string) {
	if ct == nil {
		header = []string{}
		body = [][]string{}
		return
	}
	header = ct.header
	body = ct.body
	return
}

func (ct *CsvTable) readAll(r *csv.Reader) error {
	if ct == nil {
		ct = &CsvTable{}
	}
	ct.col = map[string]int{}

	dt, err := r.ReadAll()
	if err != nil {
		return err
	}
	l := len(dt)
	if l > 0 {
		ct.header = dt[0]
		for i, e := range ct.header {
			ct.col[e] = i
		}

		if l > 1 {
			ct.body = dt[1:]
		}
	}
	return nil
}

func (ct *CsvTable) colnum(s string) int {
	if ct == nil {
		return -1
	}
	if c, ok := ct.col[s]; ok {
		return c
	}
	return -1
}

func (ct *CsvTable) getElement(row, col int) string {
	if ct == nil {
		return ""
	}
	if col < 0 || col >= ct.Cols() || row < 0 || row >= ct.Rows() {
		return ""
	}
	return ct.body[row][col]
}

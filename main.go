package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type record struct {
	Name     string
	Age      int
	Salary   float64
	Retiring bool
	Notes    *string
}

// toCSVRecord returns a slice of strings that can be fed to
// the [csv.Writer.Write] method.
func (r record) toCSVRecord() []string {
	var (
		age      = fmt.Sprintf("%d", r.Age)
		salary   = fmt.Sprintf("%g", r.Salary)
		retiring = "false"
		notes    = ""
	)
	if r.Retiring {
		retiring = "true"
	}
	if r.Notes != nil {
		notes = *r.Notes
	}
	return []string{r.Name, age, salary, retiring, notes}
}

// toJTTRow returns a JSON Typed Table row as a properly formatted
// string.
func (r record) toJTTRow() string {
	notes := "null"
	if r.Notes != nil {
		notes = fmt.Sprintf("%q", *r.Notes)
	}
	return fmt.Sprintf("%q\t%d\t%g\t%t\t%s", r.Name, r.Age, r.Salary, r.Retiring, notes)
}

var (
	header = []string{"Name", "Age", "Salary", "Retiring", "Notes"}
	ptr1   = ""
	rec1   = record{"Foo", 1, 2.1, false, &ptr1}
	rec2   = record{"Bar", 4, 6.75, true, nil}
	ptr2   = "wants to be promoted to \"Boss\", and\nget a raise"
	rec3   = record{"Baz", 2, 2.5, false, &ptr2}
)

func main() {
	usage := func() {
		fmt.Fprintf(os.Stderr, `usage main: -csv|-jtt N

Write 3 test records N times, in either CSV or JSON Typed Table format; must specify one and only one of -csv or -jtt.

`)
		flag.PrintDefaults()
		os.Exit(2)
	}

	var (
		csvFlag = flag.Bool("csv", false, "write test data as CSV; cannot be used with -jtt")
		jttFlag = flag.Bool("jtt", false, "write test data as JSON Typed Table; cannot be used with -csv")
	)

	flag.Parse()
	if len(flag.Args()) != 1 ||
		((*csvFlag && *jttFlag) ||
			!(*csvFlag || *jttFlag)) {
		usage()
	}
	nArg := flag.Arg(0)
	n, err := strconv.Atoi(nArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not parse %s as int\n", nArg)
		os.Exit(2)
	}

	switch {
	case *csvFlag:
		writeTestCSV(os.Stdout, n)
	case *jttFlag:
		writeTestJT(os.Stdout, n)
	}
}

func writeTestCSV(w io.Writer, n int) {
	writer := csv.NewWriter(w)
	writer.Write(header)
	for range n {
		writer.Write(rec1.toCSVRecord())
		writer.Write(rec2.toCSVRecord())
		writer.Write(rec3.toCSVRecord())
	}
	writer.Flush()
}

func writeTestJT(w io.Writer, n int) {
	sep := ""
	for _, field := range header {
		fmt.Fprintf(w, "%s%q", sep, field)
		sep = "\t"
	}
	fmt.Fprint(w, "\n")

	for range n {
		fmt.Fprintln(w, rec1.toJTTRow())
		fmt.Fprintln(w, rec2.toJTTRow())
		fmt.Fprintln(w, rec3.toJTTRow())
	}
}

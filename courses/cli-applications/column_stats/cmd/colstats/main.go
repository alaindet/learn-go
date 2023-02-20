package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	op := flag.String("op", "sum", "Operation to be executed")
	column := flag.Int("col", 0, "CSV column on which to execute operation")
	flag.Parse()

	err := run(flag.Args(), *op, *column, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filenames []string, op string, column int, out io.Writer) error {
	var operationFunc statsFunc

	if len(filenames) == 0 {
		return ErrNoFiles
	}

	if column < 0 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, column)
	}

	switch op {
	case "sum":
		operationFunc = sum
	case "avg":
		operationFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	consolidate := make([]float64, 0)

	for _, fname := range filenames {

		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("cannot open file: %w", err)
		}

		data, err := csv2float(f, column)
		if err != nil {
			return err
		}

		err = f.Close()
		if err != nil {
			return err
		}

		consolidate = append(consolidate, data...)
	}

	_, err := fmt.Fprintln(out, operationFunc(consolidate))
	return err
}

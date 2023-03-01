package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
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

	resCh := make(chan []float64)
	errCh := make(chan error)
	doneCh := make(chan struct{})

	wg := sync.WaitGroup{}

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

	totalData := make([]float64, 0)

	for _, fname := range filenames {

		wg.Add(1)

		go func(fname string) {
			defer wg.Done()

			// Open file
			f, err := os.Open(fname)
			if err != nil {
				errCh <- fmt.Errorf("cannot open file: %w", err)
			}

			// Extract data
			data, err := csv2float(f, column)
			if err != nil {
				errCh <- err
			}

			// Close file
			err = f.Close()
			if err != nil {
				errCh <- err
			}

			resCh <- data
		}(fname)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		doneCh <- struct{}{}
		// close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case partialData := <-resCh:
			totalData = append(totalData, partialData...)
		case <-doneCh:
			_, err := fmt.Fprintln(out, operationFunc(totalData))
			return err
		}
	}
}

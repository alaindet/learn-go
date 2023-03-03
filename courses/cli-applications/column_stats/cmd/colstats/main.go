package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
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
	case "min":
		operationFunc = min
	case "max":
		operationFunc = max
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	totalData := make([]float64, 0)

	resCh := make(chan []float64)
	errCh := make(chan error)
	doneCh := make(chan struct{})
	filesCh := make(chan string)

	wg := sync.WaitGroup{}

	go func() {
		defer close(filesCh)
		for _, fname := range filenames {
			filesCh <- fname
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for fname := range filesCh {

				// Open file
				f, err := os.Open(fname)
				if err != nil {
					errCh <- fmt.Errorf("cannot open file: %w", err)
					return
				}

				// Extract data
				data, err := csv2float(f, column)
				if err != nil {
					errCh <- err
				}

				// Close file
				if err := f.Close(); err != nil {
					errCh <- err
				}

				resCh <- data
			}
		}()
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(doneCh)
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

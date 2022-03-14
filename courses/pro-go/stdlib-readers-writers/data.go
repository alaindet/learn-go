package main

import "io"

func GenerateData(writer io.Writer) {
	data := []byte("Kayak, Lifejacket")
	writeSize := 4

	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize

		if end > len(data) {
			end = len(data)
		}

		count, err := writer.Write(data[i:end])
		p("Wrote %v byte(s): %v", count, string(data[i:end]))

		if err != nil {
			p("Error: %v", err.Error())
		}
	}

	if closer, ok := writer.(io.Closer); ok {
		closer.Close()
	}
}

func ConsumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)

	for {
		count, err := reader.Read(slice)

		if err == io.EOF {
			break
		}

		if count > 0 {
			p("Read data: %v", string(slice[0:count]))
			data = append(data, slice[0:count]...)
		}
	}

	p("Read full data: %v", string(data))
}

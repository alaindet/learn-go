# CSV Files Generator

This is a simple random CSV files generator for producing a large volume of `.csv` like this with an arbitrary number of lines

```csv
Col1,Col2
Data0,60707
Data1,25641
Data2,79731
Data3,18485
```

This is used for benchmarking the `column_stats` CLI

## Usage
```
go run ./cmd/csvgen -lines 20 -files 20 -dir ./data
```

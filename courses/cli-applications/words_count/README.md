# Words counter

## Build

```
go build .
```

## Help
```
./first_program -h
```

## Count words

```
echo "The quick brown fox jumps over the lazy dog" | ./first_program
// Outputs: 9
```

## Count lines

```
cat pelican.txt | ./first_program -l
// Outputs: 9
```

## Count bytes
```
cat pelican.txt | ./first_program -b
// Outputs: 162
```

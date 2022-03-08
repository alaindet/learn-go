# Math

The `math` package includes methods to perform mathematical calculations

## Functions
Most common functions are

- `math.Abs(val)`
- `math.Ceil(val)`
- `math.Copysign(x, y)`
- `math.Floor(val)`
- `math.Max(x, y)`
- `math.Min(x, y)`
- `math.Mod(x, y)`
- `math.Pow(x, y)`
- `math.Round(val)`
- `math.RoundToEven(val)`

## Constants

Some constants
- `math.MaxInt8`
- `math.MinInt8`
- `math.MaxInt16`
- `math.MinInt16`
- `math.MaxInt32`
- `math.MinInt32`
- `math.MaxInt64`
- `math.MinInt64`
- `math.MaxUint8`
- `math.MaxUint16`
- `math.MaxUint32`
- `math.MaxUint64`
- `math.MaxFloat32`
- `math.MaxFloat64`
- `math.SmallestNonZeroFloat32`
- `math.SmallestNonZeroFloat64`

## Randomness via `math/rand`
NOTE: By default, random functions are given the same seed and return the **same values** for every run, unless you change the seed. This is intended. Random functions from `math/rand`
- `rand.Seed(s)`
- `rand.Float32()`
- `rand.Float64()`
- `rand.Int()`
- `rand.Intn(max)`
- `rand.UInt32()`
- `rand.UInt64()`
- `rand.Shuffle(count, func)`

# Sorting
Sorting is provided by the standard `sort` package

## Functions
- `sort.Float64s([]float64)`
- `sort.Float64sAreSorted([]float64)`
- `sort.Ints([]int)`
- `sort.IntsAreSorted([]int)`
- `sort.Strings([]string)`
- `sort.StringsAreSorted([]string)`

## Searching
NOTE: These functions expect input slices to be **sorted**
NOTE: If value is not found, these functions return the index *at which you can insert the searched value* in order to maintain sorting, which is very unusual

- `sort.SearchInts([]ints, int)`
- `sort.SearchFloat64s([]float64, float64)`
- `sort.SearchStrings([]string, string)`
- `sort.Search(count, func)`

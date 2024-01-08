# Concurrent slice mapping experiment

```console
go test -bench=. -run=^# -benchtime=4s
```

**TL;DR**: Use plain old non-concurrent `Map` for simple slice mapping as its 20% to 200% faster than any of its concurrent versions (`ConcMap` and `FasterConcMap`)

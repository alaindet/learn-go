# Folders structure

- It is paramount to enable a common ground for production applications
- A great explanation is [https://peter.bourgon.org/go-best-practices-2016/](https://peter.bourgon.org/go-best-practices-2016/)

## `internal`

- Contains non-application specific code that should not be used by other applications
- By convention, only code living into `internal` folder's parent can import code from `internal`

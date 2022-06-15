# Folders structure

- It is paramount to enable a common ground for production applications
- A great explanation is [https://peter.bourgon.org/go-best-practices-2016/](https://peter.bourgon.org/go-best-practices-2016/)

## `internal`

This is a special folder for any Go project, since any package living inside `internal` can **only be imported by the parent of `internal`**, which ensures that the `internal` folder holds all the project-specific logic and no other project can access it (including third-party)

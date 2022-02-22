# Packages

- A package is a folder containing many `.go` files having the same **package statement** at the beginning
- Types of packages are
  - **executable packages**
    - generates executables files
    - predefined package name must be `main`
  - **non-executable packages** (libraries, dependencies)
    - used by other libraries or applications
    - can have any package name
    - can't be executed, only imported

# Install Hugo

- Run `go install -tags extended github.com/gohugoio/hugo@latest`
- Export your `bin/` Go folder
- Run `hugo version`, it should print
  ```
  hugo v0.109.0+extended linux/amd64 BuildDate=unknown
  ```

## Create new project
`hugo new site your_website`

## Run in development mode (watches for changes)
```
cd ./your_website
hugo server
```

## Create new page
`hugo new portfolio.md`

## Build
`hugo --cleanDestinationDir --minify`

## Create a new theme
`hugo new theme obsidian`

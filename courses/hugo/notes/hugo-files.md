# Hugo Files and Folders

## `archetypes`
- Markdown templates modeling (grouping) content of the same type
- You can create new content starting from an archetype to save time

## `config.toml`
- Contains configuration

## `content`
- Contains real content in Markdown or HTML, can have many arbitrary subdirectories based on content
- The `_index.md` file is the default content for the homepage

## `data`
- Contains yaml, json and toml files (TODO?)

## `layouts`
- Contains the actual skeletons used by your website
- The `index.html` should be the home page layout

## `static`
- Contains CSS, JS, images and assets in general

## `themes`
- Contains both 3rd-party themes and your own
- Essentially, a theme folder contains common files for any Hugo project used as a base
- A theme folder may contain an `archetypes`, a `layouts` and a `static` folder for CSS and JS (or for CSS and JS pipe outputs)
- Projects using a theme then can override theme files

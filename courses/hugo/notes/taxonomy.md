# Taxonomy

- While folder and file hierarchies define the routing structure of a Hugo website, **taxonomies** are arbitrary groups of content, managed by the user via configuration keys in the front matter of content files
- The two main taxonomies are called **Categories** and **Tags** are their meaning is completely interchangeable based on your judgement (TODO?)

An example post on a blog could have these categories and tags

```md
---
title: "First Post"
author: John Smith
categories:
- Personal
- Thoughts
tags:
- software
- html
---

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

<!-- more -->

Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat.
```

- Categories and tags are automatically exposed as list pages, for example after creating the post above you can visit `/categories` (`/tags`) with the list of categories (or tags) and then visit `/categories/personal` (or `/tags/software`) with another list of posts with the `Personal` category (or with the `software` tag)

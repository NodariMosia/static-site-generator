# Static Site Generator (SSG) built in Go

**Generate static `HTML` pages from `Markdown` files.**

This project is inspired by [boot.dev](https://www.boot.dev)'s [Build a Static Site Generator](https://www.boot.dev/courses/build-static-site-generator) course that builds the SSG in Python.
Currently, files in `./content` directory, `./static/images/rivendell.png` and `./static/index.css` files are copied from that course.

## How to use

- Place static files (images, icons, css, js, etc.) in `./static` directory. They will get copied to `./public` directory with the same directory structure.
- Link all static files (css, js, favicon, site manifest, etc.) in `template.html`.
- Place website page contents in `Markdown` file format inside `./content` directory. They will be used to generate `HTML` pages that will be added to the `./public` directory with the same directory structure.
- If you want to use different directories or different template file to generate pages, update `staticDir`, `contentDir`, `destinationDir` and `templatePath` variables in [./cmd/main.go](./cmd/main.go) file.
- To generate pages and serve them to localhost, run from project's root:

  ```bash
  ./main.sh
  # or
  make run
  # or
  go run ./cmd
  ```

## How it works

- Project copies all files and subdirectories from `./static` directory to `./public` directory.
- Reads all `Markdown` files from `./content` directory and for each of them:

  1. Extracts page title from the first heading tag (*`# some title`*) encountered in the file;
  2. Replaces title placeholder ( ***`{{ Title }}`*** ) in template with extracted title;
  3. Creates root `<div>` `HTMLNode` node to build HTML node tree;
  4. Splits markdown file contents into following blocks:
     - paragraph
     - heading
     - code
     - quote
     - unordered_list
     - ordered_list
  5. Creates associated `HTMLNode` for each block and adds it to the node tree's root node's children:

     | Block          | HTMLNode                                    |
     | -------------- | ------------------------------------------- |
     | paragraph      | `<p>` node                                  |
     | heading        | `<h1>` , ... , `<h6>` node                  |
     | code           | `<code>` node contained inside `<pre>` node |
     | quote          | `<blockquote>` node                         |
     | unordered_list | `<li>` nodes contained inside `<ul>` node   |
     | ordered_list   | `<li>` nodes contained inside `<ol>` node   |

  6. Splits each block's text (except `code` block) into following inline text nodes:
     - text
     - bold
     - italic
     - code
     - link
     - image
  7. Creates associated `HTMLNode`s for each `TextNode` and adds them to the node tree under their parent node:

     | TextNode | HTMLNode         |
     | -------- | ---------------- |
     | text     | inline text node |
     | bold     | `<b>` node       |
     | italic   | `<i>` node       |
     | code     | `<code>` node    |
     | link     | `<a>` node       |
     | image    | `<img>` node     |

  8. Generates `HTML` string by parsing previously created `HTMLNode` node tree;
  9. Replaces content placeholder ( ***`{{ Content }}`*** ) in template with generated html code;
  10. Writes final `HTML` string to the file in `./public` directory with the same name and directory structure as its source file.

# Go Web Barebone - skeleton

This repository aims to provide a project bootstrap, installing all the proposed
dependencies and a files structure. It will also sandbox the GOPATH and GOBIN
variables (using a .env file), so that different projects may not collide with
particular needs (different dependency versions etc).

A dummy main.go file is included in the

It uses `dip` to manage dependencies.

## Requirements

`go` installed on the system

## Bootstrap process

Just run `./setup.sh`. You can then decide either to delete the bootstrap files,
or keep it changing the `.gitignore` file in order to include your files,
changing the git's `origin` remote (or first delete the `.git` folder and then
`git init` again).

I usually add this line to `.gitignore` to include all but the dependencies:

`/src/*.*`

Meaning to exclude all the files and folders containing a dot in their name in
the src's root folder.

## Details

### Dependencies
All the dependencies are in the `bootstrap/dependencies` files. You can easily
change them by simply writing a different content in it. The setup file will
automatically download them and add to the dip's template file.

#### Routing
Vestigo is the routing choice. It's fairly simple and (according to the develoer)
it's the fastes router handling the standard `http.HandlerFunc` without using
custom contextes (making it fairly easy to swap in the future).

https://github.com/husobee/vestigo

#### Logging
logrus' logging framework is both simple, colorful and powerful. You can of course
use your own one as you like.

https://github.com/sirupsen/logrus

#### Testing
The best asserting library I've found so far: stretchr's testify.

https://github.com/stretchr/testify

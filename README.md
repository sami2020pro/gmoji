# gmoji 😎

<div>
  <img
      src="https://img.shields.io/badge/Build-Passing-blue"
      alt="gmoji | version 0.1.4-3 | gmoji library"
      style="max-width:100%;"
  />
  <a href="">
    <img
      src="https://img.shields.io/badge/Version-0.1.4--3-brightgreen"
      alt="gmoji | version 0.1.4-3 | gmoji library"
      style="max-width:100%;"
    />
  </a>
  <a href="https://golang.org/">
    <img
      src="https://img.shields.io/badge/Language-Go-brightgreen"
      alt="gmoji | language go | go library"
      style="max-width:100%;"
    />
  </a>
  <img
    src="https://img.shields.io/badge/Go Version-1.16-lightgrey"
    alt="gmoji | go version | go library"
    style="max-width:100%;"
  />
  <a href="https://github.com/sami2020pro/gmoji/blob/master/LICENSE">
    <img 
      src="https://img.shields.io/badge/License-MIT-lightgrey"
      alt="gmoji | MIT License | gmoji library"
      style="max-width:100%;"
    />
  </a>
</div>

[![Watch the video](data/gmoji-preview-background.jpeg)](data/gmoji-preview.mp4)

`gmoji` is a emoji library for Go. It lets you use emoji characters in strings and you can use in **Terminal**.

# Install 🤠

In `Go 1.16` to install if you have the following problem

```bash
no required module provides package github.com/sami2020pro/gmoji: working directory is not part of a module
```

you must use the following command

```bash
go env -w GO111MODULE=auto
go get -u github.com/sami2020pro/gmoji
```

if you don't have a problem, use the following command as before

```bash
go get -u github.com/sami2020pro/gmoji
```

> Note: if you want to read more about Go 1.16 modules click <a href="https://blog.golang.org/go116-module-changes">here</a>

# Usage ✊
```golang
package main

import (
        "fmt"

        "github.com/sami2020pro/gmoji"
)

func main() {
        fmt.Printf("Hello %v\n", gmoji.Fire)

        emoji := gmoji.GitHub
        fmt.Printf("I love %v\n", emoji)

        fmt.Println(gmoji.Parse("Emoji aliases are :Plus:"))
        gmoji.Println("Use fmt wrappers :Eye:")
        gmoji.Print("Golang is :HeartOutlined:\n")
}
```

Output

<div>
  <a href="data/gmoji-output.png">
    <img 
      src="data/gmoji-output.png"
      alt="gmoji usage output"
      style="max-width:100%;"
    />
  </a>
</div>

<b>fontawesome 1.0 to 3.0 is available</b>

<!--# Testing 🍷
```golang
go test
```-->
# Contributing 💻
You can send **Pull Request**

# References 📃
<ul>
  <li><a href="https://www.fontawesomecheatsheet.com/" style="text-decoration:none;">fontawesomecheatsheet</a></li>
</ul>

# Credits ⭐
<ul>
  <li><a href="https://www.github.com/sami2020pro" style="text-decoration:none;">Sami Ghasemi</a></li>
</ul>

# License 📜
The MIT License (MIT). Please see <a href="https://github.com/sami2020pro/suftime/blob/master/LICENSE">License File</a> for more information.

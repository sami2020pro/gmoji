# gmoji 😎

<div>
  <a href="https://github.com/sami2020pro/suftime/blob/main/suftime/product/version.txt">
    <img
    src="https://img.shields.io/badge/version-0.2.0-brightgreen"
    alt="coverage 100%"
    style="max-width:100%;"
    />
  </a>
  <a href="https://golang.org/">
    <img
    src="https://img.shields.io/badge/language-golang-brightgreen"
    alt="coverage 100%"
    style="max-width:100%;"
    />
  </a>
  <a href="https://github.com/sami2020pro/suftime/blob/master/LICENSE">
    <img 
    src="https://img.shields.io/badge/License-GPLv3-brightgreen"
    alt="GPLv3 License"
    style="max-width:100%;"
    />
  </a>
</div>

`gmoji` is a emoji library for Go. It lets you use emoji characters in strings and you can use in **Terminal**.

# Install 🤠

```html
go get -u github.com/sami2020pro/gmoji
```

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

/* OUTPUT

    Hello 
    I love 
    Emoji aliases are 
    Use fmt wrappers 
    Golang is 
*/

```

# Testing 🍷
```golang
go test
```

# License 📜
The MIT License (MIT). Please see <a href="https://github.com/sami2020pro/suftime/blob/master/LICENSE">License File</a> for more information.

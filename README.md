# gmoji 😎

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
The MIT License (MIT). Please see License File for more information.

# gmoji 😎

```html
Not Completed
```

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
```

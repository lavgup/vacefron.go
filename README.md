# vacefron.go
A fast wrapper for the [VAC Efron API](https://vacefron.nl/api).

## Example usage
### Simple
```go
package main

import (
    "fmt"
    vacefron "github.com/lavgup/vacefron.go"
)

func main() {    
    body, err := vacefron.Water("butch and sundance are back, baby!")
    if err != nil {
        fmt.Println(err)
    }
    
    // WARNING: This will spam your console
    fmt.Printf("Received bytes: %s", body)
}
```

### Image
```go
package main

import (
    "fmt"
    "net/url"
    vacefron "github.com/lavgup/vacefron.go"
)

func main() {    
    imageUrl := "https://cdn.discordapp.com/avatars/441164156016787486/7a9cc8980bed842503c451efc79b74f7.png"
    parsedUrl, err := url.Parse(imageUrl)
    if err != nil {
        fmt.Println(err)
    }

    body, err := vacefron.Drip(parsedUrl)
    if err != nil {
        fmt.Println(err)
    }
    
    // WARNING: This will spam your console
    fmt.Printf("Received bytes: %s", body)
}
```

The image returns are of type []byte, so to send it via discordgo, wrap it with a Buffer e.g session.ChannelFileSend(id, "image.png", bytes.NewBuffer(image)).

Endpoints which return JSON are converted to maps.
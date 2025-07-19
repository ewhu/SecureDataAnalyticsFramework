// cmd/securedataanalyticsframework/main.go
package main

import (
"flag"
"log"
"os"

"securedataanalyticsframework/internal/securedataanalyticsframework"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := securedataanalyticsframework.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}

# restwell
> Create simple rest servers from json files. Profit.

## Install
```
$ go get github.com/ru-lai/restwell
```

## Usage
```go
import (
    rw "github.com/ru-lai/restwell"
    "net/http"
    "log"
)

const jsonStr = `{"routes":[{"route":"cookies","payload":"chocolateChip"},{"route":"snacks","payload":"{\"cookies\":\"vanilla\",\"cupcakeTypes\":[\"happiness\",\"chocolateChip\"]}"}]}`

const addr = "localhost:12345"

func main() {
    mux := rw.CreateMuxFromJSON()
    s := http.Server{Handler: mux, Addr: addr}
    log.Fatal(s.ListenAndServe())
}
```

## API

### CreateMuxFromJSON(string)

Returns a pointer to a [http.ServeMux]() struct. This acts as a multiplexer that matches the url you've fed in the route and calls a pre-created handler to return your response.

*NOTE* For the handlers to map properly, the json string must be in the following format:
```json
{
    "routes": [
        {
            "route": "routeName",
            "payload": "{
                \"myJSONPayloadKey\": \"myJSONPayload\"
            }"
        },
        {
            "route": "snacks",
            "payload": "{\"cookies\":\"vanilla\",\"cupcakeTypes\":[\"happiness\",\"chocolateChip\"]}"
        }
    ]
}
```

The resulting server will have two routes:
1.  one with the path "/routeName"
1.  and another with the path "/snacks"

They will both return their payload strings as a JSON object.

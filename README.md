# uuidapikey

A Base32-Crockford encoded API Key generator, validator and converter to turn UUIDs into human readable API Keys for Golang.

## Installation
You can install this with go get.
```
go get "github.com/siesgstarena/uuidapikey"
```

## Usage
### Create a pair of keys 
```go
package main

import (
	"fmt"

	"github.com/siesgstarena/uuidapikey"
)

func main() {

	fmt.Println(uuidapikey.Create())
}

```
**Output:**
```go
{ "apiKey":"38QARV0-1ET0G6Z-2CJD9VA-2ZZAR0X", 
"uuid":"d1756360-5da0-40df-9926-a76abff5601d" }
```

### Validate your UUID
```
uuidString := "d1756360-5da0-40df-9926-a76abff5601d"
fmt.Println(uuidapikey.IsUUID(uuidString)) // returns true
```

### Validate your API Key
```go
apiKeyString := "38QARV0-1ET0G6Z-2CJD9VA-2ZZAR0X"
fmt.Println(uuidapikey.IsAPIKey(apiKeyString)) // returns true
```

### Convert UUID into API Key
```go
uuidString := "d1756360-5da0-40df-9926-a76abff5601d"
fmt.Println(uuidapikey.ToAPIKey(uuidString))
```
**Output**:
```
38QARV0-1ET0G6Z-2CJD9VA-2ZZAR0X
```

### Convert API Key into UUID
```go
apiKeyString := "38QARV0-1ET0G6Z-2CJD9VA-2ZZAR0X"
fmt.Println(uuidapikey.ToUUID(apiKeyString))
```
**Output**:
```
d1756360-5da0-40df-9926-a76abff5601d
```

## Issues
Currently, Issue Tracking is done via GitHub Issues.

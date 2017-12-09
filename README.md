# timecode-go
Golang timecode Library

## Installing
```
go get -u github.com/tsujimic/timecode-go
```

## Using the timeocode-go
```
package main

import (
	"fmt"
	"github.com/tsujimic/timecode-go"
)

func main() {
	var v int
	var r string
	var err error
	v, err = timecode.ParseInt("00:58:30.10", timecode.SMPTE2997DROP)
	r = timecode.FormatString(v, timecode.SMPTE2997DROP)
	v, err = timecode.ParseDuration("00:58:30.10", "01:00:00.00", timecode.SMPTE2997DROP)
}

```


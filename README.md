# go-whosonfirst-sources

Go package for working with Who's On First data sources

## Example

### Simple

import (
	"github.com/whosonfirst/go-whosonfirst-sources"
       "log"
)

```
log.Println(sources.IsValidSource("sfac"))
log.Println(sources.IsValidSource("chairzen"))

log.Println(sources.IsValidSourceId(404734211))

src, err := sources.GetSourceByName("mapzen")

if err != nil {
   log.Fatal(err)
}

log.Println(src.License)
```

Yields:

```
true
false
true
CC0
```

## See also

* https://github.com/whosonfirst/whosonfirst-sources/

# imago

a tool box with go

## how to use

**File**

File.Read

```
package main

import (
	"fmt"
	"log"

	"github.com/tiantour/imago"
)

func main() {
	path := "you file path"
	body, err := imago.NewFile().Read(path)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(body))
}
```

File.Write

```
package main

import (
	"log"

	"github.com/tiantour/imago"
)

func main() {
	path := "your file path"
	data := "your file data"
	err := imago.NewFile().Write(path, []byte(data))
	if err != nil {
		log.Println(err.Error())
	}
}	
```

**Random**

Random.Number

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	number := imago.NewRandom().Number(9)
	fmt.Println(number)
}
```

Random.String

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	number := imago.NewRandom().Number(9)
	fmt.Println(number)
}
```

Random.ULID

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	str := imago.NewRandom().ULID()
	fmt.Println(str)
}
```

Text.IsDigit

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	text := "1234"
	result := imago.NewText().IsDigit(text)
	fmt.Println(result)
}
```

Text.IsHan

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	text := "汉字"
	result := imago.NewText().IsHan(text)
	fmt.Println(result)
}
```
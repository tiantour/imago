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
	body, err := imago.File.Read(path)
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
	err := imago.File.Write(path, []byte(data))
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
	number := imago.Random.Number(9)
	fmt.Println(number)
}
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
	number := imago.Random.Number(9)
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
	str := imago.Random.ULID()
	fmt.Println(str)
}
```

**Crypto**

Crypto.Base64Encode

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	args := "input data"
	result := imago.Crypto.Base64Encode([]byte(args))
	fmt.Println(result)
}
```

**Crypto.Base64Decode**

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
	"qiniupkg.com/x/log.v7"
)

func main() {
	args := "input data"
	result, err := imago.Crypto.Base64Decode(args)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(result))
}
```

**Crypto.MD532**

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	args := "input data"
	result := imago.Crypto.Md532(args)
	fmt.Println(result)
}
```

**Crypto.MD516**

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	args := "input data"
	result := imago.Crypto.Md516(args)
	fmt.Println(result)
}
```

**Crypto.SHA1**

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	args := "input data"
	result := imago.Crypto.SHA1(args)
	fmt.Println(result)
}
```

**Crypto.SHA256**

```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	args := "input data"
	result := imago.Crypto.SHA256(args)
	fmt.Println(result)
}
```

**Crypto.HmacSha1**


```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	publicKey := "public key"
	privateKey := "private key"
	result := imago.Crypto.HmacSha1(publicKey, privateKey)
	fmt.Println(result)
}
```

**Crypto.Pbkdf2Sha256**


```
package main

import (
	"fmt"

	"github.com/tiantour/imago"
)

func main() {
	data := "input date"
	salt := "input salt"
	iterations := 12000
	result := imago.Crypto.Pbkdf2Sha256(data, salt, iterations)
	fmt.Println(result)
}
```

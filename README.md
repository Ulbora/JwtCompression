[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/JwtCompression)](https://goreportcard.com/report/github.com/Ulbora/JwtCompression)

A tool for compression and decompression of large JWT tokens. Uses zlib for compression.
==============

```
Compress JWT

import(
     cp "github.com/Ulbora/JwtCompression"
)

var c cp.JwtCompress
cjwt := c.CompressJwt(jwt)

```

```
Decompression



import(
     cp "github.com/Ulbora/JwtCompression"
)

var c cp.JwtCompress
var jwt = c.UnCompressJwt(cjwt)

```
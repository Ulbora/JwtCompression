//Package compresstoken ...
package compresstoken

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"encoding/base64"
	"io"
)

//JwtCompress JwtCompress
type JwtCompress struct {
}

//CompressJwt CompressJwt
func (c *JwtCompress) CompressJwt(jwt string) string {
	//compress jwt with zlib and package with base64
	var rtn string
	var b bytes.Buffer
	w, err := zlib.NewWriterLevel(&b, flate.BestCompression)
	if err == nil {
		w.Write([]byte(jwt))
		w.Close()
		rtn = base64.StdEncoding.EncodeToString(b.Bytes())
	}
	return rtn
}

//UnCompressJwt UnCompressJwt
func (c *JwtCompress) UnCompressJwt(cjwt string) string {
	//uncompress jwt with zlib after converting from base64
	var rtn string
	var b bytes.Buffer
	decoded, derr := base64.StdEncoding.DecodeString(cjwt)
	if derr == nil {
		b.Write(decoded)
		r, err := zlib.NewReader(&b)
		if err == nil {
			var out bytes.Buffer
			io.Copy(&out, r)
			r.Close()
			rtn = out.String()
		}
	}
	return rtn
}

//go mod init github.com/Ulbora/JwtCompression

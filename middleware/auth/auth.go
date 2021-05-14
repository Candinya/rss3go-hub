/*********************************************************************

rss3go: An alternative version of RSSHub for RSS3 written in go

Copyright (C) 2021 Nyawork, Candinya

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

 ********************************************************************/

package auth

import (
	"crypto/md5"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Todo: test this.

func Auth() gin.HandlerFunc {
	return func (ctx * gin.Context) {

		personaId := ctx.Param("pid")
		reqSign := ctx.Request.Header.Get("sign")

		if personaId != "" && reqSign != "" {

			personaIdBytes := []byte(personaId)
			reqSignBytes := []byte(reqSign)

			raw, err := ctx.GetRawData()

			if err != nil {
				fmt.Println(err.Error())
			}

			dataProcessed := fmt.Sprintf("%s%s", ctx.FullPath(), raw)

			dataBytes := []byte(dataProcessed)
			md5hash := md5.Sum(dataBytes)

			fmt.Printf("%x", md5hash)

			verification := secp256k1.VerifySignature(personaIdBytes, md5hash[:], reqSignBytes)

			if !verification {
				ctx.String(http.StatusUnauthorized, "Unauthorized. Missing authentication parameters.")
			}
		} else {
			ctx.String(http.StatusUnauthorized, "Unauthorized. Missing authentication parameters.")
		}

		ctx.Next()
	}
}

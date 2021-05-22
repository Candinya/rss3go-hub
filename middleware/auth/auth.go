/*********************************************************************

rss3go_hub: An alternative version of RSSHub for RSS3 written in go

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
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/gin-gonic/gin"
	"github.com/nyawork/rss3go_lib/methods"
	"io/ioutil"
	"log"
	"net/http"
)

// Todo: test this.

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		raw, err := ioutil.ReadAll(ctx.Request.Body)

		if err != nil {
			ctx.Abort()
			log.Println(err.Error())
		}

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(raw))

		personaId := ctx.Param("pid")

		if personaId == "" {
			// May POST new persona

			persona := methods.Json2RSS3(raw)

			personaId = persona.Id
		}

		reqSign := ctx.Request.Header.Get("signature")

		if personaId != "" && reqSign != "" {

			personaIdBytes := []byte(personaId)
			reqSignBytes := []byte(reqSign)

			dataProcessed := fmt.Sprintf("%s%s", ctx.FullPath(), raw)

			dataBytes := []byte(dataProcessed)
			md5hash := md5.Sum(dataBytes)

			log.Printf("%x", md5hash)

			verification := secp256k1.VerifySignature(personaIdBytes, md5hash[:], reqSignBytes)

			if !verification {
				ctx.JSON(http.StatusForbidden, gin.H{
					"code":    http.StatusForbidden,
					"ok":      false,
					"message": "Unauthorized. Invalid signature.",
				})
				ctx.Abort()
				log.Println("Signature verify failed.")
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"ok":      false,
				"message": "Unauthorized. Missing authentication parameters.",
			})
			ctx.Abort()
			log.Println("No signature provided.")
		}
	}
}

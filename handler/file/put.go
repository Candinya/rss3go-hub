/*********************************************************************

rss3go_hub: An alternative version of RSS3-Hub written in go

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

package file

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nyawork/rss3go_lib/types"
	"net/http"
	"rss3go_hub/utils/storage"
	"strings"
)

type PutRequest struct {
	Contents []interface{} `json:"contents"`
}

func PutHandler(ctx *gin.Context) {

	var req PutRequest

	var files []string

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, "Parse failed. Error: " + err.Error(), http.StatusInternalServerError)
	} else {

		for _, content := range req.Contents {
			jsonByte, err := json.Marshal(content)
			if err != nil {
				handleError(ctx, "Marshal failed. Error: " + err.Error(), http.StatusInternalServerError)
			}
			var base types.RSS3 // Why not RSS3Base? Cause I cannot use
								// dynamic types (though they are actually the same one **for now**) in go.
								// Or just use `string`? That's not quite a good idea.
			if err = json.Unmarshal(jsonByte, &base); err != nil {
				handleError(ctx, "Re-marshal failed. Error: " + err.Error(), http.StatusInternalServerError)
			}
			if strings.Contains(base.Id, "item") {
				// I've chose the wrong type..
				// So just rebind this as nothing has happened.
				var item types.RSS3Item
				if err = json.Unmarshal(jsonByte, &item); err != nil {
					handleError(ctx, "Re-marshal failed. Error: " + err.Error(), http.StatusInternalServerError)
				}

				if success, err := item.CheckSign(); err != nil {
					handleError(ctx, "Signature check failed. Error: " + err.Error(), http.StatusInternalServerError)
				} else if !success {
					handleError(ctx, "Signature doesn't match.", http.StatusUnauthorized)
				} else {
					// Time to save this
					if writeData, err := json.Marshal(item); err != nil {
						handleError(ctx, "Marshall failed. Error: " + err.Error(), http.StatusInternalServerError)
					} else if err = storage.Write(item.Id, writeData); err != nil {
						handleError(ctx, "Save failed. Error: " + err.Error(), http.StatusInternalServerError)
					} else {
						files = append(files, item.Id)
					}
				}
			} else {
				// I'm right

				if success, err := base.CheckSign(); err != nil {
					handleError(ctx, "Signature check failed. Error: " + err.Error(), http.StatusInternalServerError)
				} else if !success {
					handleError(ctx, "Signature doesn't match.", http.StatusUnauthorized)
				} else {
					// Time to save this
					if writeData, err := json.Marshal(base); err != nil {
						handleError(ctx, "Marshall failed. Error: " + err.Error(), http.StatusInternalServerError)
					} else if err = storage.Write(base.Id, writeData); err != nil {
						handleError(ctx, "Save failed. Error: " + err.Error(), http.StatusInternalServerError)
					} else {
						files = append(files, base.Id)
					}
				}
			}
		}
	}

}

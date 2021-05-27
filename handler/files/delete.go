/*********************************************************************

rss3go-hub: An alternative version of RSS3-Hub written in go

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

package files

import (
	"github.com/gin-gonic/gin"
	"github.com/nyawork/rss3go-lib/utils"
	"net/http"
	"rss3go-hub/utils/sign"
	"rss3go-hub/utils/storage"
)

type DeleteRequest struct {
	Signature string `json:"signature"`
}

func DeleteHandler(ctx *gin.Context) {

	var req DeleteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, "Parse failed. Error: " + err.Error(), http.StatusInternalServerError)
	} else if req.Signature == "" {
		handleError(ctx, "Need signature", http.StatusUnauthorized)
	} else {

		signer, err := sign.GetSigner("delete", utils.FixSign(req.Signature))
		if err != nil {
			handleError(ctx, "Can't get signer. Error: " + err.Error(), http.StatusInternalServerError)
		}

		if exist, err := storage.Exist(signer); err != nil {
			handleError(ctx, "Can't check file existence. Error: " + err.Error(), http.StatusInternalServerError)
		} else if !exist {
			handleError(ctx, signer + " not found.", http.StatusNotFound)
		} else {
			if err := storage.Rm(signer); err != nil {
				handleError(ctx, "Can't delete file. Error: " + err.Error(), http.StatusInternalServerError)
			} else {
				handleSuccess(ctx, nil, "Requested file deleted.", http.StatusOK)
			}
		}
	}

}

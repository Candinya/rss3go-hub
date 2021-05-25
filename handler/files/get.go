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

package files

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rss3go_hub/utils/storage"
)

func GetHandler(ctx *gin.Context) {

	fileId := ctx.Param("fid")

	if exist, err := storage.Exist(fileId); err != nil {
		// Storage API error
		handleError(ctx, "Can't check file existence. Error: " + err.Error(), http.StatusInternalServerError)
	} else if !exist {
		// Doesn't exist
		handleError(ctx, fileId + " not found.", http.StatusNotFound)
	} else {
		// Exists
		if fileBytes, err := storage.Read(fileId); err != nil {
			// Storage API error
			handleError(ctx, "Can't read file. Error: " + err.Error(), http.StatusInternalServerError)
		} else {
			// No error
			ctx.Data(http.StatusOK, gin.MIMEJSON, fileBytes)
		}

	}

}

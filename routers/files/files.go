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
	"rss3go_hub/handler/files"
)

func Routers (e *gin.Engine) {

	apiFiles := e.Group("/files")
	{
		apiFiles.GET("/:fid", files.GetHandler)

		apiFiles.OPTIONS("")

		apiFiles.PUT("", files.PutHandler)

		apiFiles.DELETE("", files.DeleteHandler)
	}

}

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
	"github.com/gin-gonic/gin"
	"rss3go_hub/handler/file"
)

func Routers (e *gin.Engine) {

	apiFiles := e.Group("/file")
	{
		apiFiles.POST("/:fid", file.NewHandler)

		apiFiles.GET("/:fid", file.GetHandler)

		apiFiles.PUT("/:fid", file.NewHandler)

		apiFiles.DELETE("/:fid", file.DeleteHandler)
	}

}

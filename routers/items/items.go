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

package items

import (
	"github.com/gin-gonic/gin"
	"rss3go/api/item"
	"rss3go/middleware/auth"
)

func Routers (e * gin.Engine) {

	apiItems := e.Group("/personas/:pid/items")
	{

		apiItems.GET("", items.GetHandler)

		apiItems.POST("", auth.Auth(), items.NewHandler)

		apiItemsSpecify := apiItems.Group("/:tid")
		{

			apiItemsSpecify.PATCH("", auth.Auth(), items.ModifyHandler)

			apiItemsSpecify.DELETE("", auth.Auth(), items.DeleteHandler)

		}

	}
}

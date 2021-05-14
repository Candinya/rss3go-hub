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

package persona

import (
	"github.com/gin-gonic/gin"
	"rss3go/api/persona"
)

func Routers (e * gin.Engine) {

	apiPersona := e.Group("/personas")
	{

		apiPersona.POST("", persona.NewHandler)

		apiPersonaSpecify := apiPersona.Group("/:pid")
		{

			apiPersonaSpecify.GET("", persona.GetHandler)

			apiPersonaSpecify.PATCH("", persona.ModifyHandler)

			apiPersonaSpecify.DELETE("", persona.DeleteHandler)

		}

	}
}

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

package routers

import (
	"github.com/gin-gonic/gin"
	"rss3go-hub/middleware/cors"
)

type Option func(*gin.Engine)

var options []Option

// Include : Register routers
func Include (opts ... Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Allow())

	for _, opt := range options {
		opt(r)
	}

	return r
}

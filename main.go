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

package main

import (
	"fmt"
	"rss3go/routers"
	"rss3go/routers/item"
	"rss3go/routers/link"
	"rss3go/routers/persona"
)

func main() {
	routers.Include(routers.HomeRouter, persona.Routers, item.Routers, link.Routers)

	r := routers.Init()

	if err := r.Run(); err != nil {
		fmt.Printf("Service start failed, err: %v\n", err)
	}
}

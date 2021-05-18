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
	"github.com/gin-contrib/cors"
	"log"
	"rss3go/config"
	"rss3go/routers"
	"rss3go/routers/items"
	"rss3go/routers/link"
	"rss3go/routers/page"
	"rss3go/routers/persona"
)


func main() {

	// Load config

	log.Println("Loading config...")

	if err := config.LoadConfig("config.yml"); err != nil {
		panic(err)
	}

	log.Println("Config loaded successfully.")

	// Init routers

	log.Println("Initializing routers...")

	routers.Include(page.Routers, persona.Routers, items.Routers, link.Routers)

	r := routers.Init()

	log.Println("Routers initialized successfully.")

	r.Use(cors.Default())

	log.Println("Starting gin server...")

	//gin.SetMode(gin.ReleaseMode)

	if err := r.Run(); err != nil {
		panic(err)
	}

}

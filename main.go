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
	"github.com/gin-gonic/gin"
	"log"
	"rss3go_hub/config"
	"rss3go_hub/routers"
	"rss3go_hub/routers/file"
	"rss3go_hub/routers/page"
	"rss3go_hub/routers/persona"
)


func main() {

	// Load config

	log.Println("Loading config...")

	if err := config.LoadConfig("config.yml"); err != nil {
		panic(err)
	}

	log.Println("Config loaded successfully.")

	if config.GlobalConfig.Debug {
		log.Println("Working on debug mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init routers

	log.Println("Initializing routers...")

	routers.Include(page.Routers, persona.Routers, file.Routers)

	r := routers.Init()

	log.Println("Routers initialized successfully.")

	r.Use(cors.Default())

	log.Println("Starting gin server...")

	if err := r.Run(); err != nil {
		panic(err)
	}

}

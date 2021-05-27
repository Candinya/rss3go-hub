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

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rss3go-hub/config"
	"rss3go-hub/routers"
	"rss3go-hub/routers/files"
	"rss3go-hub/routers/page"
	"rss3go-hub/utils/storage"
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

	// Init storage

	log.Println("Initializing storage...")

	storage.Init()

	log.Println("Storage initialized successfully.")

	// Init routers

	log.Println("Initializing routers...")

	routers.Include(page.Routers, files.Routers)

	r := routers.Init()

	log.Println("Routers initialized successfully.")

	log.Println("Starting gin server...")

	if err := r.Run(); err != nil {
		panic(err)
	}

}

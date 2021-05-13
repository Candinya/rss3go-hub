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

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Home

	r.GET("/", func(c * gin.Context) {
		c.String(200, "Hello, RSS3Go")
	})

	// Persona

	apiPersona := r.Group("/persona")
	{
		apiPersona.HEAD("/get/:id", func(c * gin.Context) {
			// get the metadata of a persona
		})

		apiPersona.GET("/get/:id", func(c * gin.Context) {
			// get the full content of a persona
		})


		apiPersona.POST("/add", func(c * gin.Context) {
			// add a new persona
		})


		apiPersona.OPTIONS("/set", func(c * gin.Context) {
			// get what can be changed of a persona
		})

		apiPersona.PUT("/set", func(c * gin.Context) {
			// update the persona (full object)
		})

		apiPersona.PATCH("/set", func(c * gin.Context) {
			// update the persona (changes only)
		})


		apiPersona.DELETE("/del", func(c * gin.Context) {
			// delete the persona
		})
	}

	apiItem := r.Group("/items")
	{
		apiItem.HEAD("/get/:id", func(c * gin.Context) {
			// get the metadata of an item
		})

		apiItem.GET("/get/:id", func(c * gin.Context) {
			// get the full content of an item
		})


		apiItem.POST("/add", func(c * gin.Context) {
			// add a new item
		})


		apiItem.OPTIONS("/set", func(c * gin.Context) {
			// get what can be changed of an item
		})

		apiItem.PUT("/set", func(c * gin.Context) {
			// update the item (full object)
		})

		apiItem.PATCH("/set", func(c * gin.Context) {
			// update the item (changes only)
		})


		apiItem.DELETE("/del", func(c * gin.Context) {
			// delete the item
		})
	}

	apiLink := r.Group("/link")
	{
		apiLink.HEAD("/get/:id", func(c * gin.Context) {
			// get the metadata of a link
		})

		apiLink.GET("/get/:id", func(c * gin.Context) {
			// get the full content of a link
		})


		apiLink.POST("/add", func(c * gin.Context) {
			// add a new link
		})


		apiLink.OPTIONS("/set", func(c * gin.Context) {
			// get what can be changed of an link
		})

		apiLink.PUT("/set", func(c * gin.Context) {
			// update the link (full object)
		})

		apiLink.PATCH("/set", func(c * gin.Context) {
			// update the link (changes only)
		})


		apiLink.DELETE("/del", func(c * gin.Context) {
			// delete the link
		})
	}

	r.Run()
}

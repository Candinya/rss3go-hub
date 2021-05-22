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

package links

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nyawork/rss3go_lib/methods"
	"github.com/nyawork/rss3go_lib/types"
	"net/http"
	"rss3go_hub/tools"
	"rss3go_hub/utils/storage"
)

// todo: test this

func NewHandler(ctx *gin.Context) {

	personaId := ctx.Param("pid")

	if exist, _ := storage.Exist(personaId); !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this persona doesn't exist",
		})
		return
	}

	var link types.RSS3PersonaLink

	_ = ctx.BindJSON(&link) // Ignore error

	raw, _ := storage.Read(personaId) // Ignore error
	personaStruct := methods.Json2RSS3(raw)

	var exist bool = false

	for _, i := range personaStruct.Links {
		if i.Id == link.Id {
			exist = true
			break
		}
	}

	if exist {
		// Already exists
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"ok":      false,
			"message": "Sorry, but this link already exists",
		})
	} else {
		// Doesn't exist

		// Append to persona file
		personaStruct.Links = append(personaStruct.Links, link)

		// Save
		if err := storage.Write(personaId, personaStruct.ToJson()); err != nil {
			// Save error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"ok":      false,
				"message": err.Error(),
			})
		} else {
			// No error
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"ok":      true,
				"message": "Link saved",
				"data":    link.ToJson(),
			})
		}

	}

}

func GetHandler(ctx *gin.Context) {

	personaId := ctx.Param("pid")
	linkId := ctx.Param("lid")

	if exist, _ := storage.Exist(personaId); !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this persona doesn't exist",
		})
		return
	}

	var link_p *types.RSS3PersonaLink = nil

	raw, _ := storage.Read(personaId) // Ignore error
	personaStruct := methods.Json2RSS3(raw)

	if linkId == "" {
		// return all
		all, _ := json.Marshal(personaStruct.Links)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"ok":      true,
			"message": "Item patched",
			"data":    string(all),
		})
	} else {

		for _, i := range personaStruct.Links {
			if i.Id == linkId {
				link_p = &i
				break
			}
		}

		if link_p == nil {
			// Already exists error
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"ok":      false,
				"message": "Sorry, but this link doesn't exist",
			})
		} else {
			// Exists
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"ok":      true,
				"message": "Persona found",
				"data":    link_p.ToJson(),
			})
		}

	}

}

func ModifyHandler(ctx *gin.Context) {

	personaId := ctx.Param("pid")
	linkId := ctx.Param("lid")

	if exist, _ := storage.Exist(personaId); !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this persona doesn't exist",
		})
		return
	}

	var link_p *types.RSS3PersonaLink = nil

	raw, _ := storage.Read(personaId) // Ignore error
	personaStruct := methods.Json2RSS3(raw)

	for _, i := range personaStruct.Links {
		if i.Id == linkId {
			link_p = &i
			break
		}
	}

	if link_p == nil {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this link doesn't exist",
		})
	} else {
		// Exists

		var patch interface{}

		if err := ctx.BindJSON(&patch); err != nil {
			// Parse error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"ok":      false,
				"message": err.Error(),
			})
		} else {
			// Patch parsed

			if err := tools.DeepMergePersonaLink(link_p, patch); err != nil {
				// Deep Merge Error
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"ok":      false,
					"message": err.Error(),
				})
			} else {

				// Save persona
				if err := storage.Write(personaId, personaStruct.ToJson()); err != nil {
					// Storage API error
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"code":    http.StatusInternalServerError,
						"ok":      false,
						"message": err.Error(),
					})
				} else {
					// No error
					ctx.JSON(http.StatusOK, gin.H{
						"code":    http.StatusOK,
						"ok":      true,
						"message": "Link patched",
						"data":    link_p.ToJson(),
					})
				}
			}

		}
	}

}

func DeleteHandler(ctx *gin.Context) {

	personaId := ctx.Param("pid")
	linkId := ctx.Param("lid")

	if exist, _ := storage.Exist(personaId); !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this persona doesn't exist",
		})
		return
	}

	var link_index int = -1

	raw, _ := storage.Read(personaId) // Ignore error
	personaStruct := methods.Json2RSS3(raw)

	for index, i := range personaStruct.Links {
		if i.Id == linkId {
			link_index = index
			break
		}
	}

	if link_index == -1 {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this link doesn't exist",
		})
	} else {
		// Exists

		personaStruct.Links = append(
			personaStruct.Links[:link_index],
			personaStruct.Links[link_index+1:]...,
		)

		// Save persona
		if err := storage.Write(personaId, personaStruct.ToJson()); err != nil {
			// Storage API error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"ok":      false,
				"message": err.Error(),
			})
		} else {
			// No error
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"ok":      true,
				"message": "Link deleted",
			})
		}
	}

}

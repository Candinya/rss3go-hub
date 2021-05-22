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
	"github.com/nyawork/rss3go_lib/methods"
	"github.com/nyawork/rss3go_lib/types"
	"net/http"
	"rss3go_hub/tools"
	"rss3go_hub/utils/storage"
	"time"
)

// todo: test this

func NewHandler(ctx *gin.Context) {

	var persona types.RSS3

	if err := ctx.BindJSON(&persona); err != nil {
		// Parse error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"ok":      false,
			"message": err.Error(),
		})
	} else if exist, err := storage.Exist(persona.Id); err != nil {
		// Storage API error
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"code":    http.StatusNotImplemented,
			"ok":      false,
			"message": err.Error(),
		})
	} else if exist {
		// Already exists error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"ok":      false,
			"message": "Sorry, but this persona already exists",
		})
	} else {
		// Doesn't exist

		// Update timestamps
		persona.DateCreated = time.Now().String()
		persona.DateUpdated = time.Now().String()

		// Save
		if err := storage.Write(persona.Id, persona.ToJson()); err != nil {
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
				"message": "Persona found",
				"data":    persona.ToJson(),
			})
		}

	}

}

func GetHandler(ctx *gin.Context) {

	personaId := ctx.Param("pid")

	if exist, err := storage.Exist(personaId); err != nil {
		// Storage API error
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"code":    http.StatusNotImplemented,
			"ok":      false,
			"message": err.Error(),
		})
	} else if !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this persona doesn't exist",
		})
	} else {
		// Exists
		if personaBytes, err := storage.Read(personaId); err != nil {
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
				"message": "Persona found",
				"data":    string(personaBytes),
			})
		}

	}
}

func ModifyHandler(ctx *gin.Context) {

	personaId := ctx.Param("pid")

	if exist, err := storage.Exist(personaId); err != nil {
		// Storage API error
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"code":    http.StatusNotImplemented,
			"ok":      false,
			"message": err.Error(),
		})
	} else if !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this persona doesn't exist",
		})
	} else {
		// Exists

		var patch interface{}

		if personaBytes, err := storage.Read(personaId); err != nil {
			// Storage API error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"ok":      false,
				"message": err.Error(),
			})
		} else if err := ctx.BindJSON(&patch); err != nil {
			// Parse error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"ok":      false,
				"message": err.Error(),
			})
		} else {
			// Patch parsed

			oldPersona := methods.Json2RSS3(personaBytes)

			newPersona, err := tools.DeepMergePersona(oldPersona, patch)

			if err != nil {
				// Deep Merge Error
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"ok":      false,
					"message": err.Error(),
				})
			} else {

				// Update timestamps
				newPersona.DateUpdated = time.Now().String()

				if err := storage.Write(personaId, newPersona.ToJson()); err != nil {
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
						"message": "Persona patched",
						"data":    newPersona.ToJson(),
					})
				}
			}

		}

	}

}

func DeleteHandler(ctx *gin.Context) {

	personaId := ctx.Param("pid")

	if exist, err := storage.Exist(personaId); err != nil {
		// Storage API error
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"code":    http.StatusNotImplemented,
			"ok":      false,
			"message": err.Error(),
		})
	} else if !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"ok":      false,
			"message": "Sorry, but this persona doesn't exist",
		})
	} else {
		// Exists
		nextPage := personaId
		// Delete all pages for the persona
		for nextPage != "" {
			// Recursively
			raw, _ := storage.Read(nextPage) // Ignore error
			content := methods.Json2RSS3Base(raw)

			// Delete
			_ = storage.Rm(nextPage)

			nextPage = content.ItemsNext
		}

		// No error
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"ok":      true,
			"message": "Persona deleted",
		})

	}

}

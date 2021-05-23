/*********************************************************************

rss3go_hub: An alternative version of RSS3-Hub written in go

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

package file

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"rss3go_hub/utils/storage"
)

func GetHandler (ctx *gin.Context) {

	fileId := ctx.Param("fid")

	if exist, err := storage.Exist(fileId); err != nil {
		// Storage API error
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"code": http.StatusNotImplemented,
			"ok": false,
			"message": err.Error(),
		})
	} else if !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"ok": false,
			"message": "Sorry, but this file doesn't exist",
		})
	} else {
		// Exists
		if fileBytes, err := storage.Read(fileId); err != nil {
			// Storage API error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"ok": false,
				"message": err.Error(),
			})
		} else {
			// No error
			ctx.Data(http.StatusOK, http.DetectContentType(fileBytes), fileBytes)
		}

	}

}

func NewHandler (ctx *gin.Context) {

	fileId := ctx.Param("fid")

	if exist, err := storage.Exist(fileId); err != nil {
		// Storage API error
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"code": http.StatusNotImplemented,
			"ok": false,
			"message": err.Error(),
		})
	} else if exist {
		// Already exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"ok": false,
			"message": "Sorry, but this file already exists",
		})
	} else {
		// Exists

		raw, err := ioutil.ReadAll(ctx.Request.Body)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"ok": false,
				"message": "Sorry, but we failed to read the request body.",
			})
		} else if err = storage.Write(fileId, raw); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"ok": false,
				"message": "Sorry, but we failed to save the target file.",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"ok": true,
				"message": "File has been saved.",
			})
		}

	}

}

func ReplaceHandler(ctx *gin.Context) {

	fileId := ctx.Param("fid")

	if exist, err := storage.Exist(fileId); err != nil {
		// Storage API error
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"code": http.StatusNotImplemented,
			"ok": false,
			"message": err.Error(),
		})
	} else if !exist {
		// Already exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"ok": false,
			"message": "Sorry, but this file doesn't exist",
		})
	} else {
		// Exists

		raw, err := ioutil.ReadAll(ctx.Request.Body)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"ok": false,
				"message": "Sorry, but we failed to read the request body.",
			})
		} else if err = storage.Write(fileId, raw); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"ok": false,
				"message": "Sorry, but we failed to save the target file.",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"ok": true,
				"message": "File has been saved.",
			})
		}

	}

}

func DeleteHandler (ctx *gin.Context) {

	fileId := ctx.Param("fid")

	if exist, err := storage.Exist(fileId); err != nil {
		// Storage API error
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"code": http.StatusNotImplemented,
			"ok": false,
			"message": err.Error(),
		})
	} else if !exist {
		// Doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"ok": false,
			"message": "Sorry, but this file doesn't exist",
		})
	} else {
		// Exists
		if err := storage.Rm(fileId); err != nil {
			// Storage API error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"ok": false,
				"message": err.Error(),
			})
		} else {
			// No error
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"ok": true,
				"message": "File removed",
			})
		}

	}

}

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

package storage

import (
	"github.com/nyawork/rss3go_lib/types"
	"log"
	"os"
	"rss3go_hub/config"
	"strings"
)

type TypeOfStorageUndefinedError struct {
	sType string
}

func (e *TypeOfStorageUndefinedError) Error() string {
	return "Storage type undefined: sType"
}

func Write(name types.Address, content []byte) error {

	if config.GlobalConfig.Storage.Type == "local" {

		// Change for filename save
		name = strings.ReplaceAll(name, ":", "_")

		err := os.WriteFile(config.GlobalConfig.Storage.Path+name, content, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		return err
	}
	return &TypeOfStorageUndefinedError{config.GlobalConfig.Storage.Type}
}

func Read(name types.Address) ([]byte, error) {

	if config.GlobalConfig.Storage.Type == "local" {

		// Change for filename save
		name = strings.ReplaceAll(name, ":", "_")

		data, err := os.ReadFile(config.GlobalConfig.Storage.Path + name)
		if err != nil {
			log.Fatalln(err)
		}
		return data, err
	}
	return nil, &TypeOfStorageUndefinedError{config.GlobalConfig.Storage.Type}
}

func Exist(name types.Address) (bool, error) {

	if config.GlobalConfig.Storage.Type == "local" {

		// Change for filename save
		name = strings.ReplaceAll(name, ":", "_")

		_, err := os.Stat(config.GlobalConfig.Storage.Path + name)
		fileExist := os.IsNotExist(err)
		if !fileExist && err != nil {
			log.Fatalln(err)
		}
		return fileExist, err
	}
	return false, &TypeOfStorageUndefinedError{config.GlobalConfig.Storage.Type}
}

func Rm(name types.Address) error {

	if config.GlobalConfig.Storage.Type == "local" {

		// Change for filename save
		name = strings.ReplaceAll(name, ":", "_")

		err := os.Remove(config.GlobalConfig.Storage.Path + name)
		if err != nil {
			log.Fatalln(err)
		}
		return err
	}
	return &TypeOfStorageUndefinedError{config.GlobalConfig.Storage.Type}
}

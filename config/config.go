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

package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Debug bool `yaml:"debug"`
	Storage struct {
		Type string `yaml:"type"`
		Path string `yaml:"path"`
	} `yaml:"storage"`
}

var GlobalConfig Config

func LoadConfig(filename string) error {

	data, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &GlobalConfig); err != nil {
		return err
	}

	return nil

}

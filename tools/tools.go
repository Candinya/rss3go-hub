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

package tools

import (
	"github.com/imdario/mergo"
	"rss3go/entity/types"
)

func DeepMergePersona (old types.RSS3Persona, patch interface{}) (types.RSS3Persona, error) {

	var patchPersona types.RSS3Persona

	if err := mergo.Map(&patchPersona, patch); err != nil {
		return old, err
	}
	err := mergo.Merge(&old, patchPersona, mergo.WithOverride)

	return old, err
}

func DeepMergeItems (old types.RSS3Items, patch interface{}) (types.RSS3Items, error) {

	var patchItems types.RSS3Items

	if err := mergo.Map(&patchItems, patch); err != nil {
		return old, err
	}
	err := mergo.Merge(&old, patchItems, mergo.WithOverride)

	return old, err
}

func DeepMergeLink (old types.RSS3Link, patch interface{}) (types.RSS3Link, error) {

	var patchLink types.RSS3Link

	if err := mergo.Map(&patchLink, patch); err != nil {
		return old, err
	}
	err := mergo.Merge(&old, patchLink, mergo.WithOverride)

	return old, err
}

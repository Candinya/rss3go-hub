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

package methods

import (
	"encoding/json"
	"rss3go/entity/types"
)

func Json2RSS3Persona(personaJson string) types.RSS3Persona {
	var persona types.RSS3Persona
	err := json.Unmarshal([]byte(personaJson), &persona)
	if err != nil {
		return types.RSS3Persona{}
	}
	return persona
}

//func (persona * types.RSS3Persona) toJson() []byte {
//	personaJson, _ := json.Marshal(&persona)
//	return personaJson
//}

func Json2RSS3Items(itemsJson string) types.RSS3Items {
	var items types.RSS3Items
	err := json.Unmarshal([]byte(itemsJson), &items)
	if err != nil {
		return types.RSS3Items{}
	}
	return items
}

//func (items * types.RSS3Items) toJson() []byte {
//	itemsJson, _ := json.Marshal(&items)
//	return itemsJson
//}

func Json2RSS3Link(linkJson string) types.RSS3Link {
	var link types.RSS3Link
	err := json.Unmarshal([]byte(linkJson), &link)
	if err != nil {
		return types.RSS3Link{}
	}
	return link
}

//func (link * types.RSS3Link) toJson() []byte {
//	linkJson, _ := json.Marshal(&link)
//	return linkJson
//}


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

package types

import "encoding/json"

func (base * RSS3Base) ToJson() []byte {
	baseJson, _ := json.Marshal(&base)
	return baseJson
}

func (persona * RSS3) ToJson() []byte {
	personaJson, _ := json.Marshal(&persona)
	return personaJson
}

func (items * RSS3Items) ToJson() []byte {
	itemsJson, _ := json.Marshal(&items)
	return itemsJson
}

func (list * RSS3List) ToJson() []byte {
	listJson, _ := json.Marshal(&list)
	return listJson
}

func (item * RSS3Item) ToJson() []byte {
	itemJson, _ := json.Marshal(&item)
	return itemJson
}

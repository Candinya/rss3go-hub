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

func Json2RSS3Base(baseJson []byte) types.RSS3Base {
	var base types.RSS3Base
	_ = json.Unmarshal(baseJson, &base)

	return base
}

func Json2RSS3(baseJson []byte) types.RSS3 {
	var persona types.RSS3
	_ = json.Unmarshal(baseJson, &persona)

	return persona
}

func Json2RSS3Items(itemsJson []byte) types.RSS3Items {
	var items types.RSS3Items
	_ = json.Unmarshal(itemsJson, &items)

	return items
}

func Json2RSS3List(linkJson []byte) types.RSS3List {
	var list types.RSS3List
	_ = json.Unmarshal(linkJson, &list)

	return list
}

func Json2RSS3Item(itemsJson []byte) types.RSS3Item {
	var item types.RSS3Item
	_ = json.Unmarshal(itemsJson, &item)

	return item
}


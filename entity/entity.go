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

package entity

type Address = string

// Common attributes for each file
type RSS3Base struct {
	id Address
	version string
	_type string
	date_created string
	date_updated string
	editors struct {
		blocklist []Address
		allowlist []Address
	}
	items []interface{}
	items_next Address

}

// Entrance, RSS3Persona file
type RSS3Persona struct {
	RSS3Base

	// _type extends, default to 'persona'
	// editors extends, default to nil

	profile struct {
		name string
		avatar Address
		bio string
		tag []string
	}

	links []struct {
		id Address
		name string
		tags []string
	}

	items []RSS3Item
	assets interface{}
}

func NewRSS3Persona() RSS3Persona {
	return RSS3Persona {
		RSS3Base: RSS3Base {
			id:          "",
			version:     "",
			_type:       "persona",
			date_created: "",
			date_updated: "",
			editors:    nil,
			items:      nil,
			items_next: "",
		},
		profile: struct {
			name   string
			avatar Address
			bio    string
			tag    []string
		}{},
		links:  []struct {
			id Address
			name string
			tags []string
		}{},
		items:  []RSS3Item{},
		assets: nil,
	}
}

// RSS3Items file
type RSS3Items struct {
	RSS3Base

	// _type extends, default to 'items'

	items []RSS3Item
}

func NewRSS3Items() RSS3Items {
	return RSS3Items {
		RSS3Base: RSS3Base {
			id:          "",
			version:     "",
			_type:       "items",
			date_created: "",
			date_updated: "",
			editors: struct {
				blocklist []Address
				allowlist []Address
			}{},
			items:      nil,
			items_next: "",
		},
		items:    []RSS3Item{},
	}
}

// RSS3Link file
type RSS3Link struct {
	RSS3Base

	// _type extends, default to 'relationship'

	items []RSS3OtherPersona

}

func NewRSS3Link() RSS3Link {
	return RSS3Link {
		RSS3Base: RSS3Base{
			id:          "",
			version:     "",
			_type:       "relationship",
			date_created: "",
			date_updated: "",
			editors: struct {
				blocklist []Address
				allowlist []Address
			}{},
			items:      nil,
			items_next: nil,
		},
		items:    []RSS3OtherPersona{},
	}
}

type RSS3OtherPersona struct {
	id Address
	verification string
	name string
	avatar Address
	bio string
}

type RSS3Item struct {
	id string
	authors []RSS3OtherPersona
	title string
	summary string
	tags []string
	date_published string
	date_modified string

	contents []struct {
		id Address // Link to a third party file
		mimeType string
		name string
		tags []string
		size_in_bytes string
		duration_in_seconds string
	}

	contexts []struct {
		id Address // Link to a RSS3Items file
		name string
		tags []string
	}

}

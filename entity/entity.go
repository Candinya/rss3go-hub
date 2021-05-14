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
	Id          Address `json:"id"`
	Version     string  `json:"version"`
	Type        string  `json:"type"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
	Editors     struct {
		Blocklist []Address `json:"blocklist"`
		Allowlist []Address `json:"allowlist"`
	} `json:"editors"`
	Items     []interface{} `json:"items"`
	ItemsNext Address       `json:"items_next"`

}

// Entrance, RSS3Persona file
type RSS3Persona struct {
	RSS3Base

	// _type extends, default to 'persona'
	// editors extends, default to nil

	Profile struct {
		Name   string   `json:"name"`
		Avatar Address  `json:"avatar"`
		Bio    string   `json:"bio"`
		Tag    []string `json:"tag"`
	} `json:"profile"`

	Links []struct {
		Id   Address  `json:"id"`
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	} `json:"links"`

	Items  []RSS3Item  `json:"items"`
	Assets interface{} `json:"assets"`
}

//func NewRSS3Persona() RSS3Persona {
//	return RSS3Persona {
//		RSS3Base: RSS3Base {
//			Id:          "",
//			Version:     "",
//			Type:       "persona",
//			DateCreated: "",
//			DateUpdated: "",
//			Editors:     nil,
//			Items:       nil,
//			ItemsNext:   "",
//		},
//		Profile: struct {
//			Name   string
//			Avatar Address
//			Bio    string
//			Tag    []string
//		}{},
//		Links:  []struct {
//			Id   Address
//			Name string
//			Tags []string
//		}{},
//		Items:  []RSS3Item{},
//		Assets: nil,
//	}
//}

// RSS3Items file
type RSS3Items struct {
	RSS3Base

	// _type extends, default to 'items'

	Items []RSS3Item `json:"items"`
}

//func NewRSS3Items() RSS3Items {
//	return RSS3Items {
//		RSS3Base: RSS3Base {
//			Id:          "",
//			Version:     "",
//			Type:       "items",
//			DateCreated: "",
//			DateUpdated: "",
//			Editors: struct {
//				Blocklist []Address
//				Allowlist []Address
//			}{},
//			Items:     nil,
//			ItemsNext: "",
//		},
//		Items: []RSS3Item{},
//	}
//}

// RSS3Link file
type RSS3Link struct {
	RSS3Base

	// _type extends, default to 'relationship'

	Items []RSS3OtherPersona `json:"items"`

}

//func NewRSS3Link() RSS3Link {
//	return RSS3Link {
//		RSS3Base: RSS3Base{
//			Id:          "",
//			Version:     "",
//			Type:        "relationship",
//			DateCreated: "",
//			DateUpdated: "",
//			Editors: struct {
//				Blocklist []Address
//				Allowlist []Address
//			}{},
//			Items:     nil,
//			ItemsNext: nil,
//		},
//		Items: []RSS3OtherPersona{},
//	}
//}

type RSS3OtherPersona struct {
	Id           Address `json:"id"`
	Verification string  `json:"verification"`
	Name         string  `json:"name"`
	Avatar       Address `json:"avatar"`
	Bio          string  `json:"bio"`
}

type RSS3Item struct {
	Id            string             `json:"id"`
	Authors       []RSS3OtherPersona `json:"authors"`
	Title         string             `json:"title"`
	Summary       string             `json:"summary"`
	Tags          []string           `json:"tags"`
	DatePublished string             `json:"date_published"`
	DateModified  string             `json:"date_modified"`

	contents []struct {
		Id                Address  `json:"id"` // Link to a third party file
		MimeType          string   `json:"mime_type"`
		Name              string   `json:"name"`
		Tags              []string `json:"tags"`
		SizeInBytes       string   `json:"size_in_bytes"`
		DurationInSeconds string   `json:"duration_in_seconds"`
	}

	contexts []struct {
		Id   Address  `json:"id"` // Link to a RSS3Items file
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}

}

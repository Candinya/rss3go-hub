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

type Address = string

// RSS3Base Common attributes for each file
type RSS3Base struct {
	Id          Address `json:"id"`
	Version     string  `json:"version"`
	Type        string  `json:"type"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
	Editors     *struct {
		Blocklist []Address `json:"blocklist,omitempty"`
		Allowlist []Address `json:"allowlist,omitempty"`
	} `json:"editors,omitempty"`
	//Items     []interface{} `json:"items"`
	ItemsNext Address       `json:"items_next,omitempty"`

}

// RSS3Persona file, Entrance
type RSS3Persona struct {
	*RSS3Base

	// Type extends, default to 'persona'
	// Editors extends, default to nil

	Profile struct {
		Name   string   `json:"name,omitempty"`
		Avatar Address  `json:"avatar,omitempty"`
		Bio    string   `json:"bio,omitempty"`
		Tags   []string `json:"tags,omitempty"`
	} `json:"profile"`

	Links []struct {
		Id   Address  `json:"id"`
		Name string   `json:"name"`
		Tags []string `json:"tags,omitempty"`
	} `json:"links,omitempty"`

	Items  []RSS3Item  `json:"items"`
	Assets interface{} `json:"assets,omitempty"`
}

// RSS3Items file
type RSS3Items struct {
	*RSS3Base

	// Type extends, default to 'items'

	Items []RSS3Item `json:"items"`
}

// RSS3Link file
type RSS3Link struct {
	*RSS3Base

	// Type extends, default to 'relationship'

	Items []RSS3OtherPersona `json:"items"`

}

type RSS3OtherPersona struct {
	Id           Address `json:"id"`
	Verification string  `json:"verification,omitempty"`
	Name         string  `json:"name,omitempty"`
	Avatar       Address `json:"avatar,omitempty"`
	Bio          string  `json:"bio,omitempty"`
}

type RSS3Item struct {
	Id            string             `json:"id"`
	Authors       []RSS3OtherPersona `json:"authors,omitempty"`
	Title         string             `json:"title,omitempty"`
	Summary       string             `json:"summary,omitempty"`
	Tags          []string           `json:"tags,omitempty"`
	DatePublished string             `json:"date_published,omitempty"`
	DateModified  string             `json:"date_modified,omitempty"`

	Contents []struct {
		Id                Address  `json:"id"` // Link to a third party file
		MimeType          string   `json:"mime_type"`
		Name              string   `json:"name,omitempty"`
		Tags              []string `json:"tags,omitempty"`
		SizeInBytes       string   `json:"size_in_bytes,omitempty"`
		DurationInSeconds string   `json:"duration_in_seconds,omitempty"`
	} `json:"contents,omitempty"`

	Contexts []struct {
		Id   Address  `json:"id"` // Link to a RSS3Items file
		Name string   `json:"name,omitempty"`
		Tags []string `json:"tags,omitempty"`
	} `json:"contexts,omitempty"`

}


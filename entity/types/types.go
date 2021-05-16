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
	Editors     struct {
		Blocklist []Address `json:"blocklist"`
		Allowlist []Address `json:"allowlist"`
	} `json:"editors"`
	//Items     []interface{} `json:"items"`
	ItemsNext Address       `json:"items_next"`

}

// RSS3Persona file, Entrance
type RSS3Persona struct {
	RSS3Base

	// Type extends, default to 'persona'
	// Editors extends, default to nil

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

// RSS3Items file
type RSS3Items struct {
	RSS3Base

	// Type extends, default to 'items'

	Items []RSS3Item `json:"items"`
}

// RSS3Link file
type RSS3Link struct {
	RSS3Base

	// Type extends, default to 'relationship'

	Items []RSS3OtherPersona `json:"items"`

}

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


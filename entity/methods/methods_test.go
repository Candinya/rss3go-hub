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
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestJson2RSS3Persona(t *testing.T) {

	// RSS3Persona file: `persona:diygod` - `interface RSS3Persona`
	// A persona DIYgod with a published item Never

	demo := `
{
    "id": "persona:diygod",
    "version": "rss3.io/version/v0.1.0-alpha.0",
    "type": "persona",
    "date_created": "2009-05-01T00:00:00.000Z",
    "date_updated": "2021-05-08T16:56:35.529Z",

    "profile": {
        "name": "DIYgod",
        "avatar": "dweb://diygod.jpg",
        "bio": "写代码是热爱，写到世界充满爱！",
        "tags": ["demo", "lovely", "technology"]
    },

    "links": [{
        "id": "link:diygod:followings",
        "name": "Followings"
    }, {
        "id": "link:diygod:followers",
        "name": "Followers"
    }, {
        "id": "link:diygod:blocklist",
        "name": "Blocklist"
    }],

    "items": [{
        "id": "item:diygod:never",
        "authors": ["persona:diygod"],
        "title": "Never",
        "summary": "Never stop dreaming.",
        "date_published": "2021-05-08T16:56:35.529Z",
        "date_modified": "2021-05-08T16:56:35.529Z",

        "contents": [{
            "id": "dweb://never.html",
            "mime_type": "text/html"
        }, {
            "id": "dweb://never.jpg",
            "mime_type": "image/jpeg"
        }],

        "contexts": [{
            "id": "items:diygod:never:comments",
            "name": "Comments"
        }, {
            "id": "items:diygod:never:likes",
            "name": "Likes"
        }]
    }],
    "items_next": "items:diygod:index2"
}
`

	var orig, coded interface{}

	ret := Json2RSS3Persona(demo)
	retJsonByte := ret.ToJson()

	//t.Log(string(retJsonByte))

	if err := json.Unmarshal([]byte(demo), &orig); err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(retJsonByte, &coded); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(orig, coded) {
		t.Log(cmp.Diff(orig, coded))
		t.Fail()
	}

}

func TestJson2RSS3Items(t *testing.T) {

	// RSS3Items file: `items:diygod:never:comments` - `interface RSS3Items`

	demo := `
{
    "id": "items:diygod:never:comments",
    "version": "rss3.io/version/v0.1.0-alpha.0",
    "type": "items",
    "date_created": "2009-05-01T00:00:00.000Z",
    "date_updated": "2021-05-08T16:56:35.529Z",
    "editors": {
        "allowlist": "link:diygod:followings"
    },

    "items": [{
        "id": "items:diygod:never:comments:0",
        "authors": [{
            "id": "persona:joshua",
            "verification": "xxxxx",
            "name": "Joshua",
            "avatar": "dweb://joshua.jpg"
        }],
        "title": "DIYgod is the best!",
        "date_published": "2021-05-08T16:56:35.529Z",
        "date_modified": "2021-05-08T16:56:35.529Z",

        "contents": [{
            "id": "dweb://best.jpg",
            "mime_type": "image/jpeg"
        }],

        "contexts": [{
            "id": "items:diygod:never:comments:0:sub-comments",
            "name": "Sub-Comments"
        }]
    }]
}
`

	var orig, coded interface{}

	ret := Json2RSS3Items(demo)
	retJsonByte := ret.ToJson()

	//t.Log(string(retJsonByte))

	if err := json.Unmarshal([]byte(demo), &orig); err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(retJsonByte, &coded); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(orig, coded) {
		t.Log(cmp.Diff(orig, coded))
		t.Fail()
	}
}

func TestJson2RSS3Link(t *testing.T) {

	// RSS3Link file: `link:diygod:followers` - `interface RSS3Link`

	demo := `
{
    "id": "link:diygod:followers",
    "version": "rss3.io/version/v0.1.0-alpha.0",
    "type": "link",
    "date_created": "2009-05-01T00:00:00.000Z",
    "date_updated": "2021-05-08T16:56:35.529Z",
    "editors": {
        "blocklist": "link:diygod:blocklist"
    },

    "items": ["persona:joshua", "persona:atlas", "persona:tuzi", "persona:zuia"]
}
`

	var orig, coded interface{}

	ret := Json2RSS3Link(demo)
	retJsonByte := ret.ToJson()

	//t.Log(string(retJsonByte))

	if err := json.Unmarshal([]byte(demo), &orig); err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(retJsonByte, &coded); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(orig, coded) {
		t.Log(cmp.Diff(orig, coded))
		t.Fail()
	}

}

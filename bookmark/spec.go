package bookmark

import (
	"strconv"
	"time"
)

// const refs:
// https://github.com/mozilla/gecko-dev/blob/599a15d3547d862764048ff62b74252dd41a56d3/toolkit/components/places/Bookmarks.jsm#L95
const (
	TypeBookmark  = 1
	TypeFolder    = 2
	TypeSeparator = 3

	DefaultIndex = -1

	MaxTagLength = 100

	GUIDRoot    = "root________"
	GUIDMenu    = "menu________"
	GUIDToolbar = "toolbar_____"
	GUIDUnfiled = "unfiled_____"
	GUIDMobile  = "mobile______"
	GUIDTag     = "tags________"

	GUIDVirtMenu    = "menu_______v"
	GUIDVirtToolbar = "toolbar____v"
	GUIDVirtUnfiled = "unfiled___v"
	GUIDVirtMobile  = "mobile____v"
)

// Item refs:
// https://github.com/mozilla/gecko-dev/blob/599a15d3547d862764048ff62b74252dd41a56d3/toolkit/components/places/PlacesBackups.jsm#L398
// https://github.com/mozilla/gecko-dev/blob/599a15d3547d862764048ff62b74252dd41a56d3/toolkit/components/places/Bookmarks.jsm#L7
type Item struct {
	// GUID is the globally unique identifer of the item.
	GUID string `json:"guid"`

	// ParentGUID is the globally unique identifier of the folder containing this item.
	// It will be empty for the Places root folder. Places seems to be synonymous with "Bookmarks".
	ParentGUID string `json:"parentGuid"`

	// Title is the item's title, if any.
	Title string `json:"title"`

	// Index is the zero-based position of the item in the parent folder.
	Index int `json:"index"`

	// DateAdded represents the time at which the item was added to the tree.
	DateAdded millis `json:"dateAdded"`
	// LastModified represents the time of the last modification for this item.
	LastModified millis `json:"lastModified"`

	// ID is the item's ID. Obviously.
	ID int `json:"id"`

	// TypeCode designates what type of item this is; see Type<X> consts.
	TypeCode int `json:"typeCode"`

	// Type is a human-readable string of some Mozilla bullshit.
	Type string `json:"type"`

	// Root seems to be set on all top-level items, designating them as a parent
	// for the type of items they will contain.
	Root string `json:"root"`

	// Children are the items within a TypeFolder; this won't be set on individual items.
	Children []*Item `json:"children"`

	// The following fields only apply to a subset of items.
	// Future improvements could see these being split into separate structs,
	// but frankly I don't give a damn right now.
	Annos   []Anno `json:"annos"`
	URI     string `json:"uri"`
	IconURI string `json:"iconuri"`
	Keyword string `json:"keyword"`
	Charset string `json:"charset"`
	Tags    string `json:"tags"`
}

// Not sure where this is documented at the moment.
type Anno struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Expires int    `json:"expires"`
	Flags   int    `json:"flags"`
}

type millis struct {
	time.Time
}

func (m *millis) UnmarshalJSON(v []byte) error {
	ms := string(v)

	t, _ := strconv.ParseUint(ms, 10, 64)
	m.Add(time.Millisecond * time.Duration(t))

	return nil
}

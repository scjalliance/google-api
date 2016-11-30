package contacts

import "google.golang.org/api/googleapi"

// ContactFeedResponse represents a response to a contact feed request.
type ContactFeedResponse struct {
	Version  string      `json:"version"`
	Encoding string      `json:"encoding"`
	Feed     ContactFeed `json:"feed"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// ContactFeed represents a contact feed.
type ContactFeed struct {
	IDField      `json:"id"`
	UpdatedField `json:"updated"`
	Contacts     []*Contact `json:"entry"`
}

// Contact represents an entity such as a person, venue or organization.
type Contact struct {
	IDField        `json:"id"`
	UpdatedField   `json:"updated"`
	TitleField     `json:"title"`
	Organizations  []*Organization `json:"gd$organization,omitempty"`
	EmailAddresses []*EmailAddress `json:"gd$email,omitempty"`
	PhoneNumbers   []*PhoneNumber  `json:"gd$phoneNumber,omitempty"`
}

// Organization represents an organization that a contact is affiliated with.
type Organization struct {
	NameField  `json:"gd$orgName,omitempty"`
	TitleField `json:"gd$orgTitle,omitempty"`
}

// EmailAddress represents an email address of a contact.
type EmailAddress struct {
	Value       string `json:"address"`
	DisplayName string `json:"displayName,omitempty"`
	Label       string `json:"label,omitempty"`
	Primary     bool   `json:"primary,string"`
}

// PhoneNumber represents the phone number of a contact.
type PhoneNumber struct {
	Value string `json:"$t,omitempty"`
}

/*
// Name represents the name of a contact.
type Name struct {
	Given      string `json:"givenName,omitempty"`
	Additional string `json:"additionalName,omitempty"`
	Family     string `json:"familyName,omitempty"`
	Prefix     string `json:"namePrefix,omitempty"`
	Suffix     string `json:"nameSuffix,omitempty"`
	Full       string `json:"fullName,omitempty"`
}
*/

// IDField represents an ID field.
type IDField struct {
	ID string `json:"$t"`
}

// UpdatedField represents an update timestamp field.
type UpdatedField struct {
	Updated string `json:"$t"` // Stored in ISO 8601 date and time representation
}

// TitleField represents a title field.
type TitleField struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"$t"`
}

// NameField represents a name field.
type NameField struct {
	Name string `json:"$t"`
}

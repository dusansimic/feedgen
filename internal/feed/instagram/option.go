package instagram

import (
	content "github.com/dusansimic/feedgen/internal/content/instagram"
)

// Option sets an option for the feed.
type Option func(*F)

// WithUser sets options about the user.
// fullName is the full name of user. bio is the bio of user. username is the username of user.
// profile is the url to the profile pic of user. userID is the id of user.
func WithUser(fullName, bio, username, profile, userID string) Option {
	return func(f *F) {
		f.FullName = fullName
		f.Bio = bio
		f.Username = username
		f.Profile = profile
		f.UserID = userID
	}
}

// WithContent sets the content for the feed.
func WithContent(c []content.C) Option {
	return func(f *F) {
		f.Content = c
	}
}

// WithUpdated sets the last updated time
func WithUpdated(t int) Option {
	return func(f *F) {
		f.UpdatedTime = t
	}
}

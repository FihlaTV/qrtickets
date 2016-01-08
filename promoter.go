package qrtickets

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

// Promoter - an individual, business, or group that markets events
type Promoter struct {
	Name string `json:"name" datastore:",noindex"`
	URL  string `json:"url" datastore:",noindex"`

	DatastoreKey datastore.Key `json:"promoter_id" datastore:"-"`
}

// AddPromoter - Add a promoter from form input and save to Datastore
func AddPromoter(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Create the Promoter object
	p1 := Promoter{
		Name: r.FormValue("name"),
		URL:  r.FormValue("url"),
	}

	// Add the promoter to the Datastore
	k, err := p1.Store(ctx)
	if err != nil {
		JSONError(&w, err.Error())
		return
	}

	p1.DatastoreKey = *k
	return
}

// Store - Stores the current promoter into Google datastore
func (p *Promoter) Store(ctx context.Context) (*datastore.Key, error) {
	var k *datastore.Key

	// See if a key exists, or if a new one is required
	if p.DatastoreKey.Incomplete() {
		k = datastore.NewIncompleteKey(ctx, "Promoter", nil)
	} else {
		k = &p.DatastoreKey
	}

	// Stash the entry in the datastore
	key, err := datastore.Put(ctx, k, p)
	if err != nil {
		return nil, err
	}

	return key, nil
}

package service

import "mylab/mylabdiscoveries/db"

type Dependencies struct {
	Store db.Storer
	// define other service dependencies
}

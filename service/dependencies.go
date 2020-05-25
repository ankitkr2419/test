package service

import "mylab/cpagent/db"

type Dependencies struct {
	Store db.Storer
	// define other service dependencies
}

package repository

import (
	"cygo_iris/util/log"
	"gopkg.in/mgo.v2"
)

func ErrHandle(err error) error {
	if err == nil || err.Error() == "not found" {
		return nil
	}
	log.Logger.Errorf("Database error: %v", err)
	return err
}

// Insert
// ---------------------------------------------------------------------
func Insert(collection *mgo.Collection, i interface{}) error {
	err := collection.Insert(i)
	return ErrHandle(err)
}

// Query
// ---------------------------------------------------------------------

// GetByQ - Get one document by query
func GetByQ(collection *mgo.Collection, q interface{}, i interface{}) {
	_ = collection.Find(q).One(i)
}

func Count(collection *mgo.Collection, q interface{}) (int, error) {
	cnt, err := collection.Find(q).Count()
	err = ErrHandle(err)
	return cnt, err
}

func Has(collection *mgo.Collection, q interface{}) (bool, error) {
	cnt, err := Count(collection, q)
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

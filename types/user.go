package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"` //these are maps for bson and json bson is for mongodb usage and omitempty doesnt show that argument to public in json or bson
	Firstname string `json:"firstName" bson:"firstName"`
	Lastname  string `json:"lastName" bson:"lastName"`
}

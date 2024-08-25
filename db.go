package main

type mongoServer struct {
	Host       string
	Port       uint16
	Collection string
	DB         string
}

type user struct {
	ID  string `bson:"_id"`
	UID string `bson:"uid"`
	TC  string `bson:"tc"`
	STC string `bson:"stc"`
	BG  string `bson:"bg"`
}

package models

type Profile struct {
	Confirmed bool     `json:"confirmed" json:"confirmed"`
	Visits    int64    `json:"vists" bson:"vists"`
	Likes     int64    `json:"likes" bson:"likes"`
	Fame      float64  `json:"fame" bson:"fame"`
	Propic    string   `json:"propic" bson:"propic"`
	Images    []string `json:"images" bson:"images"`
	Sex       string   `json:"sex" bson:"sex"`
	Interests []string `json:"interests" bson:"interests"`
}

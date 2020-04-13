package types

// Extra is used to store user extra data
type Extra struct {
	UserID      string `bson:"_id,omitempty,-"`
	FirstName   string `bson:"FirstName,omitempty,-"`
	LastName    string `bson:"LastName,omitempty,-"`
	Gender      string `bson:"Gender,omitempty,-"`
	BirthdayUTC int64  `bson:"BirthdayUTC,omitempty,-"`
}

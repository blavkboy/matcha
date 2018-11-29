package models

//The user file here will deal with all structures and
//methods that will work on the user and have to deal with
//the user's part in the database

//User struct will represent a user and will server
//as a storage mechanism for the server to keep track
//of user data.
type User struct {
	ID       string
	Username string
	Fname    string
	Lname    string
	Email    string
	Password string
	Bio      *Bio
}

//Bio will hold data about the user that he only needs to fill
//this data for the matching process.
type Bio struct {
	Caddress  *Address
	Oaddress  *Address
	Sexuality *Sexuality
	Hobbies   []string
	Interests []string
}

//Address will be for our records of where the user says he stays.
//This will help widen the matching criteria to something beyond
//just his current location.
type Address struct {
	Street1  string
	Street2  string
	Suburb   string
	City     string
	Province string
}

//Sexuality will represent data regarding the user's sexual preferences.
//This also allows users to be paired with others who are sexually compatible
type Sexuality struct {
	Sex         Sex
	Orientation Orientation
	Looking     Looking
	Preferences []string
}

//LookingFor will return a string telling us what kind of relationship
//the user is interested in having
func (s *Sexuality) LookingFor() string {
	if s.Looking == Fun {
		return "Fun"
	} else if s.Looking == LTR {
		return "Long Term Relationship"
	} else if s.Looking == FR {
		return "Friendship"
	} else if s.Looking == Bless {
		return "Blesser"
	}
	return ""
}

//GetOrientation will return the user's sexual orientation
func (s *Sexuality) GetOrientation() string {
	if s.Orientation == Hetero {
		return "straight"
	} else if s.Orientation == Homo {
		return "homosexual"
	} else if s.Orientation == Non {
		return "Non Comforming"
	}
	return ""
}

//GetSex will return the sex of the user
func (s *Sexuality) GetSex() string {
	if s.Sex == Male {
		return "male"
	} else if s.Sex == Female {
		return "female"
	} else if s.Sex == NonBin {
		return "Non-Binary"
	}
	return ""
}

//Sex will represent the sex of an individual
type Sex int

const (
	//Male will be the value for a male
	Male Sex = 0 + iota
	//Female will be the value for a female
	Female
	//NonBin for the non-comformist
	NonBin
)

//Orientation will serve to identify the sexual orientation of the
//user
type Orientation int

const (
	//Hetero for those who like opposites
	Hetero Orientation = 0 + iota
	//Homo for the gay ones
	Homo
	//Bi for those who go both ways
	Bi
	//Non for those who ain't got no catagory
	Non
)

//Looking will tell us what kind of relationship the user
//is interested in engaging in.
type Looking int

const (
	//LTR is for Long Term Relationship
	LTR Looking = 0 + iota
	//Fun will mean exactly that
	Fun
	//Explore will mean that the user is open to anything so far
	Explore
	//FR will represent Friendship
	FR
	//Bless is for those seeking financial relationships
	Bless
)

package orders

import "errors"

type Address struct {
	name     string
	street   string
	city     string
	postcode string
	country  string
}

func NewAddress(name, street, city, postcode, country string) (Address, error) {
	if len(name) == 0 {
		return Address{}, errors.New("empty name")
	}
	if len(street) == 0 {
		return Address{}, errors.New("empty street")
	}
	if len(city) == 0 {
		return Address{}, errors.New("empty city")
	}
	if len(postcode) == 0 {
		return Address{}, errors.New("empty postcode")
	}
	if len(country) == 0 {
		return Address{}, errors.New("empty country")
	}
	return Address{
		name:     name,
		street:   street,
		city:     city,
		postcode: postcode,
		country:  country,
	}, nil
}

func (a Address) Name() string {
	return a.name
}

func (a Address) Street() string {
	return a.street
}

func (a Address) City() string {
	return a.city
}

func (a Address) Postcode() string {
	return a.postcode
}

func (a Address) Country() string {
	return a.country
}

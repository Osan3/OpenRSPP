package ratings

// Rating represents a rating for a service provider
type Rating struct {
	ServiceProviderID int
	Value             int
}

// Ratings is a collection of ratings for a service provider
type Ratings []Rating

// AddRating adds a new rating for a service provider
func (r *Ratings) AddRating(serviceProviderID int, value int) {
	*r = append(*r, Rating{ServiceProviderID: serviceProviderID, Value: value})
}

// DeleteRating deletes a rating for a service provider
func (r *Ratings) DeleteRating(serviceProviderID int, value int) {
	for i, rating := range *r {
		if rating.ServiceProviderID == serviceProviderID && rating.Value == value {
			*r = append((*r)[:i], (*r)[i+1:]...)
			return
		}
	}
}

// UpdateRating updates an existing rating for a service provider
func (r *Ratings) UpdateRating(serviceProviderID int, oldValue int, newValue int) {
	for i, rating := range *r {
		if rating.ServiceProviderID == serviceProviderID && rating.Value == oldValue {
			(*r)[i].Value = newValue
			return
		}
	}
}
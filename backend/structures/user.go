package structures

type User struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Gender           string `json:"gender"`
	Age              int    `json:"age"`
	ActiveHoursOnApp int    `json:"active_hours_on_app"`
	HasUnsubscribed  string `json:"has_unsubscribed"`
	OrganizationId   int    `json:"organization_id"`
}

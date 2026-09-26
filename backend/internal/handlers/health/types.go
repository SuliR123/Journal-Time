package health

// GreetingOutput represents the greeting operation response.
type HealthOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

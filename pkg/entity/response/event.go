package response

type EventResponse struct {
    Code string
    Message string
    Events []Event
}

type Event struct {
}
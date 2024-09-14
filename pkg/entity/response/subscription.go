package response

type SubscriptionResponse struct {
    Code string
    Message string
    Subscriptions []Subscription
}

type Subscription struct {
}
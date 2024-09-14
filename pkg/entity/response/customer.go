package response

type CustomerResponse struct {
    Code string
    Message string
    Customers []Customer
}

type Customer struct {
}
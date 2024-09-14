package response

type AccountResponse struct {
    Code string
    Message string
    Accounts []Account
}

type Account struct {
}
package response

type BalanceResponse struct {
    Code string
    Message string
    Balances []Balance
}

type Balance struct {
}
package response

type ChargeResponse struct {
    Code string
    Message string
    Charges []Charge
}

type Charge struct {
}
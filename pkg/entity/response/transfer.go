package response

type TransferResponse struct {
    Code string
    Message string
    Transfers []Transfer
}

type Transfer struct {
}
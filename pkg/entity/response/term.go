package response

type TermResponse struct {
    Code string
    Message string
    Terms []Term
}

type Term struct {
}
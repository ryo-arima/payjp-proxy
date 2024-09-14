package response

type StatementResponse struct {
    Code string
    Message string
    Statements []Statement
}

type Statement struct {
}
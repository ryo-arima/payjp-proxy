package response

type PlanResponse struct {
    Code string
    Message string
    Plans []Plan
}

type Plan struct {
}
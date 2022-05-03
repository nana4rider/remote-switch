package controller

type ResError struct {
	Message string `json:"message"`
}

type ResBoolError struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type ResValidationError struct {
	Message          string                      `json:"message"`
	ValidationErrors []*ResValidationErrorDetail `json:"validationErrors"`
}

type ResValidationErrorDetail struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func (p *ResValidationError) Add(param, message string) *ResValidationError {
	p.ValidationErrors = append(p.ValidationErrors, &ResValidationErrorDetail{param, message})
	return p
}

func (p *ResValidationError) Has() bool {
	return len(p.ValidationErrors) > 0
}

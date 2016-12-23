package gf


type ValidationConfigSpec struct {
	Type string
	Keys []string
}

func (vc *ValidationConfig) Spec() *ValidationConfigSpec {

	spec := &ValidationConfigSpec{
		Type:		vc.Type(),
		Keys:		vc.keys,
	}

	return spec
}

type HandlerSpec struct {
	Method string
	Function string
	Endpoint string
	PayloadSchema interface{}
	ResponseSchema interface{}
	RouteParams []*ValidationConfigSpec
	IsTemplate bool
	IsTemplateFolder bool
	TemplatePath string
	TemplateMode string
}

func (handler *Handler) Spec() *HandlerSpec {

	validations := []*ValidationConfigSpec{}

	for _, vc := range handler.node.validations {

		validations = append(validations, vc.Spec())

	}

	spec := &HandlerSpec{
		Method:					handler.method,
		Function:				handler.functionKey,
		Endpoint:				handler.node.Path,
		PayloadSchema:			handler.payloadSchema,
		ResponseSchema:			handler.responseSchema,
		IsTemplate:				(handler.template != nil),
		TemplateMode:			handler.handlerType,
		TemplatePath:			handler.templatePath,
		RouteParams:			validations,
	}

	return spec
}

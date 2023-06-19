package validate

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	desc "github.com/yiwen101/CardWizards/common/descriptor"
)

type Validator interface {
	//ValidateRoute(serviceName, methodName string, req *descriptor.HTTPRequest) (string, error)
	ValidateBody(ctx context.Context, c *app.RequestContext) error
}

type validatorImplement struct{}

func NewValidator() Validator {
	return &validatorImplement{}
}

func (v *validatorImplement) ValidateBody(ctx context.Context, c *app.RequestContext) error {
	j, err := treatJsonBody(ctx, c)
	if err != nil {
		//c.String(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error in opening the json body, error message is: %s", err))
		return err
	}

	dm, err := desc.GetDescriptorManager()
	serviceName, methodName := c.Param("serviceName"), c.Param("methodName")
	desc, err := dm.GetFunctionDescriptor(serviceName, methodName)
	if err != nil {
		//c.String(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error in getting the function descriptor, error message is: %s", err))
		return err
	}

	err = validateBody(desc, j)
	if err != nil {
		//c.String(http.StatusBadRequest, fmt.Sprintf("Invalid body, error message is: %s", err))
		return err
	}

	return nil
}

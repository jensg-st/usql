package operations

import (
	"context"
	"sync"

	"github.com/direktiv/apps/go/pkg/apps"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"usql/models"
)

const (
	successKey = "success"
	resultKey  = "result"

	// http related
	statusKey  = "status"
	codeKey    = "code"
	headersKey = "headers"
)

var sm sync.Map

const (
	cmdErr = "io.direktiv.command.error"
	outErr = "io.direktiv.output.error"
	riErr  = "io.direktiv.ri.error"
)

type accParams struct {
	PostParams
	Commands []interface{}
}

type accParamsTemplate struct {
	PostBody
	Commands []interface{}
}

func PostDirektivHandle(params PostParams) middleware.Responder {
	resp := &PostOKBody{}

	var (
		err error
		ret interface{}
	)

	ri, err := apps.RequestinfoFromRequest(params.HTTPRequest)
	if err != nil {
		return generateError(riErr, err)
	}

	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	sm.Store(*params.DirektivActionID, cancel)
	defer sm.Delete(params.DirektivActionID)

	var responses []interface{}

	var paramsCollector []interface{}
	accParams := accParams{
		params,
		nil,
	}

	ret, err = runCommand0(ctx, accParams, ri)
	responses = append(responses, ret)

	if err != nil && true {
		errName := cmdErr
		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	s, err := templateString(`{
  "queries": {{  index . 0 | toJson }}
}
`, responses)
	if err != nil {
		return generateError(outErr, err)
	}

	responseBytes := []byte(s)

	// validate

	resp.UnmarshalBinary(responseBytes)
	err = resp.Validate(strfmt.Default)

	if err != nil {
		return generateError(outErr, err)
	}

	return NewPostOK().WithPayload(resp)
}

// foreach command
type LoopStruct0 struct {
	accParams
	Item interface{}
}

func runCommand0(ctx context.Context,
	params accParams, ri *apps.RequestInfo) ([]map[string]interface{}, error) {

	ri.Logger().Infof("foreach command over .Queries")

	var cmds []map[string]interface{}

	for a := range params.Body.Queries {

		ls := &LoopStruct0{
			params,
			params.Body.Queries[a],
		}

		cmd, err := templateString(`usql {{ .Body.Connection }} -c "{{ .Item.Query }}" -J --set SHOW_HOST_INFORMATION=false 
{{- range .Item.Args }} --set={{.}} {{- end}}`, ls)
		if err != nil {
			ir := make(map[string]interface{})
			ir[successKey] = false
			ir[resultKey] = err.Error()
			cmds = append(cmds, ir)
			continue
		}

		silent := convertTemplateToBool("<no value>", ls, false)
		print := convertTemplateToBool("<no value>", ls, true)
		output := ""

		envs := []string{}

		r, err := runCmd(ctx, cmd, envs, output, silent, print, ri)
		if err != nil {
			ir := make(map[string]interface{})
			ir[successKey] = false
			ir[resultKey] = err.Error()
			cmds = append(cmds, ir)
			continue
		}
		cmds = append(cmds, r)

	}

	return cmds, nil

}

// end commands

func generateError(code string, err error) *PostDefault {

	d := NewPostDefault(0).WithDirektivErrorCode(code).
		WithDirektivErrorMessage(err.Error())

	errString := err.Error()

	errResp := models.Error{
		ErrorCode:    &code,
		ErrorMessage: &errString,
	}

	d.SetPayload(&errResp)

	return d
}

func HandleShutdown() {
	// nothing for generated functions
}

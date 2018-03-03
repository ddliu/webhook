package daocloud

import (
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/task"
)

type TaskBuild struct {
}

func (t *TaskBuild) GetId() string {
	return "daocloud_build"
}

type BuildInput struct {
	Id     string
	Token  string
	Branch string
}

func (t *TaskBuild) Run(ctx *context.Context) (*context.Context, error) {
	var input BuildInput
	inputContext := ctx.GetContext("task.input")
	if err := inputContext.Unmarshal(&input); err != nil {
		return nil, err
	}

	var response map[string]string
	err := callApi("POST", "/v1/build-flows/"+ctx.Tpl(input.Id)+"/builds", ctx.Tpl(input.Token), map[string]string{
		"branch": ctx.Tpl(input.Branch),
	}, &response)

	if err != nil {
		return nil, err
	}

	return context.New(response), nil
}

func init() {
	task.RegisterTask(&TaskBuild{})
}

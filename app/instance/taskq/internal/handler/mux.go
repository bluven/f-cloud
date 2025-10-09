package handler

import (
	"github.com/hibiken/asynq"

	"github.com/bluven/f-cloud/app/instance/taskq/internal/svc"
	"github.com/bluven/f-cloud/app/instance/taskq/types"
)

func Register(svcCtx *svc.ServiceContext) *asynq.ServeMux {

	mux := asynq.NewServeMux()
	mux.Handle(types.TaskInstanceInit, NewInstanceInitHandler(svcCtx))
	mux.Handle(types.TaskInstanceDestroy, NewInstanceDestroyHandler(svcCtx))

	return mux
}

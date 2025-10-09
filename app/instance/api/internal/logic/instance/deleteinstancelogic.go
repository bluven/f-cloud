package instance

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/bluven/f-cloud/app/instance/query"
	tasktypes "github.com/bluven/f-cloud/app/instance/taskq/types"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type DeleteInstanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInstanceLogic {
	return &DeleteInstanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInstanceLogic) DeleteInstance(req *types.DeleteRequest) (resp *types.EmptyResponse, err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		instance, err := tx.Instance.GetByID(uint(req.ID))
		if err != nil {
			return err
		}

		ri, err := tx.Instance.Delete(&instance)
		switch {
		case err != nil:
			return err
		case ri.RowsAffected == 0:
			return errorx.NewNotFound("instance not found")
		default:
		}

		task := tasktypes.NewTaskInstanceDestroy(instance.ID, instance.NetworkID, instance.DiskID)
		_, err = l.svcCtx.TaskClient.Enqueue(task)
		return err
	})

	return &types.EmptyResponse{}, err
}

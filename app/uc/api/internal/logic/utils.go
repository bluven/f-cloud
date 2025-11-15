package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/model"
	"github.com/bluven/f-cloud/app/uc/query"
	"github.com/zeromicro/go-zero/core/threading"
)

func UserKey(id uint) string {
	return fmt.Sprintf("cache:user:user:%d", id)
}

func cacheGetUser(ctx context.Context, svcCtx *svc.ServiceContext, userID uint) (model.User, error) {
	var user model.User
	err := svcCtx.Cache.TakeCtx(ctx, &user, UserKey(userID), func(val any) error {
		u, err := query.User.GetByID(userID)
		if err != nil {
			return err
		}

		fmt.Printf("Value: %v, Type: %T Type: %T\n", val, val, u)

		*val.(*model.User) = u
		return nil
	})

	return user, err
}

func cacheDeleteUser(ctx context.Context, svcCtx *svc.ServiceContext, userID uint) {
	svcCtx.Cache.DelCtx(ctx, UserKey(userID))

	threading.GoSafe(func() {
		time.Sleep(1 * time.Second)
		svcCtx.Cache.DelCtx(ctx, UserKey(userID))
	})
}

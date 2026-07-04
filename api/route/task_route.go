package route

import (
	"time"

	"github.com/Prakashdurai/golang/api/controller"
	"github.com/Prakashdurai/golang/bootstrap"
	"github.com/Prakashdurai/golang/domain"
	"github.com/Prakashdurai/golang/mongo"
	"github.com/Prakashdurai/golang/repository"
	"github.com/Prakashdurai/golang/usecase"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}

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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}

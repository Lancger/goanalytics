package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/lt90s/goanalytics/conf"
	"github.com/lt90s/goanalytics/event/pubsub"
	"github.com/lt90s/goanalytics/metric/user"
	"github.com/lt90s/goanalytics/storage/mongodb"
)

func SetupMetricProcessor(subscriber pubsub.Subscriber) {
	mongoClient := mongodb.DefaultClient
	prefix := conf.GetConfString(conf.MongoDatabasePrefixKey)

	userStore := user.NewMongoStore(mongoClient, prefix)
	user.SetupProcessor(subscriber, userStore)
}

func SetupMetricApi(iRouter *gin.RouterGroup, oRouter *gin.RouterGroup, publisher pubsub.Publisher) {
	mongoClient := mongodb.DefaultClient
	prefix := conf.GetConfString(conf.MongoDatabasePrefixKey)

	userStore := user.NewMongoStore(mongoClient, prefix)
	user.SetupRoute(iRouter, oRouter, publisher, userStore)
}

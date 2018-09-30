package routers

import (
	"installer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testssh", &controllers.SingleController{}, "get:TestSSH")
	beego.Router("/checkDocker", &controllers.SingleController{}, "get:CheckDockerExisted")
	beego.Router("/pullImages", &controllers.SingleController{}, "get:PullImages")
	beego.Router("/tagImages", &controllers.SingleController{}, "get:TagImages")
	beego.Router("/removeTempImages", &controllers.SingleController{}, "get:RemoveImages")
	beego.Router("/checkKubectl", &controllers.SingleController{}, "get:CheckKubectlExisted")
	beego.Router("/checkAndUseConfig", &controllers.SingleController{}, "get:CheckAndUseConfig")
	beego.Router("/kubeProxy", &controllers.SingleController{}, "get:KubeProxy")
	beego.Router("/deployK8sDashboard", &controllers.SingleController{}, "get:DeployK8sDashboard")
}

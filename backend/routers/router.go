// @APIVersion 1.0.0
// @Title X Meme API
// @Description Submission for Crio Winter of Doing Stage 2B
// @Contact karngyan@gmail.com
// @License AGPLv3.0
// @LicenseUrl https://www.gnu.org/licenses/gpl-3.0.en.html
// @Schemes http
// @Host 127.0.0.1:8081
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"xmeme/controllers"
)

func init() {
	//ns := beego.NewNamespace("/",
	//	beego.NSNamespace("/memes",
	//		beego.NSInclude(
	//			&controllers.MemeController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
	beego.Include(&controllers.MemeController{})
}

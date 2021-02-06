package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["xmeme/controllers:MemeController"] = append(beego.GlobalControllerRouter["xmeme/controllers:MemeController"],
        beego.ControllerComments{
            Method: "GetAllMemes",
            Router: "/memes",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xmeme/controllers:MemeController"] = append(beego.GlobalControllerRouter["xmeme/controllers:MemeController"],
        beego.ControllerComments{
            Method: "CreateMeme",
            Router: "/memes",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xmeme/controllers:MemeController"] = append(beego.GlobalControllerRouter["xmeme/controllers:MemeController"],
        beego.ControllerComments{
            Method: "GetMeme",
            Router: "/memes/:mid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

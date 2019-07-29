package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetIndex",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetComment",
            Router: `/comment`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetDetails",
            Router: `/details/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetError",
            Router: `/error`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetReg",
            Router: `/reg`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["blog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "EditPage",
            Router: `/edit/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["blog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "NewPage",
            Router: `/new`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["blog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "Save",
            Router: `/save/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:UserController"] = append(beego.GlobalControllerRouter["blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:UserController"] = append(beego.GlobalControllerRouter["blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:UserController"] = append(beego.GlobalControllerRouter["blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Reg",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

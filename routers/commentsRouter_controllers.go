package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["devops/controllers:ConfigController"] = append(beego.GlobalControllerRouter["devops/controllers:ConfigController"],
        beego.ControllerComments{
            Method: "AddConfig",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:ConfigGitController"] = append(beego.GlobalControllerRouter["devops/controllers:ConfigGitController"],
        beego.ControllerComments{
            Method: "AddConfigGit",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:ConfigProcessController"] = append(beego.GlobalControllerRouter["devops/controllers:ConfigProcessController"],
        beego.ControllerComments{
            Method: "AddConfigProcess",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:ConfigsController"] = append(beego.GlobalControllerRouter["devops/controllers:ConfigsController"],
        beego.ControllerComments{
            Method: "AddConfigs",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:DefaultController"] = append(beego.GlobalControllerRouter["devops/controllers:DefaultController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"any"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["devops/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "DepartmentList",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["devops/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "UpdateDepartment",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["devops/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "DelDepartment",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["devops/controllers:DepartmentController"],
        beego.ControllerComments{
            Method: "AddDepartment",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:UserController"] = append(beego.GlobalControllerRouter["devops/controllers:UserController"],
        beego.ControllerComments{
            Method: "Auth",
            Router: `/auth`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:UserController"] = append(beego.GlobalControllerRouter["devops/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:UserController"] = append(beego.GlobalControllerRouter["devops/controllers:UserController"],
        beego.ControllerComments{
            Method: "Registered",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["devops/controllers:UserController"] = append(beego.GlobalControllerRouter["devops/controllers:UserController"],
        beego.ControllerComments{
            Method: "UploadAvatar",
            Router: `/upload_avatar`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

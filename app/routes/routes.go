// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tComments struct {}
var Comments tComments


func (_ tComments) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Comments.Index", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tTop struct {}
var Top tTop


func (_ tTop) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Top.Index", args).Url
}


type tApiV1Controller struct {}
var ApiV1Controller tApiV1Controller


func (_ tApiV1Controller) HandleBadRequestError(
		s string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "s", s)
	return revel.MainRouter.Reverse("ApiV1Controller.HandleBadRequestError", args).Url
}

func (_ tApiV1Controller) HandleNotFoundError(
		s string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "s", s)
	return revel.MainRouter.Reverse("ApiV1Controller.HandleNotFoundError", args).Url
}

func (_ tApiV1Controller) HandleInternalServerError(
		s string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "s", s)
	return revel.MainRouter.Reverse("ApiV1Controller.HandleInternalServerError", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tApiV1Comments struct {}
var ApiV1Comments tApiV1Comments


func (_ tApiV1Comments) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiV1Comments.Index", args).Url
}

func (_ tApiV1Comments) Show(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("ApiV1Comments.Show", args).Url
}

func (_ tApiV1Comments) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiV1Comments.Create", args).Url
}

func (_ tApiV1Comments) Delete(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("ApiV1Comments.Delete", args).Url
}



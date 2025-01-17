package controller

import (
	"errors"
	"github.com/1340691923/xwl_bi/platform-basic-libs/jwt"
	"github.com/1340691923/xwl_bi/platform-basic-libs/request"
	"github.com/1340691923/xwl_bi/platform-basic-libs/response"
	"github.com/1340691923/xwl_bi/platform-basic-libs/service/consumer_data"
	"github.com/1340691923/xwl_bi/platform-basic-libs/service/debug_data"
	"github.com/1340691923/xwl_bi/platform-basic-libs/service/realdata"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type RealDataController struct {
	BaseController
}

//查看实时数据列表
func (this RealDataController) List(ctx *fiber.Ctx) error {

	type ReqData struct {
		Appid    int    `json:"appid"`
		SearchKw string `json:"search_kw"`
		Date     string `json:"date"`
	}

	var reqData ReqData

	if err := ctx.BodyParser(&reqData); err != nil {
		return this.Error(ctx, err)
	}

	appid := strconv.Itoa(reqData.Appid)
	searchKw := reqData.SearchKw
	date := reqData.Date

	clientReportData := consumer_data.ClientReportData{
		TableId: appid,
	}
	res, err := clientReportData.GetList(ctx.Context(), searchKw, date)
	if err != nil {
		return this.Error(ctx, err)
	}

	return this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": res.Hits.Hits})
}

//错误数据列表
func (this RealDataController) FailDataList(ctx *fiber.Ctx) error {

	type ReqData struct {
		Appid   int `json:"appid"`
		Minutes int `json:"minutes"`
	}

	var reqData ReqData

	if err := ctx.BodyParser(&reqData); err != nil {
		return this.Error(ctx, err)
	}

	if reqData.Minutes == 0 {
		reqData.Minutes = 10
	}

	realDataService := realdata.RealDataService{}

	res,err := realDataService.FailDataList(reqData.Minutes,reqData.Appid)
	if err != nil {
		return this.Error(ctx, err)
	}

	return this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": res})
}

//抽样示例
func (this RealDataController) FailDataDesc(ctx *fiber.Ctx) error {

	type ReqData struct {
		StartTime     string `json:"start_time"`
		EndTime       string `json:"end_time"`
		Appid         int    `json:"appid"`
		ErrorReason   string `json:"error_reason"`
		ErrorHandling string `json:"error_handling"`
		ReportType    string    `json:"report_type"`
	}
	var reqData ReqData

	if err := ctx.BodyParser(&reqData); err != nil {
		return this.Error(ctx, err)
	}

	startTime := reqData.StartTime
	endTime := reqData.EndTime
	appid := strconv.Itoa(reqData.Appid)
	errorReason := reqData.ErrorReason
	errorHandling := reqData.ErrorHandling
	reportType := reqData.ReportType

	realDataService := realdata.RealDataService{}

	res,err := realDataService.FailDataDesc(appid,startTime,endTime,errorReason,errorHandling,reportType)
	if err != nil {
		return this.Error(ctx, err)
	}

	return this.Success(ctx, response.SearchSuccess, map[string]interface{}{"data": res})
}

//查看所有上报数据情况
func (this RealDataController) ReportCount(ctx *fiber.Ctx) error {

	var err error

	var reqData request.ReportCountReq
	if err := ctx.BodyParser(&reqData); err != nil {
		return this.Error(ctx, err)
	}

	startTime := reqData.StartTime
	endTime := reqData.EndTime
	appid := strconv.Itoa(reqData.Appid)

	realDataService := realdata.RealDataService{}

	res,err := realDataService.ReportCount(appid,startTime,endTime)
	if err != nil {
		return this.Error(ctx, err)
	}


	return this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": res})
}

//事件失败详情
func (this RealDataController) EventFailDesc(ctx *fiber.Ctx) error {


	var reqData request.EventFailDescReq
	if err := ctx.BodyParser(&reqData); err != nil {
		return this.Error(ctx, err)
	}

	startTime := reqData.StartTime
	endTime := reqData.EndTime
	appid := strconv.Itoa(reqData.Appid)
	dataName := reqData.DataName

	realDataService := realdata.RealDataService{}

	res,err := realDataService.EventFailDesc(appid,startTime,endTime,dataName)
	if err != nil {
		return this.Error(ctx, err)
	}

	return this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": res})
}

//添加DEBUG设备ID
func (this RealDataController) AddDebugDeviceID(ctx *fiber.Ctx) error {

	var reqData request.AddDebugDeviceIDReq
	if err := ctx.BodyParser(&reqData); err != nil {
		return this.Error(ctx, err)
	}

	appid := strconv.Itoa(reqData.Appid)
	remark := reqData.Remark
	deviceID := reqData.DeviceID

	if deviceID == "" {
		return this.Error(ctx, errors.New("设备ID不能为空"))
	}

	c, _ := jwt.ParseToken(this.GetToken(ctx))

	debugData := debug_data.DebugData{}

	err :=debugData.AddDebugDeviceID(appid,deviceID,remark,c.UserID)

	if err != nil {
		return this.Error(ctx, err)
	}

	return this.Success(ctx, response.OperateSuccess, nil)
}

//删除测试设备
func (this RealDataController) DelDebugDeviceID(ctx *fiber.Ctx) error {

	var reqData request.DelDebugDeviceIDReq
	if err := ctx.BodyParser(&reqData); err != nil {
		return this.Error(ctx, err)
	}

	appid := strconv.Itoa(reqData.Appid)

	deviceID := reqData.DeviceID

	c, _ := jwt.ParseToken(this.GetToken(ctx))

	debugData := debug_data.DebugData{}

	err :=debugData.DelDebugDeviceID(appid,deviceID,c.UserID)

	if err != nil {
		return this.Error(ctx, err)
	}

	return this.Success(ctx, response.OperateSuccess, nil)
}

//查看测试设备列表
func (this RealDataController) DebugDeviceIDList(ctx *fiber.Ctx) error {

	var reqData request.DebugDeviceIDListReq
	if err := ctx.BodyParser(&reqData); err != nil {
		return this.Error(ctx, err)
	}

	appid := reqData.Appid
	c, _ := jwt.ParseToken(this.GetToken(ctx))

	debugData := debug_data.DebugData{}

	res,err :=debugData.DebugDeviceIDList(appid,c.UserID)

	if err != nil {
		return this.Error(ctx, err)
	}
	return this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": res})
}

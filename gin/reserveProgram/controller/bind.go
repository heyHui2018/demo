package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ngaut/log"
)

func Query(c *gin.Context) {
	processId, _ := c.Get("processId")
	log.Infof(processId.(string))
	//start := time.Now()
	//processId := time.Now().Format("20060102150405") + utils.GetRandString()
	//data := make(map[string]interface{})
	//
	//bindRelation := new(models.BindingRelation)
	//bindRelation.OpenId = b.GetString("openId")
	//log.Infof("Check processId = %v,openId = %v", processId, bindRelation.OpenId)
	//
	//toCheck := map[string]string{
	//	"openId": bindRelation.OpenId,
	//}
	//if ok, _ := utils.ParamsCheck(toCheck); !ok {
	//	result := utils.Response(base.MISSING_PARAMS, base.MISSING_PARAMS_MSG, data)
	//	b.Ctx.Output.JSON(result, false, false)
	//	return
	//}
	//
	//queryBR, err := bindRelation.QueryByUniqueKey()
	//if err != nil {
	//	log.Warnf("Check Query bindRelation error,processId = %v,err = %v", processId, err)
	//	result := utils.Response(base.SYSTEM_ERR, base.SYSTEM_ERR_MSG, data)
	//	b.Ctx.Output.JSON(result, false, false)
	//	return
	//}
	//data["moretvId"] = queryBR.MoretvId
	//data["bindStatus"] = queryBR.BindStatus
	//result := utils.Response(base.SUCCESS, base.SUCCESS_MSG, data)
	//log.Infof("Check processId = %v,result = %v,耗时 %v", processId, result, time.Since(start))
	//b.Ctx.Output.JSON(result, false, false)
}

//func (b *BindController) Unbind() {
//	start := time.Now()
//	processId := time.Now().Format("20060102150405") + utils.GetRandString()
//	data := make(map[string]interface{})
//
//	bindRelation := new(models.BindingRelation)
//	bindRelation.OpenId = b.GetString("openId")
//	bindRelation.MoretvId = b.GetString("moretvId")
//	log.Infof("Unbind processId = %v,openId = %v,moretvId = %v", processId, bindRelation.OpenId, bindRelation.MoretvId)
//
//	toCheck := map[string]string{
//		"openId":   bindRelation.OpenId,
//		"moretvId": bindRelation.MoretvId,
//	}
//	if ok, _ := utils.ParamsCheck(toCheck); !ok {
//		result := utils.Response(base.MISSING_PARAMS, base.MISSING_PARAMS_MSG, data)
//		b.Ctx.Output.JSON(result, false, false)
//		return
//	}
//
//	queryBR, err := bindRelation.QueryByUniqueKey()
//	if err != nil {
//		log.Warnf("Unbind Query bindRelation error,processId = %v,err = %v", processId, err)
//		result := utils.Response(base.SYSTEM_ERR, base.SYSTEM_ERR_MSG, data)
//		b.Ctx.Output.JSON(result, false, false)
//		return
//	}
//	if queryBR.MoretvId != bindRelation.MoretvId {
//		log.Warnf("Unbind MoretvId 不匹配,processId = %v,当前绑定MoretvId = %v", processId, queryBR.MoretvId)
//		result := utils.Response(base.PARAMS_ERR, base.PARAMS_ERR_MSG, data)
//		b.Ctx.Output.JSON(result, false, false)
//		return
//	}
//
//	bindRelation.MoretvId = ""
//	bindRelation.BindStatus = 0
//	err = bindRelation.UpdateMoretvIdByUniqueKey()
//	if err != nil {
//		log.Warnf("Unbind UpdateMoretvIdByUniqueKey error, processId = %v,err = %v", processId, err)
//		result := utils.Response(base.SYSTEM_ERR, base.SYSTEM_ERR_MSG, data)
//		b.Ctx.Output.JSON(result, false, false)
//		return
//	}
//	result := utils.Response(base.SUCCESS, base.SUCCESS_MSG, data)
//	log.Infof("Check processId = %v,result = %v,耗时 %v", processId, result, time.Since(start))
//	b.Ctx.Output.JSON(result, false, false)
//}

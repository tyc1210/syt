package v1

import (
	"github.com/gin-gonic/gin"
	"syt/internal/entry/response"
	"syt/pkg/app"
	"syt/pkg/errcode"
	"syt/pkg/util"
)

type Hospital struct {
	Name  string `json:"name"`
	Level string `json:"level"`
	Tip   string `json:"tip"`
	Img   string `json:"img"`
}

var (
	hospitalList []Hospital
)

func init() {
	initHospitalList()
}

func initHospitalList() {
	hospitalList = append(hospitalList,
		Hospital{"航天中心医院1", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院2", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院3", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院4", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院5", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院6", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院7", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院8", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院9", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院10", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院11", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院12", "二级甲等", "每天 8：00 放号", "1.png"},
		Hospital{"航天中心医院13", "二级甲等", "每天 8：00 放号", "1.png"},
	)
}

func NewHospital() Hospital {
	return Hospital{}
}

// GetHospital
// @Summary 分页获取医院信息
// @Tags 医院管理
// @Produce  json
// @Param page path number true "第几页"
// @Param limit path number true "一页个数"
// @Router /api/hosp/hospital/{page}/{limit} [get]
// @Success 200 {object} app.Response "成功"
func (h Hospital) GetHospital(c *gin.Context) {
	result := app.NewCommonResult(c)
	page, err := util.StrToInt(c.Param("page"))
	limit, err := util.StrToInt(c.Param("limit"))
	if err != nil {
		result.Error(&errcode.ServerError)
		return
	}
	total := len(hospitalList)
	start := page * limit
	end := (page + 1) * limit
	data := make([]Hospital, 0, 0)
	if start < total {
		if end >= total {
			end = total
		}
		data = hospitalList[start:end]
	}
	pageData := response.NewPageData(total, data)
	result.Success(pageData)
	return
}

// GetLevel
// @Summary 获取医院等级分类
// @Tags 医院管理
// @Produce  json
// @Router /api/hosp/level [get]
// @Success 200 {object} app.Response "成功"
func (h Hospital) GetLevel(c *gin.Context) {
	result := app.NewCommonResult(c)
	items := make([]string, 0, 6)
	items = append(items,
		"三级甲等", "三级乙等", "二级甲等", "二级乙等", "一级甲等", "一级乙等",
	)
	result.Success(items)
	return
}

// GetRegion
// @Summary 获取医院地区分类
// @Tags 医院管理
// @Produce  json
// @Router /api/hosp/region [get]
// @Success 200 {object} app.Response "成功"
func (h Hospital) GetRegion(c *gin.Context) {
	result := app.NewCommonResult(c)
	items := make([]string, 0, 6)
	items = append(items,
		"东城区", "西城区", "朝阳区", "丰台区", "石景山区", "海淀区", "门头沟区", "房山区", "通州区", "顺义区",
		"昌平区", "大兴区", "怀柔区", "密云区", "延庆区",
	)
	result.Success(items)
	return
}

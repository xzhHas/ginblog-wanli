package chat_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// ChatListView 群聊记录
// @Tags 群聊管理
// @Summary 群聊记录
// @Description 群聊记录
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/chat_groups_records [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ChatModel]}
func (ChatApi) ChatListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	cr.Sort = "created_at desc"
	list, count, _ := common.ComList(models.ChatModel{ISGroup: true}, common.Option{
		PageInfo: cr,
	})

	// 判断是否为空 json-filter空值问题
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ChatModel, 0)
		res.OkWithList(list, count, c)
		return
	}
	res.OkWithList(data, count, c)
}

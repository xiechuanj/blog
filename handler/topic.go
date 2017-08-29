package handler

import (
	// "encoding/json"
	// "fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/xiechuanj/blog/models"
	"github.com/xiechuanj/blog/module"
	"gopkg.in/macaron.v1"
	"strconv"
	"time"
	// "errors"
)

func GetTopics(ctx *macaron.Context) {
	ctx.Data["IsTopic"] = true
	ctx.Data["IsLogin"] = checkAccount(ctx)
	var err error
	ctx.Data["Topics"], err = module.GetTopics()
	if err != nil {
		log.Error(err.Error())
	}
	ctx.HTML(200, "topic")
}

func TopicAdd(ctx *macaron.Context) {
	ctx.Data["IsTopic"] = true
	// var err error
	// ctx.Data["Topics"], err = module.GetTopics()
	// if err != nil {
	// 	log.Error(err.Error())
	// }
	ctx.HTML(200, "topic_add")
}

func TopicView(ctx *macaron.Context) {
	ctx.Data["IsTopic"] = true
	ctx.Data["IsLogin"] = checkAccount(ctx)
	topicId, _ := strconv.ParseInt(ctx.Params(":topicId"), 10, 64)
	var err error
	ctx.Data["Topic"], err = module.GetTopic(topicId)
	if err != nil {
		log.Info(err.Error())
		ctx.Redirect("/", 302)
		return
	}
	ctx.HTML(200, "topic_view")
}

type TopicAddForm struct {
	// Tid     string `form:"tid" binding:""`
	Title   string `form:"title" binding:"Required"`
	Content string `form:"content" binding:"Required"`
}

func TopicPost(ctx *macaron.Context, tp TopicAddForm) {
	tid := ctx.QueryInt64("tid")

	if !checkAccount(ctx) {
		ctx.Redirect("/login", 302)
		return
	}
	// tid := tp.Tid
	// fmt.Println(tid)
	title := tp.Title
	content := tp.Content
	var err error
	if tid == 0 {
		_, err := module.CreateTopic(title, content)
		if err != nil {
			log.Info(err.Error())
		}
	} else {
		topicInfo := new(models.Topic)
		err = topicInfo.GetTopic().Where("id = ?", tid).Find(&topicInfo).Error
		if err != nil {
			log.Error(err.Error())
			return
		}
		// topicInfo.ID = tid
		topicInfo.Title = title
		topicInfo.Content = content
		topicInfo.Updated = time.Now().Format("2006-01-02 15:04:05")
		err = module.UpdateTopic(*topicInfo)
		if err != nil {
			log.Info(err.Error())
		}
	}
	ctx.Redirect("/topic", 302)
	return

}

// type TopicModifyForm struct {
// 	Id      string `form:"id" binding:"Required"`
// 	Title   string `form:"title" binding:"Required"`
// 	Content string `form:"content" binding:"Required"`
// }

// func TopicPut(ctx *macaron.Context, tp TopicModifyForm) {
// 	// op := ctx.QueryInt64("op")
// 	fmt.Println(tp.Id)
// 	fmt.Println(tp.Title)
// 	fmt.Println(tp.Content)
// 	if !checkAccount(ctx) {
// 		ctx.Redirect("/login", 302)
// 		return
// 	}

// 	// topicInfo := new(models.Topic)
// 	// err = topicInfo.GetTopic().Where("id = ?", body.ID).Find(&topicInfo).Error
// 	// if err != nil {
// 	// 	result, _ = json.Marshal(map[string]string{"errMsg": "error when get topic info from db:" + err.Error()})
// 	// 	return http.StatusBadRequest, result
// 	// }

// 	// if topicInfo.ID == 0 {
// 	// 	result, _ = json.Marshal(map[string]string{"errMsg": "topic is not exist"})
// 	// 	return http.StatusBadRequest, result
// 	// }

// 	// topicInfo.Uid = body.Uid
// 	// topicInfo.Title = body.Title
// 	// topicInfo.Content = body.Content
// 	// topicInfo.Updated = time.Now().Format("2006-01-02 15:04:05")
// 	// err = module.UpdateTopic(*topicInfo)
// 	// // uid := tp.Uid
// 	// title := tp.Title
// 	// content := tp.Content
// 	// _, err := module.UpdateTopic(title, content)
// 	// if err != nil {
// 	// 	log.Info(err.Error())
// 	// }
// 	ctx.Redirect("/topic", 302)
// 	return

// }

func ModifyTopic(ctx *macaron.Context) {

	ctx.Data["IsTopic"] = true
	ctx.Data["IsLogin"] = checkAccount(ctx)
	tid := ctx.QueryInt64("tid")

	var err error
	ctx.Data["Topic"], err = module.GetTopic(tid)
	if err != nil {
		log.Info(err.Error())
	}
	ctx.Data["Tid"] = tid
	ctx.HTML(200, "topic_modify")
}

// func ViewTopic(ctx *macaron.Context) {
// 	result, _ := json.Marshal(map[string]string{"message": ""})

// 	id, _ := strconv.ParseInt(ctx.Params(":topic"), 10, 64)

// 	resultMap, err := module.GetTopic(id)
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when get topic info:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	result, _ = json.Marshal(resultMap)
// 	return http.StatusOK, result
// }

// func GetTopics(ctx *macaron.Context) (int, []byte) {
// 	result, _ := json.Marshal(map[string]string{"message": ""})
// 	topics, err := module.GetTopics()
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when get topics: " + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	result, _ = json.Marshal(map[string]interface{}{"count": len(topics), "list": topics})
// 	return http.StatusOK, result
// }

// func PostTopic(ctx *macaron.Context) (int, []byte) {
// 	result, _ := json.Marshal(map[string]string{"message": ""})

// 	body := new(struct {
// 		Uid     int64  `json:"uid"`
// 		Title   string `json:"title"`
// 		Content string `json:"content"`
// 	})

// 	reqBody, err := ctx.Req.Body().Bytes()
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when get request body:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	err = json.Unmarshal(reqBody, &body)
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when unmarshal request body:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	resultStr, err := module.CreateTopic(body.Uid, body.Title, body.Content)
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when create workflow:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	result, _ = json.Marshal(map[string]string{"message": resultStr})

// 	return http.StatusOK, result
// }

// func GetTopic(ctx *macaron.Context) (int, []byte) {
// 	result, _ := json.Marshal(map[string]string{"message": ""})

// 	id, _ := strconv.ParseInt(ctx.Params(":topic"), 10, 64)

// 	resultMap, err := module.GetTopic(id)
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when get topic info:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	result, _ = json.Marshal(resultMap)
// 	return http.StatusOK, result
// }

// func PutTopic(ctx *macaron.Context) (int, []byte) {
// 	result, _ := json.Marshal(map[string]string{"message": ""})

// 	body := new(struct {
// 		ID      int64  `json:"id"`
// 		Uid     int64  `json:"uid"`
// 		Title   string `json:"title"`
// 		Content string `json:"content"`
// 	})

// 	reqBody, err := ctx.Req.Body().Bytes()
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when get request body:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	err = json.Unmarshal(reqBody, &body)
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when unmarshal request body:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	topicInfo := new(models.Topic)
// 	err = topicInfo.GetTopic().Where("id = ?", body.ID).Find(&topicInfo).Error
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when get topic info from db:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	if topicInfo.ID == 0 {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "topic is not exist"})
// 		return http.StatusBadRequest, result
// 	}

// 	topicInfo.Uid = body.Uid
// 	topicInfo.Title = body.Title
// 	topicInfo.Content = body.Content
// 	topicInfo.Updated = time.Now().Format("2006-01-02 15:04:05")
// 	err = module.UpdateTopic(*topicInfo)

// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when save topic info:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	result, _ = json.Marshal(map[string]string{"message": "success"})

// 	return http.StatusOK, result
// }

// func DeleteTopic(ctx *macaron.Context) (int, []byte) {
// 	result := []byte("")

// 	topicID, _ := strconv.ParseInt(ctx.Params(":topic"), 10, 64)

// 	err := module.DeleteTopic(topicID)
// 	if err != nil {
// 		result, _ = json.Marshal(map[string]string{"errMsg": "error when delete topic info from db:" + err.Error()})
// 		return http.StatusBadRequest, result
// 	}

// 	result, _ = json.Marshal(map[string]string{"message": "success"})
// 	return http.StatusOK, result
// }

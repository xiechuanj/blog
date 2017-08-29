package module

import (
	// "bytes"
	// "encoding/json"
	"errors"
	// log "github.com/Sirupsen/logrus"
	"github.com/xiechuanj/blog/models"
	// "github.com/xiechuanj/blog/utils"
	// "io/ioutil"
	// "net/http"
	// "strconv"
	"strings"
	"time"
)

func GetTopics() ([]map[string]interface{}, error) {
	resultMap := make([]map[string]interface{}, 0)
	topicList := make([]models.Topic, 0)

	err := new(models.Topic).GetTopic().Order("-id").Find(&topicList).Error
	if err != nil {
		return nil, errors.New("error when get topic infos error:" + err.Error())
	}

	for _, topicInfo := range topicList {
		topicsMap := make(map[string]interface{})
		topicsMap["id"] = topicInfo.ID
		topicsMap["uid"] = topicInfo.Uid
		topicsMap["title"] = topicInfo.Title
		topicsMap["content"] = strings.Split(topicInfo.Content, "\n")[0]
		topicsMap["attachment"] = topicInfo.Attachment
		topicsMap["created"] = topicInfo.Created
		topicsMap["updated"] = topicInfo.Updated
		topicsMap["views"] = topicInfo.Views
		topicsMap["author"] = topicInfo.Author
		topicsMap["replytime"] = topicInfo.ReplyTime
		topicsMap["replycount"] = topicInfo.ReplyCount
		topicsMap["replylastuserid"] = topicInfo.ReplyLastUserId

		resultMap = append(resultMap, topicsMap)

	}

	return resultMap, nil
}

func CreateTopic(title, content string) (string, error) {
	// createTopicChan <- true
	// defer func() {
	// 	<-createTopicChan
	// }()

	var count int64
	err := new(models.Topic).GetTopic().Where("title = ?", title).Order("-id").Count(&count).Error
	if err != nil {
		return "", errors.New("error when query topic data in database:" + err.Error())
	}

	if count > 0 {
		return "", errors.New("topic name is exist!")
	}
	topicInfo := new(models.Topic)
	// topicInfo.Uid = uid
	topicInfo.Title = strings.TrimSpace(title)
	topicInfo.Content = strings.TrimSpace(content)
	topicInfo.Created = time.Now().Format("2006-01-02 15:04:05")
	topicInfo.Updated = time.Now().Format("2006-01-02 15:04:05")

	err = topicInfo.GetTopic().Save(topicInfo).Error
	if err != nil {
		return "", errors.New("error when save topic info:" + err.Error())
	}

	return "create new topic success", nil
}

func GetTopic(topicId int64) (map[string]interface{}, error) {
	// resultMap := make(map[string]interface{})
	topicInfo := new(models.Topic)

	err := topicInfo.GetTopic().Where("id = ?", topicId).First(&topicInfo).Error
	if err != nil {
		return nil, errors.New("error when get component info from db:" + err.Error())
	}

	topicsMap := make(map[string]interface{})
	topicsMap["id"] = topicInfo.ID
	topicsMap["uid"] = topicInfo.Uid
	topicsMap["title"] = topicInfo.Title
	topicsMap["content"] = topicInfo.Content
	topicsMap["attachment"] = topicInfo.Attachment
	topicsMap["created"] = topicInfo.Created
	topicsMap["updated"] = topicInfo.Updated
	topicsMap["views"] = topicInfo.Views
	topicsMap["author"] = topicInfo.Author
	topicsMap["replytime"] = topicInfo.ReplyTime
	topicsMap["replycount"] = topicInfo.ReplyCount
	topicsMap["replylastuserid"] = topicInfo.ReplyLastUserId

	return topicsMap, nil
}

func UpdateTopic(topicInfo models.Topic) error {
	return topicInfo.GetTopic().Save(&topicInfo).Error
}

func DeleteTopic(topicID int64) error {
	if topicID == 0 {
		return errors.New("topic id can't be zero")
	}

	topicInfo := new(models.Topic)
	err := topicInfo.GetTopic().Where("id = ?", topicID).First(topicInfo).Error
	if err != nil {
		return err
	}

	return topicInfo.GetTopic().Delete(&topicInfo).Error
}

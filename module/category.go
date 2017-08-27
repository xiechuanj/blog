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
	"fmt"
	"strings"
	"time"
)

func AddCategory(title string) error {
	var count int64
	err := new(models.Category).GetCategory().Where("title = ?", title).Order("-id").Count(&count).Error
	fmt.Println(count)
	if err != nil {
		return errors.New("error when query category data in database:" + err.Error())
	}

	if count > 0 {
		return errors.New("category name is exist!")
	}
	categoryInfo := new(models.Category)
	categoryInfo.Title = strings.TrimSpace(title)
	categoryInfo.Created = time.Now().Format("2006-01-02 15:04:05")

	err = categoryInfo.GetCategory().Save(categoryInfo).Error
	if err != nil {
		return errors.New("error when save category info:" + err.Error())
	}

	return nil
}

func GetAllCategories() ([]map[string]interface{}, error) {
	resultMap := make([]map[string]interface{}, 0)
	categoryList := make([]models.Category, 0)

	err := new(models.Category).GetCategory().Order("-id").Find(&categoryList).Error
	if err != nil {
		return nil, errors.New("error when get category infos error:" + err.Error())
	}

	for _, categoryInfo := range categoryList {
		categoryMap := make(map[string]interface{})
		categoryMap["id"] = categoryInfo.Id
		categoryMap["title"] = categoryInfo.Title
		categoryMap["created"] = categoryInfo.Created
		categoryMap["views"] = categoryInfo.Views
		categoryMap["topictime"] = categoryInfo.TopicTime
		categoryMap["topiccount"] = categoryInfo.TopicCount
		categoryMap["topiclastuserid"] = categoryInfo.TopicLastUserId

		resultMap = append(resultMap, categoryMap)

	}

	return resultMap, nil

}

func DeleteCategoryById(id int64) error {
	if id == 0 {
		return errors.New("category id can't be zero")
	}

	categoryInfo := new(models.Category)
	err := categoryInfo.GetCategory().Where("id = ?", id).First(categoryInfo).Error
	if err != nil {
		return err
	}

	return categoryInfo.GetCategory().Delete(&categoryInfo).Error
}

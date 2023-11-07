/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 10:40

*
*/
package main

import (
	"encoding/json"
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/article"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk().WithConfig(sdk.DefaultConfig(AccessKeyId, AccessKeySecret, []string{Endpoint})).AutoAuth().InitCloudC().InitEms().InitOss()
	s = s.AutoAuth().InitArticle()
	s.ArticleCheckStatus()
	info := &article.EditArticleParams{
		Title:     "小蝌蚪找妈妈2",
		Subtitle:  "小蝌蚪",
		Introduce: "《小蝌蚪找妈妈》根据方慧珍、盛璐德创作的同名童话改编，取材于画家齐白石创作的鱼虾等形象",
		Cover:     "",
		Content:   "青蛙妈妈下的卵慢慢地都活动起来，变成一群大脑袋长尾巴的蝌蚪，他们在水里游来游去，非常快乐。 有一天，鸭妈妈带着她的孩子到池塘中来游水。小蝌蚪看见小鸭子跟着妈妈在 ...",
		WriteId:   1,
		Author:    "taco",
		Label:     []string{"儿童故事"},
		Type:      1,
	}
	result, err := s.AddArticle(info)
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)

	articleInfo := struct {
		ID int64 `json:"ID"`
	}{}
	err = json.Unmarshal([]byte(result.Data.Data), &articleInfo)
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下文章:%+v", info)

	result, err = s.ListArticle(&article.ListArticleParams{
		PageSize: 10,
		PageNum:  1,
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)
	info.Id = articleInfo.ID
	info.Author = "test"
	result, err = s.EditArticle(info)
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)

	result, err = s.ArticleInfo(&article.ArticleInfoParams{Id: info.Id})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)
	result, err = s.DeleteArticle(&article.DeleteParams{Id: info.Id})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)

	result, err = s.ListLabel(&article.ListLabelParams{
		PageSize: 10,
		PageNum:  1,
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)

	result, err = s.UpdateLabel(&article.UpdateLabelParams{
		Id:   47,
		Name: "新闻趋势",
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)

	result, err = s.LabelInfo(&article.LabelInfoParams{Id: 47})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)

	result, err = s.DeleteLabel(&article.DeleteParams{Id: 47})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result)
	select {}
}

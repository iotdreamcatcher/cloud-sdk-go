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
	result1, err := s.ArticleAddArticle(info)
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result1)

	articleInfo := struct {
		ID int64 `json:"ID"`
	}{}
	err = json.Unmarshal([]byte(result1.Data.String()), &articleInfo)
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下文章:%+v", info)

	result2, err := s.ArticleListArticle(&article.ListArticleParams{
		PageSize: 10,
		PageNum:  1,
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result2)
	info.Id = articleInfo.ID
	info.Author = "test"
	result3, err := s.ArticleEditArticle(info)
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result3)

	result4, err := s.ArticleArticleInfo(&article.ArticleInfoParams{Id: info.Id})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result4)
	result5, err := s.ArticleDeleteArticle(&article.DeleteParams{Id: info.Id})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result5)

	result6, err := s.ArticleListLabel(&article.ListLabelParams{
		PageSize: 10,
		PageNum:  1,
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result6)

	result7, err := s.ArticleUpdateLabel(&article.UpdateLabelParams{
		Id:   47,
		Name: "新闻趋势",
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result7)

	result8, err := s.ArticleLabelInfo(&article.LabelInfoParams{Id: 47})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result8)

	result9, err := s.ArticleDeleteLabel(&article.DeleteParams{Id: 47})
	if err != nil {
		panic(err)
	}
	logx.Infof("打印一下请求的结果:%+v", result9)
	select {}
}

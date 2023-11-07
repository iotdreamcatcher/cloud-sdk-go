/*
*

	@author: taco
	@Date: 2023/10/27
	@Time: 10:29

*
*/
package sdk

import (
	"github.com/iotdreamcatcher/cloud-sdk-go/pb/article"
	"github.com/iotdreamcatcher/cloud-sdk-go/sdk/types"
	"github.com/zeromicro/go-zero/zrpc"
)

type Article struct {
	article.ArticleRpcServiceClient
	Num    int64
	Status int
	Retry  int
}

func NewArticle(RpcClientConf *zrpc.RpcClientConf) *Article {
	if len(RpcClientConf.Endpoints) == 0 {
		RpcClientConf.Endpoints = []string{
			"localhost:8080",
		}
	}
	client := article.NewArticleRpcServiceClient(zrpc.MustNewClient(*RpcClientConf).Conn())

	return &Article{
		ArticleRpcServiceClient: client,
		Num:                     1,
	}
}

func (c *Sdk) InitArticle() *Sdk {
	c.Article = NewArticle(c.Config.RpcClientConf)
	c.Article.Status = types.STATUS_READY
	c.Article.Retry += 1
	return c
}

func (c *Sdk) ArticleCheckStatus() *Sdk {
	if c.Article.Status != types.STATUS_READY {
		if c.Article.Retry >= c.Config.MaxRetryTimes {
			panic(types.ErrMaxErrTimes)
		} else {
			c.InitArticle()
		}
	}
	return c
}

func (c *Sdk) ListArticle(in *article.ListArticleParams) (*article.ArticleResp, error) {
	res, err := c.Article.ListArticle(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ArticleInfo(in *article.ArticleInfoParams) (*article.ArticleResp, error) {
	res, err := c.Article.ArticleInfo(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) AddArticle(in *article.EditArticleParams) (*article.ArticleResp, error) {
	res, err := c.Article.AddArticle(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) EditArticle(in *article.EditArticleParams) (*article.ArticleResp, error) {
	res, err := c.Article.EditArticle(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) DeleteArticle(in *article.DeleteParams) (*article.ArticleResp, error) {
	res, err := c.Article.DeleteArticle(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ReviewArticle(in *article.ReviewArticleParams) (*article.ArticleResp, error) {
	res, err := c.Article.ReviewArticle(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) ListLabel(in *article.ListLabelParams) (*article.ArticleResp, error) {
	res, err := c.Article.ListLabel(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) LabelInfo(in *article.LabelInfoParams) (*article.ArticleResp, error) {
	res, err := c.Article.LabelInfo(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Sdk) UpdateLabel(in *article.UpdateLabelParams) (*article.ArticleResp, error) {
	res, err := c.Article.UpdateLabel(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Sdk) DeleteLabel(in *article.DeleteParams) (*article.ArticleResp, error) {
	res, err := c.Article.DeleteLabel(c.SonyCtx(), in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

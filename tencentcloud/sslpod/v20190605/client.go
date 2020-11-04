// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190605

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-05"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "CreateDomain")
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过域名端口添加监控
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "DeleteDomain")
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过域名ID删除监控的域名
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardRequest() (request *DescribeDashboardRequest) {
    request = &DescribeDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeDashboard")
    return
}

func NewDescribeDashboardResponse() (response *DescribeDashboardResponse) {
    response = &DescribeDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取仪表盘数据
func (c *Client) DescribeDashboard(request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardRequest()
    }
    response = NewDescribeDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainCertsRequest() (request *DescribeDomainCertsRequest) {
    request = &DescribeDomainCertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeDomainCerts")
    return
}

func NewDescribeDomainCertsResponse() (response *DescribeDomainCertsResponse) {
    response = &DescribeDomainCertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取域名关联证书
func (c *Client) DescribeDomainCerts(request *DescribeDomainCertsRequest) (response *DescribeDomainCertsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainCertsRequest()
    }
    response = NewDescribeDomainCertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainTagsRequest() (request *DescribeDomainTagsRequest) {
    request = &DescribeDomainTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeDomainTags")
    return
}

func NewDescribeDomainTagsResponse() (response *DescribeDomainTagsResponse) {
    response = &DescribeDomainTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取账号下所有tag
func (c *Client) DescribeDomainTags(request *DescribeDomainTagsRequest) (response *DescribeDomainTagsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainTagsRequest()
    }
    response = NewDescribeDomainTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeDomains")
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过searchType搜索已经添加的域名
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNoticeInfoRequest() (request *DescribeNoticeInfoRequest) {
    request = &DescribeNoticeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "DescribeNoticeInfo")
    return
}

func NewDescribeNoticeInfoResponse() (response *DescribeNoticeInfoResponse) {
    response = &DescribeNoticeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取通知额度信息
func (c *Client) DescribeNoticeInfo(request *DescribeNoticeInfoRequest) (response *DescribeNoticeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeNoticeInfoRequest()
    }
    response = NewDescribeNoticeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainTagsRequest() (request *ModifyDomainTagsRequest) {
    request = &ModifyDomainTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "ModifyDomainTags")
    return
}

func NewModifyDomainTagsResponse() (response *ModifyDomainTagsResponse) {
    response = &ModifyDomainTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改域名tag
func (c *Client) ModifyDomainTags(request *ModifyDomainTagsRequest) (response *ModifyDomainTagsResponse, err error) {
    if request == nil {
        request = NewModifyDomainTagsRequest()
    }
    response = NewModifyDomainTagsResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshDomainRequest() (request *RefreshDomainRequest) {
    request = &RefreshDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "RefreshDomain")
    return
}

func NewRefreshDomainResponse() (response *RefreshDomainResponse) {
    response = &RefreshDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 强制重新检测域名
func (c *Client) RefreshDomain(request *RefreshDomainRequest) (response *RefreshDomainResponse, err error) {
    if request == nil {
        request = NewRefreshDomainRequest()
    }
    response = NewRefreshDomainResponse()
    err = c.Send(request, response)
    return
}

func NewResolveDomainRequest() (request *ResolveDomainRequest) {
    request = &ResolveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sslpod", APIVersion, "ResolveDomain")
    return
}

func NewResolveDomainResponse() (response *ResolveDomainResponse) {
    response = &ResolveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解析域名获得多个IP地址
func (c *Client) ResolveDomain(request *ResolveDomainRequest) (response *ResolveDomainResponse, err error) {
    if request == nil {
        request = NewResolveDomainRequest()
    }
    response = NewResolveDomainResponse()
    err = c.Send(request, response)
    return
}
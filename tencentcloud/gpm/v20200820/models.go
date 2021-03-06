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

package v20200820

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AttributeMap struct {

	// 属性字典 key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 属性字典 value
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type CancelMatchingRequest struct {
	*tchttp.BaseRequest

	// 匹配 Code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// 要取消的匹配请求 Id
	MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`
}

func (r *CancelMatchingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelMatchingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelMatchingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码
		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelMatchingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelMatchingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMatchRequest struct {
	*tchttp.BaseRequest

	// 匹配名称，[a-zA-Z0-9-\.]* 长度128
	MatchName *string `json:"MatchName,omitempty" name:"MatchName"`

	// 规则code
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// 超时时间，1-600秒
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// 是否为匹配结果请求服务器资源，0表示否，1表示请求GSE资源
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 匹配描述，最长1024
	MatchDesc *string `json:"MatchDesc,omitempty" name:"MatchDesc"`

	// 只支持https协议
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 游戏服务器队列地域
	ServerRegion *string `json:"ServerRegion,omitempty" name:"ServerRegion"`

	// 游戏服务器队列
	ServerQueue *string `json:"ServerQueue,omitempty" name:"ServerQueue"`

	// 自定义推送数据
	CustomPushData *string `json:"CustomPushData,omitempty" name:"CustomPushData"`

	// 游戏服务器会话数据
	ServerSessionData *string `json:"ServerSessionData,omitempty" name:"ServerSessionData"`

	// 游戏属性，key-value结构的数组
	GameProperties []*StringKV `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 日志开关，0表示关，1表示开
	LogSwitch *int64 `json:"LogSwitch,omitempty" name:"LogSwitch"`

	// 标签，key-value结构的数组
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateMatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配信息
		MatchInfo *MatchInfo `json:"MatchInfo,omitempty" name:"MatchInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称，[a-zA-Z0-9-\.]* 长度128
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则脚本，长度65535
	RuleScript *string `json:"RuleScript,omitempty" name:"RuleScript"`

	// 规则描述，最长1024
	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`

	// 标签，key-value结构的数组，最多关联50组标签
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则信息
		RuleInfo *RuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMatchRequest struct {
	*tchttp.BaseRequest

	// 匹配code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DeleteMatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest

	// 规则code
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDataRequest struct {
	*tchttp.BaseRequest

	// 起始时间，单位：秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间，单位：秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 时间粒度，1表示1天；2表示1小时；3表示1分钟；4表示10分钟；5表示30分钟
	TimeType *int64 `json:"TimeType,omitempty" name:"TimeType"`

	// 匹配code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DescribeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配概况
	// 注意：此字段可能返回 null，表示取不到有效值。
		OverviewData *ReportOverviewData `json:"OverviewData,omitempty" name:"OverviewData"`

		// 匹配请求次数趋势数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		TrendData *ReportTrendData `json:"TrendData,omitempty" name:"TrendData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchCodesRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMatchCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchCodesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchCodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配Code
	// 注意：此字段可能返回 null，表示取不到有效值。
		MatchCodes []*MatchCodeAttr `json:"MatchCodes,omitempty" name:"MatchCodes" list`

		// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMatchCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchCodesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchRequest struct {
	*tchttp.BaseRequest

	// 匹配code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DescribeMatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		MatchInfo *MatchInfo `json:"MatchInfo,omitempty" name:"MatchInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchesRequest struct {
	*tchttp.BaseRequest

	// 当前页号，不传则获取所有有权限的资源。
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 单页大小，不传则获取所有有权限的资源。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询类型（可选）：match表示通过matchCode或者matchName来搜索，rule表示通过ruleCode或者ruleName来搜索，其余类型不做过滤处理。
	SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

	// 查询关键词，针对SearchType进行具体过滤的内容。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 标签列表，用于过滤。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *DescribeMatchesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		MatchInfoList []*MatchInfo `json:"MatchInfoList,omitempty" name:"MatchInfoList" list`

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当前页号
		PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

		// 单页大小
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// 查询类型（可选）：matchName表示匹配名称，matchCode表示匹配code，ruleName表示规则名称，tag表示标签Key/Value
		SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

		// 查询关键词（可选）
		Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMatchesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchingProgressRequest struct {
	*tchttp.BaseRequest

	// 请求 id 列表, 列表长度 128。
	MatchTicketIds []*MTicket `json:"MatchTicketIds,omitempty" name:"MatchTicketIds" list`
}

func (r *DescribeMatchingProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchingProgressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMatchingProgressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		MatchTickets []*MatchTicket `json:"MatchTickets,omitempty" name:"MatchTickets" list`

		// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMatchingProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMatchingProgressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleRequest struct {
	*tchttp.BaseRequest

	// 规则code
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`
}

func (r *DescribeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		RuleInfo *RuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest

	// 当前页号，不传则获取所有有权限的资源。
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 单页大小，不传则获取所有有权限的资源。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询类型（可选）：match表示通过matchCode或者matchName来搜索，rule表示通过ruleCode或者ruleName来搜索，其余类型不做过滤处理。
	SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

	// 查询关键词，针对SearchType进行具体过滤的内容。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 标签列表，用于过滤。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		RuleInfoList []*RuleBriefInfo `json:"RuleInfoList,omitempty" name:"RuleInfoList" list`

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当前页号
		PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

		// 单页大小
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// 查询类型（可选）matchName表示匹配名称，matchCode表示匹配code，ruleName表示规则名称，tag表示标签Key/Value
		SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

		// 查询关键词（可选）
		Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTokenRequest struct {
	*tchttp.BaseRequest

	// 匹配code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

func (r *DescribeTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前的MatchCode对应的Token。如果当前MatchCode没有Token，该参数可能取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		MatchToken *string `json:"MatchToken,omitempty" name:"MatchToken"`

		// 当Token被替换后，GPM将兼容推送原始Token的时间（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CompatibleSpan *uint64 `json:"CompatibleSpan,omitempty" name:"CompatibleSpan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MTicket struct {

	// 匹配Code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// 请求TicketId
	MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`
}

type MatchAttribute struct {

	// 属性名 长度 128
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性类型: 0 数值; 1 string; 默认 0
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 数字属性值 默认 0.0
	NumberValue *float64 `json:"NumberValue,omitempty" name:"NumberValue"`

	// 字符串属性值 长度 1024 默认 ""
	StringValue *string `json:"StringValue,omitempty" name:"StringValue"`

	// list 属性值
	ListValue []*string `json:"ListValue,omitempty" name:"ListValue" list`

	// 字典属性值
	MapValue []*AttributeMap `json:"MapValue,omitempty" name:"MapValue" list`
}

type MatchCodeAttr struct {

	// 匹配code
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`
}

type MatchInfo struct {

	// 匹配code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// 匹配名称
	MatchName *string `json:"MatchName,omitempty" name:"MatchName"`

	// 匹配描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchDesc *string `json:"MatchDesc,omitempty" name:"MatchDesc"`

	// 规则code
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 超时时间
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// 接收通知地址
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 是否为匹配结果请求服务器资源，0否，1请求GSE资源
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 服务器队列所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerRegion *string `json:"ServerRegion,omitempty" name:"ServerRegion"`

	// 服务器队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerQueue *string `json:"ServerQueue,omitempty" name:"ServerQueue"`

	// 自定义推送数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomPushData *string `json:"CustomPushData,omitempty" name:"CustomPushData"`

	// 游戏服务器会话数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerSessionData *string `json:"ServerSessionData,omitempty" name:"ServerSessionData"`

	// 游戏属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameProperties []*StringKV `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 日志开关，0表示关，1表示开
	LogSwitch *int64 `json:"LogSwitch,omitempty" name:"LogSwitch"`

	// 日志集id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志主题id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`

	// 日志主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicName *string `json:"LogTopicName,omitempty" name:"LogTopicName"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags" list`

	// 地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 用户主账号Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 用户创建账号Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 日志状态，0表示正常，1表示日志集不存在，2表示日志主题不存在，3表示日志集和日志主题都不存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogStatus *int64 `json:"LogStatus,omitempty" name:"LogStatus"`
}

type MatchTicket struct {

	// MatchTicketId 长度 128
	Id *string `json:"Id,omitempty" name:"Id"`

	// 匹配 Code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// 根据 MatchType 取不同的结构序列化结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchResult *string `json:"MatchResult,omitempty" name:"MatchResult"`

	// 表示不同的匹配类型,NORMAL | GSE
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 玩家信息列表
	Players []*Player `json:"Players,omitempty" name:"Players" list`

	// 匹配状态: SEARCHING 匹配中; PLACING 匹配放置中; COMPLETED 匹配完成; CANCELLED 匹配取消; TIMEDOUT 匹配超时; FAILED 匹配失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 匹配状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`

	// 匹配状态原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusReason *string `json:"StatusReason,omitempty" name:"StatusReason"`

	// 收到发起匹配请求的时间 eg: "2020-08-17T08:14:38.077Z"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 匹配请求因完成、失败、超时、被取消而停止执行的时间 eg: "2020-08-17T08:14:38.077Z"
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ModifyMatchRequest struct {
	*tchttp.BaseRequest

	// 匹配名称，[a-zA-Z0-9-\.]* 长度128
	MatchName *string `json:"MatchName,omitempty" name:"MatchName"`

	// 规则code
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// 超时时间，1-600秒
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// 是否为匹配结果请求服务器资源，0表示否，1表示请求GSE资源
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 匹配code
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// 匹配描述，最长1024
	MatchDesc *string `json:"MatchDesc,omitempty" name:"MatchDesc"`

	// 只支持https协议
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 游戏服务器队列地域
	ServerRegion *string `json:"ServerRegion,omitempty" name:"ServerRegion"`

	// 游戏服务器队列
	ServerQueue *string `json:"ServerQueue,omitempty" name:"ServerQueue"`

	// 自定义推送数据
	CustomPushData *string `json:"CustomPushData,omitempty" name:"CustomPushData"`

	// 游戏服务器会话数据
	ServerSessionData *string `json:"ServerSessionData,omitempty" name:"ServerSessionData"`

	// 游戏属性，key-value结构的数组
	GameProperties []*StringKV `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 日志开关，0表示关，1表示开
	LogSwitch *int64 `json:"LogSwitch,omitempty" name:"LogSwitch"`

	// 标签，key-value结构的数组
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags" list`
}

func (r *ModifyMatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 匹配信息
		MatchInfo *MatchInfo `json:"MatchInfo,omitempty" name:"MatchInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest

	// 规则code
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则描述，最长1024
	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`

	// 标签，key-value结构的数组，最多关联50组标签
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags" list`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则信息
		RuleInfo *RuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTokenRequest struct {
	*tchttp.BaseRequest

	// 匹配Code。
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// 单位秒，取值0-1800。此参数表示当前Token被替换后，GPM将持续推送原Token的时间。在CompatibleSpan时间范围内，用户将在事件消息中收到当前和原始Token。
	CompatibleSpan *uint64 `json:"CompatibleSpan,omitempty" name:"CompatibleSpan"`

	// Token，[a-zA-Z0-9-_.], 长度0-64。如果为空，将由GPM随机生成。
	MatchToken *string `json:"MatchToken,omitempty" name:"MatchToken"`
}

func (r *ModifyTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功设置的Token。
		MatchToken *string `json:"MatchToken,omitempty" name:"MatchToken"`

		// 当前Token被替换后，GPM将持续推送原Token的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CompatibleSpan *uint64 `json:"CompatibleSpan,omitempty" name:"CompatibleSpan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Player struct {

	// 玩家 PlayerId 长度 128
	Id *string `json:"Id,omitempty" name:"Id"`

	// 玩家昵称，长度 128
	Name *string `json:"Name,omitempty" name:"Name"`

	// 玩家匹配属性，最多 10 条
	MatchAttributes []*MatchAttribute `json:"MatchAttributes,omitempty" name:"MatchAttributes" list`

	// 队伍名，可以传递不同队伍名，长度 128
	Team *string `json:"Team,omitempty" name:"Team"`

	// 自定义玩家状态 透传参数 [0, 99999]
	CustomPlayerStatus *uint64 `json:"CustomPlayerStatus,omitempty" name:"CustomPlayerStatus"`

	// 自定义玩家信息 透传参数 长度 1024
	CustomProfile *string `json:"CustomProfile,omitempty" name:"CustomProfile"`

	// 各区域延迟，最多 20 条
	RegionLatencies []*RegionLatency `json:"RegionLatencies,omitempty" name:"RegionLatencies" list`
}

type RegionLatency struct {

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 毫秒延迟 0～999999
	Latency *uint64 `json:"Latency,omitempty" name:"Latency"`
}

type ReportOverviewData struct {

	// 总次数
	TotalTimes *string `json:"TotalTimes,omitempty" name:"TotalTimes"`

	// 成功率
	SuccessPercent *float64 `json:"SuccessPercent,omitempty" name:"SuccessPercent"`

	// 超时率
	TimeoutPercent *float64 `json:"TimeoutPercent,omitempty" name:"TimeoutPercent"`

	// 失败率
	FailPercent *float64 `json:"FailPercent,omitempty" name:"FailPercent"`

	// 平均匹配时间
	AverageSec *float64 `json:"AverageSec,omitempty" name:"AverageSec"`
}

type ReportTrendData struct {

	// 总次数
	TotalList []*string `json:"TotalList,omitempty" name:"TotalList" list`

	// 被取消次数
	CancelList []*string `json:"CancelList,omitempty" name:"CancelList" list`

	// 成功次数
	SuccessList []*string `json:"SuccessList,omitempty" name:"SuccessList" list`

	// 失败次数
	FailList []*string `json:"FailList,omitempty" name:"FailList" list`

	// 超时次数
	TimeoutList []*string `json:"TimeoutList,omitempty" name:"TimeoutList" list`

	// 时间数组，单位：秒
	TimeList []*string `json:"TimeList,omitempty" name:"TimeList" list`
}

type RuleBriefInfo struct {

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 关联匹配
	MatchCodeList []*StringKV `json:"MatchCodeList,omitempty" name:"MatchCodeList" list`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则code
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`
}

type RuleInfo struct {

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDesc *string `json:"RuleDesc,omitempty" name:"RuleDesc"`

	// 规则脚本
	RuleScript *string `json:"RuleScript,omitempty" name:"RuleScript"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*StringKV `json:"Tags,omitempty" name:"Tags" list`

	// 关联匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchCodeList []*StringKV `json:"MatchCodeList,omitempty" name:"MatchCodeList" list`

	// 规则code
	RuleCode *string `json:"RuleCode,omitempty" name:"RuleCode"`

	// 地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 用户OwnerUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
}

type StartMatchingRequest struct {
	*tchttp.BaseRequest

	// 匹配 Code。
	MatchCode *string `json:"MatchCode,omitempty" name:"MatchCode"`

	// 玩家信息 最多 200 条。
	Players []*Player `json:"Players,omitempty" name:"Players" list`

	// 请求 Id 默认 "" 为空则由 GPM 自动生成 长度 128。
	MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`
}

func (r *StartMatchingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMatchingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartMatchingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码。
		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

		// 请求 id 长度 128。
		MatchTicketId *string `json:"MatchTicketId,omitempty" name:"MatchTicketId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMatchingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMatchingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StringKV struct {

	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Tag struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

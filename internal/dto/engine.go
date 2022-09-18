package dto

import ()

/**
 * dsl run request
 * url: /engine/run
 * example: {"key":"flow_abtest", "req_id":"123456", "uid":1,"features":{"feature_1":5,"feature_2":3,"feature_3":true}}
 */
type EngineRunRequest struct {
	Key      string                 `json:"key"`
	Version  string                 `json:"version"`
	ReqId    string                 `json:"req_id"`
	Uid      int64                  `json:"uid"`
	Features map[string]interface{} `json:"features"`
}

/**
 * dsl run response
 * url: /engine/run
 */
type EngineRunResponse struct {
	Key         string                   `json:"key"`
	ReqId       string                   `json:"req_id"`
	Uid         int64                    `json:"uid"`
	Features    []map[string]interface{} `json:"features"`
	Tracks      []map[string]interface{} `json:"tracks"`
	HitRules    []map[string]interface{} `json:"hit_rules"`
	NodeResults []map[string]interface{} `json:"node_results"`
	StartTime   string                   `json:"start_time"`
	EndTime     string                   `json:"end_time"`
	RunTime     int64                    `json:"run_time"`
}

/**
 * dsl list response
 * url: /engine/list
 */
type DslListResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
	Data []*Dsl `json:"data"`
}

type Dsl struct {
	Key     string `json:"key"`
	Version string `json:"version"`
	//Metadata string `json:"metadata"`
	Md5 string `json:"md5"`
}

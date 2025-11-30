package dao

import (
	"context"
	"encoding/json"

	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/model"
	"github.com/redis/go-redis/v9"
)

// buildSystemConfigKey 构建系统配置键
func buildSystemConfigKey() string {
	return "system:configs"
}

// buildAiConfigKey 构建ai配置项的key
func buildAiConfigKey(id string) string {
	return "system:ai:configs:" + id
}

type SystemDao struct {
	ctx   context.Context
	group string
	rdb   *redis.Client
}

// NewSystemDao 创建系统数据访问对象
func NewSystemDao(ctx context.Context, group string, rdb *redis.Client) *SystemDao {
	return &SystemDao{
		ctx:   ctx,
		group: group,
		rdb:   rdb,
	}
}

// 添加ai的配置项
func (s *SystemDao) AddAiConfigs(ctx context.Context, aiConfig *model.AiConfig) error {
	jsonData, err := json.Marshal(aiConfig)
	if nil != err {
		logx.Errorf("SystemDao|AddAiConfigs|Marshal|Error|%v|%s", err, aiConfig.ID)
		return err
	}
	err = s.rdb.Set(ctx, buildAiConfigKey(aiConfig.ID), jsonData, 0).Err()
	if nil != err {
		logx.Errorf("SystemDao|AddAiConfigs|Set|Error|%v|%s", err, string(jsonData))
		return err
	}
	ids, err := s.rdb.HGet(ctx, buildSystemConfigKey(), aiConfig.ID).Result()
	if nil != err {
		logx.Errorf("SystemDao|AddAiConfigs|HGet|Error|%v|%s", err, aiConfig.ID)
		return err
	}
	if len(ids) == 0 {
		jsonData, err := json.Marshal([]string{aiConfig.ID})
		if nil != err {
			logx.Errorf("SystemDao|AddAiConfigs|Marshal|Error|%v|%s", err, aiConfig.ID)
			return err
		}
		err = s.rdb.HSet(ctx, buildSystemConfigKey(), aiConfig.ID, jsonData).Err()
		if nil != err {
			logx.Errorf("SystemDao|AddAiConfigs|HSet|Error|%v|%s", err, string(jsonData))
			return err
		}
	} else {
		var Aids []string
		err := json.Unmarshal([]byte(ids), &Aids)
		if nil != err {
			logx.Errorf("SystemDao|AddAiConfigs|Unmarshal|Error|%v|%s", err, ids)
			return err
		}
		Aids = append(Aids, aiConfig.ID)
		jsonData, err := json.Marshal(Aids)
		if nil != err {
			logx.Errorf("SystemDao|AddAiConfigs|Marshal|Error|%v|%s", err, Aids)
			return err
		}
		err = s.rdb.HSet(ctx, buildSystemConfigKey(), aiConfig.ID, jsonData).Err()
		if nil != err {
			logx.Errorf("SystemDao|AddAiConfigs|HSet|Error|%v|%s", err, string(jsonData))
			return err
		}
	}
	return nil
}

package components

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/eino/components"
	"github.com/cloudwego/eino/schema"
	"github.com/zhuiye8/Lyss-chat-server/internal/domain/model"
	"github.com/zhuiye8/Lyss-chat-server/pkg/logger"
)

// ModelProvider 表示模型提供商接�?
type ModelProvider interface {
	GetModel(ctx context.Context, modelID string) (*model.Model, error)
	GetAPIKey(ctx context.Context, providerID string) (*model.APIKey, error)
	CallModel(ctx context.Context, modelID string, messages []*schema.Message) (*schema.Message, error)
	StreamModel(ctx context.Context, modelID string, messages []*schema.Message) (<-chan *schema.Message, error)
}

// ChatModelAdapter �?Eino ChatModel 组件的适配�?
type ChatModelAdapter struct {
	provider ModelProvider
	modelID  string
	logger   *logger.Logger
}

// NewChatModelAdapter 创建一个新�?ChatModel 适配�?
func NewChatModelAdapter(provider ModelProvider, modelID string, logger *logger.Logger) *ChatModelAdapter {
	return &ChatModelAdapter{
		provider: provider,
		modelID:  modelID,
		logger:   logger,
	}
}

// Call 实现 ChatModel 接口�?Call 方法
func (a *ChatModelAdapter) Call(ctx context.Context, messages []*schema.Message) (*schema.Message, error) {
	a.logger.Debug("调用模型", a.modelID)
	return a.provider.CallModel(ctx, a.modelID, messages)
}

// Stream 实现 ChatModel 接口�?Stream 方法
func (a *ChatModelAdapter) Stream(ctx context.Context, messages []*schema.Message) (<-chan *schema.Message, error) {
	a.logger.Debug("流式调用模型", a.modelID)
	return a.provider.StreamModel(ctx, a.modelID, messages)
}

// Ensure ChatModelAdapter implements components.ChatModel
var _ components.ChatModel = (*ChatModelAdapter)(nil)


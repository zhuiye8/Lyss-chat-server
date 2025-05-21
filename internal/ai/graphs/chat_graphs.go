package graphs

import (
	"context"

	"github.com/zhuiye8/Lyss-chat-server/internal/ai/components"
	"github.com/zhuiye8/Lyss-chat-server/internal/service/model"
	"github.com/zhuiye8/Lyss-chat-server/pkg/logger"
)

// ChatGraphs 包含所有聊天图�?
type ChatGraphs struct {
	Chat *ChatGraph
}

// NewChatGraphs 创建一个新的聊天图形集�?
func NewChatGraphs(ctx context.Context, modelService *model.Service, logger *logger.Logger) (*ChatGraphs, error) {
	// 创建模型提供商适配�?
	modelProvider := &components.ModelServiceAdapter{
		ModelService: modelService,
	}

	// 创建聊天模型适配�?
	chatModel := components.NewChatModelAdapter(modelProvider, "", logger)

	// 创建聊天图形
	chatGraph, err := NewChatGraph(ctx, *chatModel, logger)
	if err != nil {
		return nil, err
	}

	return &ChatGraphs{
		Chat: chatGraph,
	}, nil
}


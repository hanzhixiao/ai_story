package chat

import (
	"grandma/backend/models"
	"grandma/backend/repository"
	"grandma/backend/services"
	"grandma/backend/utils"
	"io"
	"strings"
)

type ChatService struct {
	conversationRepo *repository.ConversationRepository
	documentRepo     *repository.DocumentRepository
	config           *ChatConfig
}

type ChatConfig struct {
	OpenAIAPIKey     string
	OpenAIBaseURL    string
	AnthropicAPIKey  string
	AnthropicBaseURL string
}

func NewChatService(conversationRepo *repository.ConversationRepository, documentRepo *repository.DocumentRepository, config *ChatConfig) *ChatService {
	return &ChatService{
		conversationRepo: conversationRepo,
		documentRepo:     documentRepo,
		config:           config,
	}
}

// SendMessage 发送消息并获取流式响应
func (s *ChatService) SendMessage(req *models.ChatRequest, writer io.Writer) (string, string, error) {
	var conversationID string
	var conversation *models.Conversation
	var err error

	// 如果没有提供对话ID，创建新对话
	if req.ConversationID == "" {
		conversationID = utils.GenerateConversationID()
		conversation = &models.Conversation{
			ID:    conversationID,
			Title: s.generateTitle(req.Message),
		}
		err = s.conversationRepo.Create(conversation)
		if err != nil {
			return "", "", err
		}
	} else {
		conversationID = req.ConversationID
		conversation, err = s.conversationRepo.GetByID(conversationID)
		if err != nil {
			return "", "", err
		}
	}

	// 创建用户消息文档
	userDocID := utils.GenerateDocumentID()
	userDoc := &models.Document{
		ID:             userDocID,
		ConversationID: conversationID,
		Role:           "user",
		Content:        req.Message,
		Model:          req.Model,
	}
	err = s.documentRepo.Create(userDoc)
	if err != nil {
		return "", "", err
	}

	// 调用大模型API获取流式响应
	provider, err := services.GetProvider(
		req.Model,
		s.config.OpenAIAPIKey,
		s.config.OpenAIBaseURL,
		s.config.AnthropicAPIKey,
		s.config.AnthropicBaseURL,
	)
	if err != nil {
		return "", "", err
	}

	// 创建流式响应收集器
	responseCollector := &responseCollector{writer: writer}
	err = provider.ChatStream(req.Message, responseCollector)
	if err != nil {
		return "", "", err
	}

	// 创建助手消息文档
	assistantDocID := utils.GenerateDocumentID()
	assistantDoc := &models.Document{
		ID:             assistantDocID,
		ConversationID: conversationID,
		Role:           "assistant",
		Content:        responseCollector.content,
		Model:          req.Model,
	}
	err = s.documentRepo.Create(assistantDoc)
	if err != nil {
		return "", "", err
	}

	return conversationID, assistantDocID, nil
}

// generateTitle 从消息内容生成对话标题
func (s *ChatService) generateTitle(message string) string {
	title := strings.TrimSpace(message)
	if len(title) > 50 {
		title = title[:50] + "..."
	}
	if title == "" {
		title = "新对话"
	}
	return title
}

// responseCollector 收集流式响应内容
type responseCollector struct {
	writer  io.Writer
	content string
}

func (rc *responseCollector) Write(p []byte) (n int, err error) {
	n, err = rc.writer.Write(p)
	if err == nil {
		rc.content += string(p)
	}
	return n, err
}

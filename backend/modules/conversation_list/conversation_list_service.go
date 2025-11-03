package conversation_list

import (
	"grandma/backend/models"
	"grandma/backend/repository"
)

type ConversationListService struct {
	conversationRepo *repository.ConversationRepository
}

func NewConversationListService(conversationRepo *repository.ConversationRepository) *ConversationListService {
	return &ConversationListService{
		conversationRepo: conversationRepo,
	}
}

// GetConversationList 获取对话列表
func (s *ConversationListService) GetConversationList(page, pageSize int) (*models.ConversationListResponse, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	conversations, total, err := s.conversationRepo.List(page, pageSize)
	if err != nil {
		return nil, err
	}

	return &models.ConversationListResponse{
		Conversations: conversations,
		Total:         int(total),
		Page:          page,
		PageSize:      pageSize,
	}, nil
}


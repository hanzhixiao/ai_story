package story

import (
	"errors"
	"grandma/backend/models"
	"grandma/backend/repository"
	"grandma/backend/utils"
	"time"
)

type StoryService struct {
	storyRepo *repository.StoryRepository
}

func NewStoryService(storyRepo *repository.StoryRepository) *StoryService {
	return &StoryService{
		storyRepo: storyRepo,
	}
}

// GetStoryList 获取文档列表
func (s *StoryService) GetStoryList(guid string) (*models.StoryResponse, error) {
	var stories []models.Story
	var err error

	// 如果guid为空或为"default"，获取所有故事
	if guid == "" || guid == "default" {
		stories, err = s.storyRepo.GetAll()
	} else {
		stories, err = s.storyRepo.GetByGuid(guid)
	}

	if err != nil {
		return nil, err
	}

	return &models.StoryResponse{
		Story: stories,
		Total: len(stories),
	}, nil
}

// DeleteStory 删除文档
func (s *StoryService) DeleteStory(id string) error {
	return s.storyRepo.Delete(id)
}

// CreateStory 创建故事
func (s *StoryService) CreateStory(guid, documentID, title, content, clientContentHash string) (*models.Story, error) {
	// 如果guid为空，使用默认值"default"
	if guid == "" {
		guid = "default"
	}

	// 计算服务端内容特征值
	serverContentHash := utils.CalculateContentHash(content)

	// 验证客户端和服务端的特征值是否一致
	if clientContentHash != "" && clientContentHash != serverContentHash {
		return nil, errors.New("hash_mismatch")
	}

	// 检查是否已存在相同特征值的故事
	existingStory, err := s.storyRepo.GetByContentHash(guid, serverContentHash)
	if err != nil {
		// 如果查询出错，返回错误
		return nil, err
	}
	if existingStory != nil {
		// 故事已存在，返回重复错误
		return nil, errors.New("duplicate_story")
	}

	// 创建新故事
	story := &models.Story{
		ID:          utils.GenerateStoryId(),
		Title:       title,
		Guid:        guid,
		DocumentID:  documentID,
		Content:     content,
		ContentHash: serverContentHash,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = s.storyRepo.Create(story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

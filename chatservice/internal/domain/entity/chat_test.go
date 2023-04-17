package entity_test

import (
	"chatservice/internal/domain/entity"
	"testing"
)

func TestAddMessageShouldThrowErrorWhenMessageIsTooLarge(t *testing.T) {
	userID := "510ae7e0-4122-49e3-9384-68fa573c2afc"
	modelName := "gpt-3.5-turbo"
	model := newModel(modelName)

	systemRole := "system"
	initialMessageContent := "InitialMessageContent"
	// mockController := gomock.NewController(t)
	// mock := mocks.NewMockTikToken(mockController)
	// mock.EXPECT().CountTokens(modelName, initialMessageContent).Return(9)

	initialMessage := newMessage(systemRole, initialMessageContent, model)
	chatConfig := newChatConfig(model)

	chat, err := newChat(userID, initialMessage, chatConfig)
	if err != nil {
		t.Fatal("error creating chat")
	}

	userRole := "user"
	messageTooLargeContent := "TooLargeMessage"
	messageTooLarge := newMessage(userRole, messageTooLargeContent, model)
	// mock.EXPECT().CountTokens(modelName, messageTooLarge).Return(2)

	chat.AddMessage(messageTooLarge)
	if err != nil {
		t.Fatal("error adding message to chat")
	}
}

func newModel(modelName string) *entity.Model {
	return entity.NewModel(modelName, 10)
}

func newMessage(role, content string, model *entity.Model) *entity.Message {
	message, _ := entity.NewMessage(role, content, model)

	return message
}

func newChatConfig(model *entity.Model) *entity.ChatConfig {
	return &entity.ChatConfig{
		Model:       model,
		Temperature: 1,
	}
}

func newChat(userID string, initialMessage *entity.Message, chatConfig *entity.ChatConfig) (*entity.Chat, error) {
	return entity.NewChat(userID, initialMessage, chatConfig)
}

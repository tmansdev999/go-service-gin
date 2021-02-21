package console

import (
	"go-service-gin/application"
	"go-service-gin/domain/blogs"
	"go-service-gin/util/faker"
	"go-service-gin/util/logger"
	"go-service-gin/util/uuid"
)

// BlogConsole is
type BlogConsole struct {
	blogLogic *application.BlogApplication
}

// NewBlogConsole is
func NewBlogConsole(ba *application.BlogApplication) *BlogConsole {
	if ba == nil {
		return nil
	}

	return &BlogConsole{ba}
}

// Create is
func (bh *BlogConsole) Create() bool {
	b := &blogs.Blog{
		ID:    uuid.Generate(),
		Title: faker.Sentence(),
		Body:  faker.Paragraph(),
	}

	if err := bh.blogLogic.Create(b); err != nil {
		logger.Info(err)
		return false
	}

	return true
}

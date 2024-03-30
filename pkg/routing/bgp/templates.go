package bgp

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type Template struct{}

type Templates interface {
	ListTemplates(input *ListTemplatesInput) (*ListTemplatesOutput, error)
	GetTemplate(input *GetTemplateInput) (*GetTemplateOutput, error)
	CreateTemplate(input *CreateTemplateInput) (*CreateTemplateOutput, error)
	UpdateTemplate(input *UpdateTemplateInput) (*UpdateTemplateOutput, error)
	DeleteTemplate(input *DeleteTemplateInput) (*DeleteTemplateOutput, error)
}

type ListTemplatesInput struct{}
type ListTemplatesOutput struct{}

type GetTemplateInput struct{}
type GetTemplateOutput struct{}

type CreateTemplateInput struct{}
type CreateTemplateOutput struct{}

type UpdateTemplateInput struct{}
type UpdateTemplateOutput struct{}

type DeleteTemplateInput struct{}
type DeleteTemplateOutput struct{}

type TemplatesImpl struct {
	Client *resty.Client
}

func (c TemplatesImpl) ListTemplates(input *ListTemplatesInput) (*ListTemplatesOutput, error) {
	return nil, errors.New("not implemented")
}

func (c TemplatesImpl) GetTemplate(input *GetTemplateInput) (*GetTemplateOutput, error) {
	return nil, errors.New("not implemented")
}

func (c TemplatesImpl) CreateTemplate(input *CreateTemplateInput) (*CreateTemplateOutput, error) {
	return nil, errors.New("not implemented")
}

func (c TemplatesImpl) UpdateTemplate(input *UpdateTemplateInput) (*UpdateTemplateOutput, error) {
	return nil, errors.New("not implemented")
}

func (c TemplatesImpl) DeleteTemplate(input *DeleteTemplateInput) (*DeleteTemplateOutput, error) {
	return nil, errors.New("not implemented")
}

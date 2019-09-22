package domain_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/domain"
	"github.com/steinfletcher/new-relic-concourse-deployment-resource/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var info = domain.DeploymentInfo{
	Revision:    "v0.1.1",
	Description: "Made some changes",
	User:        "a@b.com",
}

func TestRecordDeployment_WithID(t *testing.T) {
	tests := map[string]struct {
		readerErr   error
		writerErr   error
		expectedErr error
	}{
		"reader error": {readerErr: errors.New("reader err"), writerErr: nil, expectedErr: errors.New("reader err")},
		"writer error": {readerErr: nil, writerErr: errors.New("writer err"), expectedErr: errors.New("writer err")},
		"success":      {readerErr: nil, writerErr: nil, expectedErr: nil},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			reader := mocks.NewMockDeploymentInfoReader(ctrl)
			writer := mocks.NewMockDeploymentWriter(ctrl)
			reader.EXPECT().Read().Return(info, test.readerErr)
			writer.EXPECT().Write(info, "myAppID").Return(test.writerErr)

			_, err := domain.RecordDeployment(reader, writer, "myAppName", "myAppID")

			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestRecordDeployment_WithName(t *testing.T) {
	tests := map[string]struct {
		readerErr   error
		writerErr   error
		expectedErr error
	}{
		"reader error": {readerErr: errors.New("reader err"), writerErr: nil, expectedErr: errors.New("reader err")},
		"writer error": {readerErr: nil, writerErr: errors.New("writer err"), expectedErr: errors.New("writer err")},
		"success":      {readerErr: nil, writerErr: nil, expectedErr: nil},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			reader := mocks.NewMockDeploymentInfoReader(ctrl)
			writer := mocks.NewMockDeploymentWriter(ctrl)
			reader.EXPECT().Read().Return(info, test.readerErr)
			writer.EXPECT().WriteWithName(info, "myAppName").Return(test.writerErr)

			_, err := domain.RecordDeployment(reader, writer, "myAppName", "")

			assert.Equal(t, test.expectedErr, err)
		})
	}
}

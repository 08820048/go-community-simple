package service

import (
	"communityMini/repository"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	repository.Init("../data/")
	os.Exit(m.Run())
}

// 测试通过
func TestQueryPageInfo(t *testing.T) {
	pageInfo, _ := QueryPageInfo(1)
	assert.NotEqual(t, nil, pageInfo)
	assert.Equal(t, 5, len(pageInfo.PostList))
}

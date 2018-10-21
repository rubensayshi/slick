package slick

import (
	"testing"
	"time"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

var now = time.Now()

// Skeleton of a real message
var publicSlackMessage = slack.Msg{
	Text:      "This is a test.",
	User:      "U2147483697",
	Channel:   "C2147483705",
	Timestamp: string(now.Unix()),
}
var publicMessage = Message{
	Msg:        &publicSlackMessage,
	SubMessage: nil,
	MentionsMe: false,
	IsEdit:     false,
	FromMe:     false,
}
var privateSlackMessage = slack.Msg{
	Text:      "This is a test.",
	User:      "U2147483697",
	Channel:   "D2147483705",
	Timestamp: string(now.Unix()),
}
var privateMessage = Message{
	Msg:        &privateSlackMessage,
	SubMessage: nil,
	MentionsMe: false,
	IsEdit:     false,
	FromMe:     false,
}

func TestMessage(t *testing.T) {
	assert := assert.New(t)
	// IsPrivate
	assert.False(publicMessage.IsPrivate())
	assert.True(privateMessage.IsPrivate())
	// ContainsAnyCased
	assert.True(publicMessage.ContainsAnyCased([]string{"This", "not"}))
	assert.False(publicMessage.ContainsAnyCased([]string{"this", "not"}))
	// ContainsAny
	assert.True(publicMessage.ContainsAny([]string{"This", "not"}))
	assert.True(publicMessage.ContainsAny([]string{"this", "not"}))
	assert.False(publicMessage.ContainsAny([]string{"not", "here"}))
	// ContainsAll
	assert.True(publicMessage.ContainsAll([]string{"This", "is", "a", "test"}))
	assert.False(publicMessage.ContainsAll([]string{"This", "is not", "a", "test"}))
	// HasPrefix
	assert.True(publicMessage.HasPrefix("This"))
	assert.False(publicMessage.HasPrefix("this"))
}

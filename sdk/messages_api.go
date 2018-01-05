package ciscosparkapi

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// MessagesService is the service to communicate with the Messages API endpoint
type MessagesService service

// CreateMessage Post a plain text or rich text message, and optionally, a media content attachment, to a room.
/* Post a plain text or rich text message, and optionally, a media content attachment, to a room.
The files parameter is an array, which accepts multiple values to allow for future expansion, but currently only one file may be included with the message.

 @param messageCreateRequest
 @return Message
*/
func (s *MessagesService) CreateMessage(messageCreateRequest *MessageCreateRequest) (*Message, *resty.Response, error) {

	path := "/messages/"

	response, err := RestyClient.R().
		SetBody(messageCreateRequest).
		SetResult(&Message{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Message)
	return result, response, err

}

// DeleteMessage Delete a Message.
/* Deletes a message by ID.
@param messageID Message ID.
@return
*/
func (s *MessagesService) DeleteMessage(messageID string) (*resty.Response, error) {

	path := "/messages/{messageId}"
	path = strings.Replace(path, "{"+"messageId"+"}", fmt.Sprintf("%v", messageID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetMessage Shows details for a message, by message ID.
/* Shows details for a message, by message ID.
Specify the message ID in the messageID parameter in the URI.

 @param messageID Message ID.
 @return Message
*/
func (s *MessagesService) GetMessage(messageID string) (*Message, *resty.Response, error) {

	path := "/messages/{messageId}"
	path = strings.Replace(path, "{"+"messageId"+"}", fmt.Sprintf("%v", messageID), -1)

	response, err := RestyClient.R().
		SetResult(&Message{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Message)
	return result, response, err

}

// ListMessagesQueryParams are the query params for the ListMessages API Call
type ListMessagesQueryParams struct {
	RoomID          string    `url:"roomId,omitempty"`          // List messages for a room, by ID.
	MentionedPeople string    `url:"mentionedPeople,omitempty"` // List messages where the caller is mentioned by specifying *me* or the caller personId.
	Before          time.Time `url:"before,omitempty"`          // List messages sent before a date and time, in ISO8601 format. Format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ
	BeforeMessage   string    `url:"beforeMessage,omitempty"`   // List messages sent before a message, by ID.
	Max             int32     `url:"max,omitempty"`             // Limit the maximum number of items in the response.
}

// ListMessages Lists all messages in a room. Each message will include content attachments if present.
/* Lists all messages in a room. Each message will include content attachments if present.
The list sorts the messages in descending order by creation date.
Long result sets will be split into pages.

 @param roomID List messages for a room, by ID.
 @param "mentionedPeople" (string) List messages where the caller is mentioned by specifying *me* or the caller personId.
 @param "before" (time.Time) List messages sent before a date and time, in ISO8601 format. Format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ
 @param "beforeMessage" (string) List messages sent before a message, by ID.
 @param "max" (int32) Limit the maximum number of items in the response.
 @return Messages
*/
func (s *MessagesService) ListMessages(queryParams *ListMessagesQueryParams) (*Messages, *resty.Response, error) {

	path := "/messages/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Messages{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Messages)
	return result, response, err

}

package main

type CallbackRequestQueue struct {
	Command string `json:"command"`
	MQ      struct {
		QueueName string `json:"queue_name"`
	} `json:"mq"`
	Callback  Callback `json:"callback"`
	RequestID string   `json:"request_id"`
	Message   string   `json:"message"`
	TimeStamp string   `json:"timestamp"`
}

type Callback struct {
	RequestID               string            `json:"request_id"`
	CreatedAt               int               `json:"created_at"`
	RetryCreatedAt          int               `json:"retry_created_at"`
	CallbackType            string            `json:"callback_type"`
	ContentProviderKey      string            `json:"content_provider_key"`
	MediaContentID          int               `json:"media_content_id"`
	CallbackURL             string            `json:"callback_url" validate:"required"`
	CallbackMethod          string            `json:"callback_method" validate:"required"`
	CallbackHeader          map[string]string `json:"callback_header"`
	CallbackData            CallbackDataMsg   `json:"callback_data"`
	CallbackMaxRetry        int               `json:"callback_max_retry"`
	CallbackConnectTimeout  int               `json:"callback_connect_timeout"`
	CallbackResponseTimeout int               `json:"callback_response_timeout"`
	CallbackRetry           int               `json:"callback_retry"`
	CallbackResponse        int               `json:"callback_response"`
	CallbackResponseBody    string            `json:"callback_response_body"`
}

type CallbackDataMsg struct {
	AccountId        string `json:"account_id"`
	ContentId        string `json:"content_id"`
	Error            int    `json:"error"`
	GroupNm          string `json:"group_nm"`
	MediaOutputs     string `json:"media_outputs"`
	Message          string `json:"message"`
	SnapshotOutput   string `json:"snapshot_output"`
	State            string `json:"state"`
	ThumbnailOutputs string `json:"thumbnail_outputs"`
	WaveformOutput   string `json:"waveform_output"`
}

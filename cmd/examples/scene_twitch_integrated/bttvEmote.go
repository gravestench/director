package main

type BTTVUserEmotesDescriptor struct {
	ID            string          `json:"id"`
	Bots          []interface{}   `json:"bots"`
	ChannelEmotes []ChannelEmotes `json:"channelEmotes"`
	SharedEmotes  []SharedEmotes  `json:"sharedEmotes"`
}

type ChannelEmotes struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	ImageType string `json:"imageType"`
	UserID    string `json:"userId"`
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	ProviderID  string `json:"providerId"`
}

type SharedEmotes struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	ImageType string `json:"imageType"`
	User      User   `json:"user"`
}

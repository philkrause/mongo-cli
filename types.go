package mainw

import (
	"net/http"
	"time"
)

type Data struct {
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Results []struct {
		ClusterID string `json:"clusterId"`
		Complete  bool   `json:"complete"`
		Created   struct {
			Date      time.Time `json:"date"`
			Increment int       `json:"increment"`
			Time      int       `json:"time"`
		} `json:"created"`
		DoNotDelete               bool      `json:"doNotDelete"`
		Expires                   time.Time `json:"expires"`
		GroupID                   string    `json:"groupId"`
		ID                        string    `json:"id"`
		LastOplogAppliedTimestamp struct {
			Date      time.Time `json:"date"`
			Increment int       `json:"increment"`
			Time      int       `json:"time"`
		} `json:"lastOplogAppliedTimestamp"`
		Links []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
		Parts []struct {
			ClusterID          string `json:"clusterId"`
			CompressionSetting string `json:"compressionSetting"`
			DataSizeBytes      int64  `json:"dataSizeBytes"`
			EncryptionEnabled  bool   `json:"encryptionEnabled"`
			FileSizeBytes      int64  `json:"fileSizeBytes"`
			MongodVersion      string `json:"mongodVersion"`
			ReplicaSetName     string `json:"replicaSetName"`
			StorageSizeBytes   int64  `json:"storageSizeBytes"`
			TypeName           string `json:"typeName"`
		} `json:"parts"`
	} `json:"results"`
	TotalCount int `json:"totalCount"`
}

type Transport struct {
	Username  string
	Password  string
	Transport http.RoundTripper
}

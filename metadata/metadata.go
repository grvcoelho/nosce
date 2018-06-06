package metadata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Metadata struct {
	Client   *http.Client
	Endpoint string
	Items    map[string]*Item
}

func New(endpoint string) *Metadata {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	items := make(map[string]*Item)

	items["ami-id"] = NewItem("latest/meta-data/ami-id")
	items["ami-launch-index"] = NewItem("latest/meta-data/ami-launch-index")
	items["availability-zone"] = NewItem("latest/meta-data/placement/availability-zone")
	items["hostname"] = NewItem("latest/meta-data/hostname")
	items["instance-id"] = NewItem("latest/meta-data/instance-id")
	items["instance-type"] = NewItem("latest/meta-data/instance-type")
	items["local-hostname"] = NewItem("latest/meta-data/local-hostname")
	items["local-ipv4"] = NewItem("latest/meta-data/local-ipv4")
	items["profile"] = NewItem("latest/meta-data/profile")
	items["public-hostname"] = NewItem("latest/meta-data/public-hostname")
	items["public-ipv4"] = NewItem("latest/meta-data/public-ipv4")

	items["region"] = NewItemWithTransformer(
		"latest/meta-data/placement/availability-zone",
		func(data string) (string, error) {
			length := len(data)
			return data[:length-1], nil
		},
	)

	return &Metadata{
		Client:   client,
		Endpoint: endpoint,
		Items:    items,
	}
}

func (m *Metadata) Fetch(path string) (string, error) {
	url := m.Endpoint + "/" + path
	res, err := m.Client.Get(url)
	defer res.Body.Close()

	if err != nil {
		return "", fmt.Errorf("could not GET %s", url)
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("could not GET %s", url)
	}

	data, err := ioutil.ReadAll(res.Body)

	return string(data), err
}

func (m *Metadata) Get(key string) (string, error) {
	item, ok := m.Items[key]

	if !ok {
		return "", fmt.Errorf("item %s could not be found", key)
	}

	data, err := m.Fetch(item.Path)

	if err != nil {
		return "", fmt.Errorf("item %s could not be returned: %v", key, err)
	}

	value, err := item.Transformer(data)

	return value, err
}

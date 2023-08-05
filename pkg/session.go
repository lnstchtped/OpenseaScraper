package pkg

import tlsClient "github.com/bogdanfinn/tls-client"

type Session struct {
	Client tlsClient.HttpClient
}

func NewSession(proxy string) (Session, error) {
	opts := []tlsClient.HttpClientOption{
		tlsClient.WithClientProfile(tlsClient.Firefox_110),
		tlsClient.WithProxyUrl(proxy),
	}

	client, err := tlsClient.NewHttpClient(tlsClient.NewNoopLogger(), opts...)
	if err != nil {
		return Session{}, err
	}

	return Session{Client: client}, nil
}

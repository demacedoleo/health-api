package scrapper

import "github.com/demacedoleo/health-api/internal/client"

func ScrapAFJP(data client.ScrapData) (client.UserAFJP, error) {
	return client.GetMemberDataAFJP(data)
}

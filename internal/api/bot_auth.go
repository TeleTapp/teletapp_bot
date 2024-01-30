package api

import (
	"context"
	"encoding/json"
)

type BotAuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Expires      int    `json:"expires"`
	ID           string `json:"id"`
}

func (c *Client) BotAuth(ctx context.Context, telegramID int64) (*BotAuthResponse, error) {
	data, err := c.post(ctx, "/bot/auth", map[string]int64{
		"telegram_id": telegramID,
	})
	if err != nil {
		c.logger.Error().Err(err)
		return nil, err
	}

	var result BotAuthResponse
	if err := json.Unmarshal(data, &result); err != nil {
		c.logger.Error().Err(err)
		return nil, err
	}

	return &result, nil
}

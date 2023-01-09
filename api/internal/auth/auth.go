package auth

import (
	"context"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"go.uber.org/zap"
)

type Client struct {
	logger     *zap.SugaredLogger
	authClient *firebaseAuth.Client
}

func NewClient(ctx context.Context, logger *zap.SugaredLogger) (*Client, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		logger.Fatalw("error initializing app", "error", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	logger.Info("authClient instantiated")

	return &Client{
		logger:     logger,
		authClient: authClient,
	}, nil
}

func (c *Client) VerifyToken(token string) (string, error) {
	resToken, err := c.authClient.VerifyIDTokenAndCheckRevoked(context.Background(), token)
	if err != nil {
		return "", err
	}

	return resToken.UID, nil
}

func (c *Client) GetSessionCookie(r *http.Request, token string, expiresIn time.Duration) (*http.Cookie, error) {
	cookie, err := c.authClient.SessionCookie(r.Context(), token, expiresIn)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     "session",
		Value:    cookie,
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		Secure:   true,
	}, nil
}

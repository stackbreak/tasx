package tokens

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cristalhq/jwt/v3"
)

var defaultTokenTTL = 12 * time.Hour

// alias for reusage
type Registered = jwt.RegisteredClaims

type TokenManagerHS struct {
	builder  *jwt.Builder
	verifier jwt.Verifier
}

func NewTokenMangerHS(secret []byte) (*TokenManagerHS, error) {
	signer, err := jwt.NewSignerHS(jwt.HS256, secret)
	if err != nil {
		return nil, fmt.Errorf("pkg.tokens: unable to create signer. %w", err)
	}

	builder := jwt.NewBuilder(signer)

	verifier, err := jwt.NewVerifierHS(jwt.HS256, secret)
	if err != nil {
		return nil, fmt.Errorf("pkg.tokens: unable to create verifier. %w", err)
	}

	return &TokenManagerHS{
			builder:  builder,
			verifier: verifier,
		},
		nil
}

func (tm *TokenManagerHS) CreateWithClaims(claims interface{}) (*jwt.Token, error) {
	token, err := tm.builder.Build(claims)
	if err != nil {
		return nil, fmt.Errorf("pkg.tokens: error building token. %w", err)
	}

	return token, nil
}

func (tm *TokenManagerHS) ParseToClaims(tokenStr string, claims interface{}) error {
	token, err := jwt.ParseAndVerifyString(tokenStr, tm.verifier)
	if err != nil {
		return fmt.Errorf("pkg.tokens: error parsing/verifying token. %w", err)
	}

	err = json.Unmarshal(token.RawClaims(), &claims)
	if err != nil {
		return fmt.Errorf("pkg.tokens: error extracting into claims. %w", err)
	}

	return nil
}

// -1 get defaults
func (tm *TokenManagerHS) GetRegisteredClaimsBase(tokenTTL time.Duration) *Registered {
	if tokenTTL == -1 {
		tokenTTL = defaultTokenTTL
	}

	return &Registered{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
}

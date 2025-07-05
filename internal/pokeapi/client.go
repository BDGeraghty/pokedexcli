package pokeapi

import (
	"net/http"
	"time"

	"github.com/BDGeraghty/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, cacheTTL time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheTTL), // Default cache TTL
	}
}

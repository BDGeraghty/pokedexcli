package pokeapi

import (
	"net/http"
	"time"

	//"github.com/bootdotdev/pokedexcli/internal/pokecache"
	"github.com/bootdotdev/pokedexcli/internal/pokecache" // Client -
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

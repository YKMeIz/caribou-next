package cache

import "time"

func Store(k string, v interface{}, ttl time.Duration) error {
	return client.Set(k, v, ttl).Err()
}

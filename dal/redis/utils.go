package redis

import (
	"context"
	"time"
)

// redis operation logic

// Put a K-V with no expiration
func PutKey(ctx context.Context, key string, val any) error {
	return GetRDB().Set(ctx, key, val, 0).Err()
}

// Put a K-V with expiration
func PutKeyWith(ctx context.Context, key string, val any, exp time.Duration) error {
	return GetRDB().Set(ctx, key, val, exp).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	return GetRDB().Get(ctx, key).Result()
}

func GetKeys(ctx context.Context, pattern string) ([]string, error) {
	keys, err := GetRDB().Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func DelKey(ctx context.Context, key string) error {
	_, err := GetRDB().Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
package dcron

import "time"

// UnlockScript use lua for redis use
const UnlockScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end`

type LockHasExpired interface {
	Lock(key string, value interface{}, expiration time.Duration) error
	UnLock(key string, value interface{}) (interface{}, error)
	Expire(key string, expiration time.Duration) (bool, error)
	TTL(key string) (time.Duration, error)
}

type WithoutLock struct {
}

func (d *WithoutLock) Lock(string, interface{}, time.Duration) error {
	return nil
}

func (d *WithoutLock) UnLock(string, interface{}) (interface{}, error) {
	return nil, nil
}

func (d *WithoutLock) Expire(string, time.Duration) (bool, error) {
	return true, nil
}

func (d *WithoutLock) TTL(string) (time.Duration, error) {
	return 0, nil
}

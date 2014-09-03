package miniredis

import (
	"testing"

	"github.com/garyburd/redigo/redis"
)

// Test Hash.
func TestHash(t *testing.T) {
	s, err := Run()
	ok(t, err)
	defer s.Close()
	c, err := redis.Dial("tcp", s.Addr())
	ok(t, err)

	{
		b, err := redis.Int(c.Do("HSET", "aap", "noot", "mies"))
		ok(t, err)
		equals(t, 1, b) // New field.
	}

	{
		v, err := redis.String(c.Do("HGET", "aap", "noot"))
		ok(t, err)
		equals(t, "mies", v)
		equals(t, "mies", s.HGet("aap", "noot"))
	}

	{
		b, err := redis.Int(c.Do("HSET", "aap", "noot", "mies"))
		ok(t, err)
		equals(t, 0, b) // Existing field.
	}

	// Wrong type of key
	{
		_, err := redis.String(c.Do("SET", "foo", "bar"))
		ok(t, err)
		_, err = redis.Int(c.Do("HSET", "foo", "noot", "mies"))
		assert(t, err != nil, "HSET error")
	}

	// hash exists, key doesn't.
	{
		b, err := c.Do("HGET", "aap", "nosuch")
		ok(t, err)
		equals(t, nil, b)
	}

	// hash doesn't exists.
	{
		b, err := c.Do("HGET", "nosuch", "nosuch")
		ok(t, err)
		equals(t, nil, b)
		equals(t, "", s.HGet("nosuch", "nosuch"))
	}

	// HGET on wrong type
	{
		_, err := redis.Int(c.Do("HGET", "aap"))
		assert(t, err != nil, "HGET error")
	}

	// Direct HSet()
	{
		s.HSet("wim", "zus", "jet")
		v, err := redis.String(c.Do("HGET", "wim", "zus"))
		ok(t, err)
		equals(t, "jet", v)
	}
}

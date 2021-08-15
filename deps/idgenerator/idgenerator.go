package idgeneratora

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"log"
)

type IIdGenerator interface {
	Next() int64
}

var IdGenerator IIdGenerator

type snowflakeIdGenerator struct {
	*snowflake.Snowflake
}

func (s *snowflakeIdGenerator) Next() int64 {
	return s.NextVal()
}

func init() {
	// todo: read detacenterid / workerid in env
	s, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		log.Fatalf("snowflake error: %e", err)
	}

	IdGenerator = &snowflakeIdGenerator{s}
}

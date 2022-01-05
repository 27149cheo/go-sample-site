package core

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go-sample-site/pkg/core/user/repository/orm"
	"go-sample-site/pkg/log"
	gormtool "go-sample-site/pkg/tool/gorm"
)

func Start(ctx context.Context) {
	log.Init(&log.Config{Level: "debug", Filename: "/tmp/loutest.log", SendToFile: true})

	initDatabase(ctx)
}

func initDatabase(ctx context.Context) {
	err := gormtool.Open("root", "loutest", os.Getenv("MYSQL_HOST"), "loutest")
	if err != nil {
		log.Panicf("Failed to open database %s", "loutest")
	}

	idxCtx, idxCancel := context.WithTimeout(ctx, 10*time.Minute)
	defer idxCancel()

	var wg sync.WaitGroup
	for _, r := range []tabler{
		orm.NewUserColl(),
	} {
		wg.Add(1)
		go func(r tabler) {
			defer wg.Done()
			if err := r.EnsureTable(idxCtx); err != nil {
				panic(fmt.Errorf("failed to create table for %s, err: %s", r.GetCollectionName(), err))
			}
		}(r)
	}
	wg.Wait()
}

type tabler interface {
	EnsureTable(ctx context.Context) error
	GetCollectionName() string
}

func Stop(_ context.Context) {
	gormtool.Close()
}

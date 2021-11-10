package batch

import (
	"sync"
)

/*
#include "db.h"
*/
// #cgo CXXFLAGS: -std=c++11
import "C"

type db struct {
	groupInGo  bool
	cdb        *C.DB
	mu         sync.Mutex
	cond       *sync.Cond
	committing bool
	commitSeq  uint64
	pendingSeq uint64
	pending    []*batch
}

func newDB(groupInGo bool) *db {
	db := &db{
		groupInGo: groupInGo,
		cdb:       C.NewDB(),
	}
	db.cond = sync.NewCond(&db.mu)
	return db
}

func (db *db) newBatch() *batch {
	return &batch{
		cbatch: C.NewBatch(db.cdb),
	}
}

func (db *db) commit(b *batch) int {
	if !db.groupInGo {
		return int(C.CommitBatch(db.cdb, b.cbatch))
	}

	db.mu.Lock()
	leader := len(db.pending) == 0
	seq := db.pendingSeq
	db.pending = append(db.pending, b)
	var size int

	if leader {
		// We're the leader. Wait for any running commit to finish.
		for db.committing {
			db.cond.Wait()
		}
		pending := db.pending
		db.pending = nil
		db.pendingSeq++
		db.committing = true
		db.mu.Unlock()

		for _, p := range pending[1:] {
			b.combine(p)
		}
		b.apply()
		size = len(pending)

		db.mu.Lock()
		db.committing = false
		db.commitSeq = seq
		db.cond.Broadcast()
	} else {
		// We're a follower. Wait for the commit to finish.
		for db.commitSeq < seq {
			db.cond.Wait()
		}
	}
	db.mu.Unlock()
	return size
}

type batch struct {
	cbatch *C.Batch
}

func (b *batch) combine(other *batch) {
	// In the real implementation, this combines the batches together.
}

func (b *batch) apply() {
	C.ApplyBatch(b.cbatch)
}

func (b *batch) free() {
	C.FreeBatch(b.cbatch)
}

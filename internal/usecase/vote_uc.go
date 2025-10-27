package usecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srgjo27/agora/internal/domain"
)

type voteUsecase struct {
	db         *sqlx.DB
	voteRepo   VoteRepository
	threadRepo ThreadRepository
}

func (uc *voteUsecase) VoteOnThread(ctx context.Context, userID uuid.UUID, threadID uuid.UUID, voteType int) error {
	_, err := uc.threadRepo.GetByID(ctx, threadID)
	if err != nil {
		return err
	}

	oldVote, err := uc.voteRepo.GetThreadVote(ctx, userID, threadID)
	oldVoteType := 0
	if err != nil && err != domain.ErrNotFound {
		return err
	}

	if oldVote != nil {
		oldVoteType = oldVote.VoteType
	}

	delta := voteType - oldVoteType

	tx, err := uc.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic, rolling back transaction")
			tx.Rollback()
		}
	}()

	if voteType == 0 {
		if err := uc.voteRepo.DeleteThreadVote(ctx, tx, userID, threadID); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		newVote := &domain.ThreadVote{
			UserID:   userID,
			ThreadID: threadID,
			VoteType: voteType,
		}

		if err := uc.voteRepo.UpsertThreadVote(ctx, tx, newVote); err != nil {
			tx.Rollback()
			return err
		}
	}

	if delta != 0 {
		if err := uc.threadRepo.UpdateVoteCount(ctx, tx, threadID, delta); err != nil {
			tx.Rollback()

			return err
		}
	}

	return tx.Commit()
}

func NewVoteUsecase(db *sqlx.DB, vr VoteRepository, tr ThreadRepository) VoteUsecase {
	return &voteUsecase{
		db:         db,
		voteRepo:   vr,
		threadRepo: tr,
	}
}

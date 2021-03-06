// Code generated by sqlc. DO NOT EDIT.
// source: test.sql

package db

import (
	"context"
)

const biggestVID = `-- name: BiggestVID :one
SELECT visible_id 
FROM tests 
WHERE problem_id = $1 AND orphaned = false
ORDER BY visible_id desc
LIMIT 1
`

func (q *Queries) BiggestVID(ctx context.Context, problemID int64) (int32, error) {
	row := q.queryRow(ctx, q.biggestVIDStmt, biggestVID, problemID)
	var visible_id int32
	err := row.Scan(&visible_id)
	return visible_id, err
}

const createTest = `-- name: CreateTest :exec
INSERT INTO tests (
	problem_id, visible_id, score 
) VALUES (
	$1, $2, $3
)
`

type CreateTestParams struct {
	ProblemID int64 `json:"problem_id"`
	VisibleID int32 `json:"visible_id"`
	Score     int32 `json:"score"`
}

func (q *Queries) CreateTest(ctx context.Context, arg CreateTestParams) error {
	_, err := q.exec(ctx, q.createTestStmt, createTest, arg.ProblemID, arg.VisibleID, arg.Score)
	return err
}

const setPbTestScore = `-- name: SetPbTestScore :exec
UPDATE tests 
SET score = $3
WHERE problem_id = $1 AND visible_id = $2 AND orphaned = false
`

type SetPbTestScoreParams struct {
	ProblemID int64 `json:"problem_id"`
	VisibleID int32 `json:"visible_id"`
	Score     int32 `json:"score"`
}

func (q *Queries) SetPbTestScore(ctx context.Context, arg SetPbTestScoreParams) error {
	_, err := q.exec(ctx, q.setPbTestScoreStmt, setPbTestScore, arg.ProblemID, arg.VisibleID, arg.Score)
	return err
}

const setPbTestVisibleID = `-- name: SetPbTestVisibleID :exec
UPDATE tests 
SET visible_id = $1
WHERE problem_id = $2 AND visible_id = $3 AND orphaned = false
`

type SetPbTestVisibleIDParams struct {
	NewID     int32 `json:"new_id"`
	ProblemID int64 `json:"problem_id"`
	OldID     int32 `json:"old_id"`
}

func (q *Queries) SetPbTestVisibleID(ctx context.Context, arg SetPbTestVisibleIDParams) error {
	_, err := q.exec(ctx, q.setPbTestVisibleIDStmt, setPbTestVisibleID, arg.NewID, arg.ProblemID, arg.OldID)
	return err
}

const setVisibleID = `-- name: SetVisibleID :exec
UPDATE tests 
SET visible_id = $2
WHERE id = $1 AND orphaned = false
`

type SetVisibleIDParams struct {
	ID        int64 `json:"id"`
	VisibleID int32 `json:"visible_id"`
}

func (q *Queries) SetVisibleID(ctx context.Context, arg SetVisibleIDParams) error {
	_, err := q.exec(ctx, q.setVisibleIDStmt, setVisibleID, arg.ID, arg.VisibleID)
	return err
}

const test = `-- name: Test :one
SELECT id, created_at, score, problem_id, visible_id, orphaned FROM tests 
WHERE id = $1 AND orphaned = false
`

func (q *Queries) Test(ctx context.Context, id int64) (Test, error) {
	row := q.queryRow(ctx, q.testStmt, test, id)
	var i Test
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Score,
		&i.ProblemID,
		&i.VisibleID,
		&i.Orphaned,
	)
	return i, err
}

const testVisibleID = `-- name: TestVisibleID :one
SELECT id, created_at, score, problem_id, visible_id, orphaned FROM tests 
WHERE problem_id = $1 AND visible_id = $2 AND orphaned = false
ORDER BY visible_id
`

type TestVisibleIDParams struct {
	ProblemID int64 `json:"problem_id"`
	VisibleID int32 `json:"visible_id"`
}

func (q *Queries) TestVisibleID(ctx context.Context, arg TestVisibleIDParams) (Test, error) {
	row := q.queryRow(ctx, q.testVisibleIDStmt, testVisibleID, arg.ProblemID, arg.VisibleID)
	var i Test
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Score,
		&i.ProblemID,
		&i.VisibleID,
		&i.Orphaned,
	)
	return i, err
}

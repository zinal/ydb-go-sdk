//go:build integration
// +build integration

package integration

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/version"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xtest"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/topic/topicwriter"
)

func TestTopicReadInTransaction(t *testing.T) {
	if os.Getenv("YDB_VERSION") != "nightly" && version.Lt(os.Getenv("YDB_VERSION"), "25.0") {
		t.Skip("require enables transactions for topics")
	}
	scope := newScope(t)
	ctx := scope.Ctx
	require.NoError(t, scope.TopicWriter().Write(ctx, topicwriter.Message{Data: strings.NewReader("asd")}))
	scope.Logf("topic message written")

	require.NoError(t, scope.Driver().Query().DoTx(ctx, func(ctx context.Context, tx query.TxActor) error {
		reader := scope.TopicReaderNamed("first")
		scope.Logf("trying to pop a batch")
		batch, err := reader.PopBatchTx(ctx, tx)
		scope.Logf("pop a batch result: %v", err)
		if err != nil {
			return err
		}
		content := string(xtest.Must(io.ReadAll(batch.Messages[0])))
		require.Equal(t, "asd", content)
		_ = reader.Close(ctx)
		return nil
	}))

	scope.Logf("first pop messages done")

	scope.Logf("writting second message")
	require.NoError(t, scope.TopicWriter().Write(ctx, topicwriter.Message{Data: strings.NewReader("bbb")}))

	require.NoError(t, scope.Driver().Query().DoTx(ctx, func(ctx context.Context, tx query.TxActor) error {
		reader := scope.TopicReaderNamed("second")
		// err := tx.Exec(ctx, "SELECT 1", query.WithCommit())
		err := tx.Exec(ctx, "SELECT 1")
		if err != nil {
			return err
		}

		scope.Logf("trying second pop batch")
		batch, err := reader.PopBatchTx(ctx, tx)
		scope.Logf("second pop batch result: %v", err)
		if err != nil {
			return err
		}
		content := string(xtest.Must(io.ReadAll(batch.Messages[0])))
		require.Equal(t, "bbb", content)
		return nil
	}))
}

package errs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewNonPositiveAmount(t *testing.T) {
	rs := NewNonPositiveAmount(1, "some")
	require.Equal(t, "1 of some", rs.Error())
}

func TestEmptyDataKey(t *testing.T) {
	require.EqualError(t, Extend(NewEmptyDataKey("a"), "b"), "b: a")
}

func TestDuplicatedDataKeys(t *testing.T) {
	require.EqualError(t, Extend(NewDuplicatedDataKeys("a"), "b"), "b: a")
}

func TestMistiming(t *testing.T) {
	require.EqualError(t, Extend(NewMistiming("a"), "b"), "b: a")
}

func TestUnknownAsset(t *testing.T) {
	require.EqualError(t, NewUnknownAsset("a").Extend("b"), "b: a")
}

func TestNewTooBigArray(t *testing.T) {
	require.EqualError(t, NewTooBigArray("a").Extend("b"), "b: a")
}

func TestNewInvalidName(t *testing.T) {
	require.EqualError(t, NewInvalidName("a").Extend("b"), "b: a")
}

func TestNewAccountBalanceError(t *testing.T) {
	require.EqualError(t, NewAccountBalanceError("a").Extend("b"), "b: a")
}

func TestNewToSelf(t *testing.T) {
	require.EqualError(t, NewToSelf("a").Extend("b"), "b: a")
}

func TestNewAliasTaken(t *testing.T) {
	require.EqualError(t, NewAliasTaken("a").Extend("b"), "b: a")
}

func TestNewAssetIsNotReissuable(t *testing.T) {
	require.EqualError(t, NewAssetIsNotReissuable("a").Extend("b"), "b: a")
}

func TestNewTxValidationError(t *testing.T) {
	require.EqualError(t, NewTxValidationError("a").Extend("b"), "b: a")
	require.True(t, IsValidationError(NewTxValidationError("")))
}

func TestNewAssetIssuedByOtherAddress(t *testing.T) {
	require.EqualError(t, NewAssetIssuedByOtherAddress("a").Extend("b"), "b: a")
}

func TestFeeValidation(t *testing.T) {
	require.EqualError(t, NewFeeValidation("a").Extend("b"), "b: a")
}

func TestNewAssetUpdateInterval(t *testing.T) {
	require.EqualError(t, NewAssetUpdateInterval("a").Extend("b"), "b: a")
}

func TestNewTransactionNotAllowedByScript(t *testing.T) {
	err := NewTransactionNotAllowedByScript("a", nil)
	require.EqualError(t, err.Extend("b"), "b: a")
	require.False(t, err.IsAssetScript())
	require.Len(t, err.Asset(), 0)
}

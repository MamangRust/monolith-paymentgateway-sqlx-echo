package cardrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateCardFailed is returned when creating a new card fails.
	ErrCreateCardFailed = errors.ErrInternal.WithMessage("Failed to create card")

	// ErrUpdateCardFailed is returned when updating a card fails.
	ErrUpdateCardFailed = errors.ErrInternal.WithMessage("Failed to update card")

	// ErrTrashCardFailed is returned when trashing a card fails.
	ErrTrashCardFailed = errors.ErrInternal.WithMessage("Failed to move card to trash")

	// ErrRestoreCardFailed is returned when restoring a trashed card fails.
	ErrRestoreCardFailed = errors.ErrInternal.WithMessage("Failed to restore card from trash")

	// ErrDeleteCardPermanentFailed is returned when permanently deleting a card fails.
	ErrDeleteCardPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete card")

	// ErrRestoreAllCardsFailed is returned when restoring all trashed cards fails.
	ErrRestoreAllCardsFailed = errors.ErrInternal.WithMessage("Failed to restore all cards")

	// ErrDeleteAllCardsPermanentFailed is returned when permanently deleting all cards fails.
	ErrDeleteAllCardsPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete all cards")

	// ErrUpdateCardStatusFailed is returned when updating card status fails.
	ErrUpdateCardStatusFailed = errors.ErrInternal.WithMessage("Failed to update card status")

	// ErrUpdateCreditLimitFailed is returned when updating credit limit fails.
	ErrUpdateCreditLimitFailed = errors.ErrInternal.WithMessage("Failed to update credit limit")

	// ErrAddRewardPointsFailed is returned when adding reward points fails.
	ErrAddRewardPointsFailed = errors.ErrInternal.WithMessage("Failed to add reward points")

	// ErrRedeemRewardPointsFailed is returned when redeeming reward points fails.
	ErrRedeemRewardPointsFailed = errors.ErrInternal.WithMessage("Failed to redeem reward points")

	// ErrUpdateOutstandingBalanceFailed is returned when updating outstanding balance fails.
	ErrUpdateOutstandingBalanceFailed = errors.ErrInternal.WithMessage("Failed to update outstanding balance")

	// ErrInsertBillingCycleFailed is returned when inserting billing cycle fails.
	ErrInsertBillingCycleFailed = errors.ErrInternal.WithMessage("Failed to create billing cycle")

	// ErrGetBillingCyclesFailed is returned when fetching billing cycles fails.
	ErrGetBillingCyclesFailed = errors.ErrInternal.WithMessage("Failed to get billing cycles")

	// ErrUpdateBillingCycleStatusFailed is returned when updating billing cycle status fails.
	ErrUpdateBillingCycleStatusFailed = errors.ErrInternal.WithMessage("Failed to update billing cycle status")

	// --- Auth Transaction Errors ---

	// ErrInsertAuthTransactionFailed is returned when inserting an auth transaction fails.
	ErrInsertAuthTransactionFailed = errors.ErrInternal.WithMessage("Failed to insert authorization transaction")

	// ErrApproveAuthTransactionFailed is returned when approving an auth transaction fails.
	ErrApproveAuthTransactionFailed = errors.ErrInternal.WithMessage("Failed to approve authorization transaction")

	// ErrDeclineAuthTransactionFailed is returned when declining an auth transaction fails.
	ErrDeclineAuthTransactionFailed = errors.ErrInternal.WithMessage("Failed to decline authorization transaction")

	// ErrReverseAuthTransactionFailed is returned when reversing an auth transaction fails.
	ErrReverseAuthTransactionFailed = errors.ErrInternal.WithMessage("Failed to reverse authorization transaction")

	// ErrGetAuthTransactionFailed is returned when fetching an auth transaction fails.
	ErrGetAuthTransactionFailed = errors.ErrInternal.WithMessage("Failed to get authorization transaction")

	// ErrUpdateRiskScoreFailed is returned when updating the fraud risk score fails.
	ErrUpdateRiskScoreFailed = errors.ErrInternal.WithMessage("Failed to update risk score")

	// --- Card Payment Errors ---

	// ErrInsertCardPaymentFailed is returned when creating a card payment fails.
	ErrInsertCardPaymentFailed = errors.ErrInternal.WithMessage("Failed to create card payment")

	// ErrGetCardPaymentFailed is returned when fetching card payment records fails.
	ErrGetCardPaymentFailed = errors.ErrInternal.WithMessage("Failed to get card payment records")

	// --- Reward Errors ---

	// ErrEarnRewardsFailed is returned when earning rewards fails.
	ErrEarnRewardsFailed = errors.ErrInternal.WithMessage("Failed to earn rewards")

	// ErrGetRewardBalanceFailed is returned when fetching reward balance fails.
	ErrGetRewardBalanceFailed = errors.ErrInternal.WithMessage("Failed to get reward balance")

	// ErrGetRewardHistoryFailed is returned when fetching reward history fails.
	ErrGetRewardHistoryFailed = errors.ErrInternal.WithMessage("Failed to get reward history")

	// ErrRedeemRewardsFailed is returned when redeeming rewards fails.
	ErrRedeemRewardsFailed = errors.ErrInternal.WithMessage("Failed to redeem rewards")
)

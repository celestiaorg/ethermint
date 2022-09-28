package types

// Evm module events
const (
	EventTypeEthereumTx  = TypeMsgEthereumTx
	EventTypeBlockBloom  = "block_bloom"
	EventTypeTxLog       = "tx_log"
	EventTypeTxRoot      = "tx_root"
	EventTypeReceiptRoot = "receipt_root"

	AttributeKeyContractAddress = "contract"
	AttributeKeyRecipient       = "recipient"
	AttributeKeyTxHash          = "txHash"
	AttributeKeyEthereumTxHash  = "ethereumTxHash"
	AttributeKeyTxIndex         = "txIndex"
	AttributeKeyTxGasUsed       = "txGasUsed"
	AttributeKeyTxType          = "txType"
	AttributeKeyTxLog           = "txLog"
	// tx failed in eth vm execution
	AttributeKeyEthereumTxFailed    = "ethereumTxFailed"
	AttributeValueCategory          = ModuleName
	AttributeKeyEthereumBloom       = "bloom"
	AttributeKeyEthereumTxRoot      = "ethTxRoot"
	AttributeKeyEthereumReceiptRoot = "receiptRoot"

	MetricKeyTransitionDB = "transition_db"
	MetricKeyStaticCall   = "static_call"
)

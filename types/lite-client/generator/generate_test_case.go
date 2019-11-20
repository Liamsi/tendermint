package generator

func GenerateTestCase(jsonValList string) {

	var testCases *TestCases = &TestCases{}
	valList := GetValList(jsonValList)

	// // Verify - ValidatorSet
	// CaseVerifyValidatorSetOf1(testCases, valList)
	// CaseVerifyValidatorSetOf8(testCases, valList)
	// CaseVerifyValidatorSetOf128(testCases, valList)
	// CaseVerifyValidatorSetEmpty(testCases, valList)

	// CaseVerifyValidatorSetAddTwiceVals(testCases, valList)
	// CaseVerifyValidatorSetRemoveHalfVals(testCases, valList)

	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesOneThird(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesHalf(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesTwoThirds(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesFully(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesLessThanOneThird(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesMoreThanTwoThirds(testCases, valList)

	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetWrongProposer(testCases, valList)

	// CaseVerifyValidatorSetWrongValidatorSet(testCases, valList)

	// // Verify - Commit
	// CaseVerifyCommitEmpty(testCases, valList)
	// CaseVerifyCommitWrongHeaderHash(testCases, valList)
	// CaseVerifyCommitWrongPartsHeaderCount(testCases, valList)
	// CaseVerifyCommitWrongPartsHeaderHash(testCases, valList)
	// CaseVerifyCommitWrongVoteType(testCases, valList)
	// CaseVerifyCommitWrongVoteHeight(testCases, valList)
	// CaseVerifyCommitWrongVoteRound(testCases, valList)
	// CaseVerifyCommitWrongVoteBlockID(testCases, valList)
	// CaseVerifyCommitWrongVoteTimestamp(testCases, valList)
	// CaseVerifyCommitWrongVoteSignature(testCases, valList)
	// CaseVerifyCommitWrongVoteInvalidSignature(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyCommitOneThirdValsDontSign(testCases, valList)
	// CaseVerifyCommitLessThanOneThirdValsDontSign(testCases, valList)

	// Verify - Header
	// CaseVerifyHeaderEmpty(testCases, valList)
	// CaseVerifyHeaderWrongLastCommitHash(testCases, valList)
	// CaseVerifyHeaderWrongLastResultsHash(testCases, valList)
	// CaseVerifyHeaderWrongLastBlockID(testCases, valList)
	// CaseVerifyHeaderWrongDataHash(testCases, valList)
	// CaseVerifyHeaderWrongChainID(testCases, valList)
	// CaseVerifyHeaderWrongVersion(testCases, valList)
	// CaseVerifyHeaderWrongHeight(testCases, valList)
	// CaseVerifyHeaderWrongTimestamp(testCases, valList)
	// CaseVerifyHeaderWrongNumTxs(testCases, valList)
	// CaseVerifyHeaderWrongTotalTxs(testCases, valList)
	// CaseVerifyHeaderWrongValSetHash(testCases, valList)
	CaseVerifyHeaderWrongNextValSetHash(testCases, valList)

	GenerateJSON(testCases)
}
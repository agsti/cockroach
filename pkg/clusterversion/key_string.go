// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_2-0]
	_ = x[Start22_1-1]
	_ = x[UnsplitRangesInAsyncGCJobs-2]
	_ = x[ValidateGrantOption-3]
	_ = x[PebbleFormatBlockPropertyCollector-4]
	_ = x[ProbeRequest-5]
	_ = x[SelectRPCsTakeTracingInfoInband-6]
	_ = x[PreSeedTenantSpanConfigs-7]
	_ = x[SeedTenantSpanConfigs-8]
	_ = x[PublicSchemasWithDescriptors-9]
	_ = x[EnsureSpanConfigReconciliation-10]
	_ = x[EnsureSpanConfigSubscription-11]
	_ = x[EnableSpanConfigStore-12]
	_ = x[ScanWholeRows-13]
	_ = x[SCRAMAuthentication-14]
	_ = x[UnsafeLossOfQuorumRecoveryRangeLog-15]
	_ = x[AlterSystemProtectedTimestampAddColumn-16]
	_ = x[EnableProtectedTimestampsForTenant-17]
	_ = x[DeleteCommentsWithDroppedIndexes-18]
	_ = x[RemoveIncompatibleDatabasePrivileges-19]
	_ = x[AddRaftAppliedIndexTermMigration-20]
	_ = x[PostAddRaftAppliedIndexTermMigration-21]
	_ = x[DontProposeWriteTimestampForLeaseTransfers-22]
	_ = x[EnablePebbleFormatVersionBlockProperties-23]
	_ = x[DisableSystemConfigGossipTrigger-24]
	_ = x[MVCCIndexBackfiller-25]
	_ = x[EnableLeaseHolderRemoval-26]
	_ = x[BackupResolutionInJob-27]
	_ = x[LooselyCoupledRaftLogTruncation-28]
	_ = x[ChangefeedIdleness-29]
	_ = x[BackupDoesNotOverwriteLatestAndCheckpoint-30]
	_ = x[EnableDeclarativeSchemaChanger-31]
	_ = x[RowLevelTTL-32]
	_ = x[PebbleFormatSplitUserKeysMarked-33]
	_ = x[IncrementalBackupSubdir-34]
	_ = x[DateStyleIntervalStyleCastRewrite-35]
	_ = x[EnableNewStoreRebalancer-36]
	_ = x[ClusterLocksVirtualTable-37]
	_ = x[AutoStatsTableSettings-38]
	_ = x[ForecastStats-39]
	_ = x[SuperRegions-40]
	_ = x[EnableNewChangefeedOptions-41]
	_ = x[SpanCountTable-42]
	_ = x[PreSeedSpanCountTable-43]
	_ = x[SeedSpanCountTable-44]
	_ = x[V22_1-45]
	_ = x[Start22_2-46]
	_ = x[LocalTimestamps-47]
	_ = x[EnsurePebbleFormatVersionRangeKeys-48]
	_ = x[EnablePebbleFormatVersionRangeKeys-49]
	_ = x[TrigramInvertedIndexes-50]
	_ = x[RemoveGrantPrivilege-51]
	_ = x[MVCCRangeTombstones-52]
	_ = x[UpgradeSequenceToBeReferencedByID-53]
	_ = x[SampledStmtDiagReqs-54]
	_ = x[AddSSTableTombstones-55]
	_ = x[SystemPrivilegesTable-56]
	_ = x[EnablePredicateProjectionChangefeed-57]
	_ = x[AlterSystemSQLInstancesAddLocality-58]
}

const _Key_name = "V21_2Start22_1UnsplitRangesInAsyncGCJobsValidateGrantOptionPebbleFormatBlockPropertyCollectorProbeRequestSelectRPCsTakeTracingInfoInbandPreSeedTenantSpanConfigsSeedTenantSpanConfigsPublicSchemasWithDescriptorsEnsureSpanConfigReconciliationEnsureSpanConfigSubscriptionEnableSpanConfigStoreScanWholeRowsSCRAMAuthenticationUnsafeLossOfQuorumRecoveryRangeLogAlterSystemProtectedTimestampAddColumnEnableProtectedTimestampsForTenantDeleteCommentsWithDroppedIndexesRemoveIncompatibleDatabasePrivilegesAddRaftAppliedIndexTermMigrationPostAddRaftAppliedIndexTermMigrationDontProposeWriteTimestampForLeaseTransfersEnablePebbleFormatVersionBlockPropertiesDisableSystemConfigGossipTriggerMVCCIndexBackfillerEnableLeaseHolderRemovalBackupResolutionInJobLooselyCoupledRaftLogTruncationChangefeedIdlenessBackupDoesNotOverwriteLatestAndCheckpointEnableDeclarativeSchemaChangerRowLevelTTLPebbleFormatSplitUserKeysMarkedIncrementalBackupSubdirDateStyleIntervalStyleCastRewriteEnableNewStoreRebalancerClusterLocksVirtualTableAutoStatsTableSettingsForecastStatsSuperRegionsEnableNewChangefeedOptionsSpanCountTablePreSeedSpanCountTableSeedSpanCountTableV22_1Start22_2LocalTimestampsEnsurePebbleFormatVersionRangeKeysEnablePebbleFormatVersionRangeKeysTrigramInvertedIndexesRemoveGrantPrivilegeMVCCRangeTombstonesUpgradeSequenceToBeReferencedByIDSampledStmtDiagReqsAddSSTableTombstonesSystemPrivilegesTableEnablePredicateProjectionChangefeedAlterSystemSQLInstancesAddLocality"

var _Key_index = [...]uint16{0, 5, 14, 40, 59, 93, 105, 136, 160, 181, 209, 239, 267, 288, 301, 320, 354, 392, 426, 458, 494, 526, 562, 604, 644, 676, 695, 719, 740, 771, 789, 830, 860, 871, 902, 925, 958, 982, 1006, 1028, 1041, 1053, 1079, 1093, 1114, 1132, 1137, 1146, 1161, 1195, 1229, 1251, 1271, 1290, 1323, 1342, 1362, 1383, 1418, 1452}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}

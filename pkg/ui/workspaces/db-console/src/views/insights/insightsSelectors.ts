// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

import { LocalSetting } from "src/redux/localsettings";
import { AdminUIState } from "src/redux/state";
import { createSelector } from "reselect";
import {
  defaultFilters,
  WorkloadInsightEventFilters,
  insightType,
  SchemaInsightEventFilters,
  SortSetting,
  selectID,
  selectStatementInsightDetailsCombiner,
  selectTxnInsightDetailsCombiner,
  InsightEnumToLabel,
  TxnInsightDetails,
  api,
} from "@cockroachlabs/cluster-ui";

export const filtersLocalSetting = new LocalSetting<
  AdminUIState,
  WorkloadInsightEventFilters
>("filters/InsightsPage", (state: AdminUIState) => state.localSettings, {
  app: defaultFilters.app,
  workloadInsightType: defaultFilters.workloadInsightType,
});

export const sortSettingLocalSetting = new LocalSetting<
  AdminUIState,
  SortSetting
>("sortSetting/InsightsPage", (state: AdminUIState) => state.localSettings, {
  ascending: false,
  columnTitle: "startTime",
});

export const selectTransactionInsightsLoading = (state: AdminUIState) =>
  state.cachedData.txnInsights?.inFlight &&
  (!state.cachedData.txnInsights?.valid || !state.cachedData.txnInsights?.data);

export const selectTransactionInsights = (state: AdminUIState) =>
  state.cachedData.txnInsights?.valid
    ? state.cachedData.txnInsights?.data
    : null;

const selectCachedTxnInsightDetails = createSelector(
  [(state: AdminUIState) => state.cachedData.txnInsightDetails, selectID],
  (insight, insightId): TxnInsightDetails => {
    if (!insight) {
      return null;
    }
    return insight[insightId]?.data?.result;
  },
);

const selectTxnInsight = createSelector(
  (state: AdminUIState) => state.cachedData.txnInsights?.data,
  selectID,
  (insights, execID) => {
    return insights?.find(txn => txn.transactionExecutionID === execID);
  },
);

export const selectStmtInsights = (state: AdminUIState) =>
  state.cachedData.stmtInsights?.data;

export const selectTxnInsightDetails = createSelector(
  selectTxnInsight,
  selectCachedTxnInsightDetails,
  selectStmtInsights,
  selectTxnInsightDetailsCombiner,
);

export const selectTransactionInsightDetailsError = createSelector(
  (state: AdminUIState) => state.cachedData.txnInsightDetails,
  selectID,
  (insights, insightId: string): api.TxnInsightDetailsReqErrs | null => {
    if (!insights) {
      return null;
    }
    const reqErrors = insights[insightId]?.data?.errors;
    if (insights[insightId]?.lastError) {
      Object.keys(reqErrors).forEach(
        (key: keyof api.TxnInsightDetailsReqErrs) => {
          reqErrors[key] = insights[insightId].lastError;
        },
      );
    }

    return insights[insightId]?.data?.errors;
  },
);

// Data is showed as loading when the request is in flight AND we have
// no data, (i.e. first request), or the current loaded data is
// invalid (e.g. time range changed)
export const selectStmtInsightsLoading = (state: AdminUIState) =>
  state.cachedData.stmtInsights?.inFlight &&
  (!state.cachedData.stmtInsights?.data ||
    !state.cachedData.stmtInsights.valid);

export const selectInsightTypes = () => {
  const insights: string[] = [];
  InsightEnumToLabel.forEach(insight => {
    insights.push(insight);
  });
  return insights;
};

export const selectStatementInsightDetails = createSelector(
  selectStmtInsights,
  selectID,
  selectStatementInsightDetailsCombiner,
);

export const schemaInsightsFiltersLocalSetting = new LocalSetting<
  AdminUIState,
  SchemaInsightEventFilters
>("filters/SchemaInsightsPage", (state: AdminUIState) => state.localSettings, {
  database: defaultFilters.database,
  schemaInsightType: defaultFilters.schemaInsightType,
});

export const schemaInsightsSortLocalSetting = new LocalSetting<
  AdminUIState,
  SortSetting
>(
  "sortSetting/SchemaInsightsPage",
  (state: AdminUIState) => state.localSettings,
  {
    ascending: false,
    columnTitle: "insights",
  },
);

export const selectSchemaInsights = createSelector(
  (state: AdminUIState) => state.cachedData,
  adminUiState => {
    if (!adminUiState.schemaInsights) return [];
    return adminUiState.schemaInsights.data;
  },
);

export const selectSchemaInsightsDatabases = createSelector(
  selectSchemaInsights,
  schemaInsights => {
    if (!schemaInsights) return [];
    return Array.from(
      new Set(schemaInsights.map(schemaInsight => schemaInsight.database)),
    ).sort();
  },
);

export const selectSchemaInsightsTypes = createSelector(
  selectSchemaInsights,
  schemaInsights => {
    if (!schemaInsights) return [];
    return Array.from(
      new Set(
        schemaInsights.map(schemaInsight => insightType(schemaInsight.type)),
      ),
    ).sort();
  },
);

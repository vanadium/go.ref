// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vsync

import (
	"strings"

	"v.io/v23/context"
	wire "v.io/v23/services/syncbase/nosql"
	"v.io/x/lib/vlog"
)

// getDbSchema returns the SchemaMetadata for the db.
func (iSt *initiationState) getDbSchema(ctx *context.T) (*wire.SchemaMetadata, error) {
	return iSt.config.db.GetSchemaMetadataInternal(ctx)
}

// groupConflictsByType uses CrRules of schema to group conflicts by the
// ResolverType applicable to the conflict.
func (iSt *initiationState) groupConflictsByType(schema *wire.SchemaMetadata) map[wire.ResolverType]map[string]*objConflictState {
	groupedConflicts := map[wire.ResolverType]map[string]*objConflictState{}
	for oid, conflictState := range iSt.updObjects {
		if !conflictState.isConflict {
			continue
		}
		crType := getResolutionType(oid, schema)
		crGroup := groupedConflicts[crType]
		if crGroup == nil {
			crGroup = map[string]*objConflictState{}
			groupedConflicts[crType] = crGroup
		}
		crGroup[oid] = conflictState
	}
	return groupedConflicts
}

// getResolutionType applies the CrRules present in schema on the oid and
// selects a ResolverType. All rules are matched against the row.
// If no rules match the row, we default to "LastWins". If multiple
// rules match the row, ties are broken as follows:
//  1. If one match has a longer prefix than the other, take that one.
//  2. Else, if only one match specifies a type, take that one.
//  3. Else, the two matches are identical; take the last one in the Rules array.
func getResolutionType(oid string, schema *wire.SchemaMetadata) wire.ResolverType {
	if !isRowKey(oid) {
		// This is a perms object key. Handle perms using LastWins policy till a
		// better policy is available.
		return wire.ResolverTypeLastWins
	}
	var selectedRule *wire.CrRule
	for i := range schema.Policy.Rules {
		rule := &schema.Policy.Rules[i]
		if !isRuleApplicable(oid, rule) {
			continue
		}
		if isRuleMoreSpecific(rule, selectedRule) {
			selectedRule = rule
		}
	}
	if selectedRule != nil {
		return selectedRule.Resolver
	}
	return wire.ResolverTypeLastWins
}

// isRuleApplicable ignores Type and only uses KeyPrefix to check if a rule
// applies to the given oid.
// TODO(jlodhia): Implement Type based matching.
func isRuleApplicable(oid string, rule *wire.CrRule) bool {
	tableName, rowKey := extractComponentsFromKey(oid)
	if rule.TableName != "" && tableName != rule.TableName {
		return false
	}
	if rule.KeyPrefix != "" && rule.TableName == "" {
		// TableName cannot be empty if KeyPrefix is specified
		vlog.Infof("Found illegal CrRule: %v", rule)
		return false
	}
	return strings.HasPrefix(rowKey, rule.KeyPrefix)
}

// isRuleMoreSpecific returns true only if "rule" is a subset of "other".
// For now this function ignores Type and only uses KeyPrefix.
func isRuleMoreSpecific(rule, other *wire.CrRule) bool {
	if other == nil {
		return true
	}
	if other.TableName == "" {
		// other.KeyPrefix must be empty if TableName is empty.
		return true
	}

	if rule.TableName != other.TableName {
		return false
	}
	return strings.HasPrefix(rule.KeyPrefix, other.KeyPrefix)
}
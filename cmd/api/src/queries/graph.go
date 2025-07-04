// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package queries

//go:generate go run go.uber.org/mock/mockgen -copyright_file=../../../../LICENSE.header -destination=./mocks/graph.go -package=mocks . Graph

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/specterops/bloodhound/analysis"
	"github.com/specterops/bloodhound/bhlog/measure"
	"github.com/specterops/bloodhound/cache"
	"github.com/specterops/bloodhound/cypher/analyzer"
	"github.com/specterops/bloodhound/cypher/frontend"
	"github.com/specterops/bloodhound/cypher/models/cypher/format"
	"github.com/specterops/bloodhound/dawgs/graph"
	"github.com/specterops/bloodhound/dawgs/ops"
	"github.com/specterops/bloodhound/dawgs/query"
	"github.com/specterops/bloodhound/dawgs/util"
	"github.com/specterops/bloodhound/graphschema"
	"github.com/specterops/bloodhound/graphschema/ad"
	"github.com/specterops/bloodhound/graphschema/azure"
	"github.com/specterops/bloodhound/graphschema/common"
	"github.com/specterops/bloodhound/src/api/bloodhoundgraph"
	"github.com/specterops/bloodhound/src/config"
	"github.com/specterops/bloodhound/src/model"
	"github.com/specterops/bloodhound/src/services/agi"
	"github.com/specterops/bloodhound/src/utils"
)

type SearchType = string

const (
	SearchTypeExact SearchType = "exact"
	SearchTypeFuzzy SearchType = "fuzzy"

	DefaultQueryFitnessLowerBoundSelector = -3
	DefaultQueryFitnessLowerBoundExplore  = -7
)

var (
	ErrUnsupportedDataType   = errors.New("unsupported result type for this query")
	ErrGraphUnsupported      = errors.New("type 'graph' is not supported for this endpoint")
	ErrCypherQueryTooComplex = errors.New("cypher query is too complex and is likely to result in poor or unstable database performance")
)

type EntityQueryParameters struct {
	QueryName     string
	ObjectID      string
	RequestedType model.DataType
	Skip          int
	Limit         int
	PathDelegate  any
	ListDelegate  any
}

func GetEntityObjectIDFromRequestPath(request *http.Request) (string, error) {
	if id, hasID := mux.Vars(request)["object_id"]; !hasID {
		return "", errors.New("no object ID found in request")
	} else {
		return id, nil
	}
}

func GetRequestedType(params url.Values) model.DataType {
	switch params.Get("type") {
	case "", "list":
		return model.DataTypeList
	case "graph":
		return model.DataTypeGraph
	case "count":
		return model.DataTypeCount
	default:
		return model.DataTypeCount
	}
}

func BuildEntityQueryParams(request *http.Request, queryName string, pathDelegate any, listDelegate any) (EntityQueryParameters, error) {
	var (
		requestQueryParams = request.URL.Query()
		dataType           = GetRequestedType(requestQueryParams)
	)

	if objectId, err := GetEntityObjectIDFromRequestPath(request); err != nil {
		return EntityQueryParameters{}, fmt.Errorf("error getting objectid: %w", err)
	} else if skip, limit, _, err := utils.GetPageParamsForGraphQuery(request.Context(), requestQueryParams); err != nil {
		return EntityQueryParameters{}, fmt.Errorf("error getting paging parameters: %w", err)
	} else {
		if dataType == model.DataTypeCount {
			skip = 0
			limit = 0
		}
		return EntityQueryParameters{
			QueryName:     queryName,
			ObjectID:      objectId,
			RequestedType: dataType,
			Skip:          skip,
			Limit:         limit,
			PathDelegate:  pathDelegate,
			ListDelegate:  listDelegate,
		}, nil
	}
}

type Graph interface {
	GetAssetGroupComboNode(ctx context.Context, owningObjectID string, assetGroupTag string) (map[string]any, error)
	GetAssetGroupNodes(ctx context.Context, assetGroupTag string, isSystemGroup bool) (graph.NodeSet, error)
	GetAllShortestPaths(ctx context.Context, startNodeID string, endNodeID string, filter graph.Criteria) (graph.PathSet, error)
	SearchNodesByName(ctx context.Context, nodeKinds graph.Kinds, nameQuery string, skip int, limit int) ([]model.SearchResult, error)
	SearchByNameOrObjectID(ctx context.Context, searchValue string, searchType string) (graph.NodeSet, error)
	GetADEntityQueryResult(ctx context.Context, params EntityQueryParameters, cacheEnabled bool) (any, int, error)
	GetEntityByObjectId(ctx context.Context, objectID string, kinds ...graph.Kind) (*graph.Node, error)
	GetEntityCountResults(ctx context.Context, node *graph.Node, delegates map[string]any) map[string]any
	GetNodesByKind(ctx context.Context, kinds ...graph.Kind) (graph.NodeSet, error)
	GetPrimaryNodeKindCounts(ctx context.Context, kind graph.Kind, additionalFilters ...graph.Criteria) (map[string]int, error)
	CountFilteredNodes(ctx context.Context, filterCriteria graph.Criteria) (int64, error)
	CountNodesByKind(ctx context.Context, kinds ...graph.Kind) (int64, error)
	GetFilteredAndSortedNodesPaginated(sortItems query.SortItems, filterCriteria graph.Criteria, offset, limit int) ([]*graph.Node, error)
	GetFilteredAndSortedNodes(sortItems query.SortItems, filterCriteria graph.Criteria) ([]*graph.Node, error)
	FetchNodesByObjectIDs(ctx context.Context, objectIDs ...string) (graph.NodeSet, error)
	FetchNodesByObjectIDsAndKinds(ctx context.Context, kinds graph.Kinds, objectIDs ...string) (graph.NodeSet, error)
	ValidateOUs(ctx context.Context, ous []string) ([]string, error)
	BatchNodeUpdate(ctx context.Context, nodeUpdate graph.NodeUpdate) error
	RawCypherQuery(ctx context.Context, pQuery PreparedQuery, includeProperties bool) (model.UnifiedGraph, error)
	PrepareCypherQuery(rawCypher string, queryComplexityLimit int64) (PreparedQuery, error)
	UpdateSelectorTags(ctx context.Context, db agi.AgiData, selectors model.UpdatedAssetGroupSelectors) error
	FetchNodeByGraphId(ctx context.Context, id graph.ID) (*graph.Node, error)
}

type GraphQuery struct {
	Graph                        graph.Database
	Cache                        cache.Cache
	SlowQueryThreshold           int64 // Threshold in milliseconds
	DisableCypherComplexityLimit bool
	EnableCypherMutations        bool
	cypherEmitter                format.Emitter
	strippedCypherEmitter        format.Emitter
}

func NewGraphQuery(graphDB graph.Database, cache cache.Cache, cfg config.Configuration) *GraphQuery {
	return &GraphQuery{
		Graph:                        graphDB,
		Cache:                        cache,
		SlowQueryThreshold:           cfg.SlowQueryThreshold,
		DisableCypherComplexityLimit: cfg.DisableCypherComplexityLimit,
		EnableCypherMutations:        cfg.EnableCypherMutations,
		cypherEmitter:                format.NewCypherEmitter(false),
		strippedCypherEmitter:        format.NewCypherEmitter(true),
	}
}

func (s *GraphQuery) GetAssetGroupComboNode(ctx context.Context, owningObjectID string, assetGroupTag string) (map[string]any, error) {
	var graphData = map[string]any{}

	return graphData, s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		if assetGroupNodes, err := ops.FetchNodeSet(tx.Nodes().Filterf(func() graph.Criteria {
			filters := []graph.Criteria{
				query.KindIn(query.Node(), azure.Entity, ad.Entity),
				query.StringContains(query.NodeProperty(common.SystemTags.String()), assetGroupTag),
			}

			if owningObjectID != "" {
				filters = append(filters, query.Or(
					query.Equals(query.NodeProperty(ad.DomainSID.String()), owningObjectID),
					query.Equals(query.NodeProperty(azure.TenantID.String()), owningObjectID),
				))
			}

			return query.And(filters...)
		})); err != nil {
			return err
		} else {
			if groups := assetGroupNodes.ContainingNodeKinds(ad.Group); groups.Len() > 0 {
				if groupMembershipPaths, err := analysis.ExpandGroupMembershipPaths(tx, groups); err != nil {
					return err
				} else {
					graphData = bloodhoundgraph.PathSetToBloodHoundGraph(groupMembershipPaths)

					for key := range graphData {
						// Skip the edges/relations and only evaluate the nodes.
						// Relations are prepended with "rel_" before the ID to distinguish them from edges. This was done
						// because neo4j reuses IDs across different object types, causing conflicts; adding that prefix
						// solves this issue.
						if id, err := strconv.Atoi(key); err != nil || strings.Contains(key, "rel") {
							continue
						} else {
							assetGroupNode := bloodhoundgraph.SetAssetGroupPropertiesForNode(groupMembershipPaths.AllNodes().Get(graph.ID(id)))
							graphData[key] = bloodhoundgraph.NodeToBloodHoundGraph(assetGroupNode)
						}
					}
				}
			}

			for _, node := range assetGroupNodes {
				node = bloodhoundgraph.SetAssetGroupPropertiesForNode(node)
				graphData[node.ID.String()] = bloodhoundgraph.NodeToBloodHoundGraph(node)
			}
		}

		return nil
	})
}

func (s *GraphQuery) GetAssetGroupNodes(ctx context.Context, assetGroupTag string, isSystemGroup bool) (graph.NodeSet, error) {
	var (
		assetGroupNodes graph.NodeSet
		err             error
	)

	err = s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		if assetGroupNodes, err = agi.FetchAssetGroupNodes(tx, assetGroupTag, isSystemGroup); err != nil {
			return err
		}
		return nil
	})

	return assetGroupNodes, err
}

func (s *GraphQuery) GetAllShortestPaths(ctx context.Context, startNodeID string, endNodeID string, filter graph.Criteria) (graph.PathSet, error) {
	defer measure.ContextMeasure(ctx, slog.LevelInfo, "GetAllShortestPaths")()

	var paths graph.PathSet

	return paths, s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		if startNode, err := analysis.FetchNodeByObjectID(tx, startNodeID); err != nil {
			return err
		} else if endNode, err := analysis.FetchNodeByObjectID(tx, endNodeID); err != nil {
			return err
		} else {
			criteria := []graph.Criteria{
				query.Equals(query.StartID(), startNode.ID),
				query.Equals(query.EndID(), endNode.ID),
			}

			if filter != nil {
				criteria = append(criteria, filter)
			}

			return tx.Relationships().Filter(query.And(criteria...)).FetchAllShortestPaths(func(cursor graph.Cursor[graph.Path]) error {
				for path := range cursor.Chan() {
					if len(path.Edges) > 0 {
						paths.AddPath(path)
					}
				}

				return cursor.Error()
			})
		}
	})
}

// the following negation clause matches nodes that have both ADLocalGroup and Group labels, but excludes nodes that only have the ADLocalGroup label.
// equivalent cypher: MATCH (n) WHERE NOT (n:ADLocalGroup AND NOT n:Group)
var groupFilter = query.Not(
	query.And(
		query.Kind(query.Node(), ad.LocalGroup),
		query.Not(query.Kind(query.Node(), ad.Group)),
	),
)

func SearchNodeByKindAndEqualsNameCriteria(kind graph.Kind, name string) graph.Criteria {
	return query.And(
		query.Kind(query.Node(), kind),
		query.Or(
			query.Equals(query.NodeProperty(common.Name.String()), name),
			query.Equals(query.NodeProperty(common.ObjectID.String()), name),
		),
		groupFilter,
	)
}

func searchNodeByKindAndContainsName(kind graph.Kind, name string) graph.Criteria {
	return query.And(
		query.Kind(query.Node(), kind),
		query.Or(
			query.StringContains(query.NodeProperty(common.Name.String()), name),
			query.StringContains(query.NodeProperty(common.ObjectID.String()), name),
		),
		query.Not(query.Equals(query.NodeProperty(common.Name.String()), name)),
		query.Not(query.Equals(query.NodeProperty(common.ObjectID.String()), name)),
		groupFilter,
	)
}

func formatSearchResults(exactResults []model.SearchResult, fuzzyResults []model.SearchResult, limit, skip int) []model.SearchResult {
	// Sort fuzzy results since they are all inexact matches based on the name passed in
	sort.Slice(fuzzyResults, func(i, j int) bool {
		return fuzzyResults[i].Name < fuzzyResults[j].Name
	})

	searchResults := make([]model.SearchResult, len(exactResults)+len(fuzzyResults))

	copy(searchResults, exactResults)
	copy(searchResults[len(exactResults):], fuzzyResults)

	length := len(searchResults)

	if skip > length {
		skip = length
	}

	end := skip + limit
	if end > length {
		end = length
	}

	return searchResults[skip:end]
}

func (s *GraphQuery) SearchNodesByName(ctx context.Context, nodeKinds graph.Kinds, name string, skip int, limit int) ([]model.SearchResult, error) {
	var (
		exactResults  []model.SearchResult
		fuzzyResults  []model.SearchResult
		formattedName = strings.ToUpper(name)
	)

	for _, kind := range nodeKinds {
		if err := s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
			if exactMatchNodes, err := ops.FetchNodes(tx.Nodes().Filter(SearchNodeByKindAndEqualsNameCriteria(kind, formattedName))); err != nil {
				return err
			} else {
				exactResults = append(exactResults, nodesToSearchResult(exactMatchNodes...)...)
			}

			if fuzzyMatchNodes, err := ops.FetchNodes(tx.Nodes().Filter(searchNodeByKindAndContainsName(kind, formattedName))); err != nil {
				return err
			} else {
				fuzzyResults = append(fuzzyResults, nodesToSearchResult(fuzzyMatchNodes...)...)
			}

			return nil
		}); err != nil {
			return []model.SearchResult{}, err
		}
	}

	return formatSearchResults(exactResults, fuzzyResults, limit, skip), nil
}

type PreparedQuery struct {
	query         string
	StrippedQuery string
	complexity    analyzer.ComplexityMeasure
	HasMutation   bool
}

func (s *GraphQuery) PrepareCypherQuery(rawCypher string, queryComplexityLimit int64) (PreparedQuery, error) {
	var (
		cypherFilters = []frontend.Visitor{
			&frontend.ExplicitProcedureInvocationFilter{},
			&frontend.ImplicitProcedureInvocationFilter{},
			&frontend.SpecifiedParametersFilter{},
		}
		queryBuffer         = &bytes.Buffer{}
		strippedQueryBuffer = &bytes.Buffer{}
		graphQuery          PreparedQuery
	)

	// If cypher mutations are disabled, we want to add the updating clause filter to properly error as unsupported query
	// If we are mutating, make sure our expansions aren't included in any sort of update
	if !s.EnableCypherMutations {
		cypherFilters = append(cypherFilters, &frontend.UpdatingNotAllowedClauseFilter{})
	} else {
		cypherFilters = append(cypherFilters, &frontend.UpdatingClauseFilter{})
	}

	parseCtx := frontend.NewContext(cypherFilters...)

	queryModel, err := frontend.ParseCypher(parseCtx, rawCypher)
	if err != nil {
		return graphQuery, err
	}

	graphQuery.HasMutation = parseCtx.HasMutation

	complexityMeasure, err := analyzer.QueryComplexity(queryModel)
	if err != nil {
		return graphQuery, err
	} else if err = s.strippedCypherEmitter.Write(queryModel, strippedQueryBuffer); err != nil {
		return graphQuery, err
	} else if !s.DisableCypherComplexityLimit && complexityMeasure.RelativeFitness <= queryComplexityLimit {
		// log query details if it is rejected due to poor fitness
		slog.Error(
			fmt.Sprintf("Query rejected. Query weight: %d. Maximum allowed weight: %d", complexityMeasure.RelativeFitness, queryComplexityLimit),
			"query", strippedQueryBuffer.String(),
		)

		return graphQuery, ErrCypherQueryTooComplex
	}

	graphQuery.StrippedQuery = strippedQueryBuffer.String()
	graphQuery.complexity = complexityMeasure

	if err = s.cypherEmitter.Write(queryModel, queryBuffer); err != nil {
		return graphQuery, err
	} else {
		graphQuery.query = queryBuffer.String()
	}

	return graphQuery, nil
}

// RawCypherQuery executes the given PreparedQuery and returns a model.UnifiedGraph or any error encountered during
// query execution.
func (s *GraphQuery) RawCypherQuery(ctx context.Context, pQuery PreparedQuery, includeProperties bool) (model.UnifiedGraph, error) {
	var (
		err error

		graphResponse = model.NewUnifiedGraph()
		start         = time.Now()

		txDelegate = func(tx graph.Transaction) error {
			if pathSet, err := ops.FetchPathSetByQuery(tx, pQuery.query); err != nil {
				return err
			} else {
				graphResponse.AddPathSet(pathSet, includeProperties)
			}

			return nil
		}
	)

	slog.InfoContext(
		ctx,
		"Preparing user cypher query",
		slog.String("query", pQuery.StrippedQuery),
		slog.Int64("fitness", pQuery.complexity.RelativeFitness),
	)

	if pQuery.HasMutation {
		// If the mutation is complex it is still worth spinning it into a write transaction in case it fails,
		// deadlocks or otherwise rolls back
		err = s.Graph.WriteTransaction(ctx, txDelegate)
	} else {
		err = s.Graph.ReadTransaction(ctx, txDelegate)
	}

	slog.InfoContext(
		ctx,
		"Executed user cypher query",
		slog.String("query", pQuery.StrippedQuery),
		slog.Int64("fitness", pQuery.complexity.RelativeFitness),
		slog.Duration("elapsed", time.Since(start)),
	)

	if err != nil {
		// Log query details if neo4j times out
		if util.IsNeoTimeoutError(err) {
			slog.Error("Neo4j timed out while executing cypher query",
				"query", pQuery.StrippedQuery,
				"query cost", fmt.Sprintf("%d", pQuery.complexity.RelativeFitness),
			)
		} else {
			slog.WarnContext(ctx, fmt.Sprintf("RawCypherQuery failed: %v", err))
		}
	}

	return graphResponse, err
}

func applyTimeoutReduction(queryWeight int64, availableRuntime time.Duration) (time.Duration, int64) {
	// The weight of the query is divided by 5 to get a runtime reduction factor, in a way that:
	// weights of 4 or less get the full runtime duration
	// weights of 5-9 will get 1/2 the runtime duration
	// weights of 10-15 will get 1/3 the runtime duration
	// and so on until the max weight of 50 gets 1/11 the runtime duration
	reductionFactor := 1 + (queryWeight / 5)

	availableRuntimeInt := int64(availableRuntime.Seconds())
	// reductionFactor will be the math.Floor() of the result of the division below
	availableRuntimeInt /= reductionFactor
	availableRuntime = time.Duration(availableRuntimeInt) * time.Second

	return availableRuntime, reductionFactor
}

func nodeToSearchResult(node *graph.Node) model.SearchResult {
	var (
		name, _              = node.Properties.GetWithFallback(common.Name.String(), "NO NAME", common.DisplayName.String(), common.ObjectID.String()).String()
		objectID, _          = node.Properties.GetOrDefault(common.ObjectID.String(), "NO OBJECT ID").String()
		distinguishedName, _ = node.Properties.GetOrDefault(ad.DistinguishedName.String(), "").String()
		systemTags, _        = node.Properties.GetOrDefault(common.SystemTags.String(), "").String()
	)

	return model.SearchResult{
		ObjectID:          objectID,
		Type:              analysis.GetNodeKindDisplayLabel(node),
		Name:              name,
		DistinguishedName: distinguishedName,
		SystemTags:        systemTags,
	}
}

func nodesToSearchResult(nodes ...*graph.Node) []model.SearchResult {
	searchResults := make([]model.SearchResult, len(nodes))

	for idx, node := range nodes {
		searchResults[idx] = nodeToSearchResult(node)
	}

	return searchResults
}

func (s *GraphQuery) SearchByNameOrObjectID(ctx context.Context, searchValue string, searchType SearchType) (graph.NodeSet, error) {
	var nodes = graph.NewNodeSet()

	for _, kind := range []graph.Kind{ad.Entity, azure.Entity} {
		if err := s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
			if fetchedNodes, err := ops.FetchNodeSet(tx.Nodes().Filterf(func() graph.Criteria {
				if searchType == SearchTypeExact {
					return query.And(
						query.Kind(query.Node(), kind),
						query.Or(
							query.Equals(query.NodeProperty(common.Name.String()), strings.ToUpper(searchValue)),
							query.Equals(query.NodeProperty(common.ObjectID.String()), strings.ToUpper(searchValue)),
						),
					)
				} else {
					return query.And(
						query.Kind(query.Node(), kind),
						query.Or(
							query.StringStartsWith(query.NodeProperty(common.Name.String()), strings.ToUpper(searchValue)),
							query.StringStartsWith(query.NodeProperty(common.ObjectID.String()), strings.ToUpper(searchValue)),
						),
					)
				}
			})); err != nil {
				return err
			} else {
				nodes.AddSet(fetchedNodes)
				return nil
			}
		}); err != nil {
			return nil, err
		}
	}

	return nodes, nil
}

func (s *GraphQuery) GetADEntityQueryResult(ctx context.Context, params EntityQueryParameters, cacheEnabled bool) (any, int, error) {
	if params.RequestedType == model.DataTypeGraph && params.PathDelegate == nil {
		return nil, 0, ErrGraphUnsupported
	}

	if params.RequestedType == model.DataTypeCount || params.RequestedType == model.DataTypeList && params.ListDelegate == nil {
		return nil, 0, ErrUnsupportedDataType
	}

	if node, err := s.GetEntityByObjectId(ctx, params.ObjectID, ad.Entity); err != nil {
		return nil, 0, fmt.Errorf("error getting entity node: %w", err)
	} else {
		return s.GetEntityResults(ctx, node, params, cacheEnabled)
	}
}

func (s *GraphQuery) GetEntityByObjectId(ctx context.Context, objectID string, kinds ...graph.Kind) (*graph.Node, error) {
	var (
		node *graph.Node
		err  error
	)
	if err := s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		if node, err = tx.Nodes().Filterf(func() graph.Criteria {
			return query.And(
				query.Equals(query.NodeProperty(common.ObjectID.String()), objectID),
				query.KindIn(query.Node(), kinds...),
			)
		}).First(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	} else {
		return node, nil
	}
}

func (s *GraphQuery) GetEntityCountResults(ctx context.Context, node *graph.Node, delegates map[string]any) map[string]any {
	var (
		results   = make(map[string]any)
		data      sync.Map
		waitGroup sync.WaitGroup
	)

	for delegateKey, delegate := range delegates {
		waitGroup.Add(1)

		slog.DebugContext(ctx, fmt.Sprintf("Running entity query %s", delegateKey))

		go func(delegateKey string, delegate any) {
			defer waitGroup.Done()

			if result, err := runEntityQuery(ctx, s.Graph, delegate, node, 0, 0); errors.Is(err, graph.ErrContextTimedOut) {
				slog.WarnContext(ctx, fmt.Sprintf("Running entity query for key %s: %v", delegateKey, err))
			} else if err != nil {
				slog.ErrorContext(ctx, fmt.Sprintf("Error running entity query for key %s: %v", delegateKey, err))
				data.Store(delegateKey, 0)
			} else {
				data.Store(delegateKey, result.Len())
			}
		}(delegateKey, delegate)
	}

	waitGroup.Wait()

	data.Range(func(k any, v any) bool {
		results[k.(string)] = v
		return true
	})

	results["props"] = node.Properties.Map
	return results
}

func (s *GraphQuery) CountFilteredNodes(ctx context.Context, filterCriteria graph.Criteria) (int64, error) {
	var numNodes int64

	return numNodes, s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		var err error
		numNodes, err = tx.Nodes().Filter(filterCriteria).Count()
		return err
	})
}

func (s *GraphQuery) CountNodesByKind(ctx context.Context, kinds ...graph.Kind) (int64, error) {
	return s.CountFilteredNodes(ctx, (query.KindIn(query.Node(), kinds...)))
}

func (s *GraphQuery) FetchNodeByGraphId(ctx context.Context, id graph.ID) (*graph.Node, error) {
	var node *graph.Node

	if err := s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		var err error
		node, err = ops.FetchNode(tx, id)
		return err
	}); err != nil {
		return nil, err
	} else if node == nil {
		return nil, fmt.Errorf("node not found for id: %s", id)
	} else {
		return node, err
	}
}

func (s *GraphQuery) GetPrimaryNodeKindCounts(ctx context.Context, kind graph.Kind, additionalFilters ...graph.Criteria) (map[string]int, error) {
	var (
		results = map[string]int{}
		filters = []graph.Criteria{query.KindIn(query.Node(), kind)}
	)

	if additionalFilters != nil {
		filters = append(filters, additionalFilters...)
	}

	return results, s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		return tx.Nodes().Filter(query.And(filters...)).FetchKinds(func(cursor graph.Cursor[graph.KindsResult]) error {
			for next := range cursor.Chan() {
				primaryKindStr := graphschema.PrimaryNodeKind(next.Kinds).String()
				results[primaryKindStr] += 1
			}

			return cursor.Error()
		})
	})
}

func (s *GraphQuery) GetNodesByKind(ctx context.Context, kinds ...graph.Kind) (graph.NodeSet, error) {
	var nodes graph.NodeSet

	return nodes, s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		if fetchedNodes, err := ops.FetchNodeSet(tx.Nodes().Filterf(func() graph.Criteria {
			return query.KindIn(query.Node(), kinds...)
		})); err != nil {
			return err
		} else {
			nodes = fetchedNodes
		}
		return nil
	})
}

func (s *GraphQuery) GetFilteredAndSortedNodes(sortItems query.SortItems, filterCriteria graph.Criteria) ([]*graph.Node, error) {
	return s.GetFilteredAndSortedNodesPaginated(sortItems, filterCriteria, 0, 0)
}

func (s *GraphQuery) GetFilteredAndSortedNodesPaginated(sortItems query.SortItems, filterCriteria graph.Criteria, offset, limit int) ([]*graph.Node, error) {
	var (
		nodes         []*graph.Node
		finalCriteria []graph.Criteria
	)

	return nodes, s.Graph.ReadTransaction(context.Background(), func(tx graph.Transaction) error {
		nodeQuery := tx.Nodes().Filterf(func() graph.Criteria {
			return filterCriteria
		})

		if offset > 0 {
			finalCriteria = append(finalCriteria, query.Offset(offset))
		}

		if limit > 0 {
			finalCriteria = append(finalCriteria, query.Limit(limit))
		}

		if len(sortItems) > 0 {
			finalCriteria = append(finalCriteria, sortItems.FormatCypherOrder())
		}

		return nodeQuery.Fetch(func(cursor graph.Cursor[*graph.Node]) error {
			for node := range cursor.Chan() {
				nodes = append(nodes, node)
			}
			return nil
		}, finalCriteria...)
	})
}

// FetchNodesByObjectIDs takes a list of objectIDs. Returns a graph.NodeSet for found results
// and an error for graph database errors.
func (s *GraphQuery) FetchNodesByObjectIDs(ctx context.Context, objectIDs ...string) (graph.NodeSet, error) {
	var nodes graph.NodeSet

	return nodes, s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		if fetchedNodes, err := ops.FetchNodeSet(tx.Nodes().Filterf(
			func() graph.Criteria {
				return query.And(
					query.KindIn(query.Node(), ad.Entity, azure.Entity),
					query.In(query.NodeProperty(common.ObjectID.String()), objectIDs),
				)
			}),
		); err != nil {
			return err
		} else {
			nodes = fetchedNodes
			return nil
		}
	})
}

func (s *GraphQuery) FetchNodesByObjectIDsAndKinds(ctx context.Context, kinds graph.Kinds, objectIDs ...string) (graph.NodeSet, error) {
	var nodes graph.NodeSet

	return nodes, s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
		if fetchedNodes, err := ops.FetchNodeSet(tx.Nodes().Filterf(
			func() graph.Criteria {
				return query.And(
					query.KindIn(query.Node(), kinds...),
					query.In(query.NodeProperty(common.ObjectID.String()), objectIDs),
				)
			}),
		); err != nil {
			return err
		} else {
			nodes = fetchedNodes
			return nil
		}
	})
}

func (s *GraphQuery) ValidateOUs(ctx context.Context, ous []string) ([]string, error) {
	var validated = make([]string, 0)

	for _, ou := range ous {
		if err := s.Graph.ReadTransaction(ctx, func(tx graph.Transaction) error {
			if node, err := tx.Nodes().Filterf(func() graph.Criteria {
				if strings.HasPrefix(strings.ToLower(ou), "ou=") {
					return query.And(
						query.Kind(query.Node(), ad.OU),
						query.Equals(query.NodeProperty(ad.DistinguishedName.String()), ou))
				}
				return query.And(
					query.KindIn(query.Node(), ad.Entity, azure.Entity),
					query.Equals(query.NodeProperty(common.ObjectID.String()), ou),
				)
			}).First(); err != nil {
				return err
			} else {
				if objectID, err := node.Properties.Get(common.ObjectID.String()).String(); err != nil {
					return err
				} else {
					validated = append(validated, objectID)
				}
			}
			return nil
		}); err != nil {
			if graph.IsErrNotFound(err) {
				return nil, fmt.Errorf("no record found for %s", ou)
			} else {
				return nil, err
			}
		}
	}

	return validated, nil
}

func (s *GraphQuery) BatchNodeUpdate(ctx context.Context, nodeUpdate graph.NodeUpdate) error {
	return s.Graph.WriteTransaction(ctx, func(tx graph.Transaction) error {
		updateNodeFunc := func(batch graph.Batch) error {
			return batch.UpdateNodeBy(nodeUpdate)
		}

		return s.Graph.BatchOperation(ctx, updateNodeFunc)
	})
}

func nodeSetToOrderedSlice(nodeSet graph.NodeSet) []*graph.Node {
	nodes := nodeSet.Slice()

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].ID > nodes[j].ID
	})

	return nodes
}

func (s *GraphQuery) cacheQueryResult(queryStart time.Time, cacheKey string, result graph.NodeSet) {
	queryTime := time.Since(queryStart).Milliseconds()

	// Only cache the result if it matches our criteria, including having a valid query name
	if queryTime > s.SlowQueryThreshold {
		// Using GuardedSet here even though it isn't necessary because it allows us to collect information on how often
		// we run these queries in parallel
		if set, sizeInBytes, err := s.Cache.GuardedSet(cacheKey, result); err != nil {
			slog.Error(fmt.Sprintf("[Entity Results Cache] Failed to write results to cache for key: %s", cacheKey))
		} else if !set {
			slog.Warn(fmt.Sprintf("[Entity Results Cache] Cache entry for query %s not set because it already exists", cacheKey))
		} else {
			slog.Info(fmt.Sprintf("[Entity Results Cache] Cached slow query %s (%d bytes) because it took %dms", cacheKey, sizeInBytes, queryTime))
		}
	}
}

func runEntityQuery(ctx context.Context, db graph.Database, delegate any, node *graph.Node, skip, limit int) (graph.NodeSet, error) {
	var result graph.NodeSet

	switch typedDelegate := delegate.(type) {
	case analysis.ListDelegate:
		if err := db.ReadTransaction(ctx, func(tx graph.Transaction) error {
			if fetchedResult, err := typedDelegate(tx, node, skip, limit); err != nil {
				return err
			} else {
				result = fetchedResult
			}

			return nil
		}); err != nil {
			return nil, err
		}

	case analysis.ParallelListDelegate:
		if fetchedResult, err := typedDelegate(ctx, db, node, skip, limit); err != nil {
			return nil, err
		} else {
			result = fetchedResult
		}

	default:
		return nil, fmt.Errorf("unsupported list delegate type %T", typedDelegate)
	}

	return result, nil
}

func (s *GraphQuery) runMaybeCachedEntityQuery(ctx context.Context, node *graph.Node, params EntityQueryParameters, cacheEnabled bool) (graph.NodeSet, error) {
	var (
		queryStart = time.Now()
		cacheKey   = fmt.Sprintf("ad-entity-query_%s_%s_%d", params.QueryName, params.ObjectID, params.RequestedType)

		foundResultInCache = false

		result graph.NodeSet
	)

	if cacheEnabled {
		var err error
		if foundResultInCache, err = s.Cache.Get(cacheKey, &result); err != nil {
			return nil, fmt.Errorf("error getting cache entry for %s: %w", cacheKey, err)
		}
	}

	if !cacheEnabled || !foundResultInCache {
		// Fetch the entire result for caching purposes
		if fetchedResult, err := runEntityQuery(ctx, s.Graph, params.ListDelegate, node, 0, 0); err != nil {
			return nil, err
		} else {
			result = fetchedResult
		}
	}

	if params.QueryName != "" && cacheEnabled && !foundResultInCache {
		s.cacheQueryResult(queryStart, cacheKey, result)
	}

	return result, nil
}

func (s *GraphQuery) runListQuery(ctx context.Context, node *graph.Node, params EntityQueryParameters, cacheEnabled bool) ([]model.PagedNodeListEntry, int, error) {
	var (
		skip  = params.Skip
		limit = params.Limit
	)

	if result, err := s.runMaybeCachedEntityQuery(ctx, node, params, cacheEnabled); err != nil {
		return nil, 0, err
	} else if skip > result.Len() {
		return nil, 0, fmt.Errorf(utils.ErrorInvalidSkip, skip)
	} else {
		if skip+limit > result.Len() {
			limit = result.Len() - skip
		}

		return fromGraphNodes(graph.NewNodeSet(nodeSetToOrderedSlice(result)[skip : skip+limit]...)), result.Len(), nil
	}
}

func (s *GraphQuery) runCountQuery(ctx context.Context, node *graph.Node, params EntityQueryParameters, cacheEnabled bool) (any, int, error) {
	result, err := s.runMaybeCachedEntityQuery(ctx, node, params, cacheEnabled)
	return nil, result.Len(), err
}

func runPathQuery(ctx context.Context, db graph.Database, node *graph.Node, pathDelegate any) (map[string]any, int, error) {
	var (
		result graph.PathSet
		err    error
	)

	switch typedDelegate := pathDelegate.(type) {
	case analysis.PathDelegate:
		err = db.ReadTransaction(ctx, func(tx graph.Transaction) error {
			if fetchedResult, err := typedDelegate(tx, node); err != nil {
				return err
			} else {
				result = fetchedResult
			}

			return nil
		})
	case analysis.ParallelPathDelegate:
		result, err = typedDelegate(ctx, db, node)
	default:
		err = fmt.Errorf("unsupported path delegate type %T", typedDelegate)
	}

	if err != nil {
		return nil, 0, err
	} else {
		return bloodhoundgraph.PathSetToBloodHoundGraph(result), result.Len(), nil
	}
}

func (s *GraphQuery) GetEntityResults(ctx context.Context, node *graph.Node, params EntityQueryParameters, cacheEnabled bool) (any, int, error) {
	// Graph type isn't currently under a caching model and is handled separately from other supported RequestedTypes
	switch params.RequestedType {
	case model.DataTypeGraph:
		return runPathQuery(ctx, s.Graph, node, params.PathDelegate)
	case model.DataTypeList:
		return s.runListQuery(ctx, node, params, cacheEnabled)
	case model.DataTypeCount:
		return s.runCountQuery(ctx, node, params, cacheEnabled)
	default:
		return nil, 0, fmt.Errorf("unknown data type requested")
	}
}

func fromGraphNodes(nodes graph.NodeSet) []model.PagedNodeListEntry {
	renderedNodes := make([]model.PagedNodeListEntry, 0, nodes.Len())

	for _, node := range nodes {
		var (
			nodeEntry model.PagedNodeListEntry
			props     = node.Properties
		)

		if objectId, err := props.Get(common.ObjectID.String()).String(); err != nil {
			slog.Error(fmt.Sprintf("Error getting objectid for %d: %v", node.ID, err))
			nodeEntry.ObjectID = ""
		} else {
			nodeEntry.ObjectID = objectId
		}

		if name, err := props.Get(common.Name.String()).String(); err != nil {
			slog.Error(fmt.Sprintf("Error getting name for %d: %v", node.ID, err))
			nodeEntry.Name = ""
		} else {
			nodeEntry.Name = name
		}

		nodeEntry.Label = analysis.GetNodeKindDisplayLabel(node)

		renderedNodes = append(renderedNodes, nodeEntry)
	}

	return renderedNodes
}

func (s *GraphQuery) UpdateSelectorTags(ctx context.Context, db agi.AgiData, selectors model.UpdatedAssetGroupSelectors) error {
	for _, selector := range selectors.Added {
		if err := addTagsToSelector(ctx, s, db, selector); err != nil {
			return err
		}
	}

	for _, selector := range selectors.Removed {
		if err := removeTagsFromSelector(ctx, s, db, selector); err != nil {
			return err
		}
	}
	return nil
}

func addTagsToSelector(ctx context.Context, graphQuery *GraphQuery, db agi.AgiData, selector model.AssetGroupSelector) error {
	if assetGroup, err := db.GetAssetGroup(ctx, selector.AssetGroupID); err != nil {
		return err
	} else {
		return graphQuery.Graph.WriteTransaction(ctx, func(tx graph.Transaction) error {
			tagPropertyStr := common.SystemTags.String()

			if !assetGroup.SystemGroup {
				tagPropertyStr = common.UserTags.String()
			}

			if node, err := analysis.FetchNodeByObjectID(tx, selector.Selector); err != nil {
				return err
			} else {
				if tags, err := node.Properties.Get(tagPropertyStr).String(); err != nil {
					if graph.IsErrPropertyNotFound(err) {
						node.Properties.Set(tagPropertyStr, assetGroup.Tag)
					} else {
						return err
					}
				} else if !strings.Contains(tags, assetGroup.Tag) {
					if len(tags) == 0 {
						node.Properties.Set(tagPropertyStr, assetGroup.Tag)
					} else { // add a space and append if there are existing tags
						node.Properties.Set(tagPropertyStr, tags+" "+assetGroup.Tag)
					}
				}

				if err = tx.UpdateNode(node); err != nil {
					return err
				}
			}

			return nil
		})
	}
}

func removeTagsFromSelector(ctx context.Context, graphQuery *GraphQuery, db agi.AgiData, selector model.AssetGroupSelector) error {
	if assetGroup, err := db.GetAssetGroup(ctx, selector.AssetGroupID); err != nil {
		return err
	} else {
		return graphQuery.Graph.WriteTransaction(ctx, func(tx graph.Transaction) error {
			tagPropertyStr := common.SystemTags.String()

			if !assetGroup.SystemGroup {
				tagPropertyStr = common.UserTags.String()
			}

			if node, err := analysis.FetchNodeByObjectID(tx, selector.Selector); err != nil {
				return err
			} else {
				if tags, err := node.Properties.Get(tagPropertyStr).String(); err != nil {
					if graph.IsErrPropertyNotFound(err) {
						node.Properties.Set(tagPropertyStr, assetGroup.Tag)
					} else {
						return err
					}
				} else if strings.Contains(tags, assetGroup.Tag) {
					// remove asset group tag and then remove any leftover double whitespace
					tags = strings.ReplaceAll(strings.ReplaceAll(tags, assetGroup.Tag, ""), "  ", " ")
					node.Properties.Set(tagPropertyStr, tags)
				}

				if err = tx.UpdateNode(node); err != nil {
					return err
				}
			}

			return nil
		})
	}
}

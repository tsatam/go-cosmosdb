package cosmosdb

import (
	"context"
	"net/http"
)

// Collection represents a collection
type Collection struct {
	ID                       string                    `json:"id,omitempty"`
	ResourceID               string                    `json:"_rid,omitempty"`
	Timestamp                int                       `json:"_ts,omitempty"`
	Self                     string                    `json:"_self,omitempty"`
	ETag                     string                    `json:"_etag,omitempty"`
	Documents                string                    `json:"_docs,omitempty"`
	StoredProcedures         string                    `json:"_sprocs,omitempty"`
	Triggers                 string                    `json:"_triggers,omitempty"`
	UserDefinedFunctions     string                    `json:"_udfs,omitempty"`
	Conflicts                string                    `json:"_conflicts,omitempty"`
	IndexingPolicy           *IndexingPolicy           `json:"indexingPolicy,omitempty"`
	PartitionKey             *PartitionKey             `json:"partitionKey,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `json:"uniqueKeyPolicy,omitempty"`
	ConflictResolutionPolicy *ConflictResolutionPolicy `json:"conflictResolutionPolicy,omitempty"`
	AllowMaterializedViews   bool                      `json:"allowMaterializedViews,omitempty"`
	GeospatialConfig         *GeospatialConfig         `json:"geospatialConfig,omitempty"`
}

// IndexingPolicy represents an indexing policy
type IndexingPolicy struct {
	Automatic        bool               `json:"automatic,omitempty"`
	IndexingMode     IndexingPolicyMode `json:"indexingMode,omitempty"`
	IncludedPaths    []IncludedPath     `json:"includedPaths,omitempty"`
	ExcludedPaths    []IncludedPath     `json:"excludedPaths,omitempty"`
	CompositeIndexes []CompositeIndex   `json:"compositeIndexes,omitempty"`
}

// IndexingPolicyMode represents an indexing policy mode
type IndexingPolicyMode string

// IndexingPolicyMode constants
const (
	IndexingPolicyModeConsistent IndexingPolicyMode = "Consistent"
	IndexingPolicyModeLazy       IndexingPolicyMode = "Lazy"
)

// IncludedPath represents an included path
type IncludedPath struct {
	Path    string  `json:"path,omitempty"`
	Indexes []Index `json:"indexes,omitempty"`
}

// Index represents an index
type Index struct {
	DataType  IndexDataType `json:"dataType,omitempty"`
	Kind      IndexKind     `json:"kind,omitempty"`
	Precision int           `json:"precision,omitempty"`
}

// IndexDataType represents an index data type
type IndexDataType string

// IndexDataType constants
const (
	IndexDataTypeString     IndexDataType = "String"
	IndexDataTypeNumber     IndexDataType = "Number"
	IndexDataTypePoint      IndexDataType = "Point"
	IndexDataTypePolygon    IndexDataType = "Polygon"
	IndexDataTypeLineString IndexDataType = "LineString"
)

// IndexKind represents an index kind
type IndexKind string

// IndexKind constants
const (
	IndexKindHash    IndexKind = "Hash"
	IndexKindRange   IndexKind = "Range"
	IndexKindSpatial IndexKind = "Spatial"
)

// ExcludedPath represents an excluded path
type ExcludedPath struct {
	Path string `json:"path,omitempty"`
}

// CompositeIndex represents a composite index
type CompositeIndex []struct {
	Path  string `json:"path,omitempty"`
	Order Order  `json:"order,omitempty"`
}

// Order represents an order
type Order string

// Order constants
const (
	OrderAscending  Order = "ascending"
	OrderDescending Order = "descending"
)

// PartitionKey represents a partition key
type PartitionKey struct {
	Paths   []string         `json:"paths,omitempty"`
	Kind    PartitionKeyKind `json:"kind,omitempty"`
	Version int              `json:"version,omitempty"`
}

// PartitionKeyKind represents a partition key kind
type PartitionKeyKind string

// PartitionKeyKind constants
const (
	PartitionKeyKindHash PartitionKeyKind = "Hash"
)

// UniqueKeyPolicy represents a unique key policy
type UniqueKeyPolicy struct {
	UniqueKeys []UniqueKey `json:"uniqueKeys,omitempty"`
}

// UniqueKey represents a unique key
type UniqueKey struct {
	Paths []string `json:"paths,omitempty"`
}

// ConflictResolutionPolicy represents a conflict resolution policy
type ConflictResolutionPolicy struct {
	Mode                        ConflictResolutionPolicyMode `json:"mode,omitempty"`
	ConflictResolutionPath      string                       `json:"conflictResolutionPath,omitempty"`
	ConflictResolutionProcedure string                       `json:"conflictResolutionProcedure,omitempty"`
}

// ConflictResolutionPolicyMode represents a conflict resolution policy mode
type ConflictResolutionPolicyMode string

// ConflictResolutionPolicyMode constants
const (
	ConflictResolutionPolicyModeLastWriterWins ConflictResolutionPolicyMode = "LastWriterWins"
	ConflictResolutionPolicyModeCustom         ConflictResolutionPolicyMode = "Custom"
)

// GeospatialConfig represents a geospatial config
type GeospatialConfig struct {
	Type GeospatialConfigType `json:"type,omitempty"`
}

// GeospatialConfigType represents geospatial config types
type GeospatialConfigType string

// GeospatialConfigType constants
const (
	GeospatialConfigTypeGeography GeospatialConfigType = "Geography"
)

// Collections represents collections
type Collections struct {
	Count       int           `json:"_count,omitempty"`
	ResourceID  string        `json:"_rid,omitempty"`
	Collections []*Collection `json:"DocumentCollections,omitempty"`
}

// PartitionKeyRanges represents partition key ranges
type PartitionKeyRanges struct {
	Count              int                 `json:"_count,omitempty"`
	ResourceID         string              `json:"_rid,omitempty"`
	PartitionKeyRanges []PartitionKeyRange `json:"PartitionKeyRanges,omitempty"`
}

// PartitionKeyRange represents a partition key range
type PartitionKeyRange struct {
	MissingFields

	ID                 string                  `json:"id,omitempty"`
	ResourceID         string                  `json:"_rid,omitempty"`
	Timestamp          int                     `json:"_ts,omitempty"`
	Self               string                  `json:"_self,omitempty"`
	ETag               string                  `json:"_etag,omitempty"`
	MaxExclusive       string                  `json:"maxExclusive,omitempty"`
	MinInclusive       string                  `json:"minInclusive,omitempty"`
	ResourceIDPrefix   int                     `json:"ridPrefix,omitempty"`
	ThroughputFraction float64                 `json:"throughputFraction,omitempty"`
	Status             PartitionKeyRangeStatus `json:"status,omitempty"`
	Parents            []string                `json:"parents,omitempty"`
	LSN                int                     `json:"lsn,omitempty"`
}

// PartitionKeyRangeStatus represents a partition key range status
type PartitionKeyRangeStatus string

// PartitionKeyRangeStatus constants
const (
	PartitionKeyRangeStatusOnline PartitionKeyRangeStatus = "online"
)

// MissingFields retains values that do not map to struct fields during JSON
// marshalling/unmarshalling.  MissingFields implements
// github.com/ugorji/go/codec.MissingFielder.
type MissingFields struct {
	m map[string]interface{}
}

// CodecMissingField is called to set a missing field and value pair
func (mf *MissingFields) CodecMissingField(field []byte, value interface{}) bool {
	if mf.m == nil {
		mf.m = map[string]interface{}{}
	}
	(mf.m)[string(field)] = value
	return true
}

// CodecMissingFields returns the set of fields which are not struct fields
func (mf *MissingFields) CodecMissingFields() map[string]interface{} {
	return mf.m
}

type collectionClient struct {
	*databaseClient
	path string
}

// CollectionClient is a collection client
type CollectionClient interface {
	Create(context.Context, *Collection) (*Collection, error)
	List() CollectionIterator
	ListAll(context.Context) (*Collections, error)
	Get(context.Context, string) (*Collection, error)
	Delete(context.Context, *Collection) error
	Replace(context.Context, *Collection) (*Collection, error)
	PartitionKeyRanges(context.Context, string) (*PartitionKeyRanges, error)
}

type collectionListIterator struct {
	*collectionClient
	continuation string
	done         bool
}

// CollectionIterator is a collection iterator
type CollectionIterator interface {
	Next(context.Context) (*Collections, error)
}

// NewCollectionClient returns a new collection client
func NewCollectionClient(c DatabaseClient, dbid string) CollectionClient {
	return &collectionClient{
		databaseClient: c.(*databaseClient),
		path:           "dbs/" + dbid,
	}
}

func (c *collectionClient) all(ctx context.Context, i CollectionIterator) (*Collections, error) {
	allcolls := &Collections{}

	for {
		colls, err := i.Next(ctx)
		if err != nil {
			return nil, err
		}
		if colls == nil {
			break
		}

		allcolls.Count += colls.Count
		allcolls.ResourceID = colls.ResourceID
		allcolls.Collections = append(allcolls.Collections, colls.Collections...)
	}

	return allcolls, nil
}

func (c *collectionClient) Create(ctx context.Context, newcoll *Collection) (coll *Collection, err error) {
	err = c.do(ctx, http.MethodPost, c.path+"/colls", "colls", c.path, http.StatusCreated, &newcoll, &coll, nil)
	return
}

func (c *collectionClient) List() CollectionIterator {
	return &collectionListIterator{collectionClient: c}
}

func (c *collectionClient) ListAll(ctx context.Context) (*Collections, error) {
	return c.all(ctx, c.List())
}

func (c *collectionClient) Get(ctx context.Context, collid string) (coll *Collection, err error) {
	err = c.do(ctx, http.MethodGet, c.path+"/colls/"+collid, "colls", c.path+"/colls/"+collid, http.StatusOK, nil, &coll, nil)
	return
}

func (c *collectionClient) Delete(ctx context.Context, coll *Collection) error {
	if coll.ETag == "" {
		return ErrETagRequired
	}
	headers := http.Header{}
	headers.Set("If-Match", coll.ETag)
	return c.do(ctx, http.MethodDelete, c.path+"/colls/"+coll.ID, "colls", c.path+"/colls/"+coll.ID, http.StatusNoContent, nil, nil, headers)
}

func (c *collectionClient) Replace(ctx context.Context, newcoll *Collection) (coll *Collection, err error) {
	err = c.do(ctx, http.MethodPost, c.path+"/colls/"+newcoll.ID, "colls", c.path+"/colls/"+newcoll.ID, http.StatusCreated, &newcoll, &coll, nil)
	return
}

func (c *collectionClient) PartitionKeyRanges(ctx context.Context, collid string) (pkrs *PartitionKeyRanges, err error) {
	err = c.do(ctx, http.MethodGet, c.path+"/colls/"+collid+"/pkranges", "pkranges", c.path+"/colls/"+collid, http.StatusOK, nil, &pkrs, nil)
	return
}

func (i *collectionListIterator) Next(ctx context.Context) (colls *Collections, err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.do(ctx, http.MethodGet, i.path+"/colls", "colls", i.path, http.StatusOK, nil, &colls, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}

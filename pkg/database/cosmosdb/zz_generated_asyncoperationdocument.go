// Code generated by github.com/jim-minter/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"context"
	"net/http"
	"strings"

	pkg "github.com/Azure/ARO-RP/pkg/api"
)

type asyncOperationDocumentClient struct {
	*databaseClient
	path string
}

// AsyncOperationDocumentClient is a asyncOperationDocument client
type AsyncOperationDocumentClient interface {
	Create(context.Context, string, *pkg.AsyncOperationDocument, *Options) (*pkg.AsyncOperationDocument, error)
	List(*Options) AsyncOperationDocumentRawIterator
	ListAll(context.Context, *Options) (*pkg.AsyncOperationDocuments, error)
	Get(context.Context, string, string, *Options) (*pkg.AsyncOperationDocument, error)
	Replace(context.Context, string, *pkg.AsyncOperationDocument, *Options) (*pkg.AsyncOperationDocument, error)
	Delete(context.Context, string, *pkg.AsyncOperationDocument, *Options) error
	Query(string, *Query, *Options) AsyncOperationDocumentRawIterator
	QueryAll(context.Context, string, *Query, *Options) (*pkg.AsyncOperationDocuments, error)
	ChangeFeed(*Options) AsyncOperationDocumentIterator
}

type asyncOperationDocumentChangeFeedIterator struct {
	*asyncOperationDocumentClient
	continuation string
	options      *Options
}

type asyncOperationDocumentListIterator struct {
	*asyncOperationDocumentClient
	continuation string
	done         bool
	options      *Options
}

type asyncOperationDocumentQueryIterator struct {
	*asyncOperationDocumentClient
	partitionkey string
	query        *Query
	continuation string
	done         bool
	options      *Options
}

// AsyncOperationDocumentIterator is a asyncOperationDocument iterator
type AsyncOperationDocumentIterator interface {
	Next(context.Context) (*pkg.AsyncOperationDocuments, error)
}

// AsyncOperationDocumentRawIterator is a asyncOperationDocument raw iterator
type AsyncOperationDocumentRawIterator interface {
	AsyncOperationDocumentIterator
	NextRaw(context.Context, interface{}) error
}

// NewAsyncOperationDocumentClient returns a new asyncOperationDocument client
func NewAsyncOperationDocumentClient(collc CollectionClient, collid string) AsyncOperationDocumentClient {
	return &asyncOperationDocumentClient{
		databaseClient: collc.(*collectionClient).databaseClient,
		path:           collc.(*collectionClient).path + "/colls/" + collid,
	}
}

func (c *asyncOperationDocumentClient) all(ctx context.Context, i AsyncOperationDocumentIterator) (*pkg.AsyncOperationDocuments, error) {
	allasyncOperationDocuments := &pkg.AsyncOperationDocuments{}

	for {
		asyncOperationDocuments, err := i.Next(ctx)
		if err != nil {
			return nil, err
		}
		if asyncOperationDocuments == nil {
			break
		}

		allasyncOperationDocuments.Count += asyncOperationDocuments.Count
		allasyncOperationDocuments.ResourceID = asyncOperationDocuments.ResourceID
		allasyncOperationDocuments.AsyncOperationDocuments = append(allasyncOperationDocuments.AsyncOperationDocuments, asyncOperationDocuments.AsyncOperationDocuments...)
	}

	return allasyncOperationDocuments, nil
}

func (c *asyncOperationDocumentClient) Create(ctx context.Context, partitionkey string, newasyncOperationDocument *pkg.AsyncOperationDocument, options *Options) (asyncOperationDocument *pkg.AsyncOperationDocument, err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)

	if options == nil {
		options = &Options{}
	}
	options.NoETag = true

	err = c.setOptions(options, newasyncOperationDocument, headers)
	if err != nil {
		return
	}

	err = c.do(ctx, http.MethodPost, c.path+"/docs", "docs", c.path, http.StatusCreated, &newasyncOperationDocument, &asyncOperationDocument, headers)
	return
}

func (c *asyncOperationDocumentClient) List(options *Options) AsyncOperationDocumentRawIterator {
	return &asyncOperationDocumentListIterator{asyncOperationDocumentClient: c, options: options}
}

func (c *asyncOperationDocumentClient) ListAll(ctx context.Context, options *Options) (*pkg.AsyncOperationDocuments, error) {
	return c.all(ctx, c.List(options))
}

func (c *asyncOperationDocumentClient) Get(ctx context.Context, partitionkey, asyncOperationDocumentid string, options *Options) (asyncOperationDocument *pkg.AsyncOperationDocument, err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)

	err = c.setOptions(options, nil, headers)
	if err != nil {
		return
	}

	err = c.do(ctx, http.MethodGet, c.path+"/docs/"+asyncOperationDocumentid, "docs", c.path+"/docs/"+asyncOperationDocumentid, http.StatusOK, nil, &asyncOperationDocument, headers)
	return
}

func (c *asyncOperationDocumentClient) Replace(ctx context.Context, partitionkey string, newasyncOperationDocument *pkg.AsyncOperationDocument, options *Options) (asyncOperationDocument *pkg.AsyncOperationDocument, err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)

	err = c.setOptions(options, newasyncOperationDocument, headers)
	if err != nil {
		return
	}

	err = c.do(ctx, http.MethodPut, c.path+"/docs/"+newasyncOperationDocument.ID, "docs", c.path+"/docs/"+newasyncOperationDocument.ID, http.StatusOK, &newasyncOperationDocument, &asyncOperationDocument, headers)
	return
}

func (c *asyncOperationDocumentClient) Delete(ctx context.Context, partitionkey string, asyncOperationDocument *pkg.AsyncOperationDocument, options *Options) (err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)

	err = c.setOptions(options, asyncOperationDocument, headers)
	if err != nil {
		return
	}

	err = c.do(ctx, http.MethodDelete, c.path+"/docs/"+asyncOperationDocument.ID, "docs", c.path+"/docs/"+asyncOperationDocument.ID, http.StatusNoContent, nil, nil, headers)
	return
}

func (c *asyncOperationDocumentClient) Query(partitionkey string, query *Query, options *Options) AsyncOperationDocumentRawIterator {
	return &asyncOperationDocumentQueryIterator{asyncOperationDocumentClient: c, partitionkey: partitionkey, query: query, options: options}
}

func (c *asyncOperationDocumentClient) QueryAll(ctx context.Context, partitionkey string, query *Query, options *Options) (*pkg.AsyncOperationDocuments, error) {
	return c.all(ctx, c.Query(partitionkey, query, options))
}

func (c *asyncOperationDocumentClient) ChangeFeed(options *Options) AsyncOperationDocumentIterator {
	return &asyncOperationDocumentChangeFeedIterator{asyncOperationDocumentClient: c}
}

func (c *asyncOperationDocumentClient) setOptions(options *Options, asyncOperationDocument *pkg.AsyncOperationDocument, headers http.Header) error {
	if options == nil {
		return nil
	}

	if asyncOperationDocument != nil && !options.NoETag {
		if asyncOperationDocument.ETag == "" {
			return ErrETagRequired
		}
		headers.Set("If-Match", asyncOperationDocument.ETag)
	}
	if len(options.PreTriggers) > 0 {
		headers.Set("X-Ms-Documentdb-Pre-Trigger-Include", strings.Join(options.PreTriggers, ","))
	}
	if len(options.PostTriggers) > 0 {
		headers.Set("X-Ms-Documentdb-Post-Trigger-Include", strings.Join(options.PostTriggers, ","))
	}
	if len(options.PartitionKeyRangeID) > 0 {
		headers.Set("X-Ms-Documentdb-PartitionKeyRangeID", options.PartitionKeyRangeID)
	}

	return nil
}

func (i *asyncOperationDocumentChangeFeedIterator) Next(ctx context.Context) (asyncOperationDocuments *pkg.AsyncOperationDocuments, err error) {
	headers := http.Header{}
	headers.Set("A-IM", "Incremental feed")

	headers.Set("X-Ms-Max-Item-Count", "-1")
	if i.continuation != "" {
		headers.Set("If-None-Match", i.continuation)
	}

	err = i.setOptions(i.options, nil, headers)
	if err != nil {
		return
	}

	err = i.do(ctx, http.MethodGet, i.path+"/docs", "docs", i.path, http.StatusOK, nil, &asyncOperationDocuments, headers)
	if IsErrorStatusCode(err, http.StatusNotModified) {
		err = nil
	}
	if err != nil {
		return
	}

	i.continuation = headers.Get("Etag")

	return
}

func (i *asyncOperationDocumentListIterator) Next(ctx context.Context) (asyncOperationDocuments *pkg.AsyncOperationDocuments, err error) {
	err = i.NextRaw(ctx, &asyncOperationDocuments)
	return
}

func (i *asyncOperationDocumentListIterator) NextRaw(ctx context.Context, raw interface{}) (err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	headers.Set("X-Ms-Max-Item-Count", "-1")
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.setOptions(i.options, nil, headers)
	if err != nil {
		return
	}

	err = i.do(ctx, http.MethodGet, i.path+"/docs", "docs", i.path, http.StatusOK, nil, &raw, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}

func (i *asyncOperationDocumentQueryIterator) Next(ctx context.Context) (asyncOperationDocuments *pkg.AsyncOperationDocuments, err error) {
	err = i.NextRaw(ctx, &asyncOperationDocuments)
	return
}

func (i *asyncOperationDocumentQueryIterator) NextRaw(ctx context.Context, raw interface{}) (err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	headers.Set("X-Ms-Max-Item-Count", "-1")
	headers.Set("X-Ms-Documentdb-Isquery", "True")
	headers.Set("Content-Type", "application/query+json")
	if i.partitionkey != "" {
		headers.Set("X-Ms-Documentdb-Partitionkey", `["`+i.partitionkey+`"]`)
	} else {
		headers.Set("X-Ms-Documentdb-Query-Enablecrosspartition", "True")
	}
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.setOptions(i.options, nil, headers)
	if err != nil {
		return
	}

	err = i.do(ctx, http.MethodPost, i.path+"/docs", "docs", i.path, http.StatusOK, &i.query, &raw, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}

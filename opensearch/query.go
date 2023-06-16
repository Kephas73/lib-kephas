package opensearch

import (
    "github.com/Kephas73/lib-kephas/document"
    "io"
)

func (o *OpenSearchClient) CountDocument(index []string, bodyQuery io.Reader) (document.ResponseOpenSearch, error) {
    return o.clients.CountDocument(index, bodyQuery)
}

func (o *OpenSearchClient) SearchDocument(index []string, bodyQuery io.Reader) (document.ResponseOpenSearch, error) {
    return o.clients.SearchDocument(index, bodyQuery)
}

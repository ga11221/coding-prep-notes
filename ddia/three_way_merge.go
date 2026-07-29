package main

// Multi-leader 3-way merge (CouchDB-style)
// Each document has a revision DAG. Conflicts arise from concurrent edits.
// The DB stores all conflicting leaf revisions and provides the LCA.
// The app resolves by reading (LCA, our_version, their_version) and writing
// a new revision that supersedes both conflicting versions.
// No automatic merge — only the app has domain knowledge to decide.

type Revision struct {
	id      string
	parents []string
	doc     map[string]any
}

type RevTree struct {
	revs map[string]*Revision
}

// findLCA returns the most recent common ancestor revision
func (t *RevTree) findLCA(revA, revB string) *Revision {
	return nil
}

// getConflicts returns all leaf revisions that are concurrent (no LCA of both dominates)
func (t *RevTree) getConflicts() []*Revision {
	return nil
}

// resolve: app calls this after choosing a merged value
// writes a new revision that lists all conflicting revs as parents
func (t *RevTree) resolve(conflicts []*Revision, merged map[string]any) *Revision {
	return nil
}

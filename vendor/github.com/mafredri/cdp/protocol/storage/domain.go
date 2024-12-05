// Code generated by cdpgen. DO NOT EDIT.

// Package storage implements the Storage domain.
package storage

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Storage domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Storage domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// GetStorageKeyForFrame invokes the Storage method. Returns a storage key
// given a frame id.
func (d *domainClient) GetStorageKeyForFrame(ctx context.Context, args *GetStorageKeyForFrameArgs) (reply *GetStorageKeyForFrameReply, err error) {
	reply = new(GetStorageKeyForFrameReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.getStorageKeyForFrame", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.getStorageKeyForFrame", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetStorageKeyForFrame", Err: err}
	}
	return
}

// ClearDataForOrigin invokes the Storage method. Clears storage for origin.
func (d *domainClient) ClearDataForOrigin(ctx context.Context, args *ClearDataForOriginArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.clearDataForOrigin", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.clearDataForOrigin", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "ClearDataForOrigin", Err: err}
	}
	return
}

// ClearDataForStorageKey invokes the Storage method. Clears storage for
// storage key.
func (d *domainClient) ClearDataForStorageKey(ctx context.Context, args *ClearDataForStorageKeyArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.clearDataForStorageKey", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.clearDataForStorageKey", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "ClearDataForStorageKey", Err: err}
	}
	return
}

// GetCookies invokes the Storage method. Returns all browser cookies.
func (d *domainClient) GetCookies(ctx context.Context, args *GetCookiesArgs) (reply *GetCookiesReply, err error) {
	reply = new(GetCookiesReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.getCookies", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.getCookies", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetCookies", Err: err}
	}
	return
}

// SetCookies invokes the Storage method. Sets given cookies.
func (d *domainClient) SetCookies(ctx context.Context, args *SetCookiesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.setCookies", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.setCookies", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SetCookies", Err: err}
	}
	return
}

// ClearCookies invokes the Storage method. Clears cookies.
func (d *domainClient) ClearCookies(ctx context.Context, args *ClearCookiesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.clearCookies", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.clearCookies", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "ClearCookies", Err: err}
	}
	return
}

// GetUsageAndQuota invokes the Storage method. Returns usage and quota in
// bytes.
func (d *domainClient) GetUsageAndQuota(ctx context.Context, args *GetUsageAndQuotaArgs) (reply *GetUsageAndQuotaReply, err error) {
	reply = new(GetUsageAndQuotaReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.getUsageAndQuota", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.getUsageAndQuota", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetUsageAndQuota", Err: err}
	}
	return
}

// OverrideQuotaForOrigin invokes the Storage method. Override quota for the
// specified origin
func (d *domainClient) OverrideQuotaForOrigin(ctx context.Context, args *OverrideQuotaForOriginArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.overrideQuotaForOrigin", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.overrideQuotaForOrigin", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "OverrideQuotaForOrigin", Err: err}
	}
	return
}

// TrackCacheStorageForOrigin invokes the Storage method. Registers origin to
// be notified when an update occurs to its cache storage list.
func (d *domainClient) TrackCacheStorageForOrigin(ctx context.Context, args *TrackCacheStorageForOriginArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.trackCacheStorageForOrigin", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.trackCacheStorageForOrigin", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "TrackCacheStorageForOrigin", Err: err}
	}
	return
}

// TrackCacheStorageForStorageKey invokes the Storage method. Registers
// storage key to be notified when an update occurs to its cache storage list.
func (d *domainClient) TrackCacheStorageForStorageKey(ctx context.Context, args *TrackCacheStorageForStorageKeyArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.trackCacheStorageForStorageKey", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.trackCacheStorageForStorageKey", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "TrackCacheStorageForStorageKey", Err: err}
	}
	return
}

// TrackIndexedDBForOrigin invokes the Storage method. Registers origin to be
// notified when an update occurs to its IndexedDB.
func (d *domainClient) TrackIndexedDBForOrigin(ctx context.Context, args *TrackIndexedDBForOriginArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.trackIndexedDBForOrigin", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.trackIndexedDBForOrigin", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "TrackIndexedDBForOrigin", Err: err}
	}
	return
}

// TrackIndexedDBForStorageKey invokes the Storage method. Registers storage
// key to be notified when an update occurs to its IndexedDB.
func (d *domainClient) TrackIndexedDBForStorageKey(ctx context.Context, args *TrackIndexedDBForStorageKeyArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.trackIndexedDBForStorageKey", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.trackIndexedDBForStorageKey", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "TrackIndexedDBForStorageKey", Err: err}
	}
	return
}

// UntrackCacheStorageForOrigin invokes the Storage method. Unregisters origin
// from receiving notifications for cache storage.
func (d *domainClient) UntrackCacheStorageForOrigin(ctx context.Context, args *UntrackCacheStorageForOriginArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.untrackCacheStorageForOrigin", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.untrackCacheStorageForOrigin", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "UntrackCacheStorageForOrigin", Err: err}
	}
	return
}

// UntrackCacheStorageForStorageKey invokes the Storage method. Unregisters
// storage key from receiving notifications for cache storage.
func (d *domainClient) UntrackCacheStorageForStorageKey(ctx context.Context, args *UntrackCacheStorageForStorageKeyArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.untrackCacheStorageForStorageKey", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.untrackCacheStorageForStorageKey", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "UntrackCacheStorageForStorageKey", Err: err}
	}
	return
}

// UntrackIndexedDBForOrigin invokes the Storage method. Unregisters origin
// from receiving notifications for IndexedDB.
func (d *domainClient) UntrackIndexedDBForOrigin(ctx context.Context, args *UntrackIndexedDBForOriginArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.untrackIndexedDBForOrigin", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.untrackIndexedDBForOrigin", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "UntrackIndexedDBForOrigin", Err: err}
	}
	return
}

// UntrackIndexedDBForStorageKey invokes the Storage method. Unregisters
// storage key from receiving notifications for IndexedDB.
func (d *domainClient) UntrackIndexedDBForStorageKey(ctx context.Context, args *UntrackIndexedDBForStorageKeyArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.untrackIndexedDBForStorageKey", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.untrackIndexedDBForStorageKey", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "UntrackIndexedDBForStorageKey", Err: err}
	}
	return
}

// GetTrustTokens invokes the Storage method. Returns the number of stored
// Trust Tokens per issuer for the current browsing context.
func (d *domainClient) GetTrustTokens(ctx context.Context) (reply *GetTrustTokensReply, err error) {
	reply = new(GetTrustTokensReply)
	err = rpcc.Invoke(ctx, "Storage.getTrustTokens", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetTrustTokens", Err: err}
	}
	return
}

// ClearTrustTokens invokes the Storage method. Removes all Trust Tokens
// issued by the provided issuerOrigin. Leaves other stored data, including the
// issuer's Redemption Records, intact.
func (d *domainClient) ClearTrustTokens(ctx context.Context, args *ClearTrustTokensArgs) (reply *ClearTrustTokensReply, err error) {
	reply = new(ClearTrustTokensReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.clearTrustTokens", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.clearTrustTokens", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "ClearTrustTokens", Err: err}
	}
	return
}

// GetInterestGroupDetails invokes the Storage method. Gets details for a
// named interest group.
func (d *domainClient) GetInterestGroupDetails(ctx context.Context, args *GetInterestGroupDetailsArgs) (reply *GetInterestGroupDetailsReply, err error) {
	reply = new(GetInterestGroupDetailsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.getInterestGroupDetails", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.getInterestGroupDetails", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetInterestGroupDetails", Err: err}
	}
	return
}

// SetInterestGroupTracking invokes the Storage method. Enables/Disables
// issuing of interestGroupAccessed events.
func (d *domainClient) SetInterestGroupTracking(ctx context.Context, args *SetInterestGroupTrackingArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.setInterestGroupTracking", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.setInterestGroupTracking", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SetInterestGroupTracking", Err: err}
	}
	return
}

// SetInterestGroupAuctionTracking invokes the Storage method.
// Enables/Disables issuing of interestGroupAuctionEventOccurred and
// interestGroupAuctionNetworkRequestCreated.
func (d *domainClient) SetInterestGroupAuctionTracking(ctx context.Context, args *SetInterestGroupAuctionTrackingArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.setInterestGroupAuctionTracking", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.setInterestGroupAuctionTracking", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SetInterestGroupAuctionTracking", Err: err}
	}
	return
}

// GetSharedStorageMetadata invokes the Storage method. Gets metadata for an
// origin's shared storage.
func (d *domainClient) GetSharedStorageMetadata(ctx context.Context, args *GetSharedStorageMetadataArgs) (reply *GetSharedStorageMetadataReply, err error) {
	reply = new(GetSharedStorageMetadataReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.getSharedStorageMetadata", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.getSharedStorageMetadata", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetSharedStorageMetadata", Err: err}
	}
	return
}

// GetSharedStorageEntries invokes the Storage method. Gets the entries in an
// given origin's shared storage.
func (d *domainClient) GetSharedStorageEntries(ctx context.Context, args *GetSharedStorageEntriesArgs) (reply *GetSharedStorageEntriesReply, err error) {
	reply = new(GetSharedStorageEntriesReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.getSharedStorageEntries", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.getSharedStorageEntries", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetSharedStorageEntries", Err: err}
	}
	return
}

// SetSharedStorageEntry invokes the Storage method. Sets entry with `key` and
// `value` for a given origin's shared storage.
func (d *domainClient) SetSharedStorageEntry(ctx context.Context, args *SetSharedStorageEntryArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.setSharedStorageEntry", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.setSharedStorageEntry", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SetSharedStorageEntry", Err: err}
	}
	return
}

// DeleteSharedStorageEntry invokes the Storage method. Deletes entry for
// `key` (if it exists) for a given origin's shared storage.
func (d *domainClient) DeleteSharedStorageEntry(ctx context.Context, args *DeleteSharedStorageEntryArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.deleteSharedStorageEntry", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.deleteSharedStorageEntry", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "DeleteSharedStorageEntry", Err: err}
	}
	return
}

// ClearSharedStorageEntries invokes the Storage method. Clears all entries
// for a given origin's shared storage.
func (d *domainClient) ClearSharedStorageEntries(ctx context.Context, args *ClearSharedStorageEntriesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.clearSharedStorageEntries", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.clearSharedStorageEntries", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "ClearSharedStorageEntries", Err: err}
	}
	return
}

// ResetSharedStorageBudget invokes the Storage method. Resets the budget for
// `ownerOrigin` by clearing all budget withdrawals.
func (d *domainClient) ResetSharedStorageBudget(ctx context.Context, args *ResetSharedStorageBudgetArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.resetSharedStorageBudget", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.resetSharedStorageBudget", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "ResetSharedStorageBudget", Err: err}
	}
	return
}

// SetSharedStorageTracking invokes the Storage method. Enables/disables
// issuing of sharedStorageAccessed events.
func (d *domainClient) SetSharedStorageTracking(ctx context.Context, args *SetSharedStorageTrackingArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.setSharedStorageTracking", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.setSharedStorageTracking", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SetSharedStorageTracking", Err: err}
	}
	return
}

// SetStorageBucketTracking invokes the Storage method. Set tracking for a
// storage key's buckets.
func (d *domainClient) SetStorageBucketTracking(ctx context.Context, args *SetStorageBucketTrackingArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.setStorageBucketTracking", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.setStorageBucketTracking", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SetStorageBucketTracking", Err: err}
	}
	return
}

// DeleteStorageBucket invokes the Storage method. Deletes the Storage Bucket
// with the given storage key and bucket name.
func (d *domainClient) DeleteStorageBucket(ctx context.Context, args *DeleteStorageBucketArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.deleteStorageBucket", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.deleteStorageBucket", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "DeleteStorageBucket", Err: err}
	}
	return
}

// RunBounceTrackingMitigations invokes the Storage method. Deletes state for
// sites identified as potential bounce trackers, immediately.
func (d *domainClient) RunBounceTrackingMitigations(ctx context.Context) (reply *RunBounceTrackingMitigationsReply, err error) {
	reply = new(RunBounceTrackingMitigationsReply)
	err = rpcc.Invoke(ctx, "Storage.runBounceTrackingMitigations", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "RunBounceTrackingMitigations", Err: err}
	}
	return
}

// SetAttributionReportingLocalTestingMode invokes the Storage method.
// https://wicg.github.io/attribution-reporting-api/
func (d *domainClient) SetAttributionReportingLocalTestingMode(ctx context.Context, args *SetAttributionReportingLocalTestingModeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.setAttributionReportingLocalTestingMode", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.setAttributionReportingLocalTestingMode", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SetAttributionReportingLocalTestingMode", Err: err}
	}
	return
}

// SetAttributionReportingTracking invokes the Storage method.
// Enables/disables issuing of Attribution Reporting events.
func (d *domainClient) SetAttributionReportingTracking(ctx context.Context, args *SetAttributionReportingTrackingArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.setAttributionReportingTracking", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.setAttributionReportingTracking", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SetAttributionReportingTracking", Err: err}
	}
	return
}

// SendPendingAttributionReports invokes the Storage method. Sends all pending
// Attribution Reports immediately, regardless of their scheduled report time.
func (d *domainClient) SendPendingAttributionReports(ctx context.Context) (reply *SendPendingAttributionReportsReply, err error) {
	reply = new(SendPendingAttributionReportsReply)
	err = rpcc.Invoke(ctx, "Storage.sendPendingAttributionReports", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "SendPendingAttributionReports", Err: err}
	}
	return
}

// GetRelatedWebsiteSets invokes the Storage method. Returns the effective
// Related Website Sets in use by this profile for the browser session. The
// effective Related Website Sets will not change during a browser session.
func (d *domainClient) GetRelatedWebsiteSets(ctx context.Context) (reply *GetRelatedWebsiteSetsReply, err error) {
	reply = new(GetRelatedWebsiteSetsReply)
	err = rpcc.Invoke(ctx, "Storage.getRelatedWebsiteSets", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetRelatedWebsiteSets", Err: err}
	}
	return
}

func (d *domainClient) CacheStorageContentUpdated(ctx context.Context) (CacheStorageContentUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.cacheStorageContentUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &cacheStorageContentUpdatedClient{Stream: s}, nil
}

type cacheStorageContentUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *cacheStorageContentUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *cacheStorageContentUpdatedClient) Recv() (*CacheStorageContentUpdatedReply, error) {
	event := new(CacheStorageContentUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "CacheStorageContentUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) CacheStorageListUpdated(ctx context.Context) (CacheStorageListUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.cacheStorageListUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &cacheStorageListUpdatedClient{Stream: s}, nil
}

type cacheStorageListUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *cacheStorageListUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *cacheStorageListUpdatedClient) Recv() (*CacheStorageListUpdatedReply, error) {
	event := new(CacheStorageListUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "CacheStorageListUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) IndexedDBContentUpdated(ctx context.Context) (IndexedDBContentUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.indexedDBContentUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &indexedDBContentUpdatedClient{Stream: s}, nil
}

type indexedDBContentUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *indexedDBContentUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *indexedDBContentUpdatedClient) Recv() (*IndexedDBContentUpdatedReply, error) {
	event := new(IndexedDBContentUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "IndexedDBContentUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) IndexedDBListUpdated(ctx context.Context) (IndexedDBListUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.indexedDBListUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &indexedDBListUpdatedClient{Stream: s}, nil
}

type indexedDBListUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *indexedDBListUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *indexedDBListUpdatedClient) Recv() (*IndexedDBListUpdatedReply, error) {
	event := new(IndexedDBListUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "IndexedDBListUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) InterestGroupAccessed(ctx context.Context) (InterestGroupAccessedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.interestGroupAccessed", d.conn)
	if err != nil {
		return nil, err
	}
	return &interestGroupAccessedClient{Stream: s}, nil
}

type interestGroupAccessedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *interestGroupAccessedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *interestGroupAccessedClient) Recv() (*InterestGroupAccessedReply, error) {
	event := new(InterestGroupAccessedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "InterestGroupAccessed Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) InterestGroupAuctionEventOccurred(ctx context.Context) (InterestGroupAuctionEventOccurredClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.interestGroupAuctionEventOccurred", d.conn)
	if err != nil {
		return nil, err
	}
	return &interestGroupAuctionEventOccurredClient{Stream: s}, nil
}

type interestGroupAuctionEventOccurredClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *interestGroupAuctionEventOccurredClient) GetStream() rpcc.Stream { return c.Stream }

func (c *interestGroupAuctionEventOccurredClient) Recv() (*InterestGroupAuctionEventOccurredReply, error) {
	event := new(InterestGroupAuctionEventOccurredReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "InterestGroupAuctionEventOccurred Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) InterestGroupAuctionNetworkRequestCreated(ctx context.Context) (InterestGroupAuctionNetworkRequestCreatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.interestGroupAuctionNetworkRequestCreated", d.conn)
	if err != nil {
		return nil, err
	}
	return &interestGroupAuctionNetworkRequestCreatedClient{Stream: s}, nil
}

type interestGroupAuctionNetworkRequestCreatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *interestGroupAuctionNetworkRequestCreatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *interestGroupAuctionNetworkRequestCreatedClient) Recv() (*InterestGroupAuctionNetworkRequestCreatedReply, error) {
	event := new(InterestGroupAuctionNetworkRequestCreatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "InterestGroupAuctionNetworkRequestCreated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) SharedStorageAccessed(ctx context.Context) (SharedStorageAccessedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.sharedStorageAccessed", d.conn)
	if err != nil {
		return nil, err
	}
	return &sharedStorageAccessedClient{Stream: s}, nil
}

type sharedStorageAccessedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *sharedStorageAccessedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *sharedStorageAccessedClient) Recv() (*SharedStorageAccessedReply, error) {
	event := new(SharedStorageAccessedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "SharedStorageAccessed Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) StorageBucketCreatedOrUpdated(ctx context.Context) (BucketCreatedOrUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.storageBucketCreatedOrUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &bucketCreatedOrUpdatedClient{Stream: s}, nil
}

type bucketCreatedOrUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *bucketCreatedOrUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *bucketCreatedOrUpdatedClient) Recv() (*BucketCreatedOrUpdatedReply, error) {
	event := new(BucketCreatedOrUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "StorageBucketCreatedOrUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) StorageBucketDeleted(ctx context.Context) (BucketDeletedClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.storageBucketDeleted", d.conn)
	if err != nil {
		return nil, err
	}
	return &bucketDeletedClient{Stream: s}, nil
}

type bucketDeletedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *bucketDeletedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *bucketDeletedClient) Recv() (*BucketDeletedReply, error) {
	event := new(BucketDeletedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "StorageBucketDeleted Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) AttributionReportingSourceRegistered(ctx context.Context) (AttributionReportingSourceRegisteredClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.attributionReportingSourceRegistered", d.conn)
	if err != nil {
		return nil, err
	}
	return &attributionReportingSourceRegisteredClient{Stream: s}, nil
}

type attributionReportingSourceRegisteredClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *attributionReportingSourceRegisteredClient) GetStream() rpcc.Stream { return c.Stream }

func (c *attributionReportingSourceRegisteredClient) Recv() (*AttributionReportingSourceRegisteredReply, error) {
	event := new(AttributionReportingSourceRegisteredReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "AttributionReportingSourceRegistered Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) AttributionReportingTriggerRegistered(ctx context.Context) (AttributionReportingTriggerRegisteredClient, error) {
	s, err := rpcc.NewStream(ctx, "Storage.attributionReportingTriggerRegistered", d.conn)
	if err != nil {
		return nil, err
	}
	return &attributionReportingTriggerRegisteredClient{Stream: s}, nil
}

type attributionReportingTriggerRegisteredClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *attributionReportingTriggerRegisteredClient) GetStream() rpcc.Stream { return c.Stream }

func (c *attributionReportingTriggerRegisteredClient) Recv() (*AttributionReportingTriggerRegisteredReply, error) {
	event := new(AttributionReportingTriggerRegisteredReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Storage", Op: "AttributionReportingTriggerRegistered Recv", Err: err}
	}
	return event, nil
}

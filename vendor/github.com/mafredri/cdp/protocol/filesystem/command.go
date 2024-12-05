// Code generated by cdpgen. DO NOT EDIT.

package filesystem

// GetDirectoryArgs represents the arguments for GetDirectory in the FileSystem domain.
type GetDirectoryArgs struct {
	BucketFileSystemLocator BucketFileSystemLocator `json:"bucketFileSystemLocator"` // No description.
}

// NewGetDirectoryArgs initializes GetDirectoryArgs with the required arguments.
func NewGetDirectoryArgs(bucketFileSystemLocator BucketFileSystemLocator) *GetDirectoryArgs {
	args := new(GetDirectoryArgs)
	args.BucketFileSystemLocator = bucketFileSystemLocator
	return args
}

// GetDirectoryReply represents the return values for GetDirectory in the FileSystem domain.
type GetDirectoryReply struct {
	Directory Directory `json:"directory"` // Returns the directory object at the path.
}

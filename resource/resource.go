package resource

// Resource resource is a interface for implementing an resource type.
// Defined type of resource as below:
//  * Article
//  * Image
//  * Video
//  * Navigation
//  * Text
type Resource interface {
	// Name is a identity of resource
	Name() string
	// Type is returned
	Type() string
	// List is returned to caller
	List(pageSize int) []string
	// Count how many record of resource
	Count() int
	// Search is returned results matched keyword
	Search(keyword string, pageSize int) []string
	// Content is returned that find by ID
	Content(id int) []byte
	// Modify replace old content by new content
	Modify(id int, name string, content []byte) error
	// Delete remove content matched id
	Delete(id int) error
	// Add new content is added
	Add(name string, content []byte) error
}

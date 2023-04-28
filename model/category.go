package models

type Category struct {
	ID               int          `json:"id"`
	Title            string       `json:"name"`
	ParentID         int          `json:"parent_id"`
	CategoryChildren []ChildrenID `json:"category_children,omitempty"`
}

type ChildrenID struct {
	ID         int          `json:"id"`
	Title      string       `json:"name"`
	ChildrenID []ChildrenID `json:"children_id"`
}

package models

type Category struct {
	ID               int        `json:"id,omitempty"`
	Title            string     `json:"name,omitempty"`
	ParentID         int        `json:"parent_id,omitempty"`
	CategoryChildren []Category `json:"category_children,omitempty"`
}

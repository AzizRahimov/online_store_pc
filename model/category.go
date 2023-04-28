package models

type Category struct {
	ID               int        `json:"id"`
	Title            string     `json:"name"`
	ParentID         int        `json:"parent_id"`
	CategoryChildren []Category `json:"category_children,omitempty"`
}

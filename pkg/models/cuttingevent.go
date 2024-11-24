package models

type CuttingResultEvent struct {
	CuttingResultID uint               `json:"cutting_result_id"`
	ItemID          uint               `json:"item_id"`
	Components      []ComponentPayload `json:"components"`
}

type ComponentPayload struct {
	MaterialID    uint   `json:"material_id"`
	DoorPanel     string `json:"door_panel,omitempty"`
	BackSidePanel string `json:"back_side_panel,omitempty"`
	SidePanel     string `json:"side_panel,omitempty"`
	TopPanel      string `json:"top_panel,omitempty"`
	BottomPanel   string `json:"bottom_panel,omitempty"`
	ShelvesPanel  string `json:"shelves_panel,omitempty"`
	PanelCount    int32  `json:"panel_count"`
}

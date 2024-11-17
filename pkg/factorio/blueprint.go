package factorio

import (
	"encoding/json"
	"errors"
)

type SerializedBlueprint struct {
	Index     int    `json:"index"`
	Blueprint []byte `json:"blueprint"`
}

type BlueprintBook struct {
	Item        string                `json:"item"`
	Label       string                `json:"label"`
	LabelColor  *Color                `json:"label_color,omitempty"`
	Blueprints  []SerializedBlueprint `json:"blueprints"`
	ActiveIndex int                   `json:"active_index"`
	Icons       []Icon                `json:"icons"`
	Description *string               `json:"description,omitempty"`
	Version     int64                 `json:"version"`
}

type Blueprint struct {
	Item             string     `json:"item"`
	Label            string     `json:"label"`
	LabelColor       *Color     `json:"label_color,omitempty"`
	Entities         []Entity   `json:"entities"`
	Tiles            []Tile     `json:"tiles"`
	Icons            []Icon     `json:"icons"`
	Schedules        []Schedule `json:"schedules"`
	Description      *string    `json:"description,omitempty"`
	SnapToGrid       *Position  `json:"snap-to-grid,omitempty"`
	AbsoluteSnapping *bool      `json:"absolute_snapping,omitempty"`
	PositionRel      *Position  `json:"position-relative-to-grid,omitempty"`
	Version          int64      `json:"version"`
}

type Entity struct {
	EntityNumber      int                    `json:"entity_number"`
	Name              string                 `json:"name"`
	Position          Position               `json:"position"`
	Direction         *int                   `json:"direction"`
	Orientation       *float32               `json:"orientation,omitempty"`
	Connections       map[int]Connection     `json:"connections,omitempty"`
	Neighbors         []int                  `json:"neighbors,omitempty"`
	ControlBehavior   *ControlBehavior       `json:"control_behavior,omitempty"`
	Items             map[string]int         `json:"items,omitempty"`
	Recipe            *string                `json:"recipe,omitempty"`
	Bar               *int                   `json:"bar,omitempty"`
	AmmoInventory     *Inventory             `json:"ammo_inventory,omitempty"`
	TrunkInventory    *Inventory             `json:"trunk_inventory,omitempty"`
	InfinitySettings  *InfinitySettings      `json:"infinity_settings,omitempty"`
	Type              *string                `json:"type,omitempty"`
	InputPriority     *string                `json:"input_priority,omitempty"`
	OutputPriority    *string                `json:"output_priority,omitempty"`
	Filter            *string                `json:"filter,omitempty"`
	Filters           map[string]string      `json:"filters,omitempty"`
	FilterMode        *string                `json:"filter_mode,omitempty"`
	OverrideStackSize *int                   `json:"override_stack_size,omitempty"`
	DropPosition      *Position              `json:"drop_position,omitempty"`
	PickupPosition    *Position              `json:"pickup_position,omitempty"`
	RequestFilters    []LogisticFilter       `json:"request_filters,omitempty"`
	Parameters        *SpeakerParameter      `json:"parameters,omitempty"`
	AlertParameters   *SpeakerAlertParameter `json:"alert_parameters,omitempty"`
	AutoLaunch        *bool                  `json:"auto_launch,omitempty"`
	Variation         *string                `json:"variation,omitempty"`
	Color             *Color                 `json:"color,omitempty"`
	Station           *string                `json:"station,omitempty"`
	ManualTrainsLimit *int                   `json:"manual_trains_limit,omitempty"`
	SwitchState       *bool                  `json:"switch_state,omitempty"`
	Tags              map[string]string      `json:"tags,omitempty"`
}

type Inventory struct {
	Filters []ItemFilter `json:"filters"`
	Bar     *int         `json:"bar,omitempty"`
}

type Schedule struct {
	Schedule    []ScheduleRecord `json:"schedule"`
	Locomotives []int            `json:"locomotives"`
}

type ScheduleRecord struct {
	Station        string          `json:"station"`
	WaitConditions []WaitCondition `json:"wait_conditions"`
	Temporary      *bool           `json:"temporary,omitempty"`
}

type WaitCondition struct {
	Type        string            `json:"type"`
	CompareType string            `json:"compare_type"`
	Ticks       *uint             `json:"ticks,omitempty"`
	Condition   *CircuitCondition `json:"condition,omitempty"`
}

type CircuitCondition struct {
	Comparator   *string   `json:"comparator,omitempty"`
	FirstSignal  *SignalID `json:"first_signal,omitempty"`
	SecondSignal *SignalID `json:"second_signal,omitempty"`
	Constant     *int      `json:"constant,omitempty"`
}

type Tile struct {
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Connection struct {
	One ConnectionPoint `json:"1"`
	Two ConnectionPoint `json:"2"`
}

type ConnectionPoint struct {
	Red   []ConnectionData `json:"red"`
	Green []ConnectionData `json:"green"`
}

type ConnectionData struct {
	EntityID  int `json:"entity_id"`
	CircuitID int `json:"circuit_id"`
}

type ItemFilter struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Mode  string `json:"mode"`
}

type InfinitySettings struct {
	RemoveUnfilteredItems bool             `json:"remove_unfiltered_items"`
	Filters               []InfinityFilter `json:"filters"`
}

type InfinityFilter struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Mode  string `json:"mode"`
	Index int    `json:"index"`
}

type LogisticFilter struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
	Count int    `json:"count"`
}

type SpeakerParameter struct {
	PlaybackVolume   float64 `json:"playback_volume"`
	PlaybackGlobally bool    `json:"playback_globally"`
	AllowPolyphony   bool    `json:"allow_polyphony"`
}

type SpeakerAlertParameter struct {
	ShowAlert    bool     `json:"show_alert"`
	ShowOnMap    bool     `json:"show_on_map"`
	IconSignalID SignalID `json:"icon_signal_id"`
	AlertMessage string   `json:"alert_message"`
}

type Icon struct {
	Index  int      `json:"index"`
	Signal SignalID `json:"signal"`
}

type SignalID struct {
	Name string  `json:"name"`
	Type *string `json:"type,omitempty"`
}

type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

type ControlBehavior struct {
	LogisticCondition                 *CircuitCondition  `json:"logistic_condition,omitempty"`
	ConnectToLogisticNetwork          *bool              `json:"connect_to_logistic_network,omitempty"`
	CircuitCloseSignal                *bool              `json:"circuit_close_signal,omitempty"`
	CircuitReadSignal                 *bool              `json:"circuit_read_signal,omitempty"`
	RedOutputSignal                   *SignalID          `json:"red_output_signal,omitempty"`
	OrangeOutputSignal                *SignalID          `json:"orange_output_signal,omitempty"`
	GreenOutputSignal                 *SignalID          `json:"green_output_signal,omitempty"`
	BlueOutputSignal                  *SignalID          `json:"blue_output_signal,omitempty"`
	CircuitCondition                  *CircuitCondition  `json:"circuit_condition,omitempty"`
	CircuitEnableDisable              *bool              `json:"circuit_enable_disable,omitempty"`
	SendToTrain                       *bool              `json:"send_to_train,omitempty"`
	ReadFromTrain                     *bool              `json:"read_from_train,omitempty"`
	ReadStoppedTrain                  *bool              `json:"read_stopped_train,omitempty"`
	TrainStoppedSignal                *SignalID          `json:"train_stopped_signal,omitempty"`
	SetTrainsLimit                    *bool              `json:"set_trains_limit,omitempty"`
	TrainsLimitSignal                 *SignalID          `json:"trains_limit_signal,omitempty"`
	ReadTrainsCount                   *bool              `json:"read_trains_count,omitempty"`
	TrainsCountSignal                 *SignalID          `json:"trains_count_signal,omitempty"`
	ReadLogistics                     *bool              `json:"read_logistics,omitempty"`
	ReadRobotStats                    *bool              `json:"read_robot_stats,omitempty"`
	AvailableLogisticOutputSignal     *SignalID          `json:"available_logistic_output_signal,omitempty"`
	TotalLogisticOutputSignal         *SignalID          `json:"total_logistic_output_signal,omitempty"`
	AvailableConstructionOutputSignal *SignalID          `json:"available_construction_output_signal,omitempty"`
	TotalConstructionOutputSignal     *SignalID          `json:"total_construction_output_signal,omitempty"`
	CircuitOpenGate                   *bool              `json:"circuit_open_gate,omitempty"`
	CircuitReadSensor                 *bool              `json:"circuit_read_sensor,omitempty"`
	OutputSignal                      *SignalID          `json:"output_signal,omitempty"`
	CircuitReadHandContents           *bool              `json:"circuit_read_hand_contents,omitempty"`
	CircuitContentsReadMode           *int               `json:"circuit_contents_read_mode,omitempty"`
	CircuitModeOfOperation            *int               `json:"circuit_mode_of_operation,omitempty"`
	CircuitHandReadMode               *int               `json:"circuit_hand_read_mode,omitempty"`
	CircuitSetStackSize               *bool              `json:"circuit_set_stack_size,omitempty"`
	StackControlInputSignal           *SignalID          `json:"stack_control_input_signal,omitempty"`
	CircuitReadResources              *bool              `json:"circuit_read_resources,omitempty"`
	CircuitResourceReadMode           *int               `json:"circuit_resource_read_mode,omitempty"`
	IsOn                              *bool              `json:"is_on,omitempty"`
	Filters                           []LogisticSections `json:"filters,omitempty"`
	ArithmeticConditions              *string            `json:"arithmetic_conditions,omitempty"`
	DeciderConditions                 *string            `json:"decider_conditions,omitempty"`
	CircuitParameters                 *string            `json:"circuit_parameters,omitempty"`
	UseColors                         *bool              `json:"use_colors,omitempty"`
}

type LogisticSections struct {
	Sections          []LogisticSection `json:"sections,omitempty"`
	TrashNotRequested *bool             `json:"trash_not_requested,omitempty"`
}

type LogisticSection struct {
	Index      int              `json:"index"`
	Filters    []LogisticFilter `json:"filters,omitempty"`
	Group      *string          `json:"group,omitempty"`
	Multiplier *float64         `json:"multiplier,omitempty"`
	Active     *bool            `json:"active,omitempty"`
}

type BlueprintWrapper struct {
	Blueprint *Blueprint `json:"blueprint,omitempty"`
}

type BlueprintBookWrapper struct {
	BlueprintBook *BlueprintBook `json:"blueprint-book,omitempty"`
}

func NewBlueprintFromJSON(data []byte) (*Blueprint, error) {
	var wrapper BlueprintWrapper
	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	if wrapper.Blueprint == nil {
		return nil, errors.New("JSON does not contain a blueprint")
	}

	return wrapper.Blueprint, nil
}

func NewBlueprintBookFromJSON(data []byte) (*BlueprintBook, error) {
	var wrapper BlueprintBookWrapper
	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	if wrapper.BlueprintBook == nil {
		return nil, errors.New("JSON does not contain a blueprint book")
	}

	return wrapper.BlueprintBook, nil
}

func (bp *Blueprint) ToJSON() ([]byte, error) {
	return json.Marshal(BlueprintWrapper{Blueprint: bp})
}

func (bb *BlueprintBook) ToJSON() ([]byte, error) {
	return json.Marshal(BlueprintBookWrapper{BlueprintBook: bb})
}

func (bb *BlueprintBook) AddBlueprint(bp *Blueprint) error {
	buf, err := bp.ToJSON()
	if err != nil {
		return err
	}

	bb.Blueprints = append(bb.Blueprints, SerializedBlueprint{
		Index:     len(bb.Blueprints),
		Blueprint: buf,
	})

	return nil
}

func (bb *BlueprintBook) AddBook(book *BlueprintBook) error {
	buf, err := book.ToJSON()
	if err != nil {
		return err
	}

	bb.Blueprints = append(bb.Blueprints, SerializedBlueprint{
		Index:     len(bb.Blueprints),
		Blueprint: buf,
	})

	return nil
}

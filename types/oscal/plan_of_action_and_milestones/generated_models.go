// Code generated by https://github.com/GoComply/metaschema; DO NOT EDIT.
package plan_of_action_and_milestones

import (
	"github.com/gocomply/oscalkit/types/oscal/validation_root"

	"github.com/gocomply/oscalkit/types/oscal/system_security_plan"

	"github.com/gocomply/oscalkit/types/oscal/assessment_common"
)

// A plan of action and milestones, such as those required by FedRAMP.
type PlanOfActionAndMilestones struct {

	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// A unique identifier for the system described by this system security plan.
	SystemId *SystemId `xml:"system-id,omitempty" json:"systemId,omitempty"`
	// Provides information about the publication and availability of the containing document.
	Metadata *Metadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// Used by the POA&M to import information about the system.
	ImportSsp *ImportSsp `xml:"import-ssp,omitempty" json:"importSsp,omitempty"`
	// Allows components, and inventory-items to be defined within the POA&M for circumstances where no OSCAL-based SSP exists, or is not delivered with the POA&M.
	LocalDefinitions *LocalDefinitions `xml:"local-definitions,omitempty" json:"localDefinitions,omitempty"`
	// This identifies initial and residual risks, deviations, and disposition.
	PoamItems *PoamItems `xml:"poam-items,omitempty" json:"poamItems,omitempty"`
	// A collection of citations and resource references.
	BackMatter *BackMatter `xml:"back-matter,omitempty" json:"backMatter,omitempty"`
}

// Allows components, and inventory-items to be defined within the POA&M for circumstances where no OSCAL-based SSP exists, or is not delivered with the POA&M.
type LocalDefinitions struct {

	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// Used to add any components, not defined via the System Security Plan (AR->AP->SSP)
	Components ComponentMultiplexer `xml:"component,omitempty" json:"components,omitempty"`
	// Used to add any inventory-items, not defined via the System Security Plan (AR->AP->SSP)
	InventoryItems InventoryItemMultiplexer `xml:"inventory-item,omitempty" json:"inventory-items,omitempty"`
}

// This identifies initial and residual risks, deviations, and disposition.
type PoamItems struct {

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A description supporting the parent item.
	Description *Description `xml:"description,omitempty" json:"description,omitempty"`
	// Provided as means of extending the OSCAL syntax.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// Date/time stamp identifying the start of the evidence collection reflected in these results.
	Start Start `xml:"start,omitempty" json:"start,omitempty"`
	// Date/time stamp identifying the end of the evidence collection reflected in these results. In a continuous monitoring scenario, this may be omitted or contain the same value as start if appropriate.
	End End `xml:"end,omitempty" json:"end,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// Provided as means of extending the OSCAL syntax.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// Describes an individual POA&M item.
	PoamItemGroup []PoamItem `xml:"poam-item,omitempty" json:"poam-item-group,omitempty"`
}

// Describes an individual POA&M item.
type PoamItem struct {

	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A description supporting the parent item.
	Description *Description `xml:"description,omitempty" json:"description,omitempty"`
	// Provided as means of extending the OSCAL syntax.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// Date/time stamp identifying when the finding information was collected.
	Collected Collected `xml:"collected,omitempty" json:"collected,omitempty"`
	// Date/time identifying when the finding information is out-of-date and no longer valid. Typically used with continuous assessment scenarios.
	Expires Expires `xml:"expires,omitempty" json:"expires,omitempty"`
	// Identifies the implementation statement in the SSP to which this finding is related.
	ImplementationStatementUuid ImplementationStatementUuid `xml:"implementation-statement-uuid,omitempty" json:"implementationStatementUuid,omitempty"`
	// A pointer, by ID, to an externally-defined threat.
	ThreatIds []ThreatId `xml:"threat-id,omitempty" json:"threat-ids,omitempty"`
	// The person who collected the evidence or made the observation.
	PartyUuids []PartyUuid `xml:"party-uuid,omitempty" json:"party-uuids,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// Provided as means of extending the OSCAL syntax.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// Captures an assessors conclusions as to whether an objective is fully satisfied.
	ObjectiveStatus *ObjectiveStatus `xml:"objective-status,omitempty" json:"objectiveStatus,omitempty"`
	// Describes an individual observation.
	Observations []Observation `xml:"observation,omitempty" json:"observations,omitempty"`
	// An identified risk.
	Risks []Risk `xml:"risk,omitempty" json:"risks,omitempty"`
}

type ComponentMultiplexer = system_security_plan.ComponentMultiplexer

type InventoryItemMultiplexer = system_security_plan.InventoryItemMultiplexer

type Annotation = validation_root.Annotation

type BackMatter = validation_root.BackMatter

type Collected = assessment_common.Collected

type Component = system_security_plan.Component

type Description = validation_root.Description

type End = assessment_common.End

type Expires = assessment_common.Expires

type ImplementationStatementUuid = assessment_common.ImplementationStatementUuid

type ImportSsp = assessment_common.ImportSsp

type InventoryItem = system_security_plan.InventoryItem

type Metadata = validation_root.Metadata

type ObjectiveStatus = assessment_common.ObjectiveStatus

type Observation = assessment_common.Observation

type PartyUuid = validation_root.PartyUuid

type Prop = validation_root.Prop

type Remarks = validation_root.Remarks

type Risk = assessment_common.Risk

type Start = assessment_common.Start

type SystemId = system_security_plan.SystemId

type ThreatId = assessment_common.ThreatId

type Title = validation_root.Title

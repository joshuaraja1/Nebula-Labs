package schema

// ImmigrationStatus represents a scholarship applicant's immigration category.
// It is intentionally string-backed because scholarship providers may use
// different labels while the API preserves a consistent schema.
type ImmigrationStatus string

// ScholarshipRequirements describes the eligibility criteria for a scholarship.
type ScholarshipRequirements struct {
	MatriculationStatus     string            `bson:"matriculation_status" json:"matriculation_status"`
	ImmigrationStatus      ImmigrationStatus `bson:"immigration_status" json:"immigration_status"`
	TexasResidencyStatus   string            `bson:"texas_residency_status" json:"texas_residency_status"`
	CourseLoad             string            `bson:"course_load" json:"course_load"`
	EligibleClassification string            `bson:"eligible_classification" json:"eligible_classification"`
	MinGPA                 float64           `bson:"min_gpa" json:"min_gpa"`
	Schools                []string          `bson:"schools" json:"schools"`
	Majors                 []string          `bson:"majors" json:"majors"`
	Concentrations         []string          `bson:"concentrations" json:"concentrations"`
}

// Scholarship contains the public details, application information, and
// eligibility requirements for a scholarship opportunity.
type Scholarship struct {
	ID                      int                     `bson:"_id" json:"id"`
	Name                    string                  `bson:"name" json:"name"`
	Description             string                  `bson:"description" json:"description"`
	Deadline                string                  `bson:"deadline" json:"deadline"`
	Term                    string                  `bson:"term" json:"term"`
	Contact                 string                  `bson:"contact" json:"contact"`
	ApplicationInstructions string                  `bson:"application_instructions" json:"application_instructions"`
	ApplicationLink         string                  `bson:"application_link" json:"application_link"`
	Requirements            ScholarshipRequirements `bson:"requirements" json:"requirements"`
	Renewal                 string                  `bson:"renewal" json:"renewal"`
}

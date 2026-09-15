package schema

import (
	"encoding/json"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestScholarshipJSONSchema(t *testing.T) {
	scholarship := Scholarship{
		ID:                      42,
		Name:                    "Academic Excellence Scholarship",
		Description:             "Supports students with strong academic performance.",
		Deadline:                "2027-02-01",
		Term:                    "Fall 2027",
		Contact:                 "scholarships@example.edu",
		ApplicationInstructions: "Submit the online application.",
		ApplicationLink:         "https://example.edu/apply",
		Requirements: ScholarshipRequirements{
			MatriculationStatus:     "degree-seeking",
			ImmigrationStatus:      ImmigrationStatus("eligible noncitizen"),
			TexasResidencyStatus:   "resident",
			CourseLoad:             "full-time",
			EligibleClassification: "undergraduate",
			MinGPA:                 3.5,
			Schools:                []string{"School of Engineering"},
			Majors:                 []string{"Computer Science"},
			Concentrations:         []string{"Data Science"},
		},
		Renewal: "Renewable for up to four years.",
	}

	data, err := json.Marshal(scholarship)
	if err != nil {
		t.Fatalf("marshal Scholarship: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal Scholarship JSON: %v", err)
	}

	wantTopLevelKeys := []string{
		"id", "name", "description", "deadline", "term", "contact",
		"application_instructions", "application_link", "requirements", "renewal",
	}
	for _, key := range wantTopLevelKeys {
		if _, ok := got[key]; !ok {
			t.Errorf("missing top-level JSON key %q", key)
		}
	}

	requirements, ok := got["requirements"].(map[string]any)
	if !ok {
		t.Fatalf("requirements has type %T; want object", got["requirements"])
	}
	wantRequirementKeys := []string{
		"matriculation_status", "immigration_status", "texas_residency_status",
		"course_load", "eligible_classification", "min_gpa", "schools", "majors",
		"concentrations",
	}
	for _, key := range wantRequirementKeys {
		if _, ok := requirements[key]; !ok {
			t.Errorf("missing requirements JSON key %q", key)
		}
	}
	if gotMinGPA := requirements["min_gpa"]; gotMinGPA != 3.5 {
		t.Errorf("min_gpa = %v; want numeric value 3.5", gotMinGPA)
	}
}

func TestScholarshipBSONUsesMongoID(t *testing.T) {
	data, err := bson.Marshal(Scholarship{ID: 42})
	if err != nil {
		t.Fatalf("marshal Scholarship BSON: %v", err)
	}

	var got map[string]any
	if err := bson.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal Scholarship BSON: %v", err)
	}

	if _, ok := got["_id"]; !ok {
		t.Error("missing BSON _id field")
	}
	if _, ok := got["id"]; ok {
		t.Error("unexpected BSON id field; want _id")
	}
}

func TestScholarshipSlicesRoundTrip(t *testing.T) {
	want := ScholarshipRequirements{
		Schools:        []string{"JSOM", "ECS"},
		Majors:         []string{"Information Technology and Systems"},
		Concentrations: []string{"Data Science"},
	}

	data, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal ScholarshipRequirements: %v", err)
	}

	var got ScholarshipRequirements
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal ScholarshipRequirements: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip = %#v; want %#v", got, want)
	}
}

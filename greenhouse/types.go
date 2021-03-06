/*
Copyright 2021-2022
Carnegie Robotics, LLC
4501 Hatfield Street, Pittsburgh, PA 15201
https://www.carnegierobotics.com
All rights reserved.

This file is part of greenhouse-client-go.

greenhouse-client-go is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

greenhouse-client-go is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with greenhouse-client-go. If not, see <https://www.gnu.org/licenses/>.
*/

package greenhouse

import ()

// Activity represents an object on an activity feed.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-activity-feed-object
type Activity struct {
	Body      *string `json:"body,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id        *int    `json:"id,omitempty"`
	Subject   *string `json:"subject,omitempty"`
	User      *User   `json:"user,omitempty"`
}

// ActivityFeed represents a candidate's activity feed.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-activity-feed-object
type ActivityFeed struct {
	Activities []Activity `json:"activities,omitempty"`
	Emails     []Email    `json:"emails,omitempty"`
	Id         *int       `json:"id,omitempty"`
	Notes      []Note     `json:"notes,omitempty"`
}

// Answer represents a question-answer pair on an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-demographic-answer-object
type Answer struct {
	Answer   *string `json:"answer,omitempty"`
	Question *string `json:"question,omitempty"`
}

// Application represents an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-application-object
type Application struct {
	Answers                 []Answer                    `json:"answers,omitempty"`
	AppliedAt               *string                     `json:"applied_at,omitempty"`
	Attachments             []Attachment                `json:"attachments,omitempty"`
	CandidateId             *int                        `json:"candidate_id,omitempty"`
	CreditedTo              *User                       `json:"credited_to,omitempty"`
	CurrentStage            *Stage                      `json:"current_stage,omitempty"`
	CustomFields            map[string]string           `json:"custom_fields,omitempty"`
	Id                      *int                        `json:"id,omitempty"`
	InitialStageId          *int                        `json:"initial_stage_id,omitempty"`
	JobId                   *int                        `json:"job_id,omitempty"`
	JobIds                  []int                       `json:"job_ids,omitempty"`
	Jobs                    []Job                       `json:"jobs,omitempty"`
	JobPostId               *int                        `json:"job_post_id,omitempty"`
	KeyedCustomFields       map[string]KeyedCustomField `json:"keyed_custom_fields,omitempty"`
	LastActivityAt          *string                     `json:"last_activity_at,omitempty"`
	Location                *Location                   `json:"location,omitempty"`
	Prospect                *bool                       `json:"prospect,omitempty"`
	ProspectDetail          *ProspectDetail             `json:"prospect_detail,omitempty"`
	ProspectOwnerId         *int                        `json:"prospect_owner_id,omitempty"`
	ProspectPoolId          *int                        `json:"prospect_pool_id,omitempty"`
	ProspectPoolStageId     *int                        `json:"prospect_pool_stage_id,omitempty"`
	ProspectStageId         *int                        `json:"prospect_stage_id,omitempty"`
	ProspectiveDepartment   *Department                 `json:"prospective_department,omitempty"`
	ProspectiveDepartmentId *int                        `json:"prospective_department_id,omitempty"`
	ProspectiveOffice       *Office                     `json:"prospective_office,omitempty"`
	ProspectiveOfficeId     *int                        `json:"prospective_office_id,omitempty"`
	Referrer                *TypeTypeValue              `json:"referrer,omitempty"`
	RejectedAt              *string                     `json:"rejected_at,omitempty"`
	RejectionDetails        *RejectionDetails           `json:"rejection_details,omitempty"`
	RejectionReason         *RejectionReason            `json:"rejection_reason,omitempty"`
	Source                  *Source                     `json:"source,omitempty"`
	SourceId                *int                        `json:"source_id,omitempty"`
	Status                  *string                     `json:"status,omitempty"`
}

// ApplicationHire represents the data needed to hire an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-hire-application
type ApplicationHire struct {
	CloseReasonId *int    `json:"close_reason_id,omitempty"`
	OpeningId     *int    `json:"opening_id,omitempty"`
	StartDate     *string `json:"start_date,omitempty"`
}

// ApplicationReject represents the data needed to reject an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-reject-application
type ApplicationReject struct {
	Notes             *string         `json:"notes,omitempty"`
	RejectionEmail    *RejectionEmail `json:"rejection_email,omitempty"`
	RejectionReasonId *int            `json:"rejection_reason_id,omitempty"`
}

// Approval represents an approval flow.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-approval-flow-object
type Approval struct {
	ApprovalStatus    *string         `json:"approval_status,omitempty"`
	ApprovalType      *string         `json:"approval_type,omitempty"`
	ApproverGroups    []ApproverGroup `json:"approver_groups,omitempty"`
	Id                *int            `json:"id,omitempty"`
	JobId             *int            `json:"job_id,omitempty"`
	OfferId           *int            `json:"offer_id,omitempty"`
	RequestedByUserId *int            `json:"requested_by_user_id,omitempty"`
	Sequential        *bool           `json:"sequential,omitempty"`
	Version           *int            `json:"version,omitempty"`
}

// Approver represents an approver in an approval flow.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-approval-flow-object
type Approver struct {
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmployeeId     *string  `json:"employee_id,omitempty"`
	Id             *int     `json:"id,omitempty"`
	Name           *string  `json:"name,omitempty"`
}

// ApproverGroup represents an approver group in an approval flow.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-approval-flow-object
type ApproverGroup struct {
	ApprovalsRequired *int       `json:"approvals_required,omitempty"`
	Approvers         []Approver `json:"approvers,omitempty"`
	CreatedAt         *string    `json:"created_at,omitempty"`
	Id                *int       `json:"id,omitempty"`
	JobId             *int       `json:"job_id,omitempty"`
	OfferId           *int       `json:"offer_id,omitempty"`
	Priority          *int       `json:"priority,omitempty"`
	ResolvedAt        *string    `json:"resolved_at,omitempty"`
}

// Attachment represents an attachment to a candidate or application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-attachment-to-application
type Attachment struct {
	Content     *string `json:"content,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Filename    *string `json:"filename,omitempty"`
	Type        *string `json:"type,omitempty"`
	Url         *string `json:"url,omitempty"`
	Visibility  *string `json:"visibility,omitempty"`
}

// Candidate represents a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-candidate-object
type Candidate struct {
	ActivityFeedNotes    []ActivityFeed              `json:"activity_feed_notes,omitempty"`
	Addresses            []TypeTypeValue             `json:"addresses,omitempty"`
	Application          *Application                `json:"application,omitempty"`
	ApplicationIds       []int                       `json:"application_ids,omitempty"`
	Applications         []Application               `json:"applications,omitempty"`
	Attachments          []Attachment                `json:"attachments,omitempty"`
	CanEmail             *bool                       `json:"can_email,omitempty"`
	Company              *string                     `json:"company,omitempty"`
	Coordinator          *User                       `json:"coordinator,omitempty"`
	CreatedAt            *string                     `json:"created_at,omitempty"`
	CustomFields         map[string]interface{}      `json:"custom_fields,omitempty"`
	Educations           []Education                 `json:"educations,omitempty"`
	EmailAddresses       []TypeTypeValue             `json:"email_addresses,omitempty"`
	Employments          []Employment                `json:"employments,omitempty"`
	ExposeInJobBoard     *bool                       `json:"expose_in_job_board,omitempty"`
	FirstName            *string                     `json:"first_name,omitempty"`
	KeyedCustomFields    map[string]KeyedCustomField `json:"keyed_custom_fields,omitempty"`
	Id                   *int                        `json:"id,omitempty"`
	IsPrivate            *bool                       `json:"is_private,omitempty"`
	IsProspect           *bool                       `json:"is_prospect,omitempty"`
	LastActivity         *string                     `json:"last_activity,omitempty"`
	LastName             *string                     `json:"last_name,omitempty"`
	LinkedUserIds        []int                       `json:"linked_user_ids,omitempty"`
	PhoneNumbers         []TypeTypeValue             `json:"phone_numbers,omitempty"`
	PhotoUrl             *string                     `json:"photo_url,omitempty"`
	Recruiter            *User                       `json:"recruiter,omitempty"`
	SocialMediaAddresses []TypeTypeValue             `json:"social_media_addresses,omitempty"`
	Tags                 []string                    `json:"tags,omitempty"`
	Title                *string                     `json:"title,omitempty"`
	UpdatedAt            *string                     `json:"updated_at,omitempty"`
	WebsiteAddresses     []TypeTypeValue             `json:"website_addresses,omitempty"`
}

// CandidateApplication represents the information needed to create a candidate application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-application-to-candidate-prospect
type CandidateApplication struct {
	Attachments    []Attachment   `json:"attachments,omitempty"`
	InitialStageId *int           `json:"initial_stage_id,omitempty"`
	JobId          *int           `json:"job_id,omitempty"`
	Referrer       *TypeTypeValue `json:"referrer,omitempty"`
	SourceId       *int           `json:"source_id,omitempty"`
}

// CandidateTag represents a tag.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-candidate-tag-object
type CandidateTag TypeIdName

// CloseReason represents a close reason.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-close-reason-object
type CloseReason TypeIdName

// CustomField represents a custom field.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-custom-field-object
type CustomField struct {
	Active              *bool               `json:"active,omitempty"`
	ApiOnly             *bool               `json:"api_only,omitempty"`
	CustomFieldOptions  []CustomFieldOption `json:"custom_field_options,omitempty"`
	DepartmentIds       []int               `json:"department_ids,omitempty"`
	Departments         []Department        `json:"departments,omitempty"`
	Description         *string             `json:"description,omitempty"`
	ExposeInJobBoardAPI *bool               `json:"expose_in_job_board_api"`
	FieldType           *string             `json:"field_type,omitempty"`
	GenerateEmailToken  *bool               `json:"generate_email_token,omitempty"`
	Id                  *int                `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	NameKey             *string             `json:"name_key,omitempty"`
	OfficeIds           []int               `json:"office_ids,omitempty"`
	Offices             []Office            `json:"offices,omitempty"`
	Priority            *int                `json:"priority,omitempty"`
	Private             *bool               `json:"private,omitempty"`
	RequireApproval     *bool               `json:"require_approval,omitempty"`
	Required            *bool               `json:"required,omitempty"`
	TemplateTokenString *string             `json:"template_token_string,omitempty"`
	TriggerNewVersion   *bool               `json:"trigger_new_version,omitempty"`
	ValueType           *string             `json:"value_type,omitempty"`
}

// CustomFieldOption represents a custom field option.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-custom-field-options-object
type CustomFieldOption struct {
	ExternalId *string `json:"external_id,omitempty"`
	Id         *int    `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Priority   *int    `json:"priority,omitempty"`
}

// Degree represents an education degree.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-education-objects
type Degree TypeIdNamePriority

// DemographicAnswer represents the answer to a demographic question.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-demographic-answer-object
type DemographicAnswer struct {
	ApplicationId             *int    `json:"application_id,omitempty"`
	CreatedAt                 *string `json:"created_at,omitempty"`
	DemographicAnswerOptionId *int    `json:"demographic_answer_option_id,omitempty"`
	DemographicQuestionId     *int    `json:"demographic_question_id,omitempty"`
	FreeFormText              *string `json:"free_form_text,omitempty"`
	Id                        *int    `json:"id,omitempty"`
	UpdatedAt                 *string `json:"updated_at,omitempty"`
}

// DemographicAnswerOption represents an option available in a demographic answer.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-demographic-answer-option-object
type DemographicAnswerOption struct {
	Active                *bool         `json:"active,omitempty"`
	DemographicQuestionId *int          `json:"demographic_question_id,omitempty"`
	FreeForm              *bool         `json:"free_form,omitempty"`
	Id                    *int          `json:"id,omitempty"`
	Name                  *string       `json:"name,omitempty"`
	Translations          []Translation `json:"translations,omitempty"`
}

// DemographicQuestion represents a demographic question.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-demographic-question-object
type DemographicQuestion struct {
	Active                   *bool         `json:"active,omitempty"`
	AnswerType               *string       `json:"answer_type,omitempty"`
	DemographicQuestionSetId *int          `json:"demographic_question_set_id,omitempty"`
	Id                       *int          `json:"id,omitempty"`
	Name                     *string       `json:"name,omitempty"`
	Required                 *bool         `json:"required,omitempty"`
	Translations             []Translation `json:"translations,omitempty"`
}

// DemographicQuestionSet represents a set of demographic questions.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-demographic-question-set-object
type DemographicQuestionSet struct {
	Active      *bool   `json:"active,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// Department represents a department.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-department-object
type Department struct {
	ChildDepartmentExternalIds []string `json:"child_department_external_ids,omitempty"`
	ChildIds                   []int    `json:"child_ids,omitempty"`
	ExternalId                 *string  `json:"external_id,omitempty"`
	Id                         *int     `json:"id,omitempty"`
	Name                       *string  `json:"name,omitempty"`
	ParentDepartmentExternalId *string  `json:"parent_department_external_id,omitempty"`
	ParentId                   *int     `json:"parent_id,omitempty"`
}

// Discipline represents an education discipline.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-education-objects
type Discipline TypeIdNamePriority

// Education represents a candidate's education.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-candidate
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-education
type Education struct {
	Degree       *string `json:"degree,omitempty"`
	DegreeId     *int    `json:"degree_id,omitempty"`
	Discipline   *string `json:"discipline,omitempty"`
	DisciplineId *int    `json:"discipline_id,omitempty"`
	EndDate      *string `json:"end_date,omitempty"`
	EndMonth     *string `json:"end_month,omitempty"`
	EndYear      *string `json:"end_year,omitempty"`
	Id           *int    `json:"id,omitempty"`
	SchoolId     *int    `json:"school_id,omitempty"`
	SchoolName   *string `json:"school_name,omitempty"`
	StartDate    *string `json:"start_date,omitempty"`
	StartMonth   *string `json:"start_month,omitempty"`
	StartYear    *string `json:"start_year,omitempty"`
}

// EEOC represents an EEOC object.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-eeoc-object
type EEOC struct {
	ApplicationId    *int        `json:"application_id,omitempty"`
	CandidateId      *int        `json:"candidate_id,omitempty"`
	DisabilityStatus *EEOCAnswer `json:"disability_status,omitempty"`
	Gender           *EEOCAnswer `json:"gender,omitempty"`
	Race             *EEOCAnswer `json:"race,omitempty"`
	SubmittedAt      *string     `json:"submitted_at,omitempty"`
	VeteranStatus    *EEOCAnswer `json:"veteran_status,omitempty"`
}

// EEOCAnswer represents an answer to a question in the EEOC object.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-eeoc-object
type EEOCAnswer struct {
	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Message     *string `json:"message,omitempty"`
}

// Email represents an email note in an activity feed.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-activity-feed-object
type Email struct {
	Body      *string `json:"body,omitempty"`
	Cc        *string `json:"cc,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	From      *string `json:"from,omitempty"`
	Id        *int    `json:"id,omitempty"`
	Subject   *string `json:"subject,omitempty"`
	To        *string `json:"to,omitempty"`
	UserId    *int    `json:"user_id,omitempty"`
	User      *User   `json:"user,omitempty"`
}

// EmailTemplate represents an email template.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-email-template-object
type EmailTemplate struct {
	Body        *string  `json:"body,omitempty"`
	Cc          []string `json:"cc,omitempty"`
	CreatedAt   *string  `json:"created_at,omitempty"`
	Default     *bool    `json:"default,omitempty"`
	Description *string  `json:"description,omitempty"`
	From        *string  `json:"from,omitempty"`
	HtmlBody    *string  `json:"html_body,omitempty"`
	Id          *int     `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Type        *string  `json:"type,omitempty"`
	UpdatedAt   *string  `json:"updated_at,omitempty"`
	User        *User    `json:"user,omitempty"`
}

// Employment represents a candidate's employment.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-candidate
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-employment
type Employment struct {
	CompanyName *string `json:"company_name,omitempty"`
	EndDate     *string `json:"end_date,omitempty"`
	Id          *int    `json:"id,omitempty"`
	StartDate   *string `json:"start_date,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// FutureJobPermission represents a user's future job permission.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-future-job-permission-object
type FutureJobPermission struct {
	DepartmentId         *int    `json:"department_id,omitempty"`
	ExternalDepartmentId *string `json:"external_department_id,omitempty"`
	ExternalOfficeId     *string `json:"external_office_id,omitempty"`
	Id                   *int    `json:"id,omitempty"`
	OfficeId             *int    `json:"office_id,omitempty"`
	UserRoleId           *int    `json:"user_role_id,omitempty"`
}

// HiringTeam represents a job's hiring team.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-hiring-team
type HiringTeam struct {
	Members []HiringMember `json:"members,omitempty"`
	Name    *string        `json:"name,omitempty"`
}

// HiringMember represents the member of a job's hiring team.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-hiring-team
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-hiring-team-members
type HiringMember struct {
	Active                           *bool   `json:"active,omitempty"`
	EmployeeId                       *string `json:"employee_id,omitempty"`
	FirstName                        *string `json:"first_name,omitempty"`
	Id                               *int    `json:"id,omitempty"`
	LastName                         *string `json:"last_name,omitempty"`
	Name                             *string `json:"name,omitempty"`
	Responsible                      *bool   `json:"responsible,omitempty"`
	ResponsibleForActiveCandidates   *bool   `json:"responsible_for_active_candidates,omitempty"`
	ResponsibleForFutureCandidates   *bool   `json:"responsible_for_future_candidates,omitempty"`
	ResponsibleForInactiveCandidates *bool   `json:"responsible_for_inactive_candidates,omitempty"`
	UserId                           *int    `json:"user_id,omitempty"`
}

// Interview represents an interview as part of a job stage.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-job-stage-object
type Interview struct {
	DefaultInterviewerUsers []Interviewer `json:"default_interviewer_users,omitempty"`
	EstimatedMinutes        *int          `json:"estimated_minutes,omitempty"`
	Id                      *int          `json:"id,omitempty"`
	InterviewKit            *InterviewKit `json:"interview_kit,omitempty"`
	Name                    *string       `json:"name,omitempty"`
	Schedulable             *bool         `json:"schedulable,omitempty"`
}

// Interviewer is a special User type to represent an interviewer.
type Interviewer struct {
	User
	Email          *string `json:"email,omitempty"`
	ResponseStatus *string `json:"response_status,omitempty"`
	ScorecardId    *int    `json:"scorecard_id,omitempty"`
	UserId         *int    `json:"user_id,omitempty"`
}

// InterviewKit represents a set of questions to be asked during an interview.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-job-stage-object
type InterviewKit struct {
	Content   *string             `json:"content,omitempty"`
	Id        *int                `json:"id,omitempty"`
	Questions []InterviewQuestion `json:"questions,omitempty"`
}

// InterviewQuestion represents an interview question in an interview kit.
type InterviewQuestion struct {
	Id       *int    `json:"id,omitempty"`
	Question *string `json:"question,omitempty"`
}

// Job represents a job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-job-object
type Job struct {
	ClosedAt          *string                     `json:"closed_at,omitempty"`
	Confidential      *bool                       `json:"confidential,omitempty"`
	CopiedFromId      *int                        `json:"copied_from_id,omitempty"`
	CreatedAt         *string                     `json:"created_at,omitempty"`
	CustomFields      map[string]interface{}      `json:"custom_fields,omitempty"`
	Departments       []Department                `json:"departments,omitempty"`
	HiringTeam        map[string][]HiringMember   `json:"hiring_team,omitempty"`
	Id                *int                        `json:"id,omitempty"`
	IsTemplate        *bool                       `json:"is_template,omitempty"`
	KeyedCustomFields map[string]KeyedCustomField `json:"keyed_custom_fields,omitempty"`
	Name              *string                     `json:"name,omitempty"`
	Notes             *string                     `json:"notes,omitempty"`
	Offices           []Office                    `json:"offices,omitempty"`
	OpenedAt          *string                     `json:"opened_at,omitempty"`
	Openings          []JobOpening                `json:"openings,omitempty"`
	RequisitionId     *string                     `json:"requisition_id,omitempty"`
	Status            *string                     `json:"status,omitempty"`
	UpdatedAt         *string                     `json:"updated_at,omitempty"`
}

// JobBoard represents a job board.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-tracking-link-object
type JobBoard struct {
	CompanyName *string `json:"company_name,omitempty"`
	Id          *int    `json:"id,omitempty"`
	UrlToken    *string `json:"url_token,omitempty"`
}

// JobCreateInfo represents the information needed to create a job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-create-job
type JobCreateInfo struct {
	DepartmentId         *int     `json:"department_id,omitempty"`
	ExternalDepartmentId *string  `json:"external_department_id,omitempty"`
	ExternalOfficeIds    []string `json:"external_office_ids,omitempty"`
	JobName              *string  `json:"job_name,omitempty"`
	JobPostName          *string  `json:"job_post_name,omitempty"`
	NumberOpenings       *int     `json:"number_of_openings,omitempty"`
	OfficeIds            []int    `json:"office_ids,omitempty"`
	OpeningIds           []string `json:"opening_ids,omitempty"`
	RequisitionId        *string  `json:"requisition_id,omitempty"`
	TemplateJobId        *int     `json:"template_job_id,omitempty"`
}

// JobUpdateInfo represents the information needed to update a job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-job
type JobUpdateInfo struct {
	Anywhere                *bool         `json:"anywhere,omitempty"`
	CustomFields            []CustomField `json:"custom_fields,omitempty"`
	DepartmentId            *int          `json:"department_id,omitempty"`
	ExternalDepartmentId    *string       `json:"external_department_id,omitempty"`
	ExternalOfficeIds       []string      `json:"external_office_ids,omitempty"`
	HowToSellThisJob        *string       `json:"how_to_sell_this_job,omitempty"`
	Name                    *string       `json:"name,omitempty"`
	Notes                   *string       `json:"notes,omitempty"`
	OfficeIds               []int         `json:"office_ids,omitempty"`
	RequisitionId           *string       `json:"requisition_id,omitempty"`
	TeamandResponsibilities *string       `json:"team_and_responsibilities,omitempty"`
}

// JobOpening represents a job opening.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-job-opening-object
type JobOpening struct {
	ApplicationId *int              `json:"application_id,omitempty"`
	ClosedAt      *string           `json:"closed_at,omitempty"`
	CloseReason   *CloseReason      `json:"close_reason,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
	Id            *int              `json:"id,omitempty"`
	OpenedAt      *string           `json:"opened_at,omitempty"`
	OpeningId     *string           `json:"opening_id,omitempty"`
	Status        *string           `json:"status,omitempty"`
}

// JobOpeningCreateInfo represents the information needed to create openings.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-create-new-openings
type JobOpeningCreateInfo struct {
	Openings []Opening `json:"openings,omitempty"`
}

// JobOpeningDeleteInfo represents the information needed to delete openings.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-destroy-openings
type JobOpeningDeleteInfo struct {
	Ids []int `json:"ids,omitempty"`
}

// JobOpeningUpdateInfo represents the information needed to update openings.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-edit-openings
type JobOpeningUpdateInfo struct {
	CloseReasonId *int                `json:"close_reason_id,omitempty"`
	CustomFields  []map[string]string `json:"custom_fields,omitempty"`
	Id            *int                `json:"opening_id,omitempty"`
	Status        *string             `json:"status,omitempty"`
}

// JobPost represents a job post.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-job-post-object
type JobPost struct {
	Active                   *bool                 `json:"active,omitempty"`
	Content                  *string               `json:"content,omitempty"`
	CreatedAt                *string               `json:"created_at,omitempty"`
	DemographicQuestionSetId *int                  `json:"demographic_question_set_id,omitempty"`
	External                 *bool                 `json:"external,omitempty"`
	FirstPublishedAt         *string               `json:"first_published_at,omitempty"`
	Id                       *int                  `json:"id,omitempty"`
	Internal                 *bool                 `json:"internal,omitempty"`
	InternalContent          *string               `json:"internal_content,omitempty"`
	JobId                    *int                  `json:"job_id,omitempty"`
	Live                     *bool                 `json:"live,omitempty"`
	Location                 *string               `json:"location,omitempty"`
	Questions                []DemographicQuestion `json:"questions,omitempty"`
	Title                    *string               `json:"title,omitempty"`
	UpdatedAt                *string               `json:"updated_at,omitempty"`
}

// JobPostUpdateInfo represents the information needed to update a job post.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-job-post
type JobPostUpdateInfo struct {
}

// JobStage represents a job stage.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-job-stage-object
type JobStage struct {
	CreatedAt  *string     `json:"created_at,omitempty"`
	Id         *int        `json:"id,omitempty"`
	Interviews []Interview `json:"interviews,omitempty"`
	JobId      *int        `json:"job_id,omitempty"`
	Name       *string     `json:"name,omitempty"`
	Priority   *int        `json:"priority,omitempty"`
	UpdatedAt  *string     `json:"updated_at,omitempty"`
}

// KeyedCustomField represents a keyed custom field. These are used throughout the API and
// are basically just a k-v map.
type KeyedCustomField struct {
	Name  *string      `json:"name,omitempty"`
	Type  *string      `json:"type,omitempty"`
	Value *interface{} `json:"value,omitempty"`
}

// Location represents the name of a location.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-office-object
type Location struct {
	Name *string `json:"name,omitempty"`
}

// Note represents a Note activity in the activity feed.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-activity-feed-object
type Note struct {
	Body       *string `json:"body,omitempty"`
	CreatedAt  *string `json:"created_at,omitempty"`
	Id         *int    `json:"id,omitempty"`
	Private    *bool   `json:"private,omitempty"`
	User       *User   `json:"user,omitempty"`
	UserId     *int    `json:"userid,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// Offer represents an offer made to a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-offer-object
type Offer struct {
	ApplicationId     *int                   `json:"application_id,omitempty"`
	CandidateId       *int                   `json:"candidate_id,omitempty"`
	CreatedAt         *string                `json:"created_at,omitempty"`
	CustomFields      map[string]interface{} `json:"custom_fields,omitempty"`
	Id                *int                   `json:"id,omitempty"`
	JobId             *int                   `json:"job_id,omitempty"`
	KeyedCustomFields map[string]interface{} `json:"keyed_custom_fields,omitempty"`
	Opening           *JobOpening            `json:"opening,omitempty"`
	ResolvedAt        *string                `json:"resolved_at,omitempty"`
	SentAt            *string                `json:"sent_at,omitempty"`
	StartsAt          *string                `json:"starts_at,omitempty"`
	Status            *string                `json:"status,omitempty"`
	Version           *int                   `json:"version,omitempty"`
}

// Office represents an office.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-office-object
type Office struct {
	ChildIds               []int     `json:"child_ids,omitempty"`
	ChildOfficeExternalIds []string  `json:"child_office_external_ids,omitempty"`
	ExternalId             *string   `json:"external_id,omitempty"`
	Id                     *int      `json:"id,omitempty"`
	Location               *Location `json:"location,omitempty"`
	Name                   *string   `json:"name,omitempty"`
	ParentId               *int      `json:"parent_id,omitempty"`
	ParentOfficeExternalId *string   `json:"parent_office_external_id,omitempty"`
	PrimaryContactUserId   *int      `json:"primary_contact_user_id,omitempty"`
}

// OfficeCreateInfo represents the information needed to create an office.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-office
type OfficeCreateInfo struct {
	ExternalId           *string `json:"external_id,omitempty"`
	ExternalParentId     *string `json:"external_parent_id,omitempty"`
	Location             *string `json:"location,omitempty"`
	Name                 *string `json:"name,omitempty"`
	ParentId             *int    `json:"parent_id,omitempty"`
	PrimaryContactUserId *int    `json:"primary_contact_user_id,omitempty"`
}

// OfficeUpdateInfo represents the information needed to update an office.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-edit-office
type OfficeUpdateInfo struct {
	ExternalId *string `json:"external_id,omitempty"`
	Location   *string `json:"location,omitempty"`
	Name       *string `json:"name,omitempty"`
}

// Opening represents the information needed to create a job opening.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-create-new-openings
type Opening struct {
	CustomFields []map[string]string `json:"custom_fields,omitempty"`
	Id           *int                `json:"opening_id,omitempty"`
}

// ProspectApplication represents the data required to create a prospect application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-application-to-candidate-prospect
type ProspectApplication struct {
	JobIds                  []int          `json:"job_ids,omitempty"`
	Prospect                *bool          `json:"prospect,omitempty"`
	ProspectiveDepartmentId *int           `json:"prospective_department_id,omitempty"`
	ProspectiveOfficeId     *int           `json:"prospective_office_id,omitempty"`
	ProspectOwnerId         *int           `json:"prospect_owner_id,omitempty"`
	ProspectPoolId          *int           `json:"prospect_pool_id,omitempty"`
	ProspectPoolStageId     *int           `json:"prospect_pool_stage_id,omitempty"`
	Referrer                *TypeTypeValue `json:"referrer,omitempty"`
	SourceId                *int           `json:"source_id,omitempty"`
}

// ProspectDetail represents detail about a prospect as used in an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-application
type ProspectDetail struct {
	ProspectOwner *TypeIdName `json:"prospect_owner,omitempty"`
	ProspectPool  *TypeIdName `json:"prospect_pool,omitempty"`
	ProspectStage *TypeIdName `json:"prospect_stage,omitempty"`
}

// ProspectPool represents a prospect pool.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-prospect-pools-object
type ProspectPool struct {
	Active         *bool   `json:"active,omitempty"`
	Id             *int    `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	ProspectStages []Stage `json:"prospect_stages,omitempty"`
}

// RejectionDetails represents information about an application rejection.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-application
type RejectionDetails struct {
	CustomFields      map[string]string           `json:"custom_fields,omitempty"`
	KeyedCustomFields map[string]KeyedCustomField `json:"keyed_custom_fields,omitempty"`
}

// RejectionEmail represents a rejection email.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-reject-application
type RejectionEmail struct {
	EmailTemplateId *string `json:"email_template_id,omitempty"`
	SendEmailAt     *string `json:"send_email_at,omitempty"`
}

// RejectionReason represents a rejection reason.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-reject-application
type RejectionReason struct {
	Id   *int        `json:"id,omitempty"`
	Name *string     `json:"name,omitempty"`
	Type *TypeIdName `json:"type,omitempty"`
}

// ScheduledInterview represents a scheduled interview.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-scheduled-interview-object
type ScheduledInterview struct {
	ApplicationId        *int                    `json:"application_id,omitempty"`
	End                  *ScheduledInterviewDate `json:"end,omitempty"`
	ExternalEventId      *string                 `json:"external_event_id,omitempty"`
	Id                   *int                    `json:"id,omitempty"`
	Interviewers         []Interviewer           `json:"interviewers,omitempty"`
	Location             *string                 `json:"location,omitempty"`
	Organizer            *User                   `json:"organizer,omitempty"`
	Start                *ScheduledInterviewDate `json:"start,omitempty"`
	Status               *string                 `json:"status,omitempty"`
	VideoConferencingUrl *string                 `json:"video_conferencing_url,omitempty"`
}

// ScheduledInterviewCreateInfo represents the information needed to create a scheduled interview.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-create-scheduled-interview
type ScheduledInterviewCreateInfo struct {
	ApplicationId   *int          `json:"application_id,omitempty"`
	End             *string       `json:"end,omitempty"`
	ExternalEventId *string       `json:"external_event_id,omitempty"`
	Interviewers    []Interviewer `json:"interviewers,omitempty"`
	InterviewId     *int          `json:"interview_id,omitempty"`
	Start           *string       `json:"start,omitempty"`
}

// ScheduledInterviewUpdateInfo represents the information needed to update a scheduled interview.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-scheduled-interview
type ScheduledInterviewUpdateInfo struct {
	End             *string       `json:"end,omitempty"`
	ExternalEventId *string       `json:"external_event_id,omitempty"`
	Interviewers    []Interviewer `json:"interviewers,omitempty"`
	Location        *string       `json:"location,omitempty"`
	Start           *string       `json:"start,omitempty"`
}

// ScheduledInterviewDate represents a date object for a scheduled interview.
type ScheduledInterviewDate struct {
	Date     *string `json:"date,omitempty"`
	DateTime *string `json:"date_time,omitempty"`
}

// School represents a school.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-education-objects
type School TypeIdNamePriority

// Scorecard represents a scorecard.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-scorecard-object
type Scorecard struct {
	ApplicationId         *int                 `json:"application_id,omitempty"`
	Attributes            []ScorecardAttribute `json:"attributes,omitempty"`
	CandidateId           *int                 `json:"candidate_id,omitempty"`
	CreatedAt             *string              `json:"created_at,omitempty"`
	Id                    *int                 `json:"id,omitempty"`
	Interview             *string              `json:"interview,omitempty"`
	InterviewedAt         *string              `json:"interviewed_at,omitempty"`
	InterviewStep         *Stage               `json:"interview_step,omitempty"`
	Interviewer           *User                `json:"interviewer,omitempty"`
	OverallRecommendation *string              `json:"overall_recommendation,omitempty"`
	Questions             []ScorecardQuestion  `json:"questions,omitempty"`
	Ratings               map[string][]string  `json:"ratings,omitempty"`
	SubmittedAt           *string              `json:"submitted_at,omitempty"`
	SubmittedBy           *User                `json:"submitted_by,omitempty"`
	UpdatedAt             *string              `json:"updated_at,omitempty"`
}

// ScorecardAttribute represents attributes present in a scorecard.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-scorecard-object
type ScorecardAttribute struct {
	Name   *string `json:"name,omitempty"`
	Note   *string `json:"note,omitempty"`
	Rating *string `json:"rating,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// ScorecardQuestion represents a question present in a scorecard.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-scorecard-object
type ScorecardQuestion struct {
	Answer   *string `json:"answer,omitempty"`
	Id       *int    `json:"id,omitempty"`
	Question *string `json:"question,omitempty"`
}

// Source represents a hiring source.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-source-object
type Source struct {
	TypeIdName
	PublicName *string     `json:"public_name,omitempty"`
	Type       *TypeIdName `json:"type,omitempty"`
}

// Stage represents an application stage.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-advance-application
type Stage TypeIdName

// TrackingLink represents a tracking link.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-tracking-link-object
type TrackingLink struct {
	CreatedAt       *string   `json:"created_at,omitempty"`
	CreditedTo      *User     `json:"credited_to,omitempty"`
	Id              *int      `json:"id,omitempty"`
	JobBoard        *JobBoard `json:"job_board,omitempty"`
	JobId           *int      `json:"job_id,omitempty"`
	JobPostId       *int      `json:"job_post_id,omitempty"`
	RelatedPostId   *int      `json:"related_post_id,omitempty"`
	RelatedPostType *string   `json:"related_post_type,omitempty"`
	Source          *Source   `json:"source,omitempty"`
	Token           *string   `json:"token,omitempty"`
	UpdatedAt       *string   `json:"updated_at,omitempty"`
}

// Translation represents the translation of a question in a demographic answer.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-demographic-answer-option-object
type Translation struct {
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// TypeIdNamePriority is a generic type for objects that have an ID, Name, and Priority.
type TypeIdNamePriority struct {
	TypeIdName
	Priority *int `json:"priority,omitempty"`
}

// TypeIdName is a generic type for objects that have an ID and Name.
type TypeIdName struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// TypeTypeValue is a generic type for objects that have a Type and Value.
type TypeTypeValue struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// User represents a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-user-object
type User struct {
	CreatedAt          *string  `json:"created_at,omitempty"`
	Disabled           *bool    `json:"disabled,omitempty"`
	Emails             []string `json:"emails,omitempty"`
	EmployeeId         *string  `json:"employee_id,omitempty"`
	FirstName          *string  `json:"first_name,omitempty"`
	Id                 *int     `json:"id,omitempty"`
	LastName           *string  `json:"last_name,omitempty"`
	LinkedCandidateIds []int    `json:"linked_candidate_ids,omitempty"`
	Name               *string  `json:"name,omitempty"`
	PrimaryEmail       *string  `json:"primary_email_address,omitempty"`
	SiteAdmin          *bool    `json:"site_admin,omitempty"`
	UpdatedAt          *string  `json:"updated_at,omitempty"`
}

// UserCreateInfo represents the information needed to create a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-user
type UserCreateInfo struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	SendEmail *bool   `json:"send_email_invite,omitempty"`
}

// UserEmailUpdateInfo represents the information needed to update an email address.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-e-mail-address-to-user
type UserEmailUpdateInfo struct {
	Email            *string `json:"email,omitempty"`
	SendVerification *bool   `json:"send_verification,omitempty"`
}

// UserPermission represents a user permission.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-job-permission-object
type UserPermission struct {
	Id         *int `json:"id,omitempty"`
	JobId      *int `json:"job_id,omitempty"`
	UserRoleId *int `json:"user_role_id,omitempty"`
}

// UserRole represents a user role.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#the-user-role-object
type UserRole struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// UserUpdateInfo represents the information needed to update user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-edit-user
type UserUpdateInfo struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

package greenhouse

import ()

type Activity struct {
	Body      string `json:"body,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Id        int    `json:"id,omitempty"`
	Subject   string `json:"subject,omitempty"`
	User      *User   `json:"user,omitempty"`
}

type ActivityFeed struct {
	Activities []Activity `json:"activities,omitempty"`
	Emails     []Email    `json:"emails,omitempty"`
	Id         int        `json:"id,omitempty"`
	Notes      []Note     `json:"notes,omitempty"`
}

type Answer struct {
	Answer   string `json:"answer,omitempty"`
	Question string `json:"question,omitempty"`
}

type Application struct {
	Answers                 []Answer                    `json:"answers,omitempty"`
	AppliedAt               string                      `json:"applied_at,omitempty"`
	Attachments             []Attachment                `json:"attachments,omitempty"`
	CandidateId             int                         `json:"candidate_id,omitempty"`
	CreditedTo              *User                        `json:"credited_to,omitempty"`
	CurrentStage            *Stage                       `json:"current_stage,omitempty"`
	CustomFields            map[string]string           `json:"custom_fields,omitempty"`
	Id                      int                         `json:"id,omitempty"`
	InitialStageId          int                         `json:"initial_stage_id,omitempty"`
	JobId                   int                         `json:"job_id,omitempty"`
	JobIds                  []int                       `json:"job_ids,omitempty"`
	Jobs                    []Job                       `json:"jobs,omitempty"`
	JobPostId               int                         `json:"job_post_id,omitempty"`
	KeyedCustomFields       map[string]KeyedCustomField `json:"keyed_custom_fields,omitempty"`
	LastActivityAt          string                      `json:"last_activity_at,omitempty"`
	Location                *Location                    `json:"location,omitempty"`
	Prospect                bool                        `json:"prospect,omitempty"`
	ProspectDetail          *ProspectDetail              `json:"prospect_detail,omitempty"`
	ProspectOwnerId         int                         `json:"prospect_owner_id,omitempty"`
	ProspectPoolId          int                         `json:"prospect_pool_id,omitempty"`
	ProspectPoolStageId     int                         `json:"prospect_pool_stage_id,omitempty"`
	ProspectStageId         int                         `json:"prospect_stage_id,omitempty"`
	ProspectiveDepartment   *Department                  `json:"prospective_department,omitempty"`
	ProspectiveDepartmentId int                         `json:"prospective_department_id,omitempty"`
	ProspectiveOffice       *Office                      `json:"prospective_office,omitempty"`
	ProspectiveOfficeId     int                         `json:"prospective_office_id,omitempty"`
	Referrer                *TypeTypeValue               `json:"referrer,omitempty"`
	RejectedAt              string                      `json:"rejected_at,omitempty"`
	RejectionDetails        string                      `json:"rejection_details,omitempty"`
	RejectionReason         string                      `json:"rejection_reason,omitempty"`
	Source                  *Source                      `json:"source,omitempty"`
	SourceId                int                         `json:"source_id,omitempty"`
	Status                  string                      `json:"status,omitempty"`
}

type ApplicationHire struct {
	CloseReasonId int    `json:"close_reason_id,omitempty"`
	OpeningId     int    `json:"opening_id,omitempty"`
	StartDate     string `json:"start_date,omitempty"`
}

type ApplicationReject struct {
	Notes             string         `json:"notes,omitempty"`
	RejectionEmail    *RejectionEmail `json:"rejection_email,omitempty"`
	RejectionReasonId int            `json:"rejection_reason_id,omitempty"`
}

type Approval struct {
	ApprovalStatus    string          `json:"approval_status,omitempty"`
	ApprovalType      string          `json:"approval_type,omitempty"`
	ApproverGroups    []ApproverGroup `json:"approver_groups,omitempty"`
	Id                int             `json:"id,omitempty"`
	JobId             int             `json:"job_id,omitempty"`
	OfferId           int             `json:"offer_id,omitempty"`
	RequestedByUserId int             `json:"requested_by_user_id,omitempty"`
	Sequential        bool            `json:"sequential,omitempty"`
	Version           int             `json:"version,omitempty"`
}

type Approver struct {
	EmailAddresses []string `json:"email_addresses,omitempty"`
	EmployeeId     string   `json:"employee_id,omitempty"`
	Id             int      `json:"id,omitempty"`
	Name           string   `json:"name,omitempty"`
}

type ApproverGroup struct {
	ApprovalsRequired int        `json:"approvals_required,omitempty"`
	Approvers         []Approver `json:"approvers,omitempty"`
	CreatedAt         string     `json:"created_at,omitempty"`
	Id                int        `json:"id,omitempty"`
	JobId             int        `json:"job_id,omitempty"`
	OfferId           int        `json:"offer_id,omitempty"`
	Priority          int        `json:"priority,omitempty"`
	ResolvedAt        string     `json:"resolved_at,omitempty"`
}

type Attachment struct {
	Content     string `json:"content,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Type        string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
	Visibility  string `json:"visibility,omitempty"`
}

type Candidate struct {
	ActivityFeedNotes    []ActivityFeed              `json:"activity_feed_notes,omitempty"`
	Addresses            []TypeTypeValue             `json:"addresses,omitempty"`
	Application          *Application                 `json:"application,omitempty"`
	ApplicationIds       []int                       `json:"application_ids,omitempty"`
	Applications         []Application               `json:"applications,omitempty"`
	Attachments          []Attachment                `json:"attachments,omitempty"`
	CanEmail             bool                        `json:"can_email,omitempty"`
	Company              string                      `json:"company,omitempty"`
	Coordinator          *User                        `json:"coordinator,omitempty"`
	CreatedAt            string                      `json:"created_at,omitempty"`
	CustomFields         map[string]interface{}      `json:"custom_fields,omitempty"`
	Educations           []Education                 `json:"educations,omitempty"`
	EmailAddresses       []TypeTypeValue             `json:"email_addresses,omitempty"`
	Employments          []Employment                `json:"employments,omitempty"`
	FirstName            string                      `json:"first_name,omitempty"`
	KeyedCustomFields    map[string]KeyedCustomField `json:"keyed_custom_fields,omitempty"`
	Id                   int                         `json:"id,omitempty"`
	IsPrivate            bool                        `json:"is_private,omitempty"`
	IsProspect           bool                        `json:"is_prospect,omitempty"`
	LastActivity         string                      `json:"last_activity,omitempty"`
	LastName             string                      `json:"last_name,omitempty"`
	LinkedUserIds        []int                       `json:"linked_user_ids,omitempty"`
	PhoneNumbers         []TypeTypeValue             `json:"phone_numbers,omitempty"`
	PhotoUrl             string                      `json:"photo_url,omitempty"`
	Recruiter            *User                        `json:"recruiter,omitempty"`
	SocialMediaAddresses []TypeTypeValue             `json:"social_media_addresses,omitempty"`
	Tags                 []string                    `json:"tags,omitempty"`
	Title                string                      `json:"title,omitempty"`
	UpdatedAt            string                      `json:"updated_at,omitempty"`
	WebsiteAddresses     []TypeTypeValue             `json:"website_addresses,omitempty"`
}

type CandidateApplication struct {
	Attachments    []Attachment  `json:"attachments,omitempty"`
	InitialStageId int           `json:"initial_stage_id,omitempty"`
	JobId          int           `json:"job_id,omitempty"`
	Referrer       *TypeTypeValue `json:"referrer,omitempty"`
	SourceId       int           `json:"source_id,omitempty"`
}

type CandidateTag TypeIdName

type CloseReason TypeIdName

type CustomField struct {
	Active              bool                `json:"active,omitempty"`
	ApiOnly             bool                `json:"api_only,omitempty"`
	CustomFieldOptions  []CustomFieldOption `json:"custom_field_options,omitempty"`
	DepartmentIds       []int               `json:"department_ids,omitempty"`
	Departments         []Department        `json:"departments,omitempty"`
	Description         string              `json:"description,omitempty"`
	ExposeInJobBoardAPI bool                `json:"expose_in_job_board_api"`
	FieldType           string              `json:"field_type,omitempty"`
	GenerateEmailToken  bool                `json:"generate_email_token,omitempty"`
	Id                  int                 `json:"id,omitempty"`
	Name                string              `json:"name,omitempty"`
	NameKey             string              `json:"name_key,omitempty"`
	OfficeIds           []int               `json:"office_ids,omitempty"`
	Offices             []Office            `json:"offices,omitempty"`
	Priority            int                 `json:"priority,omitempty"`
	Private             bool                `json:"private,omitempty"`
	RequireApproval     bool                `json:"require_approval,omitempty"`
	Required            bool                `json:"required,omitempty"`
	TemplateTokenString string              `json:"template_token_string,omitempty"`
	TriggerNewVersion   bool                `json:"trigger_new_version,omitempty"`
	ValueType           string              `json:"value_type,omitempty"`
}

type CustomFieldOption struct {
	ExternalId string `json:"external_id,omitempty"`
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Priority   int    `json:"priority,omitempty"`
}

type Degree TypeIdNamePriority

type DemographicAnswer struct {
	ApplicationId             int    `json:"application_id,omitempty"`
	CreatedAt                 string `json:"created_at,omitempty"`
	DemographicAnswerOptionId int    `json:"demographic_answer_option_id,omitempty"`
	DemographicQuestionId     int    `json:"demographic_question_id,omitempty"`
	FreeFormText              string `json:"free_form_text,omitempty"`
	Id                        int    `json:"id,omitempty"`
	UpdatedAt                 string `json:"updated_at,omitempty"`
}

type DemographicAnswerOption struct {
	Active                bool          `json:"active,omitempty"`
	DemographicQuestionId int           `json:"demographic_question_id,omitempty"`
	FreeForm              bool          `json:"free_form,omitempty"`
	Id                    int           `json:"id,omitempty"`
	Name                  string        `json:"name,omitempty"`
	Translations          []Translation `json:"translations,omitempty"`
}

type DemographicQuestion struct {
	Active                   bool          `json:"active,omitempty"`
	AnswerType               string        `json:"answer_type,omitempty"`
	DemographicQuestionSetId int           `json:"demographic_question_set_id,omitempty"`
	Id                       int           `json:"id,omitempty"`
	Name                     string        `json:"name,omitempty"`
	Required                 bool          `json:"required,omitempty"`
	Translations             []Translation `json:"translations,omitempty"`
}

type DemographicQuestionSet struct {
	Active      bool   `json:"active,omitempty"`
	Description string `json:"description,omitempty"`
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
}

type Department struct {
	ChildDepartmentExternalIds []string `json:"child_department_external_ids,omitempty"`
	ChildIds                   []int    `json:"child_ids,omitempty"`
	ExternalId                 string   `json:"external_id,omitempty"`
	Id                         int      `json:"id,omitempty"`
	Name                       string   `json:"name,omitempty"`
	ParentDepartmentExternalId string   `json:"parent_department_external_id,omitempty"`
	ParentId                   int      `json:"parent_id,omitempty"`
}

type DepartmentCreateInfo struct {
	Name     string `json:"name,omitempty"`
	ParentId int    `json:"parent_id,omitempty"`
}

type DepartmentUpdateInfo struct {
	Name string `json:"name,omitempty"`
}

type Discipline TypeIdNamePriority

type Education struct {
	Degree       string `json:"degree,omitempty"`
	DegreeId     int    `json:"degree_id,omitempty"`
	Discipline   string `json:"discipline,omitempty"`
	DisciplineId int    `json:"discipline_id,omitempty"`
	EndDate      string `json:"end_date,omitempty"`
	EndMonth     string `json:"end_month,omitempty"`
	EndYear      string `json:"end_year,omitempty"`
	Id           int    `json:"id,omitempty"`
	SchoolId     int    `json:"school_id,omitempty"`
	SchoolName   string `json:"school_name,omitempty"`
	StartDate    string `json:"start_date,omitempty"`
	StartMonth   int    `json:"start_month,omitempty"`
	StartYear    int    `json:"start_year,omitempty"`
}

type EEOC struct {
	ApplicationId    int        `json:"application_id,omitempty"`
	CandidateId      int        `json:"candidate_id,omitempty"`
	DisabilityStatus *EEOCAnswer `json:"disability_status,omitempty"`
	Gender           *EEOCAnswer `json:"gender,omitempty"`
	Race             *EEOCAnswer `json:"race,omitempty"`
	SubmittedAt      string     `json:"submitted_at,omitempty"`
	VeteranStatus    *EEOCAnswer `json:"veteran_status,omitempty"`
}

type EEOCAnswer struct {
	Description string `json:"description,omitempty"`
	Id          int    `json:"id,omitempty"`
	Message     string `json:"message,omitempty"`
}

type Email struct {
	Body      string `json:"body,omitempty"`
	Cc        string `json:"cc,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	From      string `json:"from,omitempty"`
	Id        int    `json:"id,omitempty"`
	Subject   string `json:"subject,omitempty"`
	To        string `json:"to,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	User      *User   `json:"user,omitempty"`
}

type EmailTemplate struct {
	Body      string   `json:"body,omitempty"`
	Cc        []string `json:"cc,omitempty"`
	CreatedAt string   `json:"created_at,omitempty"`
	Default   bool     `json:"default,omitempty"`
  Description string `json:"description,omitempty"`
	From      string   `json:"from,omitempty"`
	HtmlBody  string   `json:"html_body,omitempty"`
	Id        int      `json:"id,omitempty"`
  Name      string   `json:"name,omitempty"`
	Type      string   `json:"type,omitempty"`
	UpdatedAt string   `json:"updated_at,omitempty"`
	User      *User     `json:"user,omitempty"`
}

type Employment struct {
	CompanyName string `json:"company_name,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Id          int    `json:"id,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	Title       string `json:"title,omitempty"`
}

type FutureJobPermission struct {
	DepartmentId         int    `json:"department_id,omitempty"`
	ExternalDepartmentId string `json:"external_department_id,omitempty"`
	ExternalOfficeId     string `json:"external_office_id,omitempty"`
	Id                   int    `json:"id,omitempty"`
	OfficeId             int    `json:"office_id,omitempty"`
	UserRoleId           int    `json:"user_role_id,omitempty"`
}

type HiringTeam struct {
	Members []HiringMember `json:"members,omitempty"`
	Name    string         `json:"name,omitempty"`
}

type HiringMember struct {
	Active      bool   `json:"active,omitempty"`
	EmployeeId  string `json:"employee_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	Id          int    `json:"id,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Name        string `json:"name,omitempty"`
	Responsible bool   `json:"responsible,omitempty"`
	UserId      int    `json:"user_id,omitempty"`
}

type HiringMemberUpdateInfo struct {
	ResponsibleForActiveCandidates   bool `json:"responsible_for_active_candidates,omitempty"`
	ResponsibleForFutureCandidates   bool `json:"responsible_for_future_candidates,omitempty"`
	ResponsibleForInactiveCandidates bool `json:"responsible_for_inactive_candidates,omitempty"`
	UserId                           int  `json:"user_id,omitempty"`
}

type Interview struct {
	DefaultInterviewerUsers []Interviewer       `json:"default_interviewer_users,omitempty"`
	EstimatedMinutes        int          `json:"estimated_minutes,omitempty"`
	Id                      int          `json:"id,omitempty"`
	InterviewKit            *InterviewKit `json:"interview_kit,omitempty"`
	Name                    string       `json:"name,omitempty"`
	Schedulable             bool         `json:"schedulable,omitempty"`
}

type Interviewer struct {
	User
	Email          string `json:"email,omitempty"`
	ResponseStatus string `json:"response_status,omitempty"`
	ScorecardId    int    `json:"scorecard_id,omitempty"`
	UserId         int    `json:"user_id,omitempty"`
}

type InterviewKit struct {
	Content   string              `json:"content,omitempty"`
	Id        int                 `json:"id,omitempty"`
	Questions []InterviewQuestion `json:"questions,omitempty"`
}

type InterviewQuestion struct {
	Id       int    `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
}

type Job struct {
	ClosedAt          string                      `json:"closed_at,omitempty"`
	Confidential      bool                        `json:"confidential,omitempty"`
	CopiedFromId      int                         `json:"copied_from_id,omitempty"`
	CreatedAt         string                      `json:"created_at,omitempty"`
	CustomFields      map[string]interface{}      `json:"custom_fields,omitempty"`
	Departments       []Department                `json:"departments,omitempty"`
	HiringTeam        map[string][]HiringMember   `json:"hiring_team,omitempty"`
	Id                int                         `json:"id,omitempty"`
	IsTemplate        bool                        `json:"is_template,omitempty"`
	KeyedCustomFields map[string]KeyedCustomField `json:"keyed_custom_fields,omitempty"`
	Name              string                      `json:"name,omitempty"`
	Notes             string                      `json:"notes,omitempty"`
	Offices           []Office                    `json:"offices,omitempty"`
	OpenedAt          string                      `json:"opened_at,omitempty"`
	Openings          []JobOpening                `json:"openings,omitempty"`
	RequisitionId     string                      `json:"requisition_id,omitempty"`
	Status            string                      `json:"status,omitempty"`
	UpdatedAt         string                      `json:"updated_at,omitempty"`
}

type JobBoard struct {
	CompanyName string `json:"company_name,omitempty"`
	Id          int    `json:"id,omitempty"`
	UrlToken    string `json:"url_token,omitempty"`
}

type JobCreateInfo struct {
	DepartmentId         int      `json:"department_id,omitempty"`
	ExternalDepartmentId string   `json:"external_department_id,omitempty"`
	ExternalOfficeIds    []string `json:"external_office_ids,omitempty"`
	JobName              string   `json:"job_name,omitempty"`
	JobPostName          string   `json:"job_post_name,omitempty"`
	NumberOpenings       int      `json:"number_of_openings,omitempty"`
	OfficeIds            []int    `json:"office_ids,omitempty"`
	OpeningIds           []string `json:"opening_ids,omitempty"`
	RequisitionId        string   `json:"requisition_id,omitempty"`
	TemplateJobId        int      `json:"template_job_id,omitempty"`
}

type JobUpdateInfo struct {
	Anywhere                bool          `json:"anywhere,omitempty"`
	CustomFields            []CustomField `json:"custom_fields,omitempty"`
	DepartmentId            int           `json:"department_id,omitempty"`
	ExternalDepartmentId    string        `json:"external_department_id,omitempty"`
	ExternalOfficeIds       []string      `json:"external_office_ids,omitempty"`
	HowToSellThisJob        string        `json:"how_to_sell_this_job,omitempty"`
	Name                    string        `json:"name,omitempty"`
	Notes                   string        `json:"notes,omitempty"`
	OfficeIds               []int         `json:"office_ids,omitempty"`
	RequisitionId           string        `json:"requisition_id,omitempty"`
	TeamandResponsibilities string        `json:"team_and_responsibilities,omitempty"`
}

type JobOpening struct {
	ApplicationId int               `json:"application_id,omitempty"`
	ClosedAt      string            `json:"closed_at,omitempty"`
	CloseReason   *CloseReason       `json:"close_reason,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
	Id            int               `json:"id,omitempty"`
	OpenedAt      string            `json:"opened_at,omitempty"`
	OpeningId     string            `json:"opening_id,omitempty"`
	Status        string            `json:"status,omitempty"`
}

type JobOpeningCreateInfo struct {
	Openings []Opening `json:"openings,omitempty"`
}

type JobOpeningDeleteInfo struct {
	Ids []int `json:"ids,omitempty"`
}

type JobOpeningUpdateInfo struct {
	CloseReasonId int                 `json:"close_reason_id,omitempty"`
	CustomFields  []map[string]string `json:"custom_fields,omitempty"`
	Id            int                 `json:"opening_id,omitempty"`
	Status        string              `json:"status,omitempty"`
}

type JobPost struct {
	Active                   bool                  `json:"active,omitempty"`
	Content                  string                `json:"content,omitempty"`
	CreatedAt                string                `json:"created_at,omitempty"`
	DemographicQuestionSetId int                   `json:"demographic_question_set_id,omitempty"`
	External                 bool                  `json:"external,omitempty"`
	FirstPublishedAt         string                `json:"first_published_at,omitempty"`
	Id                       int                   `json:"id,omitempty"`
	Internal                 bool                  `json:"internal,omitempty"`
	InternalContent          string                `json:"internal_content,omitempty"`
	JobId                    int                   `json:"job_id,omitempty"`
	Live                     bool                  `json:"live,omitempty"`
	Questions                []DemographicQuestion `json:"questions,omitempty"`
	UpdatedAt                string                `json:"updated_at,omitempty"`
}

type JobPostUpdateInfo struct {
}

type JobStage struct {
	CreatedAt  string      `json:"created_at,omitempty"`
	Id         int         `json:"id,omitempty"`
	Interviews []Interview `json:"interviews,omitempty"`
	JobId      int         `json:"job_id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Priority   int         `json:"priority,omitempty"`
	UpdatedAt  string      `json:"updated_at,omitempty"`
}

type KeyedCustomField struct {
	Name  string      `json:"name,omitempty"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Location struct {
	Name string `json:"name,omitempty"`
}

type Note struct {
	Body       string `json:"body,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	Id         int    `json:"id,omitempty"`
	Private    bool   `json:"private,omitempty"`
	User       *User   `json:"user,omitempty"`
	UserId     int    `json:"userid,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type Offer struct {
	ApplicationId     int                    `json:"application_id,omitempty"`
	CandidateId       int                    `json:"candidate_id,omitempty"`
	CreatedAt         string                 `json:"created_at,omitempty"`
	CustomFields      map[string]interface{} `json:"custom_fields,omitempty"`
	Id                int                    `json:"id,omitempty"`
	JobId             int                    `json:"job_id,omitempty"`
	KeyedCustomFields map[string]interface{} `json:"keyed_custom_fields,omitempty"`
	Opening           *JobOpening             `json:"opening,omitempty"`
	ResolvedAt        string                 `json:"resolved_at,omitempty"`
	SentAt            string                 `json:"sent_at,omitempty"`
	StartsAt          string                 `json:"starts_at,omitempty"`
	Status            string                 `json:"status,omitempty"`
	Version           int                    `json:"version,omitempty"`
}

type Office struct {
	ChildIds               []int    `json:"child_ids,omitempty"`
	ChildOfficeExternalIds []string `json:"child_office_external_ids,omitempty"`
	ExternalId             string   `json:"external_id,omitempty"`
	Id                     int      `json:"id,omitempty"`
	Location               *Location `json:"location,omitempty"`
	Name                   string   `json:"name,omitempty"`
	ParentId               int      `json:"parent_id,omitempty"`
	ParentOfficeExternalId string   `json:"parent_office_external_id,omitempty"`
	PrimaryContactUserId   int      `json:"primary_contact_user_id,omitempty"`
}

type OfficeCreateInfo struct {
	ExternalId           string `json:"external_id,omitempty"`
	ExternalParentId     string `json:"external_parent_id,omitempty"`
	Location             string `json:"location,omitempty"`
	Name                 string `json:"name,omitempty"`
	ParentId             int    `json:"parent_id,omitempty"`
	PrimaryContactUserId int    `json:"primary_contact_user_id,omitempty"`
}

type OfficeUpdateInfo struct {
	ExternalId string `json:"external_id,omitempty"`
	Location   string `json:"location,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Opening struct {
	CustomFields []map[string]string `json:"custom_fields,omitempty"`
	Id           int                 `json:"opening_id,omitempty"`
}

type ProspectApplication struct {
	JobIds                  []int         `json:"job_ids,omitempty"`
	Prospect                bool          `json:"prospect,omitempty"`
	ProspectiveDepartmentId int           `json:"prospective_department_id,omitempty"`
	ProspectiveOfficeId     int           `json:"prospective_office_id,omitempty"`
	ProspectOwnerId         int           `json:"prospect_owner_id,omitempty"`
	ProspectPoolId          int           `json:"prospect_pool_id,omitempty"`
	ProspectPoolStageId     int           `json:"prospect_pool_stage_id,omitempty"`
	Referrer                *TypeTypeValue `json:"referrer,omitempty"`
	SourceId                int           `json:"source_id,omitempty"`
}

type ProspectDetail struct {
	ProspectOwner string `json:"prospect_owner,omitempty"`
	ProspectPool  string `json:"prospect_pool,omitempty"`
	ProspectStage string `json:"prospect_stage,omitempty"`
}

type ProspectPool struct {
	Active         bool    `json:"active,omitempty"`
	Id             int     `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	ProspectStages []Stage `json:"prospect_stages,omitempty"`
}

type RejectionEmail struct {
	EmailTemplateId string `json:"email_template_id,omitempty"`
	SendEmailAt     string `json:"send_email_at,omitempty"`
}

type RejectionReason struct {
	Id   int        `json:"id,omitempty"`
	Name string     `json:"name,omitempty"`
	Type TypeIdName `json:"type,omitempty"`
}

type ScheduledInterview struct {
	ApplicationId        int                    `json:"application_id,omitempty"`
	End                  *ScheduledInterviewDate `json:"end,omitempty"`
	ExternalEventId      string                 `json:"external_event_id,omitempty"`
	Id                   int                    `json:"id,omitempty"`
	Interviewers         []Interviewer          `json:"interviewers,omitempty"`
	Location             string                 `json:"location,omitempty"`
	Organizer            *User                   `json:"organizer,omitempty"`
	Start                *ScheduledInterviewDate `json:"start,omitempty"`
	Status               string                 `json:"status,omitempty"`
	VideoConferencingUrl string                 `json:"video_conferencing_url,omitempty"`
}

type ScheduledInterviewDate struct {
	Date     string `json:"date,omitempty"`
	DateTime string `json:"date_time,omitempty"`
}

type School TypeIdNamePriority

type Scorecard struct {
	ApplicationId         int                  `json:"application_id,omitempty"`
	Attributes            []ScorecardAttribute `json:"attributes,omitempty"`
	CandidateId           int                  `json:"candidate_id,omitempty"`
	CreatedAt             string               `json:"created_at,omitempty"`
	Id                    int                  `json:"id,omitempty"`
	Interview             string               `json:"interview,omitempty"`
	InterviewedAt         string               `json:"interviewed_at,omitempty"`
	InterviewStep         *Stage                `json:"interview_step,omitempty"`
	Interviewer           *User                  `json:"interviewer,omitempty"`
	OverallRecommendation string               `json:"overall_recommendation,omitempty"`
	Questions             []ScorecardQuestion  `json:"questions,omitempty"`
	Ratings               map[string][]string  `json:"ratings,omitempty"`
	SubmittedAt           string               `json:"submitted_at,omitempty"`
	SubmittedBy           *User                 `json:"submitted_by,omitempty"`
	UpdatedAt             string               `json:"updated_at,omitempty"`
}

type ScorecardAttribute struct {
	Name   string `json:"name,omitempty"`
	Note   string `json:"note,omitempty"`
	Rating string `json:"rating,omitempty"`
	Type   string `json:"type,omitempty"`
}

type ScorecardQuestion struct {
	Answer   string `json:"answer,omitempty"`
	Id       int    `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
}

type Source struct {
	TypeIdName
	PublicName string     `json:"public_name,omitempty"`
	Type       *TypeIdName `json:"type,omitempty"`
}

type Stage TypeIdName

type TrackingLink struct {
	CreatedAt       string   `json:"created_at,omitempty"`
	CreditedTo      *User     `json:"credited_to,omitempty"`
	Id              int      `json:"id,omitempty"`
	JobBoard        *JobBoard `json:"job_board,omitempty"`
	JobId           int      `json:"job_id,omitempty"`
	JobPostId       int      `json:"job_post_id,omitempty"`
	RelatedPostId   int      `json:"related_post_id,omitempty"`
	RelatedPostType string   `json:"related_post_type,omitempty"`
	Source          *Source   `json:"source,omitempty"`
	Token           string   `json:"token,omitempty"`
	UpdatedAt       string   `json:"updated_at,omitempty"`
}

type Translation struct {
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`
}

type TypeIdNamePriority struct {
	TypeIdName
	Priority int `json:"priority,omitempty"`
}

type TypeIdName struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TypeTypeValue struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type User struct {
	CreatedAt          string   `json:"created_at,omitempty"`
	Disabled           bool     `json:"disabled,omitempty"`
	Emails             []string `json:"emails,omitempty"`
	EmployeeId         string   `json:"employee_id,omitempty"`
	FirstName          string   `json:"first_name,omitempty"`
	Id                 int      `json:"id,omitempty"`
	LastName           string   `json:"last_name,omitempty"`
	LinkedCandidateIds []int    `json:"linked_candidate_ids,omitempty"`
	Name               string   `json:"name,omitempty"`
	PrimaryEmail       string   `json:"primary_email_address,omitempty"`
	SiteAdmin          bool     `json:"site_admin,omitempty"`
	UpdatedAt          string   `json:"updated_at,omitempty"`
}

type UserCreateInfo struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	SendEmail bool   `json:"send_email_invite,omitempty"`
}

type UserEmailUpdateInfo struct {
	Email            string `json:"email,omitempty"`
	SendVerification bool   `json:"send_verification,omitempty"`
}

type UserPermission struct {
	Id         int `json:"id,omitempty"`
	JobId      int `json:"job_id,omitempty"`
	UserRoleId int `json:"user_role_id,omitempty"`
}

type UserRole struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type UserUpdateInfo struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

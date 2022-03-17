package greenhouse

import ()

type Activity struct {
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	Subject   string `json:"subject"`
	User      User   `json:"user"`
}

type ActivityFeed struct {
	Activities []Activity `json:"activities"`
	Emails     []Email    `json:"emails"`
	Id         int        `json:"id"`
	Notes      []Note     `json:"notes"`
}

type Answer struct {
	Answer   string `json:"answer"`
	Question string `json:"question"`
}

type Application struct {
	Answers               []Answer               `json:"answers"`
	AppliedAt             string                 `json:"applied_at"`
	Attachments           []Attachment           `json:"attachments"`
	CandidateId           int                    `json:"candidate_id"`
	CreditedTo            User                   `json:"credited_to"`
	CurrentStage          Stage                  `json:"current_stage"`
	CustomFields          map[string]string      `json:"custom_fields"`
	Id                    int                    `json:"id"`
	Jobs                  []Job                  `json:"jobs"`
	JobPostId             int                    `json:"job_post_id"`
	KeyedCustomFields     map[string]interface{} `json:"keyed_custom_fields"`
	LastActivityAt        string                 `json:"last_activity_at"`
	Location              Location               `json:"location"`
	Prospect              bool                   `json:"prospect"`
	ProspectDetail        ProspectDetail         `json:"prospect_detail"`
	ProspectiveDepartment Department             `json:"prospective_department"`
	ProspectiveOffice     Office                 `json:"prospective_office"`
	RejectedAt            string                 `json:"rejected_at"`
	Source                Source                 `json:"source"`
	Status                string                 `json:"status"`
}

type Approval struct {
	ApprovalStatus    string          `json:"approval_status"`
	ApprovalType      string          `json:"approval_type"`
	ApproverGroups    []ApproverGroup `json:"approver_groups"`
	Id                int             `json:"id"`
	JobId             int             `json:"job_id"`
	OfferId           int             `json:"offer_id"`
	RequestedByUserId int             `json:"requested_by_user_id"`
	Sequential        bool            `json:"sequential"`
	Version           int             `json:"version"`
}

type ApproverGroup struct {
	ApprovalsRequired int        `json:"approvals_required"`
	Approvers         []Approver `json:"approvers"`
	CreatedAt         string     `json:"created_at"`
	Id                int        `json:"id"`
	JobId             int        `json:"job_id"`
	OfferId           int        `json:"offer_id"`
	Priority          int        `json:"priority"`
	ResolvedAt        string     `json:"resolved_at"`
}

type Approver struct {
	EmailAddresses []string `json:"email_addresses"`
	EmployeeId     string   `json:"employee_id"`
	Id             int      `json:"id"`
	Name           string   `json:"name"`
}

type Attachment struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Url      string `json:"url"`
}

type Candidate struct {
	Addresses            []TypeTypeValue        `json:"addresses"`
	ApplicationIds       []int                  `json:"application_ids"`
	Applications         []Application          `json:"applications"`
	CanEmail             bool                   `json:"can_email"`
	Company              string                 `json:"company"`
	Coordinator          User                   `json:"coordinator"`
	CustomFields         map[string]interface{} `json:"custom_fields"`
	Educations           []Education            `json:"educations"`
	EmailAddresses       []TypeTypeValue        `json:"email_addresses"`
	Employments          []Employment           `json:"employments"`
	KeyedCustomFields    map[string]interface{} `json:"keyed_custom_fields"`
	Id                   int                    `json:"id"`
	IsPrivate            bool                   `json:"is_private"`
	LinkedUserIds        []int                  `json:"linked_user_ids"`
	PhoneNumbers         []TypeTypeValue        `json:"phone_numbers"`
	Recruiter            User                   `json:"recruiter"`
	SocialMediaAddresses []TypeTypeValue        `json:"social_media_addresses"`
	Tags                 []string               `json:"tags"`
	Title                string                 `json:"title"`
	WebsiteAddresses     []TypeTypeValue        `json:"website_addresses"`
}

type CandidateTag TypeIdName

type CloseReason TypeIdName

type CustomField struct {
	Active              bool                `json:"active"`
	ApiOnly             bool                `json:"api_only"`
	CustomFieldOptions  []CustomFieldOption `json:"custom_field_options"`
	Departments         []Department        `json:"departments"`
	Description         string              `json:"description"`
	ExposeInJobBoardAPI bool                `json:"expose_in_job_board_api"`
	FieldType           string              `json:"field_type"`
	Id                  int                 `json:"id"`
	Name                string              `json:"name"`
	NameKey             string              `json:"name_key"`
	Offices             []Office            `json:"offices"`
	Priority            int                 `json:"priority"`
	Private             bool                `json:"private"`
	RequireApproval     bool                `json:"require_approval"`
	Required            bool                `json:"required"`
	TemplateTokenString string              `json:"template_token_string"`
	TriggerNewVersion   bool                `json:"trigger_new_version"`
	ValueType           string              `json:"value_type"`
}

type CustomFieldOption struct {
	ExternalId string `json:"external_id"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
}

type Degree TypeIdNamePriority

type DemographicAnswer struct {
	ApplicationId             int    `json:"application_id"`
	CreatedAt                 string `json:"created_at"`
	DemographicAnswerOptionId int    `json:"demographic_answer_option_id"`
	DemographicQuestionId     int    `json:"demographic_question_id"`
	FreeFormText              string `json:"free_form_text"`
	Id                        int    `json:"id"`
	UpdatedAt                 string `json:"updated_at"`
}

type DemographicAnswerOption struct {
	Active                bool          `json:"active"`
	DemographicQuestionId int           `json:"demographic_question_id"`
	FreeForm              bool          `json:"free_form"`
	Id                    int           `json:"id"`
	Name                  string        `json:"name"`
	Translations          []Translation `json:"translations"`
}

type DemographicQuestion struct {
	Active                   bool          `json:"active"`
	AnswerType               string        `json:"answer_type"`
	DemographicQuestionSetId int           `json:"demographic_question_set_id"`
	Id                       int           `json:"id"`
	Name                     string        `json:"name"`
	Required                 bool          `json:"required"`
	Translations             []Translation `json:"translations"`
}

type DemographicQuestionSet struct {
	Active      bool   `json:"active"`
	Description string `json:"description"`
	Id          int    `json:"id"`
	Title       string `json:"title"`
}

type Department struct {
	ChildDepartmentExternalIds []string `json:"child_department_external_ids"`
	ChildIds                   []int    `json:"child_ids"`
	ExternalId                 string   `json:"external_id"`
	Id                         int      `json:"id"`
	Name                       string   `json:"name"`
	ParentDepartmentExternalId string   `json:"parent_department_external_id"`
	ParentId                   int      `json:"parent_id"`
}

type DepartmentCreateInfo struct {
	Name     string `json:"name"`
	ParentId int    `json:"parent_id,omitempty"`
}

type DepartmentUpdateInfo struct {
	Name string `json:"name"`
}

type Discipline TypeIdNamePriority

type Education struct {
	Degree     string `json:"degree"`
	Discipline string `json:"discipline"`
	EndDate    string `json:"end_date"`
	Id         int    `json:"id"`
	SchoolName string `json:"school_name"`
	StartDate  string `json:"start_date"`
}

type EEOC struct {
	ApplicationId    int        `json:"application_id"`
	CandidateId      int        `json:"candidate_id"`
	DisabilityStatus EEOCAnswer `json:"disability_status"`
	Gender           EEOCAnswer `json:"gender"`
	Race             EEOCAnswer `json:"race"`
	SubmittedAt      string     `json:"submitted_at"`
	VeteranStatus    EEOCAnswer `json:"veteran_status"`
}

type EEOCAnswer struct {
	Description string `json:"description,omitempty"`
	Id          int    `json:"id"`
	Message     string `json:"message,omitempty"`
}

type Email struct {
	Body      string `json:"body"`
	Cc        string `json:"cc"`
	CreatedAt string `json:"created_at"`
	From      string `json:"from"`
	Id        int    `json:"id"`
	Subject   string `json:"subject"`
	To        string `json:"to"`
	User      User   `json:"user"`
}

type EmailTemplate struct {
	Body      string   `json:"body"`
	Cc        []string `json:"cc"`
	CreatedAt string   `json:"created_at"`
	Default   bool     `json:"default"`
	From      string   `json:"from"`
	HtmlBody  string   `json:"html_body"`
	Id        int      `json:"id"`
	Type      string   `json:"type"`
	UpdatedAt string   `json:"updated_at"`
	User      User     `json:"user"`
}

type Employment struct {
	CompanyName string `json:"company_name"`
	EndDate     string `json:"end_date"`
	Id          int    `json:"id"`
	StartDate   string `json:"start_date"`
	Title       string `json:"title"`
}

type FutureJobPermission struct {
	DepartmentId         int    `json:"department_id"`
	ExternalDepartmentId string `json:"external_department_id,omitempty"`
	ExternalOfficeId     string `json:"external_office_id,omitempty"`
	Id                   int    `json:"id"`
	OfficeId             int    `json:"office_id"`
	UserRoleId           int    `json:"user_role_id"`
}

type HiringTeam struct {
	Members []HiringMember `json:"members"`
	Name    string         `json:"name"`
}

type HiringMember struct {
	Active      bool   `json:"active"`
	EmployeeId  string `json:"employee_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	Id          int    `json:"id"`
	LastName    string `json:"last_name,omitempty"`
	Name        string `json:"name,omitempty"`
	Responsible bool   `json:"responsible"`
	UserId      int    `json:"user_id"`
}

type HiringMemberUpdateInfo struct {
	ResponsibleForActiveCandidates   bool `json:"responsible_for_active_candidates"`
	ResponsibleForFutureCandidates   bool `json:"responsible_for_future_candidates"`
	ResponsibleForInactiveCandidates bool `json:"responsible_for_inactive_candidates"`
	UserId                           int  `json:"user_id"`
}

type Interview struct {
	DefaultInterviewerUsers []User       `json:"default_interviewer_users"`
	EstimatedMinutes        int          `json:"estimated_minutes"`
	Id                      int          `json:"id"`
	InterviewKit            InterviewKit `json:"interview_kit"`
	Name                    string       `json:"name"`
	Schedulable             bool         `json:"schedulable"`
}

type Interviewer struct {
	User
	Email          string `json:"email"`
	ResponseStatus string `json:"response_status"`
	ScorecardId    int    `json:"scorecard_id"`
}

type InterviewKit struct {
	Content   string              `json:"content"`
	Id        int                 `json:"id"`
	Questions []InterviewQuestion `json:"questions"`
}

type InterviewQuestion struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
}

type Job struct {
	ClosedAt          string                            `json:"closed_at"`
	Confidential      bool                              `json:"confidential"`
	CopiedFromId      int                               `json:"copied_from_id"`
	CreatedAt         string                            `json:"created_at"`
	CustomFields      map[string]interface{}            `json:"custom_fields"`
	Departments       []Department                      `json:"departments"`
	HiringTeam        map[string][]HiringMember         `json:"hiring_team"`
	Id                int                               `json:"id"`
	IsTemplate        bool                              `json:"is_template"`
	KeyedCustomFields map[string]map[string]interface{} `json:"keyed_custom_fields"`
	Name              string                            `json:"name"`
	Notes             string                            `json:"notes"`
	Offices           []Office                          `json:"offices"`
	OpenedAt          string                            `json:"opened_at"`
	Openings          []JobOpening                      `json:"openings"`
	RequisitionId     string                            `json:"requisition_id"`
	Status            string                            `json:"status"`
	UpdatedAt         string                            `json:"updated_at"`
}

type JobBoard struct {
	CompanyName string `json:"company_name"`
	Id          int    `json:"id"`
	UrlToken    string `json:"url_token"`
}

type JobCreateInfo struct {
	DepartmentId         int      `json:"department_id,omitempty"`
	ExternalDepartmentId string   `json:"external_department_id"`
	ExternalOfficeIds    []string `json:"external_office_ids"`
	JobName              string   `json:"job_name,omitempty"`
	JobPostName          string   `json:"job_post_name,omitempty"`
	NumberOpenings       int      `json:"number_of_openings"`
	OfficeIds            []int    `json:"office_ids,omitempty"`
	OpeningIds           []string `json:"opening_ids,omitempty"`
	RequisitionId        string   `json:"requisition_id,omitempty"`
	TemplateJobId        int      `json:"template_job_id"`
}

type JobUpdateInfo struct {
	Anywhere                bool          `json:"anywhere,omitempty"`
	CustomFields            []CustomField `json:"custom_fields,omitempty"`
	DepartmentId            int           `json:"department_id,omitempty"`
	ExternalDepartmentId    string        `json:"external_department_id"`
	ExternalOfficeIds       []string      `json:"external_office_ids"`
	HowToSellThisJob        string        `json:"how_to_sell_this_job,omitempty"`
	Name                    string        `json:"name,omitempty"`
	Notes                   string        `json:"notes,omitempty"`
	OfficeIds               []int         `json:"office_ids,omitempty"`
	RequisitionId           string        `json:"requisition_id,omitempty"`
	TeamandResponsibilities string        `json:"team_and_responsibilities,omitempty"`
}

type JobOpening struct {
	ApplicationId int               `json:"application_id"`
	ClosedAt      string            `json:"closed_at"`
	CloseReason   CloseReason       `json:"close_reason"`
	CustomFields  map[string]string `json:"custom_fields"`
	Id            int               `json:"id"`
	OpenedAt      string            `json:"opened_at"`
	OpeningId     string            `json:"opening_id"`
	Status        string            `json:"status"`
}

type JobOpeningCreateInfo struct {
	Openings []Opening `json:"openings"`
}

type JobOpeningDeleteInfo struct {
	Ids []int `json:"ids"`
}

type JobOpeningUpdateInfo struct {
	CloseReasonId int                 `json:"close_reason_id,omitempty"`
	CustomFields  []map[string]string `json:"custom_fields,omitempty"`
	Id            int                 `json:"opening_id,omitempty"`
	Status        string              `json:"status,omitempty"`
}

type JobPost struct {
	Active                   bool                  `json:"active"`
	Content                  string                `json:"content"`
	CreatedAt                string                `json:"created_at"`
	DemographicQuestionSetId int                   `json:"demographic_question_set_id"`
	External                 bool                  `json:"external"`
	FirstPublishedAt         string                `json:"first_published_at"`
	Id                       int                   `json:"id"`
	Internal                 bool                  `json:"internal"`
	InternalContent          string                `json:"internal_content"`
	JobId                    int                   `json:"job_id"`
	Live                     bool                  `json:"live"`
	Questions                []DemographicQuestion `json:"questions"`
	UpdatedAt                string                `json:"updated_at"`
}

type JobPostUpdateInfo struct {
}

type JobStage struct {
	CreatedAt  string      `json:"created_at"`
	Id         int         `json:"id"`
	Interviews []Interview `json:"interviews"`
	JobId      int         `json:"job_id"`
	Name       string      `json:"name"`
	Priority   int         `json:"priority"`
	UpdatedAt  string      `json:"updated_at"`
}

type Location struct {
	Name string `json:"name"`
}

type Note struct {
	Body       string `json:"body"`
	CreatedAt  string `json:"created_at"`
	Id         int    `json:"id"`
	Private    bool   `json:"private"`
	User       User   `json:"user"`
	Visibility string `json:"visibility"`
}

type Offer struct {
	ApplicationId     int                    `json:"application_id"`
	CandidateId       int                    `json:"candidate_id"`
	CreatedAt         string                 `json:"created_at"`
	CustomFields      map[string]interface{} `json:"custom_fields"`
	Id                int                    `json:"id"`
	JobId             int                    `json:"job_id"`
	KeyedCustomFields map[string]interface{} `json:"keyed_custom_fields"`
	Opening           JobOpening             `json:"opening"`
	ResolvedAt        string                 `json:"resolved_at"`
	SentAt            string                 `json:"sent_at"`
	StartsAt          string                 `json:"starts_at"`
	Status            string                 `json:"status"`
	Version           int                    `json:"version"`
}

type Office struct {
	ChildIds               []int    `json:"child_ids"`
	ChildOfficeExternalIds []string `json:"child_office_external_ids"`
	ExternalId             string   `json:"external_id"`
	Id                     int      `json:"id"`
	Location               Location `json:"location"`
	Name                   string   `json:"name"`
	ParentId               int      `json:"parent_id"`
	ParentOfficeExternalId string   `json:"parent_office_external_id"`
	PrimaryContactUserId   int      `json:"primary_contact_user_id"`
}

type OfficeCreateInfo struct {
	ExternalId           string `json:"external_id,omitempty"`
	ExternalParentId     string `json:"external_parent_id,omitempty"`
	Location             string `json:"location,omitempty"`
	Name                 string `json:"name"`
	ParentId             int    `json:"parent_id,omitempty"`
	PrimaryContactUserId int    `json:"primary_contact_user_id,omitempty"`
}

type OfficeUpdateInfo struct {
	ExternalId string `json:"external_id,omitempty"`
	Location   string `json:"location,omitempty"`
	Name       string `json:"name"`
}

type Opening struct {
	CustomFields []map[string]string `json:"custom_fields"`
	Id           int                 `json:"opening_id"`
}

type ProspectDetail struct {
	ProspectOwner string `json:"prospect_owner"`
	ProspectPool  string `json:"prospect_pool"`
	ProspectStage string `json:"prospect_stage"`
}

type ProspectPool struct {
	Active         bool    `json:"active"`
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	ProspectStages []Stage `json:"prospect_stages"`
}

type RejectionReason struct {
	Id   int        `json:"id"`
	Name string     `json:"name"`
	Type TypeIdName `json:"type"`
}

type ScheduledInterview struct {
	ApplicationId        int                    `json:"application_id"`
	End                  ScheduledInterviewDate `json:"end"`
	ExternalEventId      string                 `json:"external_event_id"`
	Id                   int                    `json:"id"`
	Interviewers         []Interviewer          `json:"interviewers"`
	Location             string                 `json:"location"`
	Organizer            User                   `json:"organizer"`
	Start                ScheduledInterviewDate `json:"start"`
	Status               string                 `json:"status"`
	VideoConferencingUrl string                 `json:"video_conferencing_url"`
}

type ScheduledInterviewDate struct {
	Date     string `json:"date,omitempty"`
	DateTime string `json:"date_time,omitempty"`
}

type School TypeIdNamePriority

type Scorecard struct {
	ApplicationId         int                  `json:"application_id"`
	Attributes            []ScorecardAttribute `json:"attributes"`
	CandidateId           int                  `json:"candidate_id"`
	CreatedAt             string               `json:"created_at"`
	Id                    int                  `json:"id"`
	Interview             string               `json:"interview"`
	InterviewStep         Stage                `json:"interview_step"`
	Interviewer           int                  `json:"interviewer"`
	OverallRecommendation string               `json:"overall_recommendation"`
	Questions             []ScorecardQuestion  `json:"questions"`
	Ratings               map[string][]string  `json:"ratings"`
	SubmittedAt           string               `json:"submitted_at"`
	SubmittedBy           User                 `json:"submitted_by"`
	UpdatedAt             string               `json:"updated_at"`
}

type ScorecardAttribute struct {
	Name   string `json:"name"`
	Note   string `json:"note"`
	Rating string `json:"rating"`
	Type   string `json:"type"`
}

type ScorecardQuestion struct {
	Answer   string `json:"answer"`
	Id       int    `json:"id"`
	Question string `json:"question"`
}

type Source struct {
	TypeIdName
	PublicName string     `json:"public_name,omitempty"`
	Type       TypeIdName `json:"type,omitempty"`
}

type Stage TypeIdName

type TrackingLink struct {
	CreatedAt       string   `json:"created_at"`
	CreditedTo      User     `json:"credited_to"`
	Id              int      `json:"id"`
	JobBoard        JobBoard `json:"job_board"`
	JobId           int      `json:"job_id"`
	JobPostId       int      `json:"job_post_id"`
	RelatedPostId   int      `json:"related_post_id"`
	RelatedPostType string   `json:"related_post_type"`
	Source          Source   `json:"source"`
	Token           string   `json:"token"`
	UpdatedAt       string   `json:"updated_at"`
}

type Translation struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}

type TypeIdNamePriority struct {
	TypeIdName
	Priority int `json:"priority"`
}

type TypeIdName struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type TypeTypeValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type User struct {
	CreatedAt          string   `json:"created_at"`
	Disabled           bool     `json:"disabled"`
	Emails             []string `json:"emails"`
	EmployeeId         string   `json:"employee_id"`
	FirstName          string   `json:"first_name"`
	Id                 int      `json:"id"`
	LastName           string   `json:"last_name"`
	LinkedCandidateIds []int    `json:"linked_candidate_ids"`
	Name               string   `json:"name"`
	PrimaryEmail       string   `json:"primary_email_address"`
	SiteAdmin          bool     `json:"site_admin"`
	UpdatedAt          string   `json:"updated_at"`
}

type UserCreateInfo struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	SendEmail bool   `json:"send_email_invite,omitempty"`
}

type UserEmailUpdateInfo struct {
	Email            string `json:"email"`
	SendVerification bool   `json:"send_verification"`
}

type UserPermission struct {
	Id         int `json:"id"`
	JobId      int `json:"job_id"`
	UserRoleId int `json:"user_role_id"`
}

type UserRole struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type UserUpdateInfo struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

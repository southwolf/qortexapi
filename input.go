package qortexapi

import "time"

type EntryInput struct {
	Id string

	AutoGenerateId bool

	// "post": normal entry
	// "task": acknowledgement or To-Do (IsAcknowledgement == true ,IsToGroup == "2")
	// "comment": entry's comment
	EType string

	Title   string
	Content string
	GroupId string

	IsToGroup   string // “0”:Notify People ,“1”:Notify Group
	ToUserIds   string // notify users  seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	TodoUserIds string // Todo users seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	// MentionedUserIds string // @users seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	UploadGroupId string // upload file's groupid, when in my feed UploadGroupId =""
	ContentType   string // "" ,"html","markdown", when "" will use user's setting

	IsAcknowledgement bool   // if IsAcknowledgement == true, get acknowledgement from notified people(ToUserIds).
	OnlyVisibleTo     bool   // if OnlyVisibleTo == true, only notified people can see the entry
	IsToDo            bool   // IsToDo == true, will create todo for entry.
	TaskDue           string // if AddToDo == true and want to set a deadline. format:20130507
	TodoStatus        int    // set it in group setting
	Priority          int    // Now :0 ,Soon :1, Someday:2
	Label             int    // set it in group setting
	EstimateTime      string // task's time Estimate

	RootId string // required if etype == "comment"

	NewVersion   bool   // if NewVersion == true will create new version.
	OldGroupId   string // when update entry required
	LastUpdateAt string // when update entry required

	KnowledgeBase bool // if KnowledgeBase == true, this entry is KnowledgeBase.
	AnyoneCanEdit bool // if AnyoneCanEdit == true, anyone can edit this entry.
	Presentation  bool // if Presentation == true, this entry is a Presentation

	IsFromEmail bool // IsFromEmail == true, entry create through email.

	IsPublished bool   // if IsPublished == true, this entry is a public blog.
	Slug        string // required when publish blog
	Share       bool   // if Share == true, Slug is not reqired
	Email       string // Blog Comment required
	Name        string // Blog Comment required

	InlineHelp       bool
	LinkTitle        string //for qortex support knowledge base
	BaseOnEntryId    string // when share chat,BaseOnEntryId = chat entry id
	PublishedToUsers bool

	LocaleName   string
	LanguageCode string //CLD language code

	// For Creating To-Dos From Comment
	BasedPostId        string
	BasedPostLangCode  string
	GroupIdOfBasedPost string
	SelectionTextInfo  SelectionTextInfo

	PriorityWeight float64

	IsEvent        bool
	IsInviteGroup  bool
	InvitedUserIds string //seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	StartAt        string //format 20130507
	StartAtH       string //00-23
	StartAtM       string //00-59
	EndAt          string //format 20130507
	EndAtH         string //00-23
	EndAtM         string //00-59
	Resource       []string
}

type DraftInput struct {
	Id        string
	GroupId   string
	Title     string
	Content   string
	ToUserIds string
	// UserId         string
	Etype     string
	IsToGroup string
	// OrganizationId string
	ContentType string // "" ,"html","markdown", when "" will use user's setting
	IsTaskTodo  bool
	IsTaskAck   bool
	TodoUserIds string
}

type SelectionTextInfo struct {
	StartText        string
	StartOccurrences int
	EndText          string
	EndOccurrences   int
}

const (
	BT_TO_ALL_ADMINS         = "boradcast_type_to_all_admins"
	BT_TO_ALL_USERS          = "boradcast_type_to_all_users"
	BT_TO_SOME_ORGANIZATIONS = "boradcast_type_to_some_organizations"
)

type BroadcastInput struct {
	Id            string
	Title         string
	Content       string
	ToOrgIds      []string
	BroadcastType string
	RootId        string

	LocaleName string
}

// TODO: Explaination needed.
type QortexSupportInput struct {
	Id                   string
	Title                string
	Content              string
	ToOrgIds             []string
	RootId               string
	Audiance             string
	PublishQortexSupport bool
	KnowledgeBase        bool   //for qortex support knowledge base
	InlineHelp           bool   //for qortex support knowledge base
	PublishedToUsers     bool   //for qortex support knowledge base
	LinkTitle            string //for qortex support knowledge base
	ContentType          string // "" ,"html","markdown", when "" will use user's setting

	LanguageCode        string
	IsAddingTranslation bool
}

type GroupInput struct {
	Id             string
	Name           string
	Description    string
	LogoURL        string
	IconName       string
	Slug           string
	IsPrivate      bool
	IsShared       bool
	GroupOwners    []string
	InvitedOrgIds  []string
	AutoGenSlug    bool
	CollectionName string
	CollectionId   string
}

type OrgPrivilegesInput struct {
	AllowUsersCreateGroups   bool
	AllowUsersInvitePeople   bool
	RestrictSubscriptionMail bool
	Domains                  []string
}

type OrganizationInput struct {
	Id                       string
	Name                     string
	Summary                  string
	Address                  string
	Phone                    string
	Website                  string
	Country                  string
	Size                     string
	Domains                  []string
	QortexURL                string
	LogoURL                  string
	SharingToken             string
	ContactWay               string
	NeedDemo                 bool
	RestrictSubscriptionMail bool
	AnyoneCanJoin            bool
	LanguageCodes            []string
	FreeOrPro                string
}

// Like or Unlike an entry action input
type LikeInput struct {
	EntryId string
	GroupId string
	Like    string // "0" for Unlike, "1" for Like
}

type PreferencesInput struct {
	Timezone                 string
	FirstDayOfWeek           string
	TimezoneOffset           string
	PreferFullName           string
	EnterForNewLine          string
	AsideGroupsCollapse      string
	AsideOtherGroupsCollapse string
	ShowMarkUnreadThreshold  string
	AdminModeOn              string
	PreferMarkdown           string
	AutoFollowPublicGroup    string
	EnableHTML5Notification  string
	UserLocationCityName     string
}

type NewsletterInput struct {
	Email string
}

type ShareChatInput struct {
	Title         string
	Content       string
	BasedConvId   string
	BaseOnEntryId string
	GroupId       string
}

type ContactInput struct {
	Name        string
	FirstName   string
	LastName    string
	CompanyName string
	CompanySize string
	Email       string
	Phone       string
	Country     string
	City        string
	HelpContent string
	Fake        bool // always false
	IsAgreed    bool
}

type UserProfileInput struct {
	Summary       string
	Title         string
	Department    string
	Location      string
	Expertise     string
	Interests     string
	BirthMonth    string
	BirthDay      string
	WorkPhone     string
	Mobile        string
	Twitter       string
	Skype         string
	Facebook      string
	OtherWebsites []string
}

// TODO: mail-updates: remove it
type MailUpdatesInput struct {
	IndividualIsOn    bool
	SendLag           int
	AckRequest        bool
	AckConfirmation   bool
	Todo              bool
	TodoConfirmation  bool
	SystemMessage     bool
	EntryNotification bool
	Like              bool
	SendTimeIsOn      bool
	Mon               bool
	Tue               bool
	Wed               bool
	Thu               bool
	Fri               bool
	Sat               bool
	Sun               bool
	SendHoursIsOn     bool
	StartAt           int
	EndAt             int
	DailyIsOn         bool
}

type NotificationPreferenceInput struct {
	Expecting       bool
	SendInterval    int
	SendLag         int
	IsMobileEnabled bool
}

type TaskInput struct {
	TaskId                    string
	GroupId                   string
	AssigneeId                string
	TodoStatus                int
	Label                     int
	EstimateTime              string
	SpentTime                 string
	IsClaiming                bool
	ToUserIds                 string
	TimetrackingHistoryUpdate string
	TaskDue                   string // format:20130507
	EntryId                   string
	PriorityWeight            float64
}

type TaskPwMap struct {
	Id             string `bson:"-"`
	PriorityWeight float64
}

type ToDoMarkerInput struct {
	Id             string
	Label          string
	Date           time.Time `bson:",omitempty"`
	PriorityWeight float64
}

type KnowledgeOverviewInput struct {
	GroupId      string
	Title        string
	Content      string
	IsHidden     bool
	LanguageCode string
}

type SearchInput struct {
	// Value: myfeed, group, chats, user, tasks
	// In myfeed, entries, conversations, links, and attachemnts will appear in your search result
	// In chats, only conversations
	// In tasks, only tasks
	// About group and user, actually, they are mainly for pc stuff, app could ignore them
	Scope string

	GroupIds []string
	UserIds  []string

	// Sorting stuff by relevance: 	rel
	// Sorting stuff by date: 		rec
	SortBy string

	Page     int // Start from: 1
	Keywords string
}

type ZapierSubscribeInput struct {
	SubscriptionUrl string
	Event           string
	TargetUrl       string
}

type AdvancedToDoSettingsInput struct {
	Enabled            bool
	EnableTimeEstimate bool
	EnableTimeTracking bool
	TimeUnit           int
	ProjectManagerId   string

	// NOTE: for new TagIndex, its Index value must be -1
	Labels             []*TagIndex
	NotYetOpenStatuses []*TagIndex
	OpenStatuses       []*TagIndex
	ClosedStatuses     []*TagIndex
}

type NewChatInput struct {
	MessageIds []string
}

// userId string, entryType string, before string, limit int
type UserEntriesInput struct {
	OrgId         string
	UserId        string
	BeforeTime    string
	Limit         int
	IsShowPrivate bool
}

type ChangeEmailInput struct {
	OldEmail        string
	NewEmail        string
	ConfirmToken    string
	SharingToken    string
	InvitationToken string
}

type InviteInput struct {
	Emails           []string
	ToFollowGroupIds []string
	Message          string
	IgnoreError      bool
}

type AccountInput struct {
	Id              string
	OrgId           string
	Email           string
	ConfirmToken    string
	SharingToken    string
	InvitationToken string
	InvitedOrgName  string // This is not for input, for displaying

	FirstName     string // English name
	LastName      string
	FirstNameCn   string // Chinese name
	LastNameCn    string
	FirstNameJp   string // Japanese name for display
	LastNameJp    string
	FirstNameJpKa string // Japanese Katakana name for ordering
	LastNameJpKa  string

	Password        string
	ConfirmPassword string
	AvatarURL       string
	Title           string
	Department      string
	IsAgreed        bool

	LocaleName    string   // Login needs this
	LanguageCodes []string // Invitation signup needs this
}

// KOBELD: Duplicated. Can be replaced by the AccountInput
type MemberAccountInput struct {
	FirstName     string // English name
	LastName      string
	FirstNameCn   string // Chinese name
	LastNameCn    string
	FirstNameJp   string // Japanese name for display
	LastNameJp    string
	FirstNameJpKa string // Japanese Katakana name for ordering
	LastNameJpKa  string

	AvatarURL string
}

type EventInput struct {
	EventId        string
	GroupId        string
	IsInviteGroup  bool
	InvitedUserIds string //seperate with "," for example: "1234,4567" means []string{"1234", "5678"}
	StartAt        string //format 20130507
	StartAtH       string //00-23
	StartAtM       string //00-59
	EndAt          string //format 20130507
	EndAtH         string //00-23
	EndAtM         string //00-59
	Resource       []string
}

type AccessInput struct {
	Name       string // TODO: Jul 23. Should be removed
	Furigana   string // TODO: Jul 23. Should be removed
	FirstName  string
	LastName   string
	Password   string
	Email      string
	Phone      string
	OrgName    string
	LocaleName string
	Referrer   string
	Country    string
	DialCode   string
}

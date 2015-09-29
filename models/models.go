
type ActivityId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type ActivityData struct {
    Type string `json:"type"`
    Module string `json:"module"`
    ObjectModel string `json:"objectModel"`
    ObjectId string `json:"objectId"`
}

type ActivityPost struct {
  ActivityData `json:"data"`
  Trace   `json:"trace"`
}

type Activity struct {
    ActivityId
    ActivityData `json:"data"`
    Trace   `json:"trace"`
}

type CommentId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type CommentData struct {
    Message string `json:"message"`
    ObjectModel string `json:"objectModel"`
    ObjectId uint `json:"objectId"`
    SpaceId uint `json:"spaceId"`
}

type CommentPost struct {
  CommentData `json:"data"`
  Trace   `json:"trace"`
}

type Comment struct {
    CommentId
    CommentData `json:"data"`
    Trace   `json:"trace"`
}

type ContentId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type ContentData struct {
    Guid string `json:"guid"`
    ObjectModel string `json:"objectModel"`
    ObjectId uint `json:"objectId"`
    Visibility uint `json:"visibility"`
    Sticked uint `json:"sticked"`
    Archived string `json:"archived"`
    SpaceId uint `json:"spaceId"`
    UserId uint `json:"userId"`
}

type ContentPost struct {
  ContentData `json:"data"`
  Trace   `json:"trace"`
}

type Content struct {
    ContentId
    ContentData `json:"data"`
    Trace   `json:"trace"`
}

type FileId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type FileData struct {
    Guid string `json:"guid"`
    ObjectModel string `json:"objectModel"`
    ObjectId string `json:"objectId"`
    FileName string `json:"fileName"`
    Title string `json:"title"`
    MimeType string `json:"mimeType"`
    Size string `json:"size"`
}

type FilePost struct {
  FileData `json:"data"`
  Trace   `json:"trace"`
}

type File struct {
    FileId
    FileData `json:"data"`
    Trace   `json:"trace"`
}

type GroupId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type GroupData struct {
    SpaceId uint `json:"spaceId"`
    Name string `json:"name"`
    Description string `json:"description"`
    LdapDn string `json:"ldapDn"`
    CanCreatePublicSpaces uint `json:"canCreatePublicSpaces"`
    CanCreatePrivateSpaces uint `json:"canCreatePrivateSpaces"`
}

type GroupPost struct {
  GroupData `json:"data"`
  Trace   `json:"trace"`
}

type Group struct {
    GroupId
    GroupData `json:"data"`
    Trace   `json:"trace"`
}

type GroupAdminId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type GroupAdminData struct {
    UserId uint `json:"userId"`
    GroupId uint `json:"groupId"`
}

type GroupAdminPost struct {
  GroupAdminData `json:"data"`
  Trace   `json:"trace"`
}

type GroupAdmin struct {
    GroupAdminId
    GroupAdminData `json:"data"`
    Trace   `json:"trace"`
}

type LikeId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type LikeData struct {
    TargetUserId uint `json:"targetUserId"`
    ObjectModel string `json:"objectModel"`
    ObjectId uint `json:"objectId"`
}

type LikePost struct {
  LikeData `json:"data"`
  Trace   `json:"trace"`
}

type Like struct {
    LikeId
    LikeData `json:"data"`
    Trace   `json:"trace"`
}

type LoggingId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type LoggingData struct {
    Level string `json:"level"`
    Category string `json:"category"`
    Logtime uint `json:"logtime"`
    Message string `json:"message"`
}

type LoggingPost struct {
  LoggingData `json:"data"`
}

type Logging struct {
    LoggingId
    LoggingData `json:"data"`
}

type MigrationId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type MigrationData struct {
    Version string `json:"version"`
    ApplyTime uint `json:"applyTime"`
    Module string `json:"module"`
}

type MigrationPost struct {
  MigrationData `json:"data"`
}

type Migration struct {
    MigrationId
    MigrationData `json:"data"`
}

type ModuleEnabledId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type ModuleEnabledData struct {
    ModuleId string `json:"moduleId"`
}

type ModuleEnabledPost struct {
  ModuleEnabledData `json:"data"`
}

type ModuleEnabled struct {
    ModuleEnabledId
    ModuleEnabledData `json:"data"`
}

type NotificationId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type NotificationData struct {
    Class string `json:"class"`
    UserId uint `json:"userId"`
    Seen uint `json:"seen"`
    SourceObjectModel string `json:"sourceObjectModel"`
    SourceObjectId uint `json:"sourceObjectId"`
    TargetObjectModel string `json:"targetObjectModel"`
    TargetObjectId uint `json:"targetObjectId"`
    SpaceId uint `json:"spaceId"`
    Emailed uint `json:"emailed"`
    DesktopNotified uint `json:"desktopNotified"`
}

type NotificationPost struct {
  NotificationData `json:"data"`
  Trace   `json:"trace"`
}

type Notification struct {
    NotificationId
    NotificationData `json:"data"`
    Trace   `json:"trace"`
}

type PostId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type PostData struct {
    Message2trash string `json:"message2trash"`
    Message string `json:"message"`
    Url string `json:"url"`
}

type PostPost struct {
  PostData `json:"data"`
  Trace   `json:"trace"`
}

type Post struct {
    PostId
    PostData `json:"data"`
    Trace   `json:"trace"`
}

type ProfileId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type ProfileData struct {
    UserId uint `json:"userId"`
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
    Title string `json:"title"`
    Gender string `json:"gender"`
    Street string `json:"street"`
    Zip string `json:"zip"`
    City string `json:"city"`
    Country string `json:"country"`
    State string `json:"state"`
    BirthdayHideYear uint `json:"birthdayHideYear"`
    Birthday time.Time `json:"birthday"`
    About string `json:"about"`
    PhonePrivate string `json:"phonePrivate"`
    PhoneWork string `json:"phoneWork"`
    Mobile string `json:"mobile"`
    Fax string `json:"fax"`
    ImSkype string `json:"imSkype"`
    ImMsn string `json:"imMsn"`
    ImIcq uint `json:"imIcq"`
    ImXmpp string `json:"imXmpp"`
    Url string `json:"url"`
    UrlFacebook string `json:"urlFacebook"`
    UrlLinkedin string `json:"urlLinkedin"`
    UrlXing string `json:"urlXing"`
    UrlYoutube string `json:"urlYoutube"`
    UrlVimeo string `json:"urlVimeo"`
    UrlFlickr string `json:"urlFlickr"`
    UrlMyspace string `json:"urlMyspace"`
    UrlGoogleplus string `json:"urlGoogleplus"`
    UrlTwitter string `json:"urlTwitter"`
}

type ProfilePost struct {
  ProfileData `json:"data"`
}

type Profile struct {
    ProfileId
    ProfileData `json:"data"`
}

type ProfileFieldId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type ProfileFieldData struct {
    ProfileFieldCategoryId uint `json:"profileFieldCategoryId"`
    ModuleId string `json:"moduleId"`
    FieldTypeClass string `json:"fieldTypeClass"`
    FieldTypeConfig string `json:"fieldTypeConfig"`
    InternalName string `json:"internalName"`
    Title string `json:"title"`
    Description string `json:"description"`
    SortOrder uint `json:"sortOrder"`
    Required uint `json:"required"`
    ShowAtRegistration uint `json:"showAtRegistration"`
    Editable uint `json:"editable"`
    Visible uint `json:"visible"`
    LdapAttribute string `json:"ldapAttribute"`
    TranslationCategory string `json:"translationCategory"`
    IsSystem uint `json:"isSystem"`
}

type ProfileFieldPost struct {
  ProfileFieldData `json:"data"`
  Trace   `json:"trace"`
}

type ProfileField struct {
    ProfileFieldId
    ProfileFieldData `json:"data"`
    Trace   `json:"trace"`
}

type ProfileFieldCategoryId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type ProfileFieldCategoryData struct {
    Title string `json:"title"`
    Description string `json:"description"`
    SortOrder uint `json:"sortOrder"`
    ModuleId uint `json:"moduleId"`
    Visibility uint `json:"visibility"`
    TranslationCategory string `json:"translationCategory"`
    IsSystem uint `json:"isSystem"`
}

type ProfileFieldCategoryPost struct {
  ProfileFieldCategoryData `json:"data"`
  Trace   `json:"trace"`
}

type ProfileFieldCategory struct {
    ProfileFieldCategoryId
    ProfileFieldCategoryData `json:"data"`
    Trace   `json:"trace"`
}

type SettingId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type SettingData struct {
    Name string `json:"name"`
    Value string `json:"value"`
    ValueText string `json:"valueText"`
    ModuleId string `json:"moduleId"`
}

type SettingPost struct {
  SettingData `json:"data"`
  Trace   `json:"trace"`
}

type Setting struct {
    SettingId
    SettingData `json:"data"`
    Trace   `json:"trace"`
}

type SpaceId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type SpaceData struct {
    Guid string `json:"guid"`
    WallId uint `json:"wallId"`
    Name string `json:"name"`
    Description string `json:"description"`
    Website string `json:"website"`
    JoinPolicy uint `json:"joinPolicy"`
    Visibility uint `json:"visibility"`
    Status uint `json:"status"`
    Tags string `json:"tags"`
    LdapDn string `json:"ldapDn"`
    AutoAddNewMembers uint `json:"autoAddNewMembers"`
}

type SpacePost struct {
  SpaceData `json:"data"`
  Trace   `json:"trace"`
}

type Space struct {
    SpaceId
    SpaceData `json:"data"`
    Trace   `json:"trace"`
}

type SpaceMembershipId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type SpaceMembershipData struct {
    SpaceId uint `json:"spaceId"`
    UserId uint `json:"userId"`
    OriginatorUserId string `json:"originatorUserId"`
    Status uint `json:"status"`
    RequestMessage string `json:"requestMessage"`
    LastVisit time.Time `json:"lastVisit"`
    InviteRole uint `json:"inviteRole"`
    AdminRole uint `json:"adminRole"`
    ShareRole uint `json:"shareRole"`
}

type SpaceMembershipPost struct {
  SpaceMembershipData `json:"data"`
  Trace   `json:"trace"`
}

type SpaceMembership struct {
    SpaceMembershipId
    SpaceMembershipData `json:"data"`
    Trace   `json:"trace"`
}

type SpaceModuleId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type SpaceModuleData struct {
    ModuleId string `json:"moduleId"`
    SpaceId uint `json:"spaceId"`
    State uint `json:"state"`
}

type SpaceModulePost struct {
  SpaceModuleData `json:"data"`
}

type SpaceModule struct {
    SpaceModuleId
    SpaceModuleData `json:"data"`
}

type SpaceSettingId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type SpaceSettingData struct {
    SpaceId uint `json:"spaceId"`
    ModuleId string `json:"moduleId"`
    Name string `json:"name"`
    Value string `json:"value"`
}

type SpaceSettingPost struct {
  SpaceSettingData `json:"data"`
  Trace   `json:"trace"`
}

type SpaceSetting struct {
    SpaceSettingId
    SpaceSettingData `json:"data"`
    Trace   `json:"trace"`
}

type UrlOembedId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UrlOembedData struct {
    Url string `json:"url"`
    Preview string `json:"preview"`
}

type UrlOembedPost struct {
  UrlOembedData `json:"data"`
}

type UrlOembed struct {
    UrlOembedId
    UrlOembedData `json:"data"`
}

type UserId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UserData struct {
    Guid string `json:"guid"`
    WallId uint `json:"wallId"`
    GroupId uint `json:"groupId"`
    Status uint `json:"status"`
    SuperAdmin uint `json:"superAdmin"`
    Username string `json:"username"`
    Email string `json:"email"`
    AuthMode string `json:"authMode"`
    Tags string `json:"tags"`
    Language string `json:"language"`
    LastActivityEmail time.Time `json:"lastActivityEmail"`
    LastLogin time.Time `json:"lastLogin"`
    Visibility uint `json:"visibility"`
}

type UserPost struct {
  UserData `json:"data"`
  Trace   `json:"trace"`
}

type User struct {
    UserId
    UserData `json:"data"`
    Trace   `json:"trace"`
}

type UserFollowId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UserFollowData struct {
    ObjectModel string `json:"objectModel"`
    ObjectId uint `json:"objectId"`
    UserId uint `json:"userId"`
    SendNotifications uint `json:"sendNotifications"`
}

type UserFollowPost struct {
  UserFollowData `json:"data"`
}

type UserFollow struct {
    UserFollowId
    UserFollowData `json:"data"`
}

type UserHttpSessionId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UserHttpSessionData struct {
    Expire uint `json:"expire"`
    UserId uint `json:"userId"`
    Data string `json:"data"`
}

type UserHttpSessionPost struct {
  UserHttpSessionData `json:"data"`
}

type UserHttpSession struct {
    UserHttpSessionId
    UserHttpSessionData `json:"data"`
}

type UserInviteId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UserInviteData struct {
    UserOriginatorId uint `json:"userOriginatorId"`
    SpaceInviteId uint `json:"spaceInviteId"`
    Email string `json:"email"`
    Source string `json:"source"`
    Token string `json:"token"`
    Language string `json:"language"`
}

type UserInvitePost struct {
  UserInviteData `json:"data"`
  Trace   `json:"trace"`
}

type UserInvite struct {
    UserInviteId
    UserInviteData `json:"data"`
    Trace   `json:"trace"`
}

type UserMentioningId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UserMentioningData struct {
    ObjectModel string `json:"objectModel"`
    ObjectId uint `json:"objectId"`
    UserId uint `json:"userId"`
}

type UserMentioningPost struct {
  UserMentioningData `json:"data"`
}

type UserMentioning struct {
    UserMentioningId
    UserMentioningData `json:"data"`
}

type UserModuleId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UserModuleData struct {
    ModuleId string `json:"moduleId"`
    UserId uint `json:"userId"`
    State uint `json:"state"`
}

type UserModulePost struct {
  UserModuleData `json:"data"`
}

type UserModule struct {
    UserModuleId
    UserModuleData `json:"data"`
}

type UserPasswordId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UserPasswordData struct {
    UserId uint `json:"userId"`
    Algorithm string `json:"algorithm"`
    Password string `json:"password"`
    Salt string `json:"salt"`
}

type UserPasswordPost struct {
  UserPasswordData `json:"data"`
  Trace   `json:"trace"`
}

type UserPassword struct {
    UserPasswordId
    UserPasswordData `json:"data"`
    Trace   `json:"trace"`
}

type UserSettingId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type UserSettingData struct {
    UserId uint `json:"userId"`
    ModuleId string `json:"moduleId"`
    Name string `json:"name"`
    Value string `json:"value"`
}

type UserSettingPost struct {
  UserSettingData `json:"data"`
  Trace   `json:"trace"`
}

type UserSetting struct {
    UserSettingId
    UserSettingData `json:"data"`
    Trace   `json:"trace"`
}

type WallId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type WallData struct {
    ObjectModel string `json:"objectModel"`
    ObjectId uint `json:"objectId"`
}

type WallPost struct {
  WallData `json:"data"`
  Trace   `json:"trace"`
}

type Wall struct {
    WallId
    WallData `json:"data"`
    Trace   `json:"trace"`
}

type WallEntryId struct {
    Id uint64 `json:"id" sql:"AUTO_INCREMENT"`
}

type WallEntryData struct {
    WallId uint `json:"wallId"`
    ContentId uint `json:"contentId"`
}

type WallEntryPost struct {
  WallEntryData `json:"data"`
  Trace   `json:"trace"`
}

type WallEntry struct {
    WallEntryId
    WallEntryData `json:"data"`
    Trace   `json:"trace"`
}

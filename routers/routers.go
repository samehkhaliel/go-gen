
func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})
}

func InstallCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comment
	router.GET("/comment", func(c *gin.Context) {
		var comment []model.Comment
		db.Find(&comment)
		c.JSON(http.StatusOK, comment)
	})
}

func InstallContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all content
	router.GET("/content", func(c *gin.Context) {
		var content []model.Content
		db.Find(&content)
		c.JSON(http.StatusOK, content)
	})
}

func InstallFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all file
	router.GET("/file", func(c *gin.Context) {
		var file []model.File
		db.Find(&file)
		c.JSON(http.StatusOK, file)
	})
}

func InstallGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all group
	router.GET("/group", func(c *gin.Context) {
		var group []model.Group
		db.Find(&group)
		c.JSON(http.StatusOK, group)
	})
}

func InstallGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admin
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admin []model.GroupAdmin
		db.Find(&group_admin)
		c.JSON(http.StatusOK, group_admin)
	})
}

func InstallLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all like
	router.GET("/like", func(c *gin.Context) {
		var like []model.Like
		db.Find(&like)
		c.JSON(http.StatusOK, like)
	})
}

func InstallLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all logging
	router.GET("/logging", func(c *gin.Context) {
		var logging []model.Logging
		db.Find(&logging)
		c.JSON(http.StatusOK, logging)
	})
}

func InstallMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migration
	router.GET("/migration", func(c *gin.Context) {
		var migration []model.Migration
		db.Find(&migration)
		c.JSON(http.StatusOK, migration)
	})
}

func InstallModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enabled
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enabled []model.ModuleEnabled
		db.Find(&module_enabled)
		c.JSON(http.StatusOK, module_enabled)
	})
}

func InstallNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notification
	router.GET("/notification", func(c *gin.Context) {
		var notification []model.Notification
		db.Find(&notification)
		c.JSON(http.StatusOK, notification)
	})
}

func InstallPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all post
	router.GET("/post", func(c *gin.Context) {
		var post []model.Post
		db.Find(&post)
		c.JSON(http.StatusOK, post)
	})
}

func InstallProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profile
	router.GET("/profile", func(c *gin.Context) {
		var profile []model.Profile
		db.Find(&profile)
		c.JSON(http.StatusOK, profile)
	})
}

func InstallProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_field
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_field []model.ProfileField
		db.Find(&profile_field)
		c.JSON(http.StatusOK, profile_field)
	})
}

func InstallProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_category
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_category []model.ProfileFieldCategory
		db.Find(&profile_field_category)
		c.JSON(http.StatusOK, profile_field_category)
	})
}

func InstallSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all setting
	router.GET("/setting", func(c *gin.Context) {
		var setting []model.Setting
		db.Find(&setting)
		c.JSON(http.StatusOK, setting)
	})
}

func InstallSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all space
	router.GET("/space", func(c *gin.Context) {
		var space []model.Space
		db.Find(&space)
		c.JSON(http.StatusOK, space)
	})
}

func InstallSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_membership
	router.GET("/space_membership", func(c *gin.Context) {
		var space_membership []model.SpaceMembership
		db.Find(&space_membership)
		c.JSON(http.StatusOK, space_membership)
	})
}

func InstallSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_module
	router.GET("/space_module", func(c *gin.Context) {
		var space_module []model.SpaceModule
		db.Find(&space_module)
		c.JSON(http.StatusOK, space_module)
	})
}

func InstallSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_setting
	router.GET("/space_setting", func(c *gin.Context) {
		var space_setting []model.SpaceSetting
		db.Find(&space_setting)
		c.JSON(http.StatusOK, space_setting)
	})
}

func InstallUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembed
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembed []model.UrlOembed
		db.Find(&url_oembed)
		c.JSON(http.StatusOK, url_oembed)
	})
}

func InstallUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all user
	router.GET("/user", func(c *gin.Context) {
		var user []model.User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})
}

func InstallUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follow
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follow []model.UserFollow
		db.Find(&user_follow)
		c.JSON(http.StatusOK, user_follow)
	})
}

func InstallUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_session
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_session []model.UserHttpSession
		db.Find(&user_http_session)
		c.JSON(http.StatusOK, user_http_session)
	})
}

func InstallUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invite
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invite []model.UserInvite
		db.Find(&user_invite)
		c.JSON(http.StatusOK, user_invite)
	})
}

func InstallUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentioning
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentioning []model.UserMentioning
		db.Find(&user_mentioning)
		c.JSON(http.StatusOK, user_mentioning)
	})
}

func InstallUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_module
	router.GET("/user_module", func(c *gin.Context) {
		var user_module []model.UserModule
		db.Find(&user_module)
		c.JSON(http.StatusOK, user_module)
	})
}

func InstallUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_password
	router.GET("/user_password", func(c *gin.Context) {
		var user_password []model.UserPassword
		db.Find(&user_password)
		c.JSON(http.StatusOK, user_password)
	})
}

func InstallUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_setting
	router.GET("/user_setting", func(c *gin.Context) {
		var user_setting []model.UserSetting
		db.Find(&user_setting)
		c.JSON(http.StatusOK, user_setting)
	})
}

func InstallWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all wall
	router.GET("/wall", func(c *gin.Context) {
		var wall []model.Wall
		db.Find(&wall)
		c.JSON(http.StatusOK, wall)
	})
}

func InstallWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entry
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entry []model.WallEntry
		db.Find(&wall_entry)
		c.JSON(http.StatusOK, wall_entry)
	})
}

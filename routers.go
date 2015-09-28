func installActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activitys
	router.GET("/activity", func(c *gin.Context) {
		var activitys []model.Activity
		db.Find(&activitys)
		c.JSON(http.StatusOK, activitys)
	})

	// GET /activity/:guid
	// GET activity by guid
	router.GET("/activity/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var activity model.Activity
		query := &model.Activity{
			ActivityData: model.ActivityData{Guid: guid},
		}

		db.Where(&query).First(&activity)

		c.JSON(http.StatusOK, activity)
	})
}
func installCommentRouter(db gorm.DB, router *gin.Engine) {
	// GET /comment
	// GET all comments
	router.GET("/comment", func(c *gin.Context) {
		var comments []model.Comment
		db.Find(&comments)
		c.JSON(http.StatusOK, comments)
	})

	// GET /comment/:guid
	// GET comment by guid
	router.GET("/comment/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var comment model.Comment
		query := &model.Comment{
			CommentData: model.CommentData{Guid: guid},
		}

		db.Where(&query).First(&comment)

		c.JSON(http.StatusOK, comment)
	})
}
func installContentRouter(db gorm.DB, router *gin.Engine) {
	// GET /content
	// GET all contents
	router.GET("/content", func(c *gin.Context) {
		var contents []model.Content
		db.Find(&contents)
		c.JSON(http.StatusOK, contents)
	})

	// GET /content/:guid
	// GET content by guid
	router.GET("/content/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var content model.Content
		query := &model.Content{
			ContentData: model.ContentData{Guid: guid},
		}

		db.Where(&query).First(&content)

		c.JSON(http.StatusOK, content)
	})
}
func installFileRouter(db gorm.DB, router *gin.Engine) {
	// GET /file
	// GET all files
	router.GET("/file", func(c *gin.Context) {
		var files []model.File
		db.Find(&files)
		c.JSON(http.StatusOK, files)
	})

	// GET /file/:guid
	// GET file by guid
	router.GET("/file/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var file model.File
		query := &model.File{
			FileData: model.FileData{Guid: guid},
		}

		db.Where(&query).First(&file)

		c.JSON(http.StatusOK, file)
	})
}
func installGroupRouter(db gorm.DB, router *gin.Engine) {
	// GET /group
	// GET all groups
	router.GET("/group", func(c *gin.Context) {
		var groups []model.Group
		db.Find(&groups)
		c.JSON(http.StatusOK, groups)
	})

	// GET /group/:guid
	// GET group by guid
	router.GET("/group/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var group model.Group
		query := &model.Group{
			GroupData: model.GroupData{Guid: guid},
		}

		db.Where(&query).First(&group)

		c.JSON(http.StatusOK, group)
	})
}
func installGroupAdminRouter(db gorm.DB, router *gin.Engine) {
	// GET /group_admin
	// GET all group_admins
	router.GET("/group_admin", func(c *gin.Context) {
		var group_admins []model.GroupAdmin
		db.Find(&group_admins)
		c.JSON(http.StatusOK, group_admins)
	})

	// GET /group_admin/:guid
	// GET group_admin by guid
	router.GET("/group_admin/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var group_admin model.GroupAdmin
		query := &model.GroupAdmin{
			GroupAdminData: model.GroupAdminData{Guid: guid},
		}

		db.Where(&query).First(&group_admin)

		c.JSON(http.StatusOK, group_admin)
	})
}
func installLikeRouter(db gorm.DB, router *gin.Engine) {
	// GET /like
	// GET all likes
	router.GET("/like", func(c *gin.Context) {
		var likes []model.Like
		db.Find(&likes)
		c.JSON(http.StatusOK, likes)
	})

	// GET /like/:guid
	// GET like by guid
	router.GET("/like/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var like model.Like
		query := &model.Like{
			LikeData: model.LikeData{Guid: guid},
		}

		db.Where(&query).First(&like)

		c.JSON(http.StatusOK, like)
	})
}
func installLoggingRouter(db gorm.DB, router *gin.Engine) {
	// GET /logging
	// GET all loggings
	router.GET("/logging", func(c *gin.Context) {
		var loggings []model.Logging
		db.Find(&loggings)
		c.JSON(http.StatusOK, loggings)
	})

	// GET /logging/:guid
	// GET logging by guid
	router.GET("/logging/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var logging model.Logging
		query := &model.Logging{
			LoggingData: model.LoggingData{Guid: guid},
		}

		db.Where(&query).First(&logging)

		c.JSON(http.StatusOK, logging)
	})
}
func installMigrationRouter(db gorm.DB, router *gin.Engine) {
	// GET /migration
	// GET all migrations
	router.GET("/migration", func(c *gin.Context) {
		var migrations []model.Migration
		db.Find(&migrations)
		c.JSON(http.StatusOK, migrations)
	})

	// GET /migration/:guid
	// GET migration by guid
	router.GET("/migration/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var migration model.Migration
		query := &model.Migration{
			MigrationData: model.MigrationData{Guid: guid},
		}

		db.Where(&query).First(&migration)

		c.JSON(http.StatusOK, migration)
	})
}
func installModuleEnabledRouter(db gorm.DB, router *gin.Engine) {
	// GET /module_enabled
	// GET all module_enableds
	router.GET("/module_enabled", func(c *gin.Context) {
		var module_enableds []model.ModuleEnabled
		db.Find(&module_enableds)
		c.JSON(http.StatusOK, module_enableds)
	})

	// GET /module_enabled/:guid
	// GET module_enabled by guid
	router.GET("/module_enabled/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var module_enabled model.ModuleEnabled
		query := &model.ModuleEnabled{
			ModuleEnabledData: model.ModuleEnabledData{Guid: guid},
		}

		db.Where(&query).First(&module_enabled)

		c.JSON(http.StatusOK, module_enabled)
	})
}
func installNotificationRouter(db gorm.DB, router *gin.Engine) {
	// GET /notification
	// GET all notifications
	router.GET("/notification", func(c *gin.Context) {
		var notifications []model.Notification
		db.Find(&notifications)
		c.JSON(http.StatusOK, notifications)
	})

	// GET /notification/:guid
	// GET notification by guid
	router.GET("/notification/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var notification model.Notification
		query := &model.Notification{
			NotificationData: model.NotificationData{Guid: guid},
		}

		db.Where(&query).First(&notification)

		c.JSON(http.StatusOK, notification)
	})
}
func installPostRouter(db gorm.DB, router *gin.Engine) {
	// GET /post
	// GET all posts
	router.GET("/post", func(c *gin.Context) {
		var posts []model.Post
		db.Find(&posts)
		c.JSON(http.StatusOK, posts)
	})

	// GET /post/:guid
	// GET post by guid
	router.GET("/post/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var post model.Post
		query := &model.Post{
			PostData: model.PostData{Guid: guid},
		}

		db.Where(&query).First(&post)

		c.JSON(http.StatusOK, post)
	})
}
func installProfileRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile
	// GET all profiles
	router.GET("/profile", func(c *gin.Context) {
		var profiles []model.Profile
		db.Find(&profiles)
		c.JSON(http.StatusOK, profiles)
	})

	// GET /profile/:guid
	// GET profile by guid
	router.GET("/profile/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var profile model.Profile
		query := &model.Profile{
			ProfileData: model.ProfileData{Guid: guid},
		}

		db.Where(&query).First(&profile)

		c.JSON(http.StatusOK, profile)
	})
}
func installProfileFieldRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field
	// GET all profile_fields
	router.GET("/profile_field", func(c *gin.Context) {
		var profile_fields []model.ProfileField
		db.Find(&profile_fields)
		c.JSON(http.StatusOK, profile_fields)
	})

	// GET /profile_field/:guid
	// GET profile_field by guid
	router.GET("/profile_field/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var profile_field model.ProfileField
		query := &model.ProfileField{
			ProfileFieldData: model.ProfileFieldData{Guid: guid},
		}

		db.Where(&query).First(&profile_field)

		c.JSON(http.StatusOK, profile_field)
	})
}
func installProfileFieldCategoryRouter(db gorm.DB, router *gin.Engine) {
	// GET /profile_field_category
	// GET all profile_field_categorys
	router.GET("/profile_field_category", func(c *gin.Context) {
		var profile_field_categorys []model.ProfileFieldCategory
		db.Find(&profile_field_categorys)
		c.JSON(http.StatusOK, profile_field_categorys)
	})

	// GET /profile_field_category/:guid
	// GET profile_field_category by guid
	router.GET("/profile_field_category/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var profile_field_category model.ProfileFieldCategory
		query := &model.ProfileFieldCategory{
			ProfileFieldCategoryData: model.ProfileFieldCategoryData{Guid: guid},
		}

		db.Where(&query).First(&profile_field_category)

		c.JSON(http.StatusOK, profile_field_category)
	})
}
func installSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /setting
	// GET all settings
	router.GET("/setting", func(c *gin.Context) {
		var settings []model.Setting
		db.Find(&settings)
		c.JSON(http.StatusOK, settings)
	})

	// GET /setting/:guid
	// GET setting by guid
	router.GET("/setting/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var setting model.Setting
		query := &model.Setting{
			SettingData: model.SettingData{Guid: guid},
		}

		db.Where(&query).First(&setting)

		c.JSON(http.StatusOK, setting)
	})
}
func installSpaceRouter(db gorm.DB, router *gin.Engine) {
	// GET /space
	// GET all spaces
	router.GET("/space", func(c *gin.Context) {
		var spaces []model.Space
		db.Find(&spaces)
		c.JSON(http.StatusOK, spaces)
	})

	// GET /space/:guid
	// GET space by guid
	router.GET("/space/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var space model.Space
		query := &model.Space{
			SpaceData: model.SpaceData{Guid: guid},
		}

		db.Where(&query).First(&space)

		c.JSON(http.StatusOK, space)
	})
}
func installSpaceMembershipRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_membership
	// GET all space_memberships
	router.GET("/space_membership", func(c *gin.Context) {
		var space_memberships []model.SpaceMembership
		db.Find(&space_memberships)
		c.JSON(http.StatusOK, space_memberships)
	})

	// GET /space_membership/:guid
	// GET space_membership by guid
	router.GET("/space_membership/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var space_membership model.SpaceMembership
		query := &model.SpaceMembership{
			SpaceMembershipData: model.SpaceMembershipData{Guid: guid},
		}

		db.Where(&query).First(&space_membership)

		c.JSON(http.StatusOK, space_membership)
	})
}
func installSpaceModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_module
	// GET all space_modules
	router.GET("/space_module", func(c *gin.Context) {
		var space_modules []model.SpaceModule
		db.Find(&space_modules)
		c.JSON(http.StatusOK, space_modules)
	})

	// GET /space_module/:guid
	// GET space_module by guid
	router.GET("/space_module/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var space_module model.SpaceModule
		query := &model.SpaceModule{
			SpaceModuleData: model.SpaceModuleData{Guid: guid},
		}

		db.Where(&query).First(&space_module)

		c.JSON(http.StatusOK, space_module)
	})
}
func installSpaceSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /space_setting
	// GET all space_settings
	router.GET("/space_setting", func(c *gin.Context) {
		var space_settings []model.SpaceSetting
		db.Find(&space_settings)
		c.JSON(http.StatusOK, space_settings)
	})

	// GET /space_setting/:guid
	// GET space_setting by guid
	router.GET("/space_setting/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var space_setting model.SpaceSetting
		query := &model.SpaceSetting{
			SpaceSettingData: model.SpaceSettingData{Guid: guid},
		}

		db.Where(&query).First(&space_setting)

		c.JSON(http.StatusOK, space_setting)
	})
}
func installUrlOembedRouter(db gorm.DB, router *gin.Engine) {
	// GET /url_oembed
	// GET all url_oembeds
	router.GET("/url_oembed", func(c *gin.Context) {
		var url_oembeds []model.UrlOembed
		db.Find(&url_oembeds)
		c.JSON(http.StatusOK, url_oembeds)
	})

	// GET /url_oembed/:guid
	// GET url_oembed by guid
	router.GET("/url_oembed/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var url_oembed model.UrlOembed
		query := &model.UrlOembed{
			UrlOembedData: model.UrlOembedData{Guid: guid},
		}

		db.Where(&query).First(&url_oembed)

		c.JSON(http.StatusOK, url_oembed)
	})
}
func installUserRouter(db gorm.DB, router *gin.Engine) {
	// GET /user
	// GET all users
	router.GET("/user", func(c *gin.Context) {
		var users []model.User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	// GET /user/:guid
	// GET user by guid
	router.GET("/user/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var user model.User
		query := &model.User{
			UserData: model.UserData{Guid: guid},
		}

		db.Where(&query).First(&user)

		c.JSON(http.StatusOK, user)
	})
}
func installUserFollowRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_follow
	// GET all user_follows
	router.GET("/user_follow", func(c *gin.Context) {
		var user_follows []model.UserFollow
		db.Find(&user_follows)
		c.JSON(http.StatusOK, user_follows)
	})

	// GET /user_follow/:guid
	// GET user_follow by guid
	router.GET("/user_follow/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var user_follow model.UserFollow
		query := &model.UserFollow{
			UserFollowData: model.UserFollowData{Guid: guid},
		}

		db.Where(&query).First(&user_follow)

		c.JSON(http.StatusOK, user_follow)
	})
}
func installUserHttpSessionRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_http_session
	// GET all user_http_sessions
	router.GET("/user_http_session", func(c *gin.Context) {
		var user_http_sessions []model.UserHttpSession
		db.Find(&user_http_sessions)
		c.JSON(http.StatusOK, user_http_sessions)
	})

	// GET /user_http_session/:guid
	// GET user_http_session by guid
	router.GET("/user_http_session/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var user_http_session model.UserHttpSession
		query := &model.UserHttpSession{
			UserHttpSessionData: model.UserHttpSessionData{Guid: guid},
		}

		db.Where(&query).First(&user_http_session)

		c.JSON(http.StatusOK, user_http_session)
	})
}
func installUserInviteRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_invite
	// GET all user_invites
	router.GET("/user_invite", func(c *gin.Context) {
		var user_invites []model.UserInvite
		db.Find(&user_invites)
		c.JSON(http.StatusOK, user_invites)
	})

	// GET /user_invite/:guid
	// GET user_invite by guid
	router.GET("/user_invite/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var user_invite model.UserInvite
		query := &model.UserInvite{
			UserInviteData: model.UserInviteData{Guid: guid},
		}

		db.Where(&query).First(&user_invite)

		c.JSON(http.StatusOK, user_invite)
	})
}
func installUserMentioningRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_mentioning
	// GET all user_mentionings
	router.GET("/user_mentioning", func(c *gin.Context) {
		var user_mentionings []model.UserMentioning
		db.Find(&user_mentionings)
		c.JSON(http.StatusOK, user_mentionings)
	})

	// GET /user_mentioning/:guid
	// GET user_mentioning by guid
	router.GET("/user_mentioning/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var user_mentioning model.UserMentioning
		query := &model.UserMentioning{
			UserMentioningData: model.UserMentioningData{Guid: guid},
		}

		db.Where(&query).First(&user_mentioning)

		c.JSON(http.StatusOK, user_mentioning)
	})
}
func installUserModuleRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_module
	// GET all user_modules
	router.GET("/user_module", func(c *gin.Context) {
		var user_modules []model.UserModule
		db.Find(&user_modules)
		c.JSON(http.StatusOK, user_modules)
	})

	// GET /user_module/:guid
	// GET user_module by guid
	router.GET("/user_module/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var user_module model.UserModule
		query := &model.UserModule{
			UserModuleData: model.UserModuleData{Guid: guid},
		}

		db.Where(&query).First(&user_module)

		c.JSON(http.StatusOK, user_module)
	})
}
func installUserPasswordRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_password
	// GET all user_passwords
	router.GET("/user_password", func(c *gin.Context) {
		var user_passwords []model.UserPassword
		db.Find(&user_passwords)
		c.JSON(http.StatusOK, user_passwords)
	})

	// GET /user_password/:guid
	// GET user_password by guid
	router.GET("/user_password/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var user_password model.UserPassword
		query := &model.UserPassword{
			UserPasswordData: model.UserPasswordData{Guid: guid},
		}

		db.Where(&query).First(&user_password)

		c.JSON(http.StatusOK, user_password)
	})
}
func installUserSettingRouter(db gorm.DB, router *gin.Engine) {
	// GET /user_setting
	// GET all user_settings
	router.GET("/user_setting", func(c *gin.Context) {
		var user_settings []model.UserSetting
		db.Find(&user_settings)
		c.JSON(http.StatusOK, user_settings)
	})

	// GET /user_setting/:guid
	// GET user_setting by guid
	router.GET("/user_setting/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var user_setting model.UserSetting
		query := &model.UserSetting{
			UserSettingData: model.UserSettingData{Guid: guid},
		}

		db.Where(&query).First(&user_setting)

		c.JSON(http.StatusOK, user_setting)
	})
}
func installWallRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall
	// GET all walls
	router.GET("/wall", func(c *gin.Context) {
		var walls []model.Wall
		db.Find(&walls)
		c.JSON(http.StatusOK, walls)
	})

	// GET /wall/:guid
	// GET wall by guid
	router.GET("/wall/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var wall model.Wall
		query := &model.Wall{
			WallData: model.WallData{Guid: guid},
		}

		db.Where(&query).First(&wall)

		c.JSON(http.StatusOK, wall)
	})
}
func installWallEntryRouter(db gorm.DB, router *gin.Engine) {
	// GET /wall_entry
	// GET all wall_entrys
	router.GET("/wall_entry", func(c *gin.Context) {
		var wall_entrys []model.WallEntry
		db.Find(&wall_entrys)
		c.JSON(http.StatusOK, wall_entrys)
	})

	// GET /wall_entry/:guid
	// GET wall_entry by guid
	router.GET("/wall_entry/:guid", func(c *gin.Context) {
		guid := c.Param("guid")
		var wall_entry model.WallEntry
		query := &model.WallEntry{
			WallEntryData: model.WallEntryData{Guid: guid},
		}

		db.Where(&query).First(&wall_entry)

		c.JSON(http.StatusOK, wall_entry)
	})
}

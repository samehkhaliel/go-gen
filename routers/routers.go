
func InstallActivityRouter(db gorm.DB, router *gin.Engine) {
	// GET /activity
	// GET all activity
	router.GET("/activity", func(c *gin.Context) {
		var activity []model.Activity
		db.Find(&activity)
		c.JSON(http.StatusOK, activity)
	})

	// GET /activity/:id
	// GET activity by id
	router.GET("/activity/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var activity model.Activity
			query := &model.Activity{
				ActivityId: model.ActivityId{Id: id},
			}
			db.Where(&query).First(&activity)
			c.JSON(http.StatusOK, activity)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /comment/:id
	// GET comment by id
	router.GET("/comment/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var comment model.Comment
			query := &model.Comment{
				CommentId: model.CommentId{Id: id},
			}
			db.Where(&query).First(&comment)
			c.JSON(http.StatusOK, comment)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /content/:id
	// GET content by id
	router.GET("/content/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var content model.Content
			query := &model.Content{
				ContentId: model.ContentId{Id: id},
			}
			db.Where(&query).First(&content)
			c.JSON(http.StatusOK, content)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /file/:id
	// GET file by id
	router.GET("/file/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var file model.File
			query := &model.File{
				FileId: model.FileId{Id: id},
			}
			db.Where(&query).First(&file)
			c.JSON(http.StatusOK, file)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /group/:id
	// GET group by id
	router.GET("/group/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var group model.Group
			query := &model.Group{
				GroupId: model.GroupId{Id: id},
			}
			db.Where(&query).First(&group)
			c.JSON(http.StatusOK, group)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /group_admin/:id
	// GET group_admin by id
	router.GET("/group_admin/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var group_admin model.GroupAdmin
			query := &model.GroupAdmin{
				GroupAdminId: model.GroupAdminId{Id: id},
			}
			db.Where(&query).First(&group_admin)
			c.JSON(http.StatusOK, group_admin)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /like/:id
	// GET like by id
	router.GET("/like/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var like model.Like
			query := &model.Like{
				LikeId: model.LikeId{Id: id},
			}
			db.Where(&query).First(&like)
			c.JSON(http.StatusOK, like)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /logging/:id
	// GET logging by id
	router.GET("/logging/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var logging model.Logging
			query := &model.Logging{
				LoggingId: model.LoggingId{Id: id},
			}
			db.Where(&query).First(&logging)
			c.JSON(http.StatusOK, logging)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /migration/:id
	// GET migration by id
	router.GET("/migration/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var migration model.Migration
			query := &model.Migration{
				MigrationId: model.MigrationId{Id: id},
			}
			db.Where(&query).First(&migration)
			c.JSON(http.StatusOK, migration)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /module_enabled/:id
	// GET module_enabled by id
	router.GET("/module_enabled/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var module_enabled model.ModuleEnabled
			query := &model.ModuleEnabled{
				ModuleEnabledId: model.ModuleEnabledId{Id: id},
			}
			db.Where(&query).First(&module_enabled)
			c.JSON(http.StatusOK, module_enabled)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /notification/:id
	// GET notification by id
	router.GET("/notification/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var notification model.Notification
			query := &model.Notification{
				NotificationId: model.NotificationId{Id: id},
			}
			db.Where(&query).First(&notification)
			c.JSON(http.StatusOK, notification)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /post/:id
	// GET post by id
	router.GET("/post/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var post model.Post
			query := &model.Post{
				PostId: model.PostId{Id: id},
			}
			db.Where(&query).First(&post)
			c.JSON(http.StatusOK, post)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /profile/:id
	// GET profile by id
	router.GET("/profile/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var profile model.Profile
			query := &model.Profile{
				ProfileId: model.ProfileId{Id: id},
			}
			db.Where(&query).First(&profile)
			c.JSON(http.StatusOK, profile)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /profile_field/:id
	// GET profile_field by id
	router.GET("/profile_field/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var profile_field model.ProfileField
			query := &model.ProfileField{
				ProfileFieldId: model.ProfileFieldId{Id: id},
			}
			db.Where(&query).First(&profile_field)
			c.JSON(http.StatusOK, profile_field)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /profile_field_category/:id
	// GET profile_field_category by id
	router.GET("/profile_field_category/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var profile_field_category model.ProfileFieldCategory
			query := &model.ProfileFieldCategory{
				ProfileFieldCategoryId: model.ProfileFieldCategoryId{Id: id},
			}
			db.Where(&query).First(&profile_field_category)
			c.JSON(http.StatusOK, profile_field_category)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /setting/:id
	// GET setting by id
	router.GET("/setting/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var setting model.Setting
			query := &model.Setting{
				SettingId: model.SettingId{Id: id},
			}
			db.Where(&query).First(&setting)
			c.JSON(http.StatusOK, setting)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /space/:id
	// GET space by id
	router.GET("/space/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var space model.Space
			query := &model.Space{
				SpaceId: model.SpaceId{Id: id},
			}
			db.Where(&query).First(&space)
			c.JSON(http.StatusOK, space)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /space_membership/:id
	// GET space_membership by id
	router.GET("/space_membership/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var space_membership model.SpaceMembership
			query := &model.SpaceMembership{
				SpaceMembershipId: model.SpaceMembershipId{Id: id},
			}
			db.Where(&query).First(&space_membership)
			c.JSON(http.StatusOK, space_membership)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /space_module/:id
	// GET space_module by id
	router.GET("/space_module/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var space_module model.SpaceModule
			query := &model.SpaceModule{
				SpaceModuleId: model.SpaceModuleId{Id: id},
			}
			db.Where(&query).First(&space_module)
			c.JSON(http.StatusOK, space_module)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /space_setting/:id
	// GET space_setting by id
	router.GET("/space_setting/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var space_setting model.SpaceSetting
			query := &model.SpaceSetting{
				SpaceSettingId: model.SpaceSettingId{Id: id},
			}
			db.Where(&query).First(&space_setting)
			c.JSON(http.StatusOK, space_setting)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /url_oembed/:id
	// GET url_oembed by id
	router.GET("/url_oembed/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var url_oembed model.UrlOembed
			query := &model.UrlOembed{
				UrlOembedId: model.UrlOembedId{Id: id},
			}
			db.Where(&query).First(&url_oembed)
			c.JSON(http.StatusOK, url_oembed)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /user/:id
	// GET user by id
	router.GET("/user/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var user model.User
			query := &model.User{
				UserId: model.UserId{Id: id},
			}
			db.Where(&query).First(&user)
			c.JSON(http.StatusOK, user)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /user_follow/:id
	// GET user_follow by id
	router.GET("/user_follow/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var user_follow model.UserFollow
			query := &model.UserFollow{
				UserFollowId: model.UserFollowId{Id: id},
			}
			db.Where(&query).First(&user_follow)
			c.JSON(http.StatusOK, user_follow)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /user_http_session/:id
	// GET user_http_session by id
	router.GET("/user_http_session/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var user_http_session model.UserHttpSession
			query := &model.UserHttpSession{
				UserHttpSessionId: model.UserHttpSessionId{Id: id},
			}
			db.Where(&query).First(&user_http_session)
			c.JSON(http.StatusOK, user_http_session)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /user_invite/:id
	// GET user_invite by id
	router.GET("/user_invite/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var user_invite model.UserInvite
			query := &model.UserInvite{
				UserInviteId: model.UserInviteId{Id: id},
			}
			db.Where(&query).First(&user_invite)
			c.JSON(http.StatusOK, user_invite)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /user_mentioning/:id
	// GET user_mentioning by id
	router.GET("/user_mentioning/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var user_mentioning model.UserMentioning
			query := &model.UserMentioning{
				UserMentioningId: model.UserMentioningId{Id: id},
			}
			db.Where(&query).First(&user_mentioning)
			c.JSON(http.StatusOK, user_mentioning)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /user_module/:id
	// GET user_module by id
	router.GET("/user_module/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var user_module model.UserModule
			query := &model.UserModule{
				UserModuleId: model.UserModuleId{Id: id},
			}
			db.Where(&query).First(&user_module)
			c.JSON(http.StatusOK, user_module)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /user_password/:id
	// GET user_password by id
	router.GET("/user_password/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var user_password model.UserPassword
			query := &model.UserPassword{
				UserPasswordId: model.UserPasswordId{Id: id},
			}
			db.Where(&query).First(&user_password)
			c.JSON(http.StatusOK, user_password)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /user_setting/:id
	// GET user_setting by id
	router.GET("/user_setting/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var user_setting model.UserSetting
			query := &model.UserSetting{
				UserSettingId: model.UserSettingId{Id: id},
			}
			db.Where(&query).First(&user_setting)
			c.JSON(http.StatusOK, user_setting)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /wall/:id
	// GET wall by id
	router.GET("/wall/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var wall model.Wall
			query := &model.Wall{
				WallId: model.WallId{Id: id},
			}
			db.Where(&query).First(&wall)
			c.JSON(http.StatusOK, wall)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
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

	// GET /wall_entry/:id
	// GET wall_entry by id
	router.GET("/wall_entry/:id", func(c *gin.Context) {
		_id := c.Param("id")
		if id, err := strconv.ParseUint(_id, 10, 64); err == nil {
			var wall_entry model.WallEntry
			query := &model.WallEntry{
				WallEntryId: model.WallEntryId{Id: id},
			}
			db.Where(&query).First(&wall_entry)
			c.JSON(http.StatusOK, wall_entry)
		} else {
			log.Print(err)
			c.AbortWithError(http.StatusBadRequest, err)
		}
	})
}

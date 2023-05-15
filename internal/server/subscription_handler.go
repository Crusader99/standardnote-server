package server

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type subscription struct {
	filesServerUrl string
}

func (h *subscription) SubscriptionV1(c echo.Context) error {
	user := currentUser(c)

	// The official Standard Notes client has a race condition,
	// the features endpoint will only be called when delaying response...
	time.Sleep(1 * time.Second)

	return c.JSON(http.StatusOK, echo.Map{
		"meta": echo.Map{
			"auth": echo.Map{
				"userUuid": user.ID,
				"roles": []echo.Map{
					{
						"uuid": "8047edbb-a10a-4ff8-8d53-c2cae600a8e8",
						"name": "PRO_USER",
					},
					{
						"uuid": "8802d6a3-b97c-4b25-968a-8fb21c65c3a1",
						"name": "CORE_USER",
					},
				},
			},
			"server": echo.Map{
				"filesServerUrl": h.filesServerUrl,
			},
		},
		"data": echo.Map{
			"success": true,
			"user": echo.Map{
				"uuid":  user.ID,
				"email": user.Email,
			},
			"subscription": echo.Map{
				"uuid":             "d4a65722-4f02-11ed-b7e0-0242ac12000a",
				"planName":         "PRO_PLAN",
				"endsAt":           8640000000000000,
				"createdAt":        0,
				"updatedAt":        0,
				"cancelled":        0,
				"subscriptionId":   1,
				"subscriptionType": "",
			},
		},
	})
}

func (h *subscription) Features(c echo.Context) error {
	user := currentUser(c)

	return c.JSON(http.StatusOK, echo.Map{
		"meta": echo.Map{
			"auth": echo.Map{
				"userUuid": user.ID,
				"roles": []echo.Map{
					{
						"uuid": "8047edbb-a10a-4ff8-8d53-c2cae600a8e8",
						"name": "PRO_USER",
					},
					{
						"uuid": "8802d6a3-b97c-4b25-968a-8fb21c65c3a1",
						"name": "CORE_USER",
					},
				},
			},
			"server": echo.Map{
				"filesServerUrl": h.filesServerUrl,
			},
		},
		"data": echo.Map{
			"success":  true,
			"userUuid": user.ID,
			"features": []interface{}{
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":             "Dark",
					"identifier":       "org.standardnotes.theme-focus",
					"permission_name":  "theme:focused",
					"clientControlled": true,
					"isDark":           true,
					"dock_icon": map[string]interface{}{
						"type":             "circle",
						"background_color": "#a464c2",
						"foreground_color": "#ffffff",
						"border_color":     "#a464c2",
					},
					"index_path":   "index.css",
					"content_type": "SN|Theme",
					"area":         "themes",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":            "Sign-in email alerts",
					"identifier":      "com.standardnotes.sign-in-alerts",
					"permission_name": "server:sign-in-alerts",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":            "Spreadsheet",
					"identifier":      "org.standardnotes.standard-sheets",
					"note_type":       "spreadsheet",
					"file_type":       "json",
					"interchangeable": false,
					"permission_name": "editor:sheets",
					"description":     "A powerful spreadsheet editor with formatting and formula support. Not recommended for large data sets, as encryption of such data may decrease editor performance.",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/spreadsheets.png",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path": "dist/index.html",
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type": "SN|Component",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"name":            "Autobiography",
					"identifier":      "org.standardnotes.theme-autobiography",
					"permission_name": "theme:autobiography",
					"dock_icon": map[string]interface{}{
						"type":             "circle",
						"background_color": "#9D7441",
						"foreground_color": "#ECE4DB",
						"border_color":     "#9D7441",
					},
					"index_path":   "index.css",
					"content_type": "SN|Theme",
					"area":         "themes",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":              "Minimal Markdown",
					"identifier":        "org.standardnotes.minimal-markdown-editor",
					"note_type":         "markdown",
					"file_type":         "md",
					"index_path":        "index.html",
					"permission_name":   "editor:markdown-minimist",
					"spellcheckControl": true,
					"deprecated":        true,
					"description":       "A minimal Markdown editor with live rendering and in-text search via Ctrl/Cmd + F",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/min-markdown.jpg",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PRO_PLAN",
					},
					"identifier":      "org.standardnotes.subscription-sharing",
					"permission_name": "server:subscription-sharing",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"name": "Tag Nesting",
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"identifier":      "org.standardnotes.tag-nesting",
					"permission_name": "app:tag-nesting",
					"description":     "Organize your tags into folders.",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":            "Midnight",
					"identifier":      "org.standardnotes.theme-midnight",
					"permission_name": "theme:midnight",
					"isDark":          true,
					"dock_icon": map[string]interface{}{
						"type":             "circle",
						"background_color": "#086DD6",
						"foreground_color": "#ffffff",
						"border_color":     "#086DD6",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path":   "index.css",
					"content_type": "SN|Theme",
					"area":         "themes",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"name":            "Dynamic Panels",
					"identifier":      "org.standardnotes.theme-dynamic",
					"permission_name": "theme:dynamic",
					"layerable":       true,
					"no_mobile":       true,
					"index_path":      "index.css",
					"content_type":    "SN|Theme",
					"area":            "themes",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":              "Markdown with Math",
					"identifier":        "org.standardnotes.fancy-markdown-editor",
					"spellcheckControl": true,
					"permission_name":   "editor:markdown-math",
					"note_type":         "markdown",
					"file_type":         "md",
					"deprecated":        true,
					"index_path":        "index.html",
					"description":       "A beautiful split-pane Markdown editor with synced-scroll, LaTeX support, and colorful syntax.",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/fancy-markdown.jpg",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PRO_PLAN",
					},
					"identifier":      "org.standardnotes.files-max-storage-tier",
					"permission_name": "server:files-max-storage-tier",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"name":            "Solarized Dark",
					"identifier":      "org.standardnotes.theme-solarized-dark",
					"permission_name": "theme:solarized-dark",
					"isDark":          true,
					"dock_icon": map[string]interface{}{
						"type":             "circle",
						"background_color": "#2AA198",
						"foreground_color": "#ffffff",
						"border_color":     "#2AA198",
					},
					"index_path":   "index.css",
					"content_type": "SN|Theme",
					"area":         "themes",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":              "Markdown",
					"identifier":        "org.standardnotes.advanced-markdown-editor",
					"note_type":         "markdown",
					"file_type":         "md",
					"permission_name":   "editor:markdown-pro",
					"spellcheckControl": true,
					"description":       "A fully featured Markdown editor that supports live preview, a styling toolbar, and split pane support.",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/adv-markdown.jpg",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path": "dist/index.html",
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"name":            "Futura",
					"identifier":      "org.standardnotes.theme-futura",
					"permission_name": "theme:futura",
					"isDark":          true,
					"dock_icon": map[string]interface{}{
						"type":             "circle",
						"background_color": "#fca429",
						"foreground_color": "#ffffff",
						"border_color":     "#fca429",
					},
					"index_path":   "index.css",
					"content_type": "SN|Theme",
					"area":         "themes",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"name":       "Super Notes",
					"identifier": "com.standardnotes.super-editor",
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"permission_name": "editor:super-editor",
					"description":     "A new way to edit notes. Type / to bring up the block selection menu, or @ to embed images or link other tags and notes. Type - then space to start a list, or [] then space to start a checklist. Drag and drop an image or file to embed it in your note. Cmd/Ctrl + F to bring up search and replace.",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"name": "Smart Filters",
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"identifier":      "org.standardnotes.smart-filters",
					"permission_name": "app:smart-filters",
					"description":     "Create smart filters for viewing notes matching specific criteria.",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":            "Listed Custom Domain",
					"identifier":      "org.standardnotes.listed-custom-domain",
					"permission_name": "listed:custom-domain",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":            "Authenticator",
					"note_type":       "authentication",
					"file_type":       "json",
					"interchangeable": false,
					"identifier":      "org.standardnotes.token-vault",
					"permission_name": "editor:token-vault",
					"description":     "Encrypt and protect your 2FA secrets for all your internet accounts. Authenticator handles your 2FA secrets so that you never lose them again, or have to start over when you get a new device.",
					"thumbnail_url":   "https://standard-notes.s3.amazonaws.com/screenshots/models/editors/token-vault.png",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path": "dist/index.html",
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type": "SN|Component",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":              "Markdown Alternative",
					"identifier":        "org.standardnotes.markdown-visual-editor",
					"note_type":         "markdown",
					"file_type":         "md",
					"deprecated":        true,
					"permission_name":   "editor:markdown-visual",
					"spellcheckControl": true,
					"description":       "A WYSIWYG-style Markdown editor that renders Markdown in preview-mode while you type without displaying any syntax.",
					"index_path":        "build/index.html",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":       "FileSafe",
					"identifier": "org.standardnotes.file-safe",
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
						map[string]interface{}{
							"name": "stream-items",
							"content_types": []interface{}{
								"SN|FileSafe|Credentials",
								"SN|FileSafe|FileMetadata",
								"SN|FileSafe|Integration",
							},
						},
					},
					"permission_name": "component:filesafe",
					"area":            "editor-stack",
					"deprecated":      true,
					"description":     "Encrypted attachments for your notes using your Dropbox, Google Drive, or WebDAV server. Limited to 50MB per file.",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/FileSafe-banner.png",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path":      "dist/index.html",
					"content_type":    "SN|Component",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"name": "Encrypted files",
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"identifier":      "org.standardnotes.files",
					"permission_name": "app:files",
					"description":     "",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"name":            "Titanium",
					"identifier":      "org.standardnotes.theme-titanium",
					"permission_name": "theme:titanium",
					"dock_icon": map[string]interface{}{
						"type":             "circle",
						"background_color": "#6e2b9e",
						"foreground_color": "#ffffff",
						"border_color":     "#6e2b9e",
					},
					"index_path":   "index.css",
					"content_type": "SN|Theme",
					"area":         "themes",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"identifier":      "org.standardnotes.daily-gdrive-backup",
					"permission_name": "server:daily-gdrive-backup",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"identifier":      "org.standardnotes.daily-onedrive-backup",
					"permission_name": "server:daily-onedrive-backup",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":            "Two factor authentication",
					"identifier":      "org.standardnotes.two-factor-auth",
					"permission_name": "server:two-factor-auth",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":              "Rich Text",
					"note_type":         "rich-text",
					"file_type":         "html",
					"identifier":        "org.standardnotes.plus-editor",
					"permission_name":   "editor:plus",
					"spellcheckControl": true,
					"description":       "From highlighting to custom font sizes and colors, to tables and lists, this editor is perfect for crafting any document.",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/plus-editor.jpg",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path": "dist/index.html",
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":              "Checklist",
					"identifier":        "org.standardnotes.simple-task-editor",
					"note_type":         "task",
					"spellcheckControl": true,
					"file_type":         "md",
					"interchangeable":   false,
					"permission_name":   "editor:task-editor",
					"description":       "A great way to manage short-term and long-term to-do\"s. You can mark tasks as completed, change their order, and edit the text naturally in place.",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/task-editor.jpg",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path": "dist/index.html",
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type": "SN|Component",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PRO_PLAN",
					},
					"name":            "Unlimited note history",
					"identifier":      "org.standardnotes.note-history-unlimited",
					"permission_name": "server:note-history-unlimited",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"identifier":      "org.standardnotes.daily-dropbox-backup",
					"permission_name": "server:daily-dropbox-backup",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":              "Code",
					"spellcheckControl": true,
					"identifier":        "org.standardnotes.code-editor",
					"permission_name":   "editor:code-editor",
					"note_type":         "code",
					"file_type":         "txt",
					"interchangeable":   true,
					"index_path":        "index.html",
					"description":       "Syntax highlighting and convenient keyboard shortcuts for over 120 programming languages. Ideal for code snippets and procedures.",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/code.jpg",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type": "SN|Component",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"no_expire":    false,
					"role_name":    "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":       "Alternative Rich Text",
					"identifier": "org.standardnotes.bold-editor",
					"note_type":  "rich-text",
					"file_type":  "html",
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
						map[string]interface{}{
							"name": "stream-items",
							"content_types": []interface{}{
								"SN|FileSafe|Credentials",
								"SN|FileSafe|FileMetadata",
								"SN|FileSafe|Integration",
							},
						},
					},
					"spellcheckControl": true,
					"deprecated":        true,
					"permission_name":   "editor:bold",
					"description":       "A simple and peaceful rich editor that helps you write and think clearly.",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/bold.jpg",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path":      "dist/index.html",
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":              "Basic Markdown",
					"identifier":        "org.standardnotes.simple-markdown-editor",
					"note_type":         "markdown",
					"spellcheckControl": true,
					"file_type":         "md",
					"deprecated":        true,
					"permission_name":   "editor:markdown-basic",
					"description":       "A Markdown editor with dynamic split-pane preview.",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/simple-markdown.jpg",
					"availableInRoles": []interface{}{
						"PLUS_USER",
						"PRO_USER",
					},
					"index_path": "dist/index.html",
					"component_permissions": []interface{}{
						map[string]interface{}{
							"name": "stream-context-item",
							"content_types": []interface{}{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
				map[string]interface{}{
					"availableInSubscriptions": []interface{}{
						"PLUS_PLAN",
						"PRO_PLAN",
					},
					"name":            "Email backups",
					"identifier":      "org.standardnotes.daily-email-backup",
					"permission_name": "server:daily-email-backup",
					"expires_at":      8640000000000000,
					"no_expire":       false,
					"role_name":       "PRO_USER",
				},
			},
		},
	})
}

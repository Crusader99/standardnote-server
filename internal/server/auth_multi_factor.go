package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mdouchement/standardfile/internal/database"
	"github.com/mdouchement/standardfile/internal/sferror"
)

type auth_multi_factor struct {
	db database.Client
}

// UpdateSettingParams are used to enable multi factor auth.
type UpdateSettingParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (h *auth_multi_factor) EnableMFA(c echo.Context) error {
	user := currentUser(c)
	if user == nil {
		return c.JSON(http.StatusUnauthorized, sferror.New("Unauthorized"))
	}

	var params UpdateSettingParams
	err := json.NewDecoder(c.Request().Body).Decode(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, sferror.New("Unable to parse JSON"))
	}

	if params.Name != "MFA_SECRET" {
		return c.JSON(http.StatusBadRequest, sferror.New("Unsupported setting: "+params.Name))
	}

	user.MultiFactorSecret = params.Value
	if err := h.db.Save(user); err != nil {
		return c.JSON(http.StatusBadRequest, sferror.New("Could not persist MultiFactorSecret"))
	}

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
		},
		"data": echo.Map{
			"setting": echo.Map{
				"uuid":      "7258d88c-5821-493f-ad4f-b88f6d06ee07",
				"name":      params.Name,
				"value":     "{\"version\":1,\"encrypted\":{\"iv\":\"e3a19317aceb476373b0ee2d137c23d9\",\"tag\":\"02c59a262e72a3cf0b9573962c46365e\",\"aad\":\"\",\"ciphertext\":\"fPd8KvxCr01lKr4x2vXfCeHy3TmovB46cDPxXmkIpmI=\",\"encoding\":\"utf-8\"}}",
				"createdAt": time.Now().UnixMilli(),
				"updatedAt": time.Now().UnixMilli(),
				"sensitive": 1,
			},
		},
	})
}

func (h *auth_multi_factor) CheckHasEnabledMFA(c echo.Context) error {
	user := currentUser(c)

	if user.MultiFactorSecret == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
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
			},
			"data": echo.Map{
				"success": false,
				"error": echo.Map{
					"message": "Setting mfa_secret for user " + user.ID + " not found!",
				},
			},
		})
	}
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
		},
		"data": echo.Map{
			"success":   true,
			"sensitive": true,
		},
	})
}

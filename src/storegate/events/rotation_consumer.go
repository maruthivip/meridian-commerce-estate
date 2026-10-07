// Package events consumes promo rotation notifications in storegate.
package events

import "encoding/json"

// Subscribes to topic promo-rotation-events published by promobeam so the
// storefront can invalidate its cached promo slots on rotation changes.
const promoRotationTopic = "promo-rotation-events"

type rotationEvent struct {
	RotationID string   `json:"rotation_id"`
	AdIDs      []string `json:"ad_ids"`
}

// HandlePromoRotation processes one message from promo-rotation-events.
func HandlePromoRotation(payload []byte) (string, error) {
	var ev rotationEvent
	if err := json.Unmarshal(payload, &ev); err != nil {
		return "", err
	}
	return ev.RotationID, nil
}

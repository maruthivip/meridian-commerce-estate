// Package events publishes promo rotation notifications for promobeam.
package events

import "encoding/json"

// Topic promo-rotation-events: emitted whenever the active ad rotation
// changes. Subscribers refresh their cached promo slots.
const PromoRotationTopic = "promo-rotation-events"

type RotationEvent struct {
	RotationID string   `json:"rotation_id"`
	AdIDs      []string `json:"ad_ids"`
}

// PublishRotation serialises the event for the promo-rotation-events topic.
func PublishRotation(ev RotationEvent) ([]byte, string, error) {
	payload, err := json.Marshal(ev)
	return payload, PromoRotationTopic, err
}

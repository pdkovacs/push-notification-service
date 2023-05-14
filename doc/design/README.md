# Push Nofitication Service - PNS

Conceptually, PNS consists of one or multiple processing nodes. Each node is managing a subset of the entire set of websocket-connections which have been created on behalf (in the context) of the application of which PNS is a subsystem.

Users can connect and remain connected to PNS with multiple devices. (The "device" in this context translates in practice into distinct "cookied" sessions, i.e. each distinct client session counts as a distinct device.) A user is notified of the availability of new notifications on each of its active devices ("active" meaning connected to the application). The user may request to see the data of the new notifications (or can wait until some more are available). Once the new notification data is delivered to the user at any of their active devices,
1. the new notifications are recategorized to be "old";
2. the new status of the new notifications (typically "there are no new notifications right now") is, ideally, propagated to all active devices.

Notifications must be kept available until the user requests to see them (i.e. potentially for an infinite period of time) and even for some configurable period thereafter. 

## New client connection

When a user connects with a device to a PNS node, the node

1. authenticates the user
2. sends the current `new-notifications-count` to the user's newly connected device

The PNS nodes must use the same authentication system as the application. 

## Notification

Notifications are created by the application following application-specific logic/requirements. Creating a notification consists of the following steps:

The application

1. creates a record in the `notifications` table with the following fields:
   * `notification_id`
   * `notification_data`
   * `user_id`
   *   `creation_date`
   * `seen_date` (set NULL initially)
2. sends a notification about the new notification to all PNS nodes (via, most probably, a *message broker* of some sort) with the `notification_id`.

Each of the PNS nodes

1. checks whether at least one device of the user to whom the notification is destined is connected to it,
2. if it is,
   1. looks up the number of notifications still unseen by the user and
   2. sends the unseen count via each WS connection maintained by the node for the given user

## Client re-connection

The same as a new client connection

## Client requests unseen notification data

1. The PNS node collects the data
2. sends the data the client
3. once the data is delivered the `seen_date` of the corresponding notification record is set to the current date-time.

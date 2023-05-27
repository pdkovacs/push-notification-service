## Overview

In order to make the service highly available, PNS is composed of one or multiple processing nodes. Each node is managing
a subset of the entire set of websocket-connections which have been created on behalf (in the context) of the
application of which PNS is a subsystem. Asynchronous communications between the PSN nodes are supported by some event
delivery system (via the `unseen-notification-count-changed` event).

## Subscription requests from a user device

From the user devices' perspective the connection is made to the PSN. Connecting the user devices to specific nodes
behind the service is handled by whatever _Service_ facility is available with the underlying clustering/cloud provider
(e.g. Service in K8S, Target Group/ALB in AWS ECS/EKS).

When a user device connects to PNS, the PNS node to which the _the platform_ assigned the client's connection performs
the following steps:

1. authenticates the user in a deployment-dependent manner (see [Deployment options](#deployment-options));
2. sends the `newNotificationCountChanged` message to the newly connected device

---
**NOTE**

User device _re-connection_ / _re-subscription_ is handled in the same manner as a new subcription.

---

## Notification

Notifications are created by the application following application-specific logic/requirements. Creating a notification
consists of the following steps:

The application

1. creates a message with the following fields:
   * `notification_id`
   * `notification_data`
   * `user_id`
   * `creation_date`
   * `seen_date` (set `NULL` initially)
2. sends the message to PNS in a deployment-dependent way (see [Deployment options](#deployment-options));
3. the PNS node receiving the message
      1. persists the message (inserts it into `notifications` table)
      2. publishes an `unseen-notification-count-changed` event (along with the `user_id` of the user to whom the notification is addressed)
      3. performs the same step as if it had itself received the `unseen-notification-count-changed` event.

Each of the PNS nodes receiving the `unseen-notification-count-changed` event performs the following steps:
      1. looks up, in the persistent store, the number of notifications not-yet-seen by the user associated with the
         `user_id` included in the even,
      3. sends the `newNotificationCountChanged` message to the devices of the target user it is connected with.

## Client request for not-yet-seen notification data

The PNS node receiving the request

1. collects the data into a `newNotificationData` message;
2. sends the `newNotificationData` message to the client;
3. sets the `seen_date` of the corresponding notification record in the persistent store to the current date-time;
4. publishes an `unseen-notification-count-changed` event with the user's `user_id`;
5. the PNS nodes receiving the `unseen-notification-count-changed` event
   1. count the number of those notifications sent destined to the user that have their `seen_date` property set to
      `NULL` and send a `newNotificationCountChanged` message to their clients with count just obtained.
6. counts the number of those notifications sent destined to the user that have their `seen_date` property set to `NULL`
   and send a `newNotificationCountChanged` message to their clients with count just obtained.

## PNS node comes on-line

Each PNS node coming on-line

1. Subscribe to the `unseen-notification-count-changed` event.

## PNS node goes off-line

PNS nodes going off-line will be implicitly handled by the devices connected to them: the devices will (try to) reconnect to the PNS.

## Deployment options 

dPSN can be deployed as a separate service or as a component embedded into the application. In the latter case, the application nodes also serve as PSN nodes.

### PSN nodes embedded in the application nodes

1. Send notification

   The application calls the `notify()` method on the PNS library's API.

2. Authenticate user device subcription requests

   The application's `/subscribe` endpoint will do authentication as the application's all other endpoints. The implementation of the `/subscribe` endpoint
   will call the `subscribe()` method on the PNS library's API.

### PSN as a separate service

1. Send notification

   The application calls the `/notify` PNS endpoint in case PNS is deployed as a seperate service. PNS supports
   authenticating the application using OAuth2 client-credentials flow.

2. Authenticate user device subcription request

   PNS calls the application's `/subscribe` endpoint with the same headers (including cookie headers) as PNS was called
   by the user device. In this deployment mode, the application's `/subscribe` endpoint stops processing the request
   after authentication.


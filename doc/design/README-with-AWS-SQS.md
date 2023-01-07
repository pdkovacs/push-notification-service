# Push Nofitication Service - PNS / With AWS SQS
---
*Provide solution for making **pending** notifications available indefinitely*
---
---


## New client connection

The PNS nodes must use the same authentication system as the application. When a user connects with a device to a PNS node, the node
1. authenticates the user
2. generates a UUID called `connection_id`
3. creates a record in the `notification_registry` composed of
   * the `connection_id`,
   * the `user_id` established during authentication and
   * the `node_uri` which is PNS node's URL or URI or some similar identifier suitable to locate the node in the given operating environment
4. includes the `connection_id` in the `Set-Cookie` response header for the PNS cookie
5. ***creates a queue for the `connection_id`***
6. includes the `connection_id` in the `Set-Cookie` response header for the PNS cookie

The node maintains (cashes) a mapping of the `connection_id` to websocket-connection application objects.

## Notification

After the application creates a notification as an application-object, it

1. ~~creates a record of~~
    * ~~`notification_id`~~
    * ~~`notification_data`~~
    ~~in the `notification_data` table~~
2. looks up the `connection_id`s registered for the user in the `notification_regisry`
3. ***sends the notification data to the queue associated with the `connection_id`.***
4. ~~creates a record of~~
   * ~~`connection_id`~~
   * ~~`notification_id`~~
   ~~in the `pending_notifications` table using the `notification_id` field from the `notification_data` table and the `connection_id`; the record is composed of~~
5. ~~looks up the `node_uri`s associated with the given user in the `notification_registry`~~
6. ~~sends a notification about the new notification to the URIs found (using the new notification's ID)~~


Each of the PNS nodes read all queues associated with the `connection_id`s mainteined by each node respectively.

1. ~~looks up the notification data by the `notification_id` in the `notification_data` table,~~
2. ~~looks up the relevant `connection_id`s in the `pending_notifications` table by the `notification_id`~~
3. sends the notification data via the websocket for each `connection_id` ~~which it (the given node) maintains~~ ***queue from which it receives a message***
4. removes ~~the record with the processed `notification_id` and `connection_id` from the `pending_notifications` table~~ ***the processed messages from the queues***.

For reference, the `pending_notifications` table includes the following fields:

## Client re-connection

1. In case the node to which the client re-connected didn't maintain the `connection_id` in the `Cookie` header, the node updates the corresponding record of the `notification_registry`.
2. For each record in the `pending_notifications` table having the same `connection_id` as the connection being re-established,
   
   1. looks up the notification data corresponding to the `notification_id` of the pending notification in the `notification_data` table
   2. sends the notification data via the websocket for each connection_id which it (the given node) maintains
   3.  removes ~~the record with the processed `notification_id` and `connection_id` from the `pending_notifications` table~~ ***the processed messages from the queues***.

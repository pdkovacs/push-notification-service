# Push Notification Service - PNS

PNS is an HTTP _service_ accepting (a) WS-connection requests from an _application_'s HTTP-clients and (b) notifications to be instantantly delivered (_pushed_) from the application to the application's _users_.

## User/client-facing functions

* User devices can _connect_ at the `POST /subscribe` endpoint and remain connected to PNS with one _device_ or multiple devices
* Each connected device of each user _is notified_ (by PNS) of the availability of new notifications from the application via the
    
    ```
    { "newNotificationCountChanged": number }
    ```
    
    message. (The content (data) of the new notifications is not sent at this point.)
* User devices can send _request_ to PNS to see the data of the new notifications (or can wait until some more new notifications are available) using the
    
    ```
    unseenNotificationDataRequest
    ```

    message. The response is sent to devices via the 

    ```
    { "newNotificationData": [{ notification_data: string, create_date: string }] }
    ```

    message

* User devices can request (a range of) already seen notification data using the

    ```
    { seenNotificationsRequest: { receivedAfter: string } }
    ```

    message. The response is sent via the

    ```
    { "seenNotificationData": [{ notification_data: string, create_date: string, seen_date: string }] }
    ```

    message.


---
**NOTE**

"device" in this context translates in practice into distinct "cookied" sessions, i.e. each distinct client session
counts as a distinct device.

---

---
**NOTE**

Notifications must be kept available until the user requests to see them (i.e. potentially for an infinite period of
time) and even for some configurable period thereafter. 

---


## Application-facing functions

* PNS delegates requests to authenticate a connection request to the application. 
For successful authentications, the application returns in its response the `user_id` of the authenticated user to whom the connecting device belongs.
* The application requests PNS to send notifications to a user's connected devices.

The service can be deployed either as a service separate from the application or embedded in the application as library.


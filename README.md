# Push Notification Service - PNS

PNS is an HTTP _service_ accepting

* WS-connection requests from an _application_'s HTTP-clients and
* notifications to be instantantly delivered (_pushed_) from the application to the application's _users_.

## User/client-facing functions

* User devices can _connect_ at the `POST /subscribe` endpoint and remain connected to PNS with one _device_ or multiple devices
  * the POST request carries a token (session-cookie by default) associated with the credentials of the user which is checked by the application;
  * in case the token refers to valid credentials of a user who is allowed to use the notification service, the endpoint returns the `user_id`
* Each connected device of each user _is notified_ (by PNS) of the availability of new notifications from the application via the
    
    ```
    { "newNotificationCountChanged": number }
    ```
    
    message. (The content (data) of the new notifications is not sent at this point.)
* User devices can request the notification data using the

    ```
    { notificationDataRequest: { receivedAfter: string } }
    ```

    message. If the `receivedAfter` property is empty, the response will include only data for notications not yet seen. The response is sent via the

    ```
    { "notificationData": [{ notification_data: string, created_date: string, seen_date: string }] }
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

* The application can request PNS to send notifications to a user's connected devices using the
  
  ```
  { user_id: string, notifications: [{ data: string, created_date: string }] }
  ```

message.

The service can be deployed either as a service separate from the application or embedded in the application as library.


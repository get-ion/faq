**A session can be defined as a server-side storage of information that is desired to persist throughout the user's interaction with the web application.**

Instead of storing large and constantly changing data via cookies in the user's browser (i.e. CookieStore), **only a unique identifier is stored on the client side called a "session id"**. This session id is passed to the web server on every request. The web application uses the session id as the key for retrieving the stored data from the database/memory. The session data is then available inside the session storage, memory or/and backend.

------

Example Code:

- https://github.com/get-ion/sessions/tree/master/_examples/overview
- https://github.com/get-ion/sessions/tree/master/_examples/flash-messages
- https://github.com/get-ion/sessions/tree/master/_examples/securecookie

> Look more at: https://github.com/get-ion/sessions/tree/master/_examples
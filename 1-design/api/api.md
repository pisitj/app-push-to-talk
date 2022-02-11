# API
### Auth
- POST /signup
- POST /login
- GET /check-login
- GET /check-admin
- GET /check-whitelist/push-to-talk
- GET /logout

### Backend
- GET /chat
- POST /chat
- GET /chat/{chat_id}/message
- POST /chat/{chat_id}/message
- GET /push-to-talk/chat/{chat_id}/message/{message_id}
- POST /push-to-talk/chat/{chat_id}/message
- GET /push-to-talk/report
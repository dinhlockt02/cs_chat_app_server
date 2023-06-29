# List of system message
## GroupCreated
- "system_event": "group-created"
- "sender": the person who created the group
## MemberJoined
- "system_event": "member-joined"
- "sender": the person who joined the group
## MemberLeaved
- "system_event": "member-leaved"
- "sender": the person who leaved the group
## Example
```json
{
  "id": "649d955fd42ab84c02e4641f",
  "type": "system",
  "sender": {
    "id": "649d864a58481de7979c7506",
    "avatar": "",
    "email": "test@gmail.com",
    "name": ""
  },
  "group": {
    "id": "649d955fd42ab84c02e4641e",
    "name": "group-name-1",
    "image_url": "https://file-examples.com/wp-content/storage/2017/10/file_example_JPG_100kB.jpg",
    "type": "group"
  },
  "message": "",
  "system_event": "group-created",
  "created_at": "2023-06-29T14:29:51.066Z",
  "is_me": true
}
```

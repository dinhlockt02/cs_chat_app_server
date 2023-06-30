# List of notification type
## AcceptFriendRequest
- id = 1
- actionType = "accept-request"
- channel = "basic_channel"
- action = []
- object relation: the subject accept the indirect (aka owner)'s friend request
## ReceiveFriendRequest
- id = 2
- actionType = "receive-friend-request"
- channel = "basic_channel"
- action = ["accept", "deny"]
- object relation: the Subject (aka owner) received the friend request (Direct) from Prep's
## ReceiveGroupRequest
- id = 3
- actionType = "receive-group-request"
- channel = "basic_channel"
- action = ["accept", "deny"]
- object relation: the Subject (aka owner) received the group request (Direct) to Group (Indirect) from Prep's

# ObjectType
- "user"
- "request"
- "group"

# 
 


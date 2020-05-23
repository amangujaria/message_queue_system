# Message Queue system

**Usage**
- Start a server as `./start.sh server`
- Start a client as `./start.sh client`

Following functionalities are exposed at client and can be executed via the shell thrown by start command:
- `subscribe TopicName`
    - subscribe to messages flowing in for `TopicName` topic
- `publish TopicName Message` or `publish TopicName Long Message Separated by Spaces`
    - publishes the message to the clients subscribed to `TopicName` topic
- `list`
    - shows the list of topics the client is currently subscribed to

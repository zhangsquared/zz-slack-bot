# zz-slack-bot

Slack bot
1. When "ping" is captured, reply "pong"

![Ping](./doc/ping.png)

2. When the bot is mentioned, also reply

![Mention](./doc/mention.png)

## local devlopment

```bash
go tool air
```

### Slack setup

> Step 1: Create an App on Slack

Sign in to the [Slack API](https://api.slack.com/apps) site.</br>
Click "Create New App" and choose "From scratch".</br>
App name: zz-chat-bot</br>
Workspace: zhangsquared</br>

> Step 2: Configure Permissions (Scopes)

In the left menu, go to "OAuth & Permissions". In the "Scopes" section add the following scopes:
- `chat:write` — allows the bot to send messages.
- `app_mentions:read` — allows the bot to see when it's mentioned.
- `channels:history` — allows the bot to read channel messages (if needed).

Got a Bot User OAuth Token (starts with `xoxb-`)

> Step 3: Obtain an App-Level Token

In the left menu, click "Basic Information".</br>
Find "App-Level Tokens" and click "Generate Token and Scopes".</br>
Give it any name and add the `connections:write` scope.</br>
The generated token will start with `xapp-`; this is required for Socket Mode.

> Step 4: Enable Socket Mode

To allow our local machine to receive Slack events without configuring a public URL (webhook):

In the Slack App settings left menu click "Socket Mode" and set it to Enable.</br>
On the "Event Subscriptions" page set "Enable Events" to On.</br>
On the same page, expand "Subscribe to bot events" and add `app_mention` and `message.channels`.</br>

> Step 5: Start the bot program

Since the connection is using WebSocket, we just spin up the local process. Invite the bot to a channel.

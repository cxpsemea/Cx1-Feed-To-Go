# Cx1-Feed-To-Go
Simple golang webhook listener accepting Slack-format messages from CheckmarxOne

This is a very minimal example of a webserver listening for Slack-format triggers from a Feedback App in CheckmarxOne. It does not do any logic to process the events, it will simply output to the console whatever data that is sent by the feedback app. The Slack-format feedback app in CheckmarxOne was selected for this due to the simple-to-parse JSON format message.

If you want to use this for testing with a public cloud deployment of CheckmarxOne (like in the public multi-tenant environment) you will need to find a way to expose port 80 to the internet, a simple way to do that for testing purposes is via ngrok. A "secret" is needed to connect everything, in this example the secret is "herpaderp" but can be changed in the code.

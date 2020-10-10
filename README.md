## Hermes

A slackbot for use at home to manage interactions with smart devices, integration with trello to do lists and tracking, reminders and more. 

### Feature - Food Shopping Trello Integration

To track our grocery shopping list, we use Trello - but to remind us to eat healthily, we also have a Slack reminder in our home workspace. It would be handy if we could respond in the channel/thread where this reminder appears, with dinner choices for the week, and a Slackbot autopopulated our shopping list based on those choices. 

This slack bot should be able to: 
- Add a specific item to the shopping list. 
- Remove a specific item from the shopping list. 
- Given a recipe name, lookup recipe in backend storage service and populate shopping list based on ingredients listed there. 
- Print out the current contents of the shopping list. 

## Example Usage

There are some required Environment Variables: 
```
export TRELLO_APPKEY=<INSERT_YOUR_DEVELOPER_KEY_HERE>
export TRELLO_TOKEN=<INSERT_YOUR_TRELLO_AUTH_TOKEN_HERE>
```

Build the CLI
```
make build
```

To get more info on available commands: 
```
./leto help
```

To get more info on a specific subcommand:
```
./leto trello help
```

To get a specific board from Trello
```
/leto trello --user=pamelakelly7 board --name="History"
```